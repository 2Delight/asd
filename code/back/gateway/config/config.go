package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		ConnectionString string `yaml:"connectionString"`
	} `yaml:"database"`
	HttpPort string `yaml:"httpPort"`
	GrpcPort string `yaml:"grpcPort"`
}

type Service struct {
	Host     string `yaml:"host"`
	HttpPort string `yaml:"httpPort"`
	GrpcPort string `yaml:"grpcPort"`
}

const (
	pathToConfig = "config.yaml"

	AppName = "gateway-api"
)

func LoadConfig() (*Config, error) {
	rawYaml, err := os.ReadFile(pathToConfig)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = yaml.Unmarshal(rawYaml, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) GetCheckoutPgConnectionString() string {
	return c.Database.ConnectionString
}

func (c *Config) GetHTTPPort() string {
	return c.HttpPort
}

func (c *Config) GetGPRCPort() string {
	return c.GrpcPort
}
