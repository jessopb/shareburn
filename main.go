package main

import (
	"net/http"
	"shareburn/controllers"
	"shareburn/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("/frontend/index.html")
	r.Static("/static", "/frontend/styles/")

	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/s", controllers.GetStrings)
	r.GET("/s/:key", controllers.GetString)
	r.POST("/s", controllers.PutString)
	// r.Use(cors.Default())
	r.Run("0.0.0.0:8091")
}
