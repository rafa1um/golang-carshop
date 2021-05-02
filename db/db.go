package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectToDb() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	connection := "user=" + os.Getenv("DB_USERNAME") + " dbname=" + os.Getenv("DB_NAME") + " password=" + os.Getenv("DB_PASSWORD") + " host=" + os.Getenv("DB_HOST") + " sslmode=disable"
	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}
	return db
}
