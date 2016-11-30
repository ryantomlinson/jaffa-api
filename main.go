package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ryantomlinson/jaffa-api/controllers"
	"github.com/ryantomlinson/jaffa-api/db"
)

var router *gin.Engine

func main() {
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Initialize the routes
	v1 := router.Group("/v1")
	{
		feature := new(controllers.FeatureController)

		v1.GET("/features", feature.GetAllFeatures)
	}

	db.Init()

	// Start serving the application
	router.Run(":9000")
}