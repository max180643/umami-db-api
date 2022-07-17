package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"umami-db-api/src/controllers"
	"umami-db-api/src/models"
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

	route.GET("/", func(c *gin.Context) {
		r := models.DefaultResponse{
			Status:   "success",
			Response: "Please go to https://github.com/max180643/umami-db-api for API usage.",
		}
		c.JSON(http.StatusOK, r)
	})

	route.POST("/track", func(c *gin.Context) {
		var requestBody models.TrackRequest
		err := c.ShouldBindJSON(&requestBody)

		// fallback
		defer func() {
			if rv := recover(); rv != nil {
				r := models.DefaultResponse{
					Status:   "failure",
					Response: rv.(string),
				}

				c.AbortWithStatusJSON(http.StatusInternalServerError, r)
			}
		}()

		if err != nil {
			log.Panicln("[/track] BindJSON: ", err)
		} else {
			websiteName := *requestBody.WebsiteName
			hostname := *requestBody.Hostname
			ip := requestBody.Ip
			userAgent := requestBody.UserAgent
			url := requestBody.Url

			// default ip
			if ip == "" {
				ip = c.Request.Header.Get("X-Forwarded-For")
				if ip == "" {
					ip = c.ClientIP()
				}
			}
			// default user-agent
			if userAgent == "" {
				userAgent = c.Request.Header.Get("User-Agent")
			}

			// default url
			if url == "" {
				url = "/"
			}

			r := controllers.Track(websiteName, hostname, ip, userAgent, url)

			c.JSON(http.StatusOK, r)
		}
	})

	route.GET("/health", func(c *gin.Context) {
		r := models.DefaultResponse{
			Status:   "success",
			Response: "OK",
		}
		c.JSON(http.StatusOK, r)
	})

	route.GET("/:route", func(c *gin.Context) {
		r := models.DefaultResponse{
			Status:   "failure",
			Response: "Route not found.",
		}
		c.JSON(http.StatusNotFound, r)
	})

	PORT := fmt.Sprintf(":%s", CONFIG.SERVER.PORT)
	route.Run(PORT)
}
