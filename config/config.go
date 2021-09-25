package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	Domain string
	Server struct {
		Port uint16
	}
	SQL struct {
		Addr     string
		Username string
		Password string
		Table    string
	}
}

func New(port uint16) (c *Config, err error) {
	v := initViper()

	if port > 0 {
		v.Set("server_port", port)
	}

	if err = v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warn().Msg("No config file found in search paths, using default values")
		}
	}

	err = v.Unmarshal(&c)
	return
}

func initViper() *viper.Viper {
	v := viper.NewWithOptions(viper.KeyDelimiter("_"))

	v.SetConfigFile(".env")
	v.AutomaticEnv()

	// Server
	viper.SetDefault("server_port", 4000)

	// SQL Database
	v.SetDefault("sql_addr", "127.0.0.1:5432")
	v.SetDefault("sql_username", "root")
	v.SetDefault("sql_password", "root")

	return v
}
