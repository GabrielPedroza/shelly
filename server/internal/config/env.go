// This file gives autocompletion when importing environment variables. Since we are handling these variables
// on the server, these are safe to import in other files. Add more as it requires. It should be pretty straight-foward. Follow these
// numbered steps below

package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// 1. Add the new env var name here
type AppConfig struct {
	GoogleClientID        string
	GoogleClientSecret    string
	DiscordClientID       string
	DiscordClientSecretID string
	DiscordRedirectURI    string
	JWTSecret             string
	JWTSecretRefresh      string
	TursoConnectionURL    string
	TursoDatabaseName     string
	TursoAuthToken        string
}

func LoadEnv() AppConfig {
	// Dynamically determine the path to the .env file
	envPath := func() string {
		if path := os.Getenv("ENV_PATH"); path != "" {
			return path
		}
		return filepath.Join(os.Getenv("WORKDIR"), "../../../.env") // Default fallback for local development
	}()

	// Load the .env file
	err := godotenv.Load(envPath)
	if err != nil {
		log.Printf("No .env file found at %s, relying on system environment variables", envPath)
	}
	// 2. Properly import the env var. Reference how the other env vars are imported.
	config := AppConfig{
		GoogleClientID:        os.Getenv("GOOGLE_CLIENT_ID"),
		GoogleClientSecret:    os.Getenv("GOOGLE_CLIENT_SECRET"),
		DiscordClientID:       os.Getenv("DISCORD_CLIENT_ID"),
		DiscordClientSecretID: os.Getenv("DISCORD_CLIENT_SECRET"),
		DiscordRedirectURI:    os.Getenv("DISCORD_REDIRECT_URI"),
		JWTSecret:             os.Getenv("JWT_SECRET"),
		JWTSecretRefresh:      os.Getenv("JWT_SECRET_REFRESH"),
		TursoConnectionURL:    os.Getenv("TURSO_CONNECTION_URL"),
		TursoDatabaseName:     os.Getenv("TURSO_DATABASE_NAME"),
		TursoAuthToken:        os.Getenv("TURSO_AUTH_TOKEN"),
	}

	// 3. Have an If check if it is needed but ideally, you and your teammates will have the same .env file contents
	// Do note that it is preferred to add it to the if statement for everyone to be on the same page.
	if config.GoogleClientID == "" ||
		config.GoogleClientSecret == "" ||
		config.DiscordClientID == "" ||
		config.DiscordClientSecretID == "" ||
		config.DiscordRedirectURI == "" ||
		config.JWTSecret == "" ||
		config.JWTSecretRefresh == "" ||
		config.TursoConnectionURL == "" ||
		config.TursoDatabaseName == "" ||
		config.TursoAuthToken == "" {
		log.Fatal("Missing required environment variables")
	}

	return config
}
