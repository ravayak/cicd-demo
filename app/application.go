package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApp Runs the app
func StartApp() {
	port := "8081"

	MapUrls()
	fmt.Println("Application starting on port: ", port)
	router.Run(":" + port)

}
