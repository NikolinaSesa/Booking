package config

import "os"

type Config struct {
	//Port       string
	//UserDBHost string
	//UserDBPort string
	Address string
}

func NewConfig() *Config {
	return &Config{
		//Port:       os.Getenv("USER_SERVICE_PORT"),
		//UserDBHost: os.Getenv("USER_DB_HOST"),
		//UserDBPort: os.Getenv("USER_DB_PORT"),
		Address: os.Getenv("USER_SERVICE_ADDRESS"),
	}
}
