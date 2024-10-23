package config

import (
	"fmt"

	"github.com/vrischmann/envconfig"
)

type Config struct {
	SMPT
	Consumer
	URLs
}

type SMPT struct {
	Host     string
	Port     string
	OrgEmail string
	Password string
}

type Consumer struct {
	Brokers string   `yaml:"brokers" env-required:"true"`
	Topics  []string `yaml:"topics" env-required:"true"`
	GroupId string   `yaml:"group_id" env-required:"true"`
	Offset  string   `yaml:"offset" env-required:"true"`
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
