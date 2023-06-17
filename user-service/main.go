package main

import (
	"fmt"

	"github.com/NikolinaSesa/Booking/user-service/startup"
	cfg "github.com/NikolinaSesa/Booking/user-service/startup/config"
)

func main() {
	fmt.Println("Hello")

	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
