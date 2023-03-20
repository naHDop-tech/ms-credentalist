package utils

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBPort              string        `mapstructure:"DB_PORT"`
	DBHost              string        `mapstructure:"DB_HOST"`
	DBUser              string        `mapstructure:"DB_USER"`
	DBPassword          string        `mapstructure:"DB_PASSWORD"`
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBName              string        `mapstructure:"DB_NAME"`
	AppHost             string        `mapstructure:"HOST"`
	AppPort             string        `mapstructure:"PORT"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	AllowedMethods      []string      `mapstructure:"ALLOWED_METHODS"`
	AllowedHeaders      []string      `mapstructure:"ALLOWED_HEADERS"`
	AllowedOrigin       []string      `mapstructure:"ALLOWED_ORIGIN"`
	SmtpHost            string        `mapstructure:"SMTP_HOST"`
	SmtpPort            string        `mapstructure:"SMTP_PORT"`
	SmtpPassword        string        `mapstructure:"SMTP_PASSWORD"`
	SmtpLogin           string        `mapstructure:"SMTP_LOGIN"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
