package configs

import (
	"log"
	"os"
)

type Config struct {
	Port string
}

func LoadConfig() Config {
	port := os.Getenv("APP_PORT")
	if port == "" {
		log.Println("APP_PORT not set, defaulting to 8080")
		port = "8080"
	}

	return Config{
		Port: port,
	}
}
