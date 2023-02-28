package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/MickStanciu/potter-api/potterapi/internal/gen/potter"
	"github.com/MickStanciu/potter-api/potterapi/internal/server"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen to port 9000: %v", err)
	}

	grpcServer := grpc.NewServer()
	potter.RegisterHogwartsServiceServer(grpcServer, &server.Server{})
	reflection.Register(grpcServer)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to start grpc server: %v", err)
		}
	}()

	//GRPC Gateway
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:9000",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial grpc server: %v", err)
	}

	mux := runtime.NewServeMux()
	if err := potter.RegisterHogwartsServiceHandler(context.Background(), mux, conn); err != nil {
		log.Fatalf("failed to register: %v", err)
	}

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
