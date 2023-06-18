package main

import (
	"Booking/flight-service/startup"
	cfg "Booking/flight-service/startup/config"
	"fmt"
)

func main() {

	fmt.Println("Hello flight")

	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
