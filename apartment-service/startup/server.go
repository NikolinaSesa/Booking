package startup

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"Booking/apartment-service/application"
	"Booking/apartment-service/domain"
	"Booking/apartment-service/infrastructure/api"
	"Booking/apartment-service/infrastructure/persistence"
	"Booking/apartment-service/startup/config"

	"go.mongodb.org/mongo-driver/mongo"

	apartment "Booking/apartment-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	client, err := persistence.GetClient("", "")
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (s *Server) initApartmentStore(client *mongo.Client) domain.ApartmentStore {
	store := persistence.NewApartmentMongoDBStore(client)
	store.DeleteAll()

	for _, apartment := range apartments {
		err := store.Insert(apartment)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (s *Server) initApartmentService(store domain.ApartmentStore) *application.ApartmentService {
	return application.NewApartmentService(store)
}

func (s *Server) initApartmentHandler(service *application.ApartmentService) *api.ApartmentHandler {
	return api.NewApartmentHandler(service)
}

func (server *Server) startGrpcServer(apartmentHandler *api.ApartmentHandler) {
	/*
		listener, err := net.Listen("tcp", fmt.Sprintf(":#{server.config.Port}"))
		if err != nil {
			log.Fatalf("Failed to listen: #{err}")
		}
		grpcServer := grpc.NewServer()
		proto.RegisterApartmentServiceServer(grpcServer, apartmentHandler)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: #{err}")
		}
	*/
	listener, err := net.Listen("tcp", server.config.Address)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(listener)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	apartment.RegisterApartmentServiceServer(grpcServer, apartmentHandler)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatal("server error: ", err)
		}
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh

	grpcServer.Stop()
}
