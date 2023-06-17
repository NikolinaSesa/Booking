package apartment_service

import (
	"Booking/apartment-service/startup"
	cfg "Booking/apartment-service/startup/config"
	"fmt"
)

func main() {
	fmt.Println("Hello")

	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
