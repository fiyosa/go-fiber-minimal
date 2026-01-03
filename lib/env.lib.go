package lib

import (
	"fmt"
	"go-fiber-ddd/config"
	"os"

	"github.com/joho/godotenv"
)

var Env envManager

type envManager struct{}

func (envManager) Init() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf("Error loading .env file: %v \n\n", err.Error())
		os.Exit(1)
	}
	config.Env.LoadEnv()
}
