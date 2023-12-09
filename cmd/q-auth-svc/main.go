package main

import (
	"github.com/MeM0rd/q-auth-svc/internal/auth"
	"github.com/MeM0rd/q-auth-svc/pkg/client/postgres"
	"github.com/MeM0rd/q-auth-svc/pkg/logger"
	authPbService "github.com/MeM0rd/q-auth-svc/pkg/pb/auth"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"net"
	"os"
)

func init() {
	godotenv.Load()

	postgres.Open()
}

func main() {
	l := logger.New()

	listener, err := net.Listen("tcp", os.Getenv("PORT"))
	if err != nil {
		l.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	authPbService.RegisterAuthPbServiceServer(s, &auth.Server{
		Logger: l,
	})
	l.Infof("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		l.Fatalf("failed to serve: %v", err)
	}
}
