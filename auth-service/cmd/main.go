package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	grpcadapter "task-tracker/auth-service/internal/adapter/grpc"
	"task-tracker/auth-service/internal/adapter/postgres"
	"task-tracker/auth-service/internal/config"
	"task-tracker/auth-service/internal/db"
	"task-tracker/auth-service/internal/service"
	"task-tracker/gen/proto/auth"
	pkgdb "task-tracker/pkg/db"

	"google.golang.org/grpc"
)

func main() {
	conf := config.NewConfig()
	ctx := context.Background()

	pool, err := pkgdb.NewPool(ctx, conf.DatabaseURL())
	if err != nil {
		log.Fatal(err)
	}

	defer pool.Close()
	queries := db.New(pool)

	credRepo := postgres.NewCredentialRepository(queries)
	tokenRepo := postgres.NewTokenRepository(queries)

	authService := service.NewAuthService(credRepo, tokenRepo, conf.JWTSecret, conf.JWTAccessTTL, conf.JWTRefreshTTL)
	server := grpcadapter.NewAuthServer(authService)

	grpcServer := grpc.NewServer()

	auth.RegisterAuthServiceServer(grpcServer, server)

	lis, err := net.Listen("tcp", conf.GRPCPort)
	if err != nil {
		log.Fatal(err)
	}

	errChan := make(chan error, 1)

	go func(errChan chan<- error) {
		err := grpcServer.Serve(lis)
		if err != nil {
			errChan <- err
		}
	}(errChan)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-sigChan:
		grpcServer.GracefulStop()
		return
	case err := <-errChan:
		log.Fatal(err)
	}
}
