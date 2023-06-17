package startup

import (
	"fmt"
	"log"
	"net"

	"github.com/NikolinaSesa/Booking/user-service/application"
	"github.com/NikolinaSesa/Booking/user-service/domain"
	"github.com/NikolinaSesa/Booking/user-service/infrastructure/api"
	"github.com/NikolinaSesa/Booking/user-service/infrastructure/persistence"
	"github.com/NikolinaSesa/Booking/user-service/proto"
	"github.com/NikolinaSesa/Booking/user-service/startup/config"
	"go.mongodb.org/mongo-driver/mongo"

	"google.golang.org/grpc"
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
	userStore := s.initUserStore(mongoClient)

	userService := s.initUserService(userStore)

	userHandler := s.initUserHandler(userService)

	s.startGrpcServer(userHandler)
}

func (s *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(s.config.UserDBHost, s.config.UserDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (s *Server) initUserStore(client *mongo.Client) domain.UserStore {
	store := persistence.NewUserMongoDBStore(client)
	store.DeleteAll()

	for _, user := range users {
		err := store.Insert(user)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (s *Server) initUserService(store domain.UserStore) *application.UserService {
	return application.NewUserService(store)
}

func (s *Server) initUserHandler(service *application.UserService) *api.UserHandler {
	return api.NewUserHandler(service)
}

func (server *Server) startGrpcServer(userHandler *api.UserHandler) {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	fmt.Println("***************************************Tu sam", listener, err)
	if err != nil {
		log.Fatalf("Failed to listen: #{err}")
	}
	grpcServer := grpc.NewServer()
	proto.RegisterUserServiceServer(grpcServer, userHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: #{err}")
	}

}
