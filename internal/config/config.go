package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port     string `mapstructure:"port"`
	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DbName   string `mapstructure:"name"`
	}
}

func (c *Config) GetDbUrl() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.DbName,
	)

}

func GetConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	// Set default values
	viper.SetDefault("port", "8080")

	var config Config

	if err := viper.ReadInConfig(); err != nil {
		return &config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return &config, err
	}

	return &config, nil
}
