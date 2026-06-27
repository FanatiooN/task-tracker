package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	grpcadapter "task-tracker/gateway/internal/adapter/grpc"
	httpadapter "task-tracker/gateway/internal/adapter/http"
	"task-tracker/gateway/internal/config"
)

func main() {
	conf := config.NewConfig()

	userClient, userConn, err := grpcadapter.NewUserClient(conf.UserServiceAddr)
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
	}

	taskClient, taskConn, err := grpcadapter.NewTaskClient(conf.TaskServiceAddr)
	if err != nil {
		log.Fatalf("failed to connect to task service: %v", err)
	}

	authClient, authConn, err := grpcadapter.NewAuthClient(conf.AuthServiceAddr)
	if err != nil {
		log.Fatalf("failed to connect to auth service: %v", err)
	}

	router := httpadapter.NewRouter(httpadapter.Handlers{
		User:  httpadapter.NewUserHandler(userClient),
		Task:  httpadapter.NewTaskHandler(taskClient),
		Auth:  httpadapter.NewAuthHandler(authClient),
		OAuth: httpadapter.NewOAuthHandler(authClient, conf.OAuth.GoogleClientID, conf.OAuth.GoogleRedirectURI, conf.FrontendAddr),
	}, authClient)

	log.Printf("gateway started, port = %s", conf.Port)

	server := http.Server{
		Addr:    conf.Port,
		Handler: router,
	}

	errChan := make(chan error, 1)

	go func(errChan chan<- error) {
		err := server.ListenAndServe()
		if err != nil {
			errChan <- err
		}
	}(errChan)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-sigChan:
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		_ = server.Shutdown(ctx)
		_ = userConn.Close()
		_ = taskConn.Close()
		_ = authConn.Close()

		return
	case err := <-errChan:
		log.Fatal(err)
	}
}
