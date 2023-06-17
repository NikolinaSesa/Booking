package config

import "os"

type Config struct {
	//Port            string
	//ApartmentDBHost string
	//ApartmentDBPort string
	Address string
}

func NewConfig() *Config {
	return &Config{
		//Port:            os.Getenv("APARTMENT_SERVICE_PORT"),
		//ApartmentDBHost: os.Getenv("APARTMENT_DB_HOST"),
		//ApartmentDBPort: os.Getenv("APARTMENT_DB_PORT"),
		Address: os.Getenv("APARTMENT_SERVICE_ADDRESS"),
	}
}
