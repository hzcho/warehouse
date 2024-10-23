package config

import (
	"fmt"
	"time"

	"github.com/vrischmann/envconfig"
)

type Config struct {
	Server
	Mongo
	Storage
}

type Server struct {
	Host      string
	Port      string
	ReadTime  time.Duration
	WriteTime time.Duration
}

type Mongo struct {
	Username string
	Host     string
	Port     string
	DBName   string
	Password string
}

type Storage struct {
	UploadDir string
}

func InitConfig(prefix string) (*Config, error) {
	conf := &Config{}
	if err := envconfig.InitWithPrefix(conf, prefix); err != nil {
		return nil, fmt.Errorf("init config error: %w", err)
	}

	return conf, nil
}
