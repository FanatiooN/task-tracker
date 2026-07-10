package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"task-tracker/gen/proto/task"
	kafkatopics "task-tracker/kafka"
	pkgdb "task-tracker/pkg/db"
	pkgkafka "task-tracker/pkg/kafka"
	grpcadapter "task-tracker/task-service/internal/adapter/grpc"
	"task-tracker/task-service/internal/adapter/kafka"
	"task-tracker/task-service/internal/adapter/postgres"
	"task-tracker/task-service/internal/config"
	"task-tracker/task-service/internal/db"
	"task-tracker/task-service/internal/service"

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

	repo := postgres.NewTaskRepository(queries)

	err = pkgkafka.CreateTopic(kafkatopics.SendNotificationTopic, conf.KafkaBrokerAddr)
	if err != nil {
		log.Fatal(err)
	}

	producer := pkgkafka.NewProducer(kafkatopics.SendNotificationTopic, conf.KafkaBrokerAddr)
	reportProducer := kafka.NewReportProducer(producer)

	defer producer.Close()

	taskService := service.NewTaskService(repo, reportProducer)
	server := grpcadapter.NewTaskServer(taskService)

	grpcServer := grpc.NewServer()

	task.RegisterTaskServiceServer(grpcServer, server)

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
