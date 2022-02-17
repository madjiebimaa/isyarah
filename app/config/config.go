package config

import (
	"log"

	"github.com/joho/godotenv"
)

func GetEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("environment variables is not configure")
	}
}
