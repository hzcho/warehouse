package config

import (
	"fmt"
	"time"

	"github.com/vrischmann/envconfig"
)

type Config struct {
	Server
	URL
	Auth
}

type Server struct {
	Port      string
	ReadTime  time.Duration
	WriteTime time.Duration
}

type URL struct {
	Auth      string
	Warehouse string
}

type Auth struct {
	ATDuration     time.Duration
	RFDuration     time.Duration
	PrivateKeyPath string
	PublicKeyPath  string
}

func InitConfig(prefix string) (*Config, error) {
	conf := &Config{}
	if err := envconfig.InitWithPrefix(conf, prefix); err != nil {
		return nil, fmt.Errorf("init config error: %w", err)
	}

	return conf, nil
}
