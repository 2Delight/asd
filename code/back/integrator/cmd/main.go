package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"integrator-api/internal/app/integrator"
	"integrator-api/internal/pkg/server"
	pbintegrator "integrator-api/pkg/integrator"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	httPort  = "8082"
	grpcPort = "8083"
)

func main() {
	ctx := context.Background()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalf("error when listen grpc: %s", err.Error())
	}
	defer listener.Close()

	grpcServer := grpc.NewServer()
	integratorHandler := integrator.NewIntegratorHandler()
	pbintegrator.RegisterWorkerServiceServer(grpcServer, integratorHandler)
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
	if err = pbintegrator.RegisterWorkerServiceHandler(ctx, mux, gatewayConn); err != nil {
		log.Fatalf("failed to register handler: %s", err.Error())
	}

	s := server.NewHTTPServer(mux, httPort)
	log.Println("launched server at", httPort)
	log.Fatal(s.Launch())
}
