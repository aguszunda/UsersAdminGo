package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func ConnDb() {
	server := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")
	motor := os.Getenv("DB_MOTOR")

	connString := fmt.Sprintf("server=%s;port=%s;database=%s;trusted_connection=yes", server, port, database)

	conn, err := sql.Open(motor, connString)
	if err != nil {
		log.Fatal("connection error: ", err.Error())
	}
	fmt.Println("DataBase connected")
	defer conn.Close()
}
