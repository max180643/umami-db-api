package handlers

import (
	"log"

	"net/http"
	"umami-db-api/src/controllers"
	"umami-db-api/src/models"

	"github.com/gin-gonic/gin"
)

func DefaultHandler(c *gin.Context) {
	r := models.DefaultResponse{
		Status:   "success",
		Response: "Please go to https://github.com/max180643/umami-db-api for API usage.",
	}
	c.JSON(http.StatusOK, r)
}

func HealthCheckHandler(c *gin.Context) {
	r := models.DefaultResponse{
		Status:   "success",
		Response: "OK",
	}
	c.JSON(http.StatusOK, r)
}

func NotFoundHandler(c *gin.Context) {
	r := models.DefaultResponse{
		Status:   "failure",
		Response: "Route not found.",
	}
	c.JSON(http.StatusNotFound, r)
}

func TrackHandler(c *gin.Context) {
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
}
