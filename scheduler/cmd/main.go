package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	grpcadapter "task-tracker/scheduler/internal/adapter/grpc"
	"task-tracker/scheduler/internal/config"

	"github.com/robfig/cron/v3"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	conf := config.NewConfig()
	ctx := context.Background()

	taskClient, taskConn, err := grpcadapter.NewTaskClient(conf.TaskServiceAddr)
	if err != nil {
		log.Fatalf("failed to connect to task service: %v", err)
	}

	sendReport := func() {
		_, err := taskClient.SendDailyReport(ctx, &emptypb.Empty{})
		if err != nil {
			log.Printf("send daily report: %v", err)
		}
	}

	c := cron.New()
	c.AddFunc(conf.DailyReportCron, sendReport)
	c.Start()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	<-sigChan

	c.Stop()
	_ = taskConn.Close()
}
