package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

type Config struct {
	Server ServerConfig `yaml:"server"`
}

var (
	Conf *Config
)

func Initialize(env string) error {
	viper.SetConfigName(fmt.Sprintf("%s.config", env))
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return nil
}
