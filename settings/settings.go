package settings

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Settings struct {
	Server   string
	Port     string
	Database string
	Motor    string
}

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func NewSettings() *Settings {
	var s Settings

	s.Server = os.Getenv("DB_HOST")
	s.Port = os.Getenv("DB_PORT")
	s.Database = os.Getenv("DB_NAME")
	s.Motor = os.Getenv("DB_MOTOR")

	return &s

}
