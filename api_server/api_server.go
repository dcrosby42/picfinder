package api_server

import (
	"fmt"
	"net"

	"github.com/dcrosby42/picfinder/dbutil"
	"github.com/dcrosby42/picfinder/fileinfo"
	picfinder_grpc "github.com/dcrosby42/picfinder/grpc"
	"github.com/jmoiron/sqlx"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Implements the picfinder_grpc.PicfinderServer interface:
type apiServer struct {
	db *sqlx.DB
}

func (me *apiServer) AddFile(ctx context.Context, request *picfinder_grpc.AddFileRequest) (*picfinder_grpc.AddFileResponse, error) {
	fmt.Printf("!!!! AddFile request.FileInfo%#v\n", request.FileInfo)
	if request.FileInfo != nil {
		fileInfo := fileinfo.FromGrpcFileInfo(request.FileInfo)
		_, err := fileinfo.Insert(me.db, &fileInfo)
		if err != nil {
			fmt.Printf("!!!! --> ERR! %s\n", err)
		}
	}

	resp := &picfinder_grpc.AddFileResponse{}
	return resp, nil
}

func BuildAndListen(bindAddr string) error {
	fmt.Printf("Picfinder GRPC Api Server startup, bindAddr=%q\n", bindAddr)

	db, err := dbutil.ConnectDatabase()
	if err != nil {
		return fmt.Errorf("Failed to connect database err=%s", err)
	}

	listener, err := net.Listen("tcp", bindAddr)
	if err != nil {
		return fmt.Errorf("Failed to bind listener: %v", err)
	}

	grpcServer := grpc.NewServer()
	picfinder_grpc.RegisterPicfinderServer(grpcServer, &apiServer{db})
	grpcServer.Serve(listener)
	return nil
}
