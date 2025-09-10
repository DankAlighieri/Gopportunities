package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Initialize Gin router
	router := gin.Default()

	// Initilialize routes
	initializeRoutes(router)

	// Run the server
	router.Run(":8080")
}
