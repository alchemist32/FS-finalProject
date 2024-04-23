package config

import (
	"os"
)

var ServerPort string

func LoadEnv() {
	ServerPort = os.Getenv("SERVER_PORT")
}
