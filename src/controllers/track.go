package controllers

import (
	"strconv"
	"umami-db-api/src/db"
	"umami-db-api/src/models"
)

func Track(websiteName string, hostname string, ip string, userAgent string, url string) models.DefaultResponse {
	// get website id
	websiteId := strconv.Itoa(db.GetWebsiteIdByName(websiteName))

	// new session
	sessionId := strconv.Itoa(db.CreateSession(websiteId, hostname, ip, userAgent))

	// new pageview
	r := db.AddPageview(websiteId, sessionId, url)

	return r
}
