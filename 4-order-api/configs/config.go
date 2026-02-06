package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Dsn string
}

func LoadConfig() *Config{
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return &Config{
		Dsn: os.Getenv("DSN"),
	}
}