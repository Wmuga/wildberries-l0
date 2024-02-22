package config

import (
	"encoding/json"
	"os"
	"path"
)

// Full application configuration.
// Create new instance with NewConfig()
type Config struct {
	HTTP `json:"http"`
	Nats `json:"nats"`
	DB   `json:"db"`
}

// Database connection configuration
type DB struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// Nats-streaming connection configuration
type Nats struct {
	Port int `json:"port"`
}

// HTTP host configuration
type HTTP struct {
	Port int `json:"port"`
}

// Reads config.json and creates new instance of Config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	file := path.Join(wd, "config", "config.json")
	bytes, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
