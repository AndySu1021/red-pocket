package config

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"red-packet/util/db"
	"red-packet/util/gin"
	"red-packet/util/redis"
	zlog "red-packet/util/zerolog"
)

// AppConfig APP設定
type AppConfig struct {
	fx.Out

	Http     *gin.Config   `mapstructure:"http"`
	Log      *zlog.Config  `mapstructure:"log"`
	Database *db.Config    `mapstructure:"database"`
	Redis    *redis.Config `mapstructure:"redis"`
}

// NewConfig Initiate config
func NewConfig() (AppConfig, error) {
	viper.AutomaticEnv()
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	var config = AppConfig{}

	if err := viper.ReadInConfig(); err != nil {
		return AppConfig{}, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return AppConfig{}, err
	}

	return config, nil
}
