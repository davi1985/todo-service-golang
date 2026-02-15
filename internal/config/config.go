package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	Database DatabaseConfig
}

type DatabaseConfig struct {
	DSN string
}

func NewConfig() *Config {
	return &Config{
		Database: DatabaseConfig{
			DSN: getDatabasePath(),
		},
	}
}

func getDatabasePath() string {
	dataDir := "data"
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return "todos.db"
	}
	
	return filepath.Join(dataDir, "todos.db")
}
