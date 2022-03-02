package config

import (
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
