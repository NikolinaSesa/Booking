package startup

import (
	"context"
	"fmt"
	"log"
	"net/http"

	cfg "github.com/NikolinaSesa/Booking/api-gateway/startup/config"
	user "github.com/NikolinaSesa/Booking/user-service/proto"
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

	fmt.Println("***************************************Tu sam")

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	userEndpoint := fmt.Sprintf("%s:%s", s.config.UserHost, s.config.UserPort)

	fmt.Println("***************************************Tu sam", userEndpoint)

	err := user.RegisterUserServiceHandlerFromEndpoint(context.TODO(), s.mux, userEndpoint, opts)
	if err != nil {
		panic(err)
	}
}

func (s *Server) Start() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", s.config.Port), s.mux))
}
