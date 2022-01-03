package db

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DBClient *sqlx.DB

func Open() (*sqlx.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env")
		return DBClient, err
	}

	dbString := os.ExpandEnv("host=${DB_HOST} user=${DB_USER} dbname=${DB_NAME} sslmode=disable password=${DB_PASSWORD}")

	DBClient, err = sqlx.Open("postgres", dbString)
	if err != nil {
		panic("Failed to connect to db")
	}

	return DBClient, nil
}
