package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"shareburn/server/controllers"
	"shareburn/server/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func port() string {
	port := os.Getenv("SERVE_PORT")
	if len(port) == 0 {
		port = "8092"
	}
	return ":" + port
}

func main() {

	enverr := godotenv.Load()
	if enverr != nil {
		fmt.Println("Failed to load env")
	}

	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	r := gin.Default()
	r.LoadHTMLFiles(filepath.Join(mydir, "frontend", "index.html"))
	r.Static("/static", filepath.Join(mydir, "frontend", "styles"))

	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/s", controllers.GetStrings)
	r.GET("/s/:key", controllers.GetString)
	r.POST("/s", controllers.PutString)
	// r.Use(cors.Default())
	r.Run(port())
}
