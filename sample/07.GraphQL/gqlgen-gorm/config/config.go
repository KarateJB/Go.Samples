package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config: Config
type Config struct {
	Port string `mapstructure:"PORT"`
	DB   string `mapstructure:"DB"`
}

// Init: initialize the config file to struct
func Init() *Config {
	var config Config
	viper.SetConfigName("config.yml")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Failed to load config file: %w \n", err))
	}
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("Failed to decode config file into struct, %w \n", err))
	}

	return &config
}
