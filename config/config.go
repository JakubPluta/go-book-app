package config

import (
	"log"
	"time"

	"github.com/joeshaw/envdecode"
)


type dbConfig struct {
	Host string `env:"DB_HOST,required"`
	Port     int    `env:"DB_PORT,required"`
    Username string `env:"DB_USER,required"`
    Password string `env:"DB_PASS,required"`
    DbName   string `env:"DB_NAME,required"`
}

type Config struct {
	Debug bool `env:"DEBUG,required"`
	Server serverConfig
	Db dbConfig
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
