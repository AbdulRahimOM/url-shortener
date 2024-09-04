package config

import (
	//viper

	"log"

	viper "github.com/spf13/viper"
)

var Db struct {
	Host     string `mapstructure:"DB_HOST"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Name     string `mapstructure:"DB_NAME"`
	Port     string `mapstructure:"DB_PORT"`
}

var Project struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
}

var Url struct {
	Length int `mapstructure:"SHORT_URL_LENGTH"`
}

func init() {
	getEnvValues()
	connectToDB()
	migrateTables()
}

func getEnvValues() {
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("env")
	viper.ReadInConfig()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("error occured while reading env variables, error:", err)
	}

	err = viper.Unmarshal(&Db)
	if err != nil {
		log.Fatalln("error occured while writing env values onto variables, error:", err)
	}

	err = viper.Unmarshal(&Project)
	if err != nil {
		log.Fatalln("error occured while writing env values onto variables, error:", err)
	}

	err = viper.Unmarshal(&Url)
	if err != nil {
		log.Fatalln("error occured while writing env values onto variables, error:", err)
	}
}
