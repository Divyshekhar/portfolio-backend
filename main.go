package main

import (
	"os"

	"github.com/Divyshekhar/portfolio-backend/controllers"
	"github.com/Divyshekhar/portfolio-backend/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	router := gin.Default()

	router.POST("/contact", controllers.MailController)
	port := os.Getenv("PORT")
	router.Run(":" + port)
}
