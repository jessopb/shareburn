package main

import (
	"net/http"
	"poof/controllers"
	"poof/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/poofs", controllers.GetStrings)
	r.GET("/poofs/:key", controllers.GetString)
	r.POST("/poof", controllers.PutString)
	r.Run("localhost:8091")

}
