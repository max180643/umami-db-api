package db

import (
	"context"
	"fmt"
	"log"
	"umami-db-api/src/models"
	"umami-db-api/src/utils"
)

func AddPageview(websiteId string, sessionId string, url string) models.DefaultResponse {
	var viewId int
	var conn = utils.DatabaseConnect()

	// fallback
	defer func() {
		conn.Close(context.Background())
	}()

	// add pageview
	query := fmt.Sprintf("INSERT INTO pageview (website_id, session_id, url) VALUES ('%s', '%s', '%s') RETURNING view_id", websiteId, sessionId, url)
	err := conn.QueryRow(context.Background(), query).Scan(&viewId)

	if err != nil {
		log.Panicln("[CreatePageview] QueryRow failed: ", err)
	}

	r := models.DefaultResponse{
		Status:   "success",
		Response: "Save successfully.",
	}

	return r
}
