package controllers

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/farrelahmady/my-gram-hacktiv8/database"
	"github.com/farrelahmady/my-gram-hacktiv8/models"
	"github.com/gin-gonic/gin"
)

func GetAllPhotos(c *gin.Context) {
	db := database.GetDB()
	Photo := []models.Photo{}

	err := db.Debug().Find(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)
}

func GetPhoto(c *gin.Context) {
	db := database.GetDB()
	Photo := models.Photo{}

	err := db.Debug().Where("id = ?", c.Param("id")).Take(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)
}

func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	Photo := models.Photo{}
	UserID := uint(c.MustGet("user").(jwt.MapClaims)["id"].(float64))

	if c.Request.Header.Get("Content-Type") == appJson {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = UserID
	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Photo)
}

func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	Photo := models.Photo{}

	if c.Request.Header.Get("Content-Type") == appJson {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err := db.Debug().Where("id = ?", c.Param("id")).Updates(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)
}

func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	Photo := models.Photo{}

	err := db.Debug().Where("id = ?", c.Param("id")).Delete(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Photo deleted",
	})
}
