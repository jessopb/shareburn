package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"shareburn/controllers"
	"shareburn/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
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
	r.Run(fmt.Sprintf("%s:%s", os.Getenv("SERVE_IP"), os.Getenv("SERVE_PORT")))
}
