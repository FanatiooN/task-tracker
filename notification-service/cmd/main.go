package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	kafkatopics "task-tracker/kafka"
	"task-tracker/notification-service/internal/adapter/email"
	"task-tracker/notification-service/internal/adapter/kafka"
	"task-tracker/notification-service/internal/adapter/postgres"
	"task-tracker/notification-service/internal/config"
	"task-tracker/notification-service/internal/db"
	"task-tracker/notification-service/internal/service"
	pkgdb "task-tracker/pkg/db"
	pkgkafka "task-tracker/pkg/kafka"
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

	userContactRepo := postgres.NewUserContactRepository(queries)
	notificationRepo := postgres.NewNotificationRepository(queries)

	emailProvider := email.NewNotificationProvider(conf.EmailProviderToken, conf.SenderEmail)
	emailService := service.NewNotificationService(userContactRepo, notificationRepo, emailProvider)

	groupID := "notification-service"

	contactConsumer := pkgkafka.NewConsumer(kafkatopics.UserContactLinkedTopic, groupID, conf.KafkaBrokerAddr)
	go kafka.NewContactConsumer(contactConsumer, userContactRepo).Start(ctx)

	notificationConsumer := pkgkafka.NewConsumer(kafkatopics.SendNotificationTopic, groupID, conf.KafkaBrokerAddr)
	go kafka.NewNotificationConsumer(notificationConsumer, emailService).Start(ctx)

	errChan := make(chan error, 1)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-sigChan:
		return
	case err := <-errChan:
		log.Fatal(err)
	}
}
