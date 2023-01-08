package main

import (
	"net/http"
	"poof/controllers"
	"poof/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.LoadHTMLGlob("frontend/*.html")
	r.LoadHTMLFiles("frontend/index.html")
	r.Static("/static", "frontend/styles/")

	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/poofs", controllers.GetStrings)
	r.GET("/poofs/:key", controllers.GetString)
	r.POST("/poof", controllers.PutString)
	// r.Use(cors.Default())
	r.Run("localhost:8091")

}
