package utils

import (
	"github.com/spf13/viper"
)

// Config contains the configuration for the application.
// the values are read by viper from a config file or environment variables.
type Config struct {
	DSN               string `mapstructure:"DSN"`
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
}

// LoadConfig loads the configuration from a config file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile("app.env")
	// viper.SetConfigName("app")
	// viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
