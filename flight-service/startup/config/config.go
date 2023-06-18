package config

import "os"

type Config struct {
	Address string
}

func NewConfig() *Config {
	return &Config{
		Address: os.Getenv("FLIGHT_SERVICE_ADDRESS"),
	}
}
