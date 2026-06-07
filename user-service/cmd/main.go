package main

import (
	"context"
	"log"
	"net"
	"task-tracker/gen/proto/user"
	pkgdb "task-tracker/pkg/db"
	grpcadapter "task-tracker/user-service/internal/adapter/grpc"
	"task-tracker/user-service/internal/adapter/postgres"
	"task-tracker/user-service/internal/config"
	"task-tracker/user-service/internal/db"
	"task-tracker/user-service/internal/service"

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

	repo := postgres.NewUserRepository(queries)
	userService := service.NewUserService(repo)
	server := grpcadapter.NewUserServer(userService)

	grpcServer := grpc.NewServer()

	user.RegisterUserServiceServer(grpcServer, server)

	lis, err := net.Listen("tcp", conf.GRPCPort)
	if err != nil {
		log.Fatal(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
