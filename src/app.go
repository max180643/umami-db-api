package main

import (
	"fmt"
	"strconv"
	"umami-db-api/src/handlers"
	"umami-db-api/src/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.InitializeConfig()
	CONFIG := utils.GetConfig()

	if debug, _ := strconv.ParseBool(CONFIG.SERVER.DEBUG); !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	route := gin.New()
	route.Use(gin.LoggerWithConfig(gin.LoggerConfig{SkipPaths: []string{"/health"}}))
	route.Use(gin.Recovery())
	route.Use(cors.Default())
	route.NoRoute(handlers.NotFoundHandler)

	// routes
	server := route.Group("/")
	{
		server.GET("/", handlers.DefaultHandler)
		server.POST("/track", handlers.TrackHandler)
		server.GET("/health", handlers.HealthCheckHandler)
	}

	PORT := fmt.Sprintf(":%s", CONFIG.SERVER.PORT)
	route.Run(PORT)
}
