package config

import (
	"bytes"
	"github.com/abobacode/endpoint/pkg/database"
	"os"

	"gopkg.in/yaml.v3"
)

type Server struct {
	Database database.Opt `yaml:"database"`
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
