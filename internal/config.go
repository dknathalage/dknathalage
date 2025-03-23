package internal

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Filename string `mapstructure:"README_FILENAME"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile("config.env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	}

	log.Printf("%+v\n", config)

	return &config, nil
}
