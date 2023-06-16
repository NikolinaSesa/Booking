package main

import (
	"Booking/user-service/startup"
	cfg "Booking/user-service/startup/config"
	"fmt"
)

func main() {
	fmt.Println("Hello")

	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
