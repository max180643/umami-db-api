package db

import (
	"context"
	"fmt"
	"log"
	"umami-db-api/src/utils"
)

func CreateSession(websiteId string, hostname string, ip string, userAgent string) int {
	// generate session uuid
	sessionUuid := utils.GenerateUuid(websiteId, hostname, ip, userAgent)

	// check uuid is exists
	var sessionId int
	var conn = utils.DatabaseConnect()

	// fallback
	defer func() {
		conn.Close(context.Background())
	}()

	// get session by uuid
	query := fmt.Sprintf("SELECT session_id FROM session WHERE session_uuid = '%s'", sessionUuid)
	err := conn.QueryRow(context.Background(), query).Scan(&sessionId)

	if err != nil {
		log.Panicln("[GetSessionByUuid] QueryRow failed: ", err)
	}

	if sessionId == 0 {
		// session not found - create new session id
		query := fmt.Sprintf("INSERT INTO session (session_uuid, website_id, hostname, browser, os, language, country, device) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s') RETURNING session_id", sessionUuid, websiteId, hostname, utils.GetBrowser(userAgent), utils.GetOs(userAgent), "en-US", utils.GetCountry(ip), utils.GetDevice(userAgent))
		err := conn.QueryRow(context.Background(), query).Scan(&sessionId)

		if err != nil {
			log.Panicln("[CreateSession] QueryRow failed: ", err)
		}
	}

	// session found
	return sessionId
}
