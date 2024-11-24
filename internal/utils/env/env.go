package env

import (
	"log"

	"github.com/spf13/viper"
)

// Setup config file
type Env struct {
	AppName                string `mapstructure:"APP_NAME"`
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerEnv              string `mapstructure:"SERVER_ENV"`
	ServerURL              string `mapstructure:"SERVER_URL"`
	ServerPort             string `mapstructure:"SERVER_PORT"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBDriver               string `mapstructure:"DB_DRIVER"`
	DBHost                 string `mapstructure:"DB_HOST"`
	DBPort                 string `mapstructure:"DB_PORT"`
	DBUser                 string `mapstructure:"DB_USER"`
	DBPass                 string `mapstructure:"DB_PASS"`
	DBName                 string `mapstructure:"DB_NAME"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
	GinMode                string `mapstructure:"GIN_MODE"`
	RabbitCon              string `mapstructure:"RABBITCONNECT"`
	RabbitQueue            string `mapstructure:"RABBITQUEUE"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile("./config/config.env")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	} else {
		log.Println("The App is running in Production env")
	}

	return &env
}
