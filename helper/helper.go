package helper

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Address string `yaml:"address"`
		Port    string `yaml:"port"`
		Timeout int8   `yaml:"timeout"`
	} `yaml:"server"`

	DB struct {
		Host     string `yaml:"host"`
		Port     int16  `yaml:"port"`
		DBName   string `yaml:"db_name"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		SSLMode  string `yaml:"ssl_mode"`
	} `yaml:"db"`

	Google struct {
		ClientID    string `yaml:"client_id"`
		Secret      string `yaml:"secret"`
		RedirectURI string `yaml:"redirect_uri"`
	} `yaml:"google"`
}

func NewConfig(file string) (*Config, error) {

	config := Config{}

	reader, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(reader, &config); err != nil {
		return nil, err
	}

	return &config, nil

}