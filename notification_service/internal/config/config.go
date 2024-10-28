package config

import (
	"fmt"

	"github.com/vrischmann/envconfig"
)

type Config struct {
	SMTP
	Consumer
	URLs
}

type SMTP struct {
	Host     string
	Port     string
	OrgEmail string
	Password string
}

type Consumer struct {
	Brokers string
	Topics  []string
	GroupId string
	Offset  string
}

type URLs struct {
	Auth string
}

func InitConfig(prefix string) (*Config, error) {
	conf := &Config{}
	if err := envconfig.InitWithPrefix(conf, prefix); err != nil {
		return nil, fmt.Errorf("init config error: %w", err)
	}

	return conf, nil
}
