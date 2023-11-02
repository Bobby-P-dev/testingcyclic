package controllers

import (
	"net/http"

	"github.com/Bobby-P-dev/testingcyclic/database"
	"github.com/Bobby-P-dev/testingcyclic/helpers"
	"github.com/Bobby-P-dev/testingcyclic/models"
	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func CreateData(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	Product := models.Product{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}
	err := db.Debug().Create(&Product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "bad request",
			"mesage": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": Product,
	})
}

func GetData(c *gin.Context) {
	db := database.GetDB()
	var Product []models.Product

	err := db.Find(&Product).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "data not found",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Product,
	})
}
