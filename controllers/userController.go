package controllers

import (
	"net/http"

	"github.com/farrelahmady/my-gram-hacktiv8/database"
	"github.com/farrelahmady/my-gram-hacktiv8/helpers"
	"github.com/farrelahmady/my-gram-hacktiv8/models"
	"github.com/gin-gonic/gin"
)

var (
	appJson = "application/json"
)

// Login godoc
// @Summary Login
// @Description Login
// @Tags User
// @Accept  json
// @Produce  json
// @Param email formData string true "Email"
// @Param password formData string true "Password"
// @Success 200 {object} models.User
// @Failure 401 {object} string
// @Router /auth/login [post]
func Login(c *gin.Context) {
	db := database.GetDB()
	contentType := c.Request.Header.Get("Content-Type")
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJson {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthenticated",
			"message": "Invalid email or password",
		})
		return
	}

	if !helpers.ComparePass([]byte(User.Password), []byte(password)) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthenticated",
			"message": "Invalid email or password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// Register godoc
// @Summary Register
// @Description Register
// @Tags User
// @Accept  json
// @Produce  json
// @Param email formData string true "Email"
// @Param password formData string true "Password"
// @Param username formData string true "Username"
// @Param age formData int true "Age"
// @Success 200 {object} models.User
// @Failure 400  {object}  string
// @Router /auth/register [post]
func Register(c *gin.Context) {
	db := database.GetDB()
	contentType := c.Request.Header.Get("Content-Type")
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJson {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  map[string]interface{}{"id": User.ID, "email": User.Email, "username": User.Username, "age": User.Age},
	})
}
