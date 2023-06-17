package startup

import (
	"Booking/apartment-service/application"
	"Booking/apartment-service/domain"
	"Booking/apartment-service/infrastructure/api"
	"Booking/apartment-service/infrastructure/persistence"
	"Booking/apartment-service/proto"
	"Booking/apartment-service/startup/config"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
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
	apartmentStore := s.initApartmentStore(mongoClient)

	apartmentService := s.initApartmentService(apartmentStore)

	apartmentHandler := s.initApartmentHandler(apartmentService)

	s.startGrpcServer(apartmentHandler)
}

func (s *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(s.config.ApartmentDBHost, s.config.ApartmentDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (s *Server) initApartmentStore(client *mongo.Client) domain.ApartmentStore {
	store := persistence.NewApartmentMongoDBStore(client)
	return store
}

func (s *Server) initApartmentService(store domain.ApartmentStore) *application.ApartmentService {
	return application.NewApartmentService(store)
}

func (s *Server) initApartmentHandler(service *application.ApartmentService) *api.ApartmentHandler {
	return api.NewApartmentHandler(service)
}

func (server *Server) startGrpcServer(apartmentHandler *api.ApartmentHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":#{server.config.Port}"))
	if err != nil {
		log.Fatalf("Failed to listen: #{err}")
	}
	grpcServer := grpc.NewServer()
	proto.RegisterApartmentServiceServer(grpcServer, apartmentHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: #{err}")
	}
}
