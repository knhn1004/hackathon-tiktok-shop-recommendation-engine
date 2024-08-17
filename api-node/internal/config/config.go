package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	Target  *string
	KeyName string
}

var (
	// OpenAI
	OpenAIKey string

	configs = []Configuration{
		{&OpenAIKey, "OPENAI_API_KEY"},
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

	return nil
}