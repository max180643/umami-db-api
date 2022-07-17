package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

func DatabaseConnect() *pgx.Conn {
	CONFIG := GetConfig()
	ADDRESS := CONFIG.DB.ADDRESS
	PORT := CONFIG.DB.PORT
	USER := CONFIG.DB.USER
	PASSWORD := CONFIG.DB.PASSWORD
	DBNAME := CONFIG.DB.DBNAME

	ConnectionString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", USER, PASSWORD, ADDRESS, PORT, DBNAME)
	conn, err := pgx.Connect(context.Background(), ConnectionString)

	if err != nil {
		log.Panicln("Unable to connect to database")
	}

	return conn
}
