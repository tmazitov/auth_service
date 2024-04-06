package config

import (
	"encoding/json"
	"os"

	cond "github.com/tmazitov/auth_service.git/pkg/conductor"
	"github.com/tmazitov/auth_service.git/pkg/service"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type StorageConfig struct {
	Addr     string `json:"addr"`
	User     string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	URL      string `json:"url"`
	SSL      bool   `json:"ssl"`
}

type GoogleOathConfig struct {
	ClientID     string   `json:"clientID"`
	ClientSecret string   `json:"clientSecret"`
	RedirectURL  string   `json:"redirectURL"`
	Scopes       []string `json:"scopes"`
}

type RedisConfig struct {
	Addr string `json:"addr"`
	DB   int    `json:"db"`
}

type DocsConfig struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type JwtConfig struct {
	Secret  string `json:"secret"`
	Access  int    `json:"accessMinutes"`
	Refresh int    `json:"refreshDays"`
}

func (c *StorageConfig) Validate() bool {
	return (c.Addr != "" && c.User != "" && c.Password != "" && c.Database != "") || c.URL != ""
}

type Config struct {
	Service   *service.ServiceConfig
	Docs      *DocsConfig
	GoogleRaw *GoogleOathConfig `json:"google"`
	Google    *oauth2.Config
	Conductor *cond.ConductorConfig `json:"conductor"`
	Jwt       *JwtConfig            `json:"jwt"`
	Storage   *StorageConfig
	Redis     *RedisConfig
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

	config.Google = &oauth2.Config{
		ClientID:     config.GoogleRaw.ClientID,
		ClientSecret: config.GoogleRaw.ClientSecret,
		RedirectURL:  config.GoogleRaw.RedirectURL,
		Scopes:       config.GoogleRaw.Scopes,
		Endpoint:     google.Endpoint,
	}

	return &config, nil
}
