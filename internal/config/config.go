package config

import (
	"encoding/json"
	"io/ioutil"

	cond "github.com/tmazitov/auth_service.git/pkg/conductor"
)

type Config struct {
	Conductor *cond.ConductorConfig `json:"conductor"`
}

func NewConfig(path string) (*Config, error) {
	// Open the JSON file
	file, err := ioutil.ReadFile(path)
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