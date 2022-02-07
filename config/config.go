package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func LoadingConfig() (string, error) {
	var err error

	if err = godotenv.Load(); err != nil {
		return "", errors.New("err")
	}
	databaseUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NOME"),
	)

	return databaseUrl, nil
}
