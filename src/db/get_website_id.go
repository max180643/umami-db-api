package db

import (
	"context"
	"fmt"
	"log"
	"umami-db-api/src/utils"
)

func GetWebsiteIdByName(name string) int {
	var websiteId int
	var conn = utils.DatabaseConnect()

	// fallback
	defer func() {
		conn.Close(context.Background())
	}()

	query := fmt.Sprintf("SELECT website_id FROM website WHERE name = '%s'", name)
	err := conn.QueryRow(context.Background(), query).Scan(&websiteId)

	if err != nil {
		log.Panicln("[GetWebsiteId] QueryRow failed: ", err)
	}

	return (websiteId)
}
