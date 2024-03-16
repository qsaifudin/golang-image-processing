package main

import (
	"image-processing-service/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
    // Initialize Gin router
    router := gin.Default()

    // CORS middleware
    config := cors.DefaultConfig()
    config.AllowAllOrigins = true
    router.Use(cors.New(config))

    // Register routes
    handlers.RegisterRoutes(router)

    // Start the server
    router.Run(":5000")
}
