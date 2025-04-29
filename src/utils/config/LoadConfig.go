package config

import (
	"ABCD/src/constants"
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig[T any](configPath, configName, configType string) (*T, error) {
	var config T

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf(constants.ReadConfigFailedError, err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf(constants.UnmarshalFailedError, err)
	}

	return &config, nil
}
