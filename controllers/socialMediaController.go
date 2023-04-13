package controllers

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/farrelahmady/my-gram-hacktiv8/database"
	"github.com/farrelahmady/my-gram-hacktiv8/models"
	"github.com/gin-gonic/gin"
)

// GetAllSocialMedias godoc
// @Summary Get all social media
// @Description Get all social media
// @Tags SocialMedias
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.SocialMedia
// @Failure 400 {object} string
// @Router /socialMedias [get]
func GetAllSocialMedias(c *gin.Context) {
	db := database.GetDB()
	SocialMedia := []models.SocialMedia{}

	err := db.Debug().Find(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

// GetSocialMedia godoc
// @Summary Get a social media
// @Description Get a social media
// @Tags SocialMedias
// @Accept  json
// @Produce  json
// @Param id path int true "SocialMedia ID"
// @Success 200 {object} models.SocialMedia
// @Failure 400 {object} string
// @Router /socialMedias/{id} [get]
func GetSocialMedia(c *gin.Context) {
	db := database.GetDB()
	SocialMedia := models.SocialMedia{}

	err := db.Debug().Where("id = ?", c.Param("id")).Take(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

// CreateSocialMedia godoc
// @Summary Create a social media
// @Description Create a social media
// @Tags SocialMedias
// @Accept  json
// @Produce  json
// @Param socialMedia body models.SocialMedia true "SocialMedia"
// @Success 200 {object} models.SocialMedia
// @Failure 400 {object} string
// @Router /socialMedias [post]
func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	SocialMedia := models.SocialMedia{}
	UserID := uint(c.MustGet("user").(jwt.MapClaims)["id"].(float64))

	if c.Request.Header.Get("Content-Type") == appJson {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = UserID
	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

// UpdateSocialMedia godoc
// @Summary Update a social media
// @Description Update a social media
// @Tags SocialMedias
// @Accept  json
// @Produce  json
// @Param id path int true "SocialMedia ID"
// @Param socialMedia body models.SocialMedia true "SocialMedia"
// @Success 200 {object} models.SocialMedia
// @Failure 400 {object} string
// @Router /socialMedias/{id} [put]
func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	SocialMedia := models.SocialMedia{}

	err := db.Debug().Where("id = ?", c.Param("id")).Take(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if c.Request.Header.Get("Content-Type") == appJson {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	err = db.Debug().Save(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

// DeleteSocialMedia godoc
// @Summary Delete a social media
// @Description Delete a social media
// @Tags SocialMedias
// @Accept  json
// @Produce  json
// @Param id path int true "SocialMedia ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /socialMedias/{id} [delete]
func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	SocialMedia := models.SocialMedia{}

	err := db.Debug().Where("id = ?", c.Param("id")).Take(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	err = db.Debug().Delete(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "SocialMedia deleted",
	})
}
