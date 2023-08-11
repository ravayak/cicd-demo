package app

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	testsCtrl "github.com/ravayak/copee_backend/apis/controllers/tests"
)

func mapTestsUrls() {
	testsRoutes := router.Group("/tests")
	testsRoutes.GET("", testsCtrl.Test)
}


// MapUrls maps urls to the associated controllers
func MapUrls() {
	// No Route handler
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, "route not found")
	})
	// No proxies trusted
	router.SetTrustedProxies(nil)
	// set release mode for deploy version
	gin.SetMode(gin.ReleaseMode)

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS")
	corsConfig.AddAllowHeaders("X-Requested-By")
	router.Use(cors.New(corsConfig))


	mapTestsUrls()
	
}
