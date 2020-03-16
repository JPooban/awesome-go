package configs

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/joho/godotenv"
)

// Init is using to initialize the configs
func Init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	viper.SetConfigName("default")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}
