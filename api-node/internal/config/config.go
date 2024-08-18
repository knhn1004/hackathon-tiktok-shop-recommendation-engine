package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Configuration struct {
	Target  *string
	KeyName string
}

var (
	// OpenAI
	OpenAIKey string

	// Database
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     int

	// Clerk
	ClerkPublicKey string

	configs = []Configuration{
		{&OpenAIKey, "OPENAI_API_KEY"},
		{&DBHost, "DB_HOST"},
		{&DBUser, "DB_USER"},
		{&DBPassword, "DB_PASSWORD"},
		{&DBName, "DB_NAME"},
		{&ClerkPublicKey, "CLERK_PUBLIC_KEY"},
	}

	// Web
	ServerAddr string = "localhost:8080"
)


// Load reads config from .env file
func Load() error {
	// Load .env file
	if err := godotenv.Load(".env.local"); err != nil {
		return err
	}

	for _, config := range configs {
		if value, ok := os.LookupEnv(config.KeyName); !ok {
			return errors.New(config.KeyName + " required")
		} else {
			*config.Target = value
		}
	}

	// Handle DBPort separately
	portStr, ok := os.LookupEnv("DB_PORT")
	if !ok {
		return errors.New("DB_PORT required")
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return fmt.Errorf("invalid DB_PORT: %v", err)
	}
	DBPort = port

	return nil
}