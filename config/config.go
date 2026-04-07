package config

import (
	"APPOINMENT_BOOKING_SYSTEM/models"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var AppConfig models.AppConfig

func GetConfigurations() {
	var configPath string

	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	if os.Getenv("GO_ENV") == "local" {
		configPath = filepath.Join(basepath, "data")
	} else {
		configPath = "/etc/config"
	}

	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Config file error: %v", err)
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatalf("Error unmarshalling config json: %v", err)
	}

}
