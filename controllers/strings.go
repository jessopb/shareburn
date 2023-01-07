package controllers

import (
	"net/http"
	"poof/models"

	"github.com/Pallinder/go-randomdata"
	"github.com/gin-gonic/gin"
)

type PutStringInput struct {
	StringValue string `json:"string_value" binding:"required"`
}

func GetStrings(c *gin.Context) {
	var allstrings []models.SaveString
	models.DB.Find(&allstrings)

	c.JSON(http.StatusOK, gin.H{"data": allstrings})
}

func GetString(c *gin.Context) {
	var savedString models.SaveString

	if err := models.DB.Where("string_key = ?", c.Param("key")).First(&savedString).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not found"})
		return
	}
	models.DB.Where("string_key = ?", c.Param("key")).Delete(&savedString)
	c.IndentedJSON(http.StatusOK, gin.H{"data": savedString})
}

func PutString(c *gin.Context) {
	var input PutStringInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not found"})
		return
	}

	key := randomdata.SillyName()
	newString := models.SaveString{StringKey: key, StringValue: input.StringValue}
	models.DB.Create(&newString)
	c.String(http.StatusOK, key+"\n")
}
