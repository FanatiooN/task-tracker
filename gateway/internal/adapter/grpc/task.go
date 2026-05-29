package grpc

import (
	taskpb "task-tracker/gen/proto/task"

	grpclib "google.golang.org/grpc"
)

func NewTaskClient(addr string) (taskpb.TaskServiceClient, error) {
	conn, err := grpclib.NewClient(addr)
	if err != nil {
		return nil, err
	}

	client := taskpb.NewTaskServiceClient(conn)

	return client, nil
}
