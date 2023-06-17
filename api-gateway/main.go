package main

import (
	"fmt"

	"github.com/NikolinaSesa/Booking/api-gateway/startup"
	"github.com/NikolinaSesa/Booking/api-gateway/startup/config"
)

func main() {

	fmt.Println("hello api_gateway")

	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
