package config

import (
	"demo/util/db"
	"demo/util/gin"
	"demo/util/redis"
	zlog "demo/util/zerolog"
	"github.com/spf13/viper"
	"go.uber.org/fx"
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
	viper.SetConfigName("app")
	viper.SetConfigType("properties")

	var config = AppConfig{}

	if err := viper.ReadInConfig(); err != nil {
		return AppConfig{}, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return AppConfig{}, err
	}

	return config, nil
}
