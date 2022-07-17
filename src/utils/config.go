package utils

import (
	"log"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	DEBUG string `mapstructure:"DEBUG"`
	PORT  string `mapstructure:"PORT"`
}

type DatabaseConfig struct {
	ADDRESS  string `mapstructure:"ADDRESS"`
	PORT     string `mapstructure:"PORT"`
	USER     string `mapstructure:"USER"`
	PASSWORD string `mapstructure:"PASSWORD"`
	DBNAME   string `mapstructure:"DBNAME"`
}

type Config struct {
	SERVER ServerConfig   `mapstructure:"SERVER"`
	DB     DatabaseConfig `mapstructure:"DB"`
}

var vp *viper.Viper
var loadConfig Config

func LoadConfig() (Config, error) {
	vp = viper.New()
	var config Config

	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath(".")

	if err := vp.ReadInConfig(); err != nil {
		return Config{}, err
	}

	if err := vp.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func InitializeConfig() {
	// Load Configuration - config.json
	config, err := LoadConfig()

	if err != nil {
		log.Panicln("Could not load configuration file")
	}

	loadConfig = config
}

func GetConfig() Config {
	return loadConfig
}
