package main

import (
	"Booking/apartment-service/startup"
	cfg "Booking/apartment-service/startup/config"
	"fmt"
)

func main() {

	fmt.Println("Hello apartment")

	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()

}
