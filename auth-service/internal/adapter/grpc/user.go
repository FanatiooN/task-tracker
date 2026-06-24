package grpc

import (
	userpb "task-tracker/gen/proto/user"

	grpclib "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserClient(addr string) (userpb.UserServiceClient, *grpclib.ClientConn, error) {
	conn, err := grpclib.NewClient(addr, grpclib.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}

	client := userpb.NewUserServiceClient(conn)

	return client, conn, nil
}
