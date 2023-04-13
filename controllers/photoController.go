package controllers

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/farrelahmady/my-gram-hacktiv8/database"
	"github.com/farrelahmady/my-gram-hacktiv8/models"
	"github.com/gin-gonic/gin"
)

// GetAllPhotos godoc
// @Summary Get all photos
// @Description Get all photos
// @Tags Photos
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Photo
// @Failure 400 {object} string
// @Router /photos [get]
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

// GetPhoto godoc
// @Summary Get a photo
// @Description Get a photo
// @Tags Photos
// @Accept  json
// @Produce  json
// @Param id path int true "Photo ID"
// @Success 200 {object} models.Photo
// @Failure 400 {object} string
// @Router /photos/{id} [get]
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

// CreatePhoto godoc
// @Summary Create a photo
// @Description Create a photo
// @Tags Photos
// @Accept  json
// @Produce  json
// @Param title formData string true "Photo Title"
// @Param photo_url formData string true "Photo URL"
// @Param caption formData string true "Photo Caption"
// @Success 201 {object} models.Photo
// @Failure 400 {object} string
// @Router /photos [post]
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

// UpdatePhoto godoc
// @Summary Update a photo
// @Description Update a photo
// @Tags Photos
// @Accept  json
// @Produce  json
// @Param id path int true "Photo ID"
// @Param title formData string true "Photo Title"
// @Param photo_url formData string true "Photo URL"
// @Param caption formData string true "Photo Caption"
// @Success 200 {object} models.Photo
// @Failure 400 {object} string
// @Router /photos/{id} [put]
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

// DeletePhoto godoc
// @Summary Delete a photo
// @Description Delete a photo
// @Tags Photos
// @Accept  json
// @Produce  json
// @Param id path int true "Photo ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /photos/{id} [delete]
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
