package utils

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() error {

	if err := godotenv.Load(); err != nil {
		fmt.Println("no .env file found")
	}
	return nil
}
