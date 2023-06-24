package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ConnDb() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	server := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")
	motor := os.Getenv("sqlserver")

	connString := fmt.Sprintf("server=%s;port=%s;database=%s;integrated security=true", server, port, database)

	conn, err := sql.Open(motor, connString)
	if err != nil {
		log.Fatal("connection error: ", err)
	}
	fmt.Println("DataBase connected")
	defer conn.Close()
}
