package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"gateway-api/internal/app/gateway"
	"gateway-api/internal/pkg/clients/integrator"
	"gateway-api/internal/pkg/clients/validator"
	"gateway-api/internal/pkg/database"
	"gateway-api/internal/pkg/server"
	"gateway-api/internal/pkg/services/specification"
	pbgateway "gateway-api/pkg/gateway"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	databaseConnStr = "postgres://user:password@gateway-db:5432/specifications?sslmode=disable"

	integratorConnStr = "integrator:8083"
	validatorConnStr  = "validator:8085"

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

	pool, err := pgxpool.Connect(ctx, databaseConnStr)
	if err != nil {
		log.Fatalf("connect to db: %s", err.Error())
	}
	defer pool.Close()

	specRepo := database.NewSpecRepository(pool)

	integratorConn, err := grpc.Dial(integratorConnStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error connect to intergrator:", err)
	}
	defer integratorConn.Close()
	integratorCli := integrator.NewClient(integratorConn)

	validatorConn, err := grpc.Dial(validatorConnStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error connect to validator:", err)
	}
	defer validatorConn.Close()
	validatorCli := validator.NewClient(validatorConn)

	specService := specification.NewService(integratorCli, validatorCli, specRepo)

	grpcServer := grpc.NewServer()
	specHandler := gateway.NewSpecHandler(specService)
	pbgateway.RegisterGatewayServiceServer(grpcServer, specHandler)
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
	if err = pbgateway.RegisterGatewayServiceHandler(ctx, mux, gatewayConn); err != nil {
		log.Fatalf("failed to register handler: %s", err.Error())
	}

	s := server.NewHTTPServer(mux, httPort)
	log.Println("launched server at", httPort)
	log.Fatal(s.Launch())
}
