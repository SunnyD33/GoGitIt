package utils

import (
	"bufio"
	"errors"
	//"fmt"
	"io"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

func LoadEnv() error {
	err := godotenv.Load()

	if err != nil {
		return errors.New("error loading .env file. Please check and/or set .env file location")
	}

	return nil
}

func SetEnv(r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", errors.New("a file location was not set")
	}

	envLocation := scanner.Text()

	return envLocation, nil
}