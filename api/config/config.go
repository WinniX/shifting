package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type OAuthProvider struct {
	ClientID     string   `yaml:"client_id"`
	ClientSecret string   `yaml:"client_secret"`
	AuthUrl      string   `yaml:"auth_url"`
	TokenUrl     string   `yaml:"token_url"`
	Scopes       []string `yaml:"scopes"`
}

type Config struct {
	Port          int
	Host          string
	DSN           string
	JwtSigningKey string                   `yaml:"jwt_signing_key"`
	OAuth         map[string]OAuthProvider `yaml:"oauth"`
}

func LoadConfig(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
