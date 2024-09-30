package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sqlx.DB

func Connect() error {
	err := godotenv.Load("config.env")
	if err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)

	DB, err = sqlx.Connect("mysql", connectionString)
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}

	fmt.Println("Connected to the database")
	return nil
}
