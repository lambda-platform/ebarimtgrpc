package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"sync"
)

var HTTPPort string
var Port string
var SoFilesPath string

var onceConfig sync.Once

func init() {
	onceConfig.Do(func() {

		_ = godotenv.Overload()

		Port = GetEnvVariable("GRPC_PORT")
		HTTPPort = GetEnvVariable("HTTP_PORT")
		SoFilesPath = GetEnvVariable("SO_FILES_PATH")

	})
}
func GetEnvVariable(key string) string {
	value := os.Getenv(key)
	if value == "" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Error loading .env file")
			return ""
		}
		value = os.Getenv(key)
	}
	return value
}
