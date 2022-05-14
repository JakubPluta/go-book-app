package config

import (
	"log"
	"time"

	"github.com/joeshaw/envdecode"
)


type Config struct {
	Server serverConfig
}

type serverConfig struct {
	Port int `env:"SERVER_PORT,required"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ,required"`
    TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE,required"`
    TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE,required"`


}


func AppConfig() *Config{
	var c Config
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode %s", err)
	}
	return &c
}