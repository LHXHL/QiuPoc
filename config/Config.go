package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func Config() string {
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return ""
	}

	dnslog := viper.GetString("dnslog.host")
	return dnslog
}
