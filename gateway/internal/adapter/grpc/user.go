package grpc

import (
	userpb "task-tracker/gen/proto/user"

	grpclib "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserClient(addr string) (userpb.UserServiceClient, error) {
	conn, err := grpclib.NewClient(addr, grpclib.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := userpb.NewUserServiceClient(conn)

	return client, nil
}
