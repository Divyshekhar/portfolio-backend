package main

import (
	"os"

	"github.com/Divyshekhar/portfolio-backend/controllers"
	"github.com/Divyshekhar/portfolio-backend/initializers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173", "https://divyshekhar.vercel.app"},
		AllowMethods: []string{"POST", "GET", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))

	router.POST("/contact", controllers.MailController)
	port := os.Getenv("PORT")
	router.Run(":" + port)
}
