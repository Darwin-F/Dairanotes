package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"net/url"
	"os"
)

func ConnectDatabase() (db *sqlx.DB, err error) {
	err = godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	db, err = sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		url.QueryEscape(os.Getenv("DB_LOCATION"))),
	)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Failed to connect to ftth_v3 database")
		return
	}
	return
}
