package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

const envFilesPath = "./env/dev/"

func GoDotEnvVariable(key string) (string, error) {
	err := godotenv.Load(fmt.Sprintf("%s.env", envFilesPath))

	if err != nil {
		return "", fmt.Errorf("error loading .env file: %s", err.Error())
	}

	return os.Getenv(key), nil
}
