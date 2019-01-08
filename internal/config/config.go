package config

import (
	"errors"

	"github.com/spf13/viper"
)

// Constants are configuration settings
type Constants struct {
	PORT       string
	DBUser     string
	DBPassword string
	DBName     string
}

// Config is passed around the app
type Config struct {
	Constants
}

// New is used to generate a configuration instance which will be passed around the codebase
func New() (*Config, error) {
	config := Config{}
	constants, err := initViper()
	// make sure the database settings are set
	if constants.DBUser == "" || constants.DBName == "" || constants.DBPassword == "" {
		return &config, errors.New("Database settings are not set")
	}
	config.Constants = constants
	if err != nil {
		return &config, err
	}
	return &config, err
}

func initViper() (Constants, error) {
	viper.SetConfigName("service.config")
	viper.AddConfigPath(".")    // search the root directory for the configuration file
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		return Constants{}, err
	}
	viper.SetDefault("SERVER_PORT", "8080")

	var constants Constants
	err = viper.Unmarshal(&constants)
	return constants, err
}
