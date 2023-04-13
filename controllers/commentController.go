package controllers

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/farrelahmady/my-gram-hacktiv8/database"
	"github.com/farrelahmady/my-gram-hacktiv8/models"
	"github.com/gin-gonic/gin"
)

// GetAllComments godoc
// @Summary Get all comments
// @Description Get all comments
// @Tags Comments
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Comment
// @Failure 400 {object} string
// @Router /comments [get]
func GetAllComments(c *gin.Context) {
	db := database.GetDB()
	Comment := []models.Comment{}

	err := db.Debug().Find(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)
}

// GetComment godoc
// @Summary Get a comment
// @Description Get a comment
// @Tags Comments
// @Accept  json
// @Produce  json
// @Param id path int true "Comment ID"
// @Success 200 {object} models.Comment
// @Failure 400 {object} string
// @Router /comments/{id} [get]
func GetComment(c *gin.Context) {
	db := database.GetDB()
	Comment := models.Comment{}

	err := db.Debug().Where("id = ?", c.Param("id")).Take(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)
}

// CreateComment godoc
// @Summary Create a comment
// @Description Create a comment
// @Tags Comments
// @Accept  json
// @Produce  json
// @Param comment body models.Comment true "Comment"
// @Success 201 {object} models.Comment
// @Failure 400 {object} string
// @Router /comments [post]
func CreateComment(c *gin.Context) {
	db := database.GetDB()
	Comment := models.Comment{}
	UserID := uint(c.MustGet("user").(jwt.MapClaims)["id"].(float64))

	if c.Request.Header.Get("Content-Type") == appJson {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = UserID
	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Comment)
}

// UpdateComment godoc
// @Summary Update a comment
// @Description Update a comment
// @Tags Comments
// @Accept  json
// @Produce  json
// @Param id path int true "Comment ID"
// @Param comment body models.Comment true "Comment"
// @Success 200 {object} models.Comment
// @Failure 400 {object} string
// @Router /comments/{id} [put]
func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	Comment := models.Comment{}

	if c.Request.Header.Get("Content-Type") == appJson {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	err := db.Debug().Where("id = ?", c.Param("id")).Updates(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)
}

// DeleteComment godoc
// @Summary Delete a comment
// @Description Delete a comment
// @Tags Comments
// @Accept  json
// @Produce  json
// @Param id path int true "Comment ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /comments/{id} [delete]
func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	Comment := models.Comment{}

	err := db.Debug().Where("id = ?", c.Param("id")).Delete(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comment deleted",
	})
}
