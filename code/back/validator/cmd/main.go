package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	validator "validator-api/internal/app"
	"validator-api/internal/pkg/server"
	pbvalidator "validator-api/pkg/validator"
)

const (
	httPort  = "8080"
	grpcPort = "8081"
)

func main() {
	ctx := context.Background()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalf("error when listen grpc: %s", err.Error())
	}
	defer listener.Close()

	grpcServer := grpc.NewServer()
	integratorHandler := validator.NewValidatorHandler()
	pbvalidator.RegisterValidatorServiceServer(grpcServer, integratorHandler)
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed when serve: %s", err.Error())
		}
	}()

	gatewayConn, err := grpc.NewClient(
		listener.Addr().String(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect: %s", err.Error())
	}
	defer gatewayConn.Close()

	mux := runtime.NewServeMux()
	if err = pbvalidator.RegisterValidatorServiceHandler(ctx, mux, gatewayConn); err != nil {
		log.Fatalf("failed to register handler: %s", err.Error())
	}

	s := server.NewHTTPServer(mux, httPort)
	log.Println("launched server at", httPort)
	log.Fatal(s.Launch())
}
