package grpc

import (
	authpb "task-tracker/gen/proto/auth"

	grpclib "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAuthClient(addr string) (authpb.AuthServiceClient, *grpclib.ClientConn, error) {
	conn, err := grpclib.NewClient(addr, grpclib.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}

	client := authpb.NewAuthServiceClient(conn)

	return client, conn, nil
}
