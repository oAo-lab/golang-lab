package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type (
	AppConfig struct {
		API    APIConfig    `yaml:"api"`
		Server ServerConfig `yaml:"server"`
	}
	APIConfig struct {
		APIKey string `yaml:"apiKey"`
		APIUrl string `yaml:"apiUrl"`
	}

	ServerConfig struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}
)

// LoadConfig 采用泛型读取任意配置的结构
func LoadConfig[T any](configFile string) (configData T, err error) {
	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return *new(T), fmt.Errorf("failed to read config file: %w", err)
	}

	if err := viper.Unmarshal(&configData); err != nil {
		return configData, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return configData, nil
}
