package utils

import (
	"bufio"
	"errors"
	"fmt"

	//"fmt"
	"io"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

func LoadEnv(filepath string) error {
	err := godotenv.Load(filepath)

	if err != nil {
		return errors.New("error loading .env file in current filepath")
	}

	return nil
}

func SetEnv(r io.Reader, w io.Writer) (string, error) {
	msg := "Please enter the full filepath for your .env file \n"
	fmt.Fprint(w, msg)
	scanner := bufio.NewScanner(r)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", err
	}

	envLocation := scanner.Text()

	return envLocation, nil
}
