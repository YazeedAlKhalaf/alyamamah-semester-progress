package internal

import (
	"github.com/spf13/viper"
)

type Config struct {
	ConsumerKey       string `mapstructure:"CONSUMER_KEY"`
	ConsumerKeySecret string `mapstructure:"CONSUMER_KEY_SECRET"`
	AccessToken       string `mapstructure:"ACCESS_TOKEN"`
	AccessTokenSecret string `mapstructure:"ACCESS_TOKEN_SECRET"`
}

func LoadConfig(path string) (Config, error) {
	var config Config

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
