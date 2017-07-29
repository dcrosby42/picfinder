package api_client

import (
	"net"

	"google.golang.org/grpc"

	picfinder_grpc "github.com/dcrosby42/picfinder/grpc"
)

func NewClient_HostPort(host, port string) (picfinder_grpc.PicfinderClient, func() error, error) {
	return NewClient_Addr(net.JoinHostPort(host, port))
}

func NewClient_Addr(addr string) (picfinder_grpc.PicfinderClient, func() error, error) {
	gconn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, func() error { return nil }, err
	}
	gclient := picfinder_grpc.NewPicfinderClient(gconn)
	return gclient, gconn.Close, nil
}
