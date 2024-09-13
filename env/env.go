package env

import (
	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load(".env", ".env.secret")
	if err != nil {
		panic("no env file found")
	}
}
