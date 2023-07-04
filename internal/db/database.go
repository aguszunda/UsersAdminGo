package db

import (
	"agustinzunda/usersadmingo/settings"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func ConnDb() {
	setting := settings.NewSettings()
	fmt.Println(setting)

	server := setting.Server
	port := setting.Port
	database := setting.Database
	motor := setting.Motor

	connString := fmt.Sprintf("server=%s;port=%s;database=%s;trusted_connection=yes", server, port, database)

	conn, err := sql.Open(motor, connString)
	if err != nil {
		log.Fatal("connection error: ", err.Error())
	}

	fmt.Println("DataBase connected")
	defer conn.Close()
}
