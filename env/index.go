package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var AdminPassword string

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error: Cannot read .env file")
		panic(err.Error())
	}
	AdminPassword = os.Getenv("ADMIN_PASSWORD")
}
