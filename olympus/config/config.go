package config

import (
	"os"
	"strconv"
)

type PostgresDbConfig struct {
	PostgresDbHost     string
	PostgresDbPort     int
	PostgresDbUser     string
	PostgresDbPassword string
	PostgresDbName     string
}

type Config struct {
	Postgres   PostgresDbConfig
	DebugMode  bool
	ServerPort int
}

// New returns a new Config struct
func New() *Config {
	return &Config{
		Postgres: PostgresDbConfig{
			PostgresDbHost:     getEnv("POSTGRES_DB_HOST", ""),
			PostgresDbPort:     getEnvAsInt("POSTGRES_DB_PORT", 0),
			PostgresDbUser:     getEnv("POSTGRES_DB_USER", ""),
			PostgresDbPassword: getEnv("POSTGRES_DB_PASSWORD", ""),
			PostgresDbName:     getEnv("POSTGRES_DB_NAME", ""),
		},
		DebugMode:  getEnvAsBool("DEBUG_MODE", true),
		ServerPort: getEnvAsInt("SERVER_PORT", 8080),
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// Helper to read an environment variable into a bool or return default value
func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}
