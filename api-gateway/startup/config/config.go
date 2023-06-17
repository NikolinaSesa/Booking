package config

import "os"

type Config struct {
	//Port     string
	//UserHost string
	//UserPort string
	Address                 string
	UserServiceAddress      string
	ApartmentServiceAddress string
}

func NewConfig() *Config {
	return &Config{
		//Port:     os.Getenv("GATEWAY_PORT"),
		//UserHost: os.Getenv("USER_SERVICE_HOST"),
		//UserPort: os.Getenv("USER_SERVICE_PORT"),
		UserServiceAddress:      os.Getenv("USER_SERVICE_ADDRESS"),
		Address:                 os.Getenv("GATEWAY_ADDRESS"),
		ApartmentServiceAddress: os.Getenv("APARTMENT_SERVICE_ADDRESS"),
	}
}
