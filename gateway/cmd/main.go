package main

import (
	"log"
	"net/http"

	grpcadapter "task-tracker/gateway/internal/adapter/grpc"
	httpadapter "task-tracker/gateway/internal/adapter/http"
	"task-tracker/gateway/internal/config"
)

func main() {
	conf := config.NewConfig()

	userClient, err := grpcadapter.NewUserClient(conf.UserServiceAddr)
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
	}

	taskClient, err := grpcadapter.NewTaskClient(conf.TaskServiceAddr)
	if err != nil {
		log.Fatalf("failed to connect to task service: %v", err)
	}

	router := httpadapter.NewRouter(httpadapter.Handlers{
		User: httpadapter.NewUserHandler(userClient),
		Task: httpadapter.NewTaskHandler(taskClient),
	})

	log.Printf("gateway started, port = %s", conf.Port)

	if err := http.ListenAndServe(conf.Port, router); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
