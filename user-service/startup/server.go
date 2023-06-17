package startup

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/NikolinaSesa/Booking/user-service/application"
	"github.com/NikolinaSesa/Booking/user-service/domain"
	"github.com/NikolinaSesa/Booking/user-service/infrastructure/api"
	"github.com/NikolinaSesa/Booking/user-service/infrastructure/persistence"
	"github.com/NikolinaSesa/Booking/user-service/startup/config"
	"go.mongodb.org/mongo-driver/mongo"

	user "github.com/NikolinaSesa/Booking/user-service/proto"

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
	userStore := s.initUserStore(mongoClient)

	userService := s.initUserService(userStore)

	userHandler := s.initUserHandler(userService)

	s.startGrpcServer(userHandler)
}

func (s *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient("", "")
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

	/*
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

	user.RegisterUserServiceServer(grpcServer, userHandler)

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
