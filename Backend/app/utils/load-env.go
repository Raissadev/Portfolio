package utils

import (
	"log"

	"github.com/joho/godotenv"
)

type LoadInterface interface {
	LoadEnv()
}

type LoadEnv struct {
	//
}

func (l *LoadEnv) New() *LoadEnv {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return l
}
