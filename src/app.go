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

	route := gin.Default()
	route.Use(cors.Default())

	// routes
	server := route.Group("/")
	{
		server.GET("/", handlers.DefaultHandler)
		server.POST("/track", handlers.TrackHandler)
		server.GET("/health", handlers.HealthCheckHandler)
		server.Any("/:route", handlers.NotFoundHandler)
	}

	PORT := fmt.Sprintf(":%s", CONFIG.SERVER.PORT)
	route.Run(PORT)
}
