package config

import (
	"github.com/spf13/viper"
)

const (
	ConfigDefaultPath     string = "configs"
	ConfigDefaultFilename string = "default-config"
)

var (
	BlankConfig Config = Config{}
)

type (
	Config struct {
		App  AppConfig
		HTTP HTTPConfig
	}

	AppConfig struct {
		CookieKey string `mapstructure:"cookie-key"`
	}

	HTTPConfig struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
	}
)

func InitConfiguration() (Config, error) {
	viper.AddConfigPath(ConfigDefaultPath)
	viper.SetConfigName(ConfigDefaultFilename)
	if err := viper.ReadInConfig(); err != nil {
		return BlankConfig, err
	}

	var cfg Config
	if err := viper.UnmarshalKey("app", &cfg.HTTP); err != nil {
		return BlankConfig, err
	}
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return BlankConfig, err
	}

	return cfg, nil
}
