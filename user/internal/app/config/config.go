package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI       string
	MongoDbName    string
	SaltServiceUrl string
}

func LoadMongoConfig() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var config Config
	config.MongoURI = os.Getenv("MONGODB_URI")
	config.MongoDbName = os.Getenv("MONGODB_NAME")
	config.SaltServiceUrl = os.Getenv("SALT_SERVICE_URL")
	log.Println("confid loaded", config)
	return config
}
