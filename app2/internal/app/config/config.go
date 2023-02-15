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
	config.URI = os.Getenv("MONGO_URI")
	config.DbName = os.Getenv("MONGO_DB_NAME")
	return config
}
