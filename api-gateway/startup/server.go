package startup

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	cfg "github.com/NikolinaSesa/Booking/api-gateway/startup/config"
	userProto "github.com/NikolinaSesa/Booking/user-service/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux
}

func NewServer(config *cfg.Config) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}
	server.initHandlers()
	return server
}

func (s *Server) initHandlers() {

	/*
		fmt.Println("***************************************Tu sam")

		opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
		userEndpoint := fmt.Sprintf("%s", s.config.UserServiceAddress)

		fmt.Println("***************************************Tu sam", userEndpoint)

		err := user.RegisterUserServiceHandlerFromEndpoint(context.TODO(), s.mux, userEndpoint, opts)
		if err != nil {
			panic(err)
		}
	*/
	gwmux := runtime.NewServeMux()

	//user-service
	conn, err := grpc.DialContext(context.Background(), s.config.UserServiceAddress, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("Failed to dial server: ", err)
	}

	client := userProto.NewUserServiceClient(conn)
	err = userProto.RegisterUserServiceHandlerClient(context.Background(), gwmux, client)

	if err != nil {
		log.Fatal("Failed to register gateway: ", err)
	}

	//apartment-service

	_, err2 := grpc.DialContext(context.Background(), s.config.ApartmentServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err2 != nil {
		log.Fatal("Failed to dial server 2: ", err2)
	}

	//
	gwServer := &http.Server{
		Addr:    s.config.Address,
		Handler: gwmux,
	}
	go func() {
		if err := gwServer.ListenAndServe(); err != nil {
			log.Fatal("server error: ", err)
		}
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh

	if err = gwServer.Close(); err != nil {
		log.Fatalln("error while stopping server: ", err)
	}

}

//func (s *Server) Start() {
//	log.Fatal(http.ListenAndServe(s.config.Address, s.mux))
//}
