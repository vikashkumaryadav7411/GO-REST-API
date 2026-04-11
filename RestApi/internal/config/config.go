package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI   string
	MongoDB    string
	ServerPort string
}

func Load() (Config, error) {

	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("Failed to Load .env")
	}

	MongoURI, err := extractEnv("MONGO_URI")
	if err != nil {
		return Config{}, err
	}

	MongoDB, err := extractEnv("MONGO_DB_NAME")
	if err != nil {
		return Config{}, err
	}

	port, err := extractEnv("PORT")
	if err != nil {
		return Config{}, err
	}

	return Config{
		MongoURI:   MongoURI,
		MongoDB:    MongoDB,
		ServerPort: port,
	}, nil
}

func extractEnv(key string) (string, error) {

	val := os.Getenv(key)

	if val == "" {
		return "", fmt.Errorf("missing req env")
	}
	return val, nil
}
