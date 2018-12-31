package config

import (
	"github.com/spf13/viper"
)

// Constants are configuration settings
type Constants struct {
	PORT string
}

// Config is passed around the app
type Config struct {
	Constants
}

// New is used to generate a configuration instance which will be passed around the codebase
func New() (*Config, error) {
	config := Config{}
	constants, err := initViper()
	config.Constants = constants
	if err != nil {
		return &config, err
	}
	return &config, err
}

func initViper() (Constants, error) {
	viper.SetConfigName("example.config")
	viper.AddConfigPath(".")    // search the root directory for the configuration file
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		return Constants{}, err
	}
	viper.SetDefault("PORT", "8080")

	var constants Constants
	err = viper.Unmarshal(&constants)
	return constants, err
}
