package config

import (
	"github.com/joho/godotenv"
)

func LoadENV() error {
	err := godotenv.Load()

	if err != nil {
		return err
	}

	return nil
}
