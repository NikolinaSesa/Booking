package startup

import (
	"Booking/flight-service/application"
	"Booking/flight-service/startup/config"
	"log"

	"Booking/flight-service/domain"
	"Booking/flight-service/infrastructure/persistence"

	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (s *Server) Start() {
	mongoClient := s.initMongoClient()
	userStore := s.initFlightStore(mongoClient)

	_ = s.initFlightService(userStore)
}

func (s *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient()
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (s *Server) initFlightStore(client *mongo.Client) domain.FlightStore {
	store := persistence.NewFlightMongoDBStore(client)
	store.DeleteAllFlights()

	for _, flight := range flights {
		err := store.InsertFlight(flight)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (s *Server) initFlightService(store domain.FlightStore) *application.FlightService {
	return application.NewFlightService(store)
}
