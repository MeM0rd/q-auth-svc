package main

import (
	"context"
	"github.com/MeM0rd/q-api-gateway/pkg/logger"
	authPbService "github.com/MeM0rd/q-api-gateway/pkg/pb/auth"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	l := logger.New()

	listener, err := net.Listen("tcp", os.Getenv("PORT"))
	if err != nil {
		l.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	authPbService.RegisterAuthPbServiceServer(s, &server{})
	l.Infof("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		l.Fatalf("failed to serve: %v", err)
	}
}

type server struct {
	authPbService.UnimplementedAuthPbServiceServer
}

func (s *server) Register(ctx context.Context, in *authPbService.RegisterRequest) (*authPbService.RegisterResponse, error) {
	log.Printf("Received: %v", in.GetEmail())
	return &authPbService.RegisterResponse{Status: "OK", Err: ""}, nil
}
