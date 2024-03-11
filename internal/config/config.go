package config

import (
	"encoding/json"
	"os"

	cond "github.com/tmazitov/auth_service.git/pkg/conductor"
)

type StorageConfig struct {
	Addr     string `json:"addr"`
	User     string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	URL      string `json:"url"`
}

func (c *StorageConfig) Validate() bool {
	return (c.Addr != "" && c.User != "" && c.Password != "" && c.Database != "") || c.URL != ""
}

type Config struct {
	Conductor  *cond.ConductorConfig `json:"conductor"`
	JwtSecret  string                `json:"jwtSecret"`
	JwtAccess  int                   `json:"jwtAccessMinutes"` // in minutes
	JwtRefresh int                   `json:"jwtRefreshDays"`   // in days
	DB         *StorageConfig        `json:"db"`
}

func NewConfig(path string) (*Config, error) {
	// Open the JSON file
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into the Config struct
	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
