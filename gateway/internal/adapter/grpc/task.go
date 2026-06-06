package grpc

import (
	taskpb "task-tracker/gen/proto/task"

	grpclib "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewTaskClient(addr string) (taskpb.TaskServiceClient, error) {
	conn, err := grpclib.NewClient(addr, grpclib.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := taskpb.NewTaskServiceClient(conn)

	return client, nil
}
