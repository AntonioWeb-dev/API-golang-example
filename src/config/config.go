package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	SqlConfig = ""
	Port      = 0
	// Token key
	SecretKey []byte
)

// Config - Load all configs
func Config() {
	var error error

	if error = godotenv.Load(); error != nil {
		log.Fatal(error)
	}
	Port, error = strconv.Atoi(os.Getenv("API_PORT"))
	if error != nil {
		Port = 3000
	}

	SqlConfig = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
	)
	SecretKey = []byte(os.Getenv("JWT_SECRET"))

}
