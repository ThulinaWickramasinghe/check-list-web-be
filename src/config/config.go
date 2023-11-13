package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port                  string `mapstructure:"PORT"`
	MongoConnectionString string `mapstructure:"MONGO_CONNECTION_STRING"`
	JWTSecret             string `mapstructure:"JWT_SECRET"`
}

var Env *Config

func Load() {
	Env = load()
}

func setDefaults() {
	viper.SetDefault("PORT", "3000")
	viper.SetDefault("MONGO_CONNECTION_STRING", "mongodb+srv://thulina:thulina@cluster0.g7dlwk4.mongodb.net/?retryWrites=true&w=majority")
	viper.SetDefault("JWT_SECRET", "SECRET")
}

func load() (config *Config) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	setDefaults()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return
}
