package grpc

import (
	userpb "task-tracker/gen/proto/user"

	grpclib "google.golang.org/grpc"
)

func NewUserClient(addr string) (userpb.UserServiceClient, error) {
	conn, err := grpclib.NewClient(addr)
	if err != nil {
		return nil, err
	}

	client := userpb.NewUserServiceClient(conn)

	return client, nil
}
