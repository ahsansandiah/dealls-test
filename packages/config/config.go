package config

import (
	"os"

	"github.com/spf13/viper"
)

type ContextKey string

type Config struct {
	AppEnv                     string `mapstructure:"APP_ENV"`
	AppTz                      string `mapstructure:"APP_TZ"`
	AppCurrency                string `mapstructure:"APP_CURRENCY"`
	AppIsDev                   bool
	DatabaseDriver             string `mapstructure:"DATABASE_DRIVER"`
	DatabaseDSN                string `mapstructure:"DATABASE_DNS"`
	DatabaseMaxOpenConnections int    `mapstructure:"DATABASE_MAX_OPEN_CONNECTIONS"`
	DatabaseMaxIdleConnections int    `mapstructure:"DATABASE_MAX_IDLE_CONNECTIONS"`
	RedisConnection            string `mapstructure:"REDIS_CONNECTION"`
	RedisAddress               string `mapstructure:"REDIS_ADDRESS"`
	RedisUsername              string `mapstructure:"REDIS_USERNAME"`
	RedisPassword              string `mapstructure:"REDIS_PASSWORD"`
	RedisDatabase              int    `mapstructure:"REDIS_DATABASE"`
	JwtSecretKey               string `mapstructure:"JWT_SECRET_KEY"`
	JwtAccessTokenDuration     int    `mapstructure:"JWT_ACCESS_TOKEN_DURATION_SECONDS"`
	PortHttpServer             string `mapstructure:"PORT_HTTP_SERVER"`
	ServerHTTPReadTimeout      int    `mapstructure:"SERVER_HTTP_READ_TIMEOUT"`
}

func NewConfig() (*Config, error) {
	env := os.Getenv("APPENV")
	if env == "" {
		env = "local"
	}

	viper.AddConfigPath(".")
	viper.AddConfigPath("packages/config")
	viper.SetConfigName(env)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			viper.SetConfigName("placeholder")

			if err := viper.ReadInConfig(); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	cfg := &Config{}
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	cfg.AppIsDev = cfg.AppEnv == "staging" || cfg.AppEnv == "local"

	return cfg, nil
}
