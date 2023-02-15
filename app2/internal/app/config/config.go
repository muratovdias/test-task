package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type MongoConfig struct {
	URI    string
	DbName string
}

func LoadMongoConfig() MongoConfig {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var config MongoConfig
	config.URI = os.Getenv("MONGODB_URI")
	config.DbName = os.Getenv("MONGODB_NAME")
	log.Println("confid loaded", config)
	return config
}
