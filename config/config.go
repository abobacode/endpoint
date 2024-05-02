package config

import (
	"bytes"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/abobacode/endpoint/pkg/database"
)

type Server struct {
	Database database.Opt `yaml:"database"`
	YouTube  YouTube      `yaml:"youtube"`
}

type YouTube struct {
	ApiKey string `yaml:"api_key"`
}

type Config struct {
	Server `yaml:"server"`
}

func New(filepath string) (*Config, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	d := yaml.NewDecoder(bytes.NewReader(content))
	if err = d.Decode(&config); err != nil {
		return nil, err
	}
	return config, nil
}
