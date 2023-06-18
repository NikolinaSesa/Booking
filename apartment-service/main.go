package main

import (
	"fmt"

	"github.com/NikolinaSesa/Booking/apartment-service/startup"
	cfg "github.com/NikolinaSesa/Booking/apartment-service/startup/config"
)

func main() {

	fmt.Println("Hello apartment")

	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()

}
