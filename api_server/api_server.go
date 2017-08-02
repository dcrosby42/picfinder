package api_server

import (
	"fmt"
	"io"
	"net"
	"time"

	"github.com/dcrosby42/picfinder/config"
	"github.com/dcrosby42/picfinder/dbutil"
	"github.com/dcrosby42/picfinder/fileinfo"
	picfinder_grpc "github.com/dcrosby42/picfinder/grpc"
	"github.com/jmoiron/sqlx"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

const ErrGeneral = 1
const ErrBadParams = 2
const ErrDbConn = 3

// Implements the picfinder_grpc.PicfinderServer interface:
type apiServer struct {
	db *sqlx.DB
}

func (me *apiServer) Ping(ctx context.Context, request *picfinder_grpc.PingRequest) (*picfinder_grpc.PingResponse, error) {
	fmt.Printf("Received ping request\n")
	resp := &picfinder_grpc.PingResponse{
		Header: &picfinder_grpc.ResponseHeader{
			Status:  0,
			Message: "PONG",
		},
	}
	return resp, nil
}

func (me *apiServer) AddFile(ctx context.Context, request *picfinder_grpc.AddFileRequest) (*picfinder_grpc.AddFileResponse, error) {
	resp := &picfinder_grpc.AddFileResponse{}
	resp.Header = &picfinder_grpc.ResponseHeader{}
	if request.FileInfo == nil {
		fmt.Printf("!!! ERR AddFile() invoked with nil FileInfo\n")
		resp.Header.Status = ErrBadParams
		resp.Header.Message = "AddFileRequest.FileInfo must not be bil"
		return resp, nil
	}
	if request.FileInfo != nil {
		fileInfo := fileinfo.FromGrpcFileInfo(request.FileInfo)

		fileInfo, updateAction, err := fileinfo.InsertOrUpdateByHostPath(me.db, fileInfo)
		resp.UpdateAction = string(updateAction)
		if err != nil {
			fmt.Printf("!!! ERR AddFile(): host=%s path=%s contentHashLower32=%d updateAction=%s, InsertOrUpdateByHostPath err=%s\n", fileInfo.Host, fileInfo.PathString(), fileInfo.ContentHashLower32, updateAction, err)
			resp.Header.Status = ErrDbConn
			resp.Header.Message = fmt.Sprintf("Database error during InsertOrUpdateByHostPath")
			return resp, nil
		}
		fmt.Printf("AddFile(): id=%d host=%s path=%s contentHashLower32=%d updateAction=%s\n", fileInfo.Id, fileInfo.Host, fileInfo.PathString(), fileInfo.ContentHashLower32, updateAction)
	}

	return resp, nil
}

func (me *apiServer) Sandbox_GetData(stream picfinder_grpc.Picfinder_Sandbox_GetDataServer) error {
	fmt.Printf("Sandbox_GetData invoked.")

	clientDone := make(chan bool)
	resp := &picfinder_grpc.Sandbox_GetDataResponse{}
	go func() {
		timer := time.NewTicker(1 * time.Second)
	Dance:
		for {
			select {
			case <-timer.C:
				resp.Data = []byte("This is the data")
				stream.Send(resp)

			case <-clientDone:
				break Dance
			}
		}
	}()

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			fmt.Printf("Sandbox_GetData EOF, exit.")
			close(clientDone)
			return nil
		}
		if err != nil {
			fmt.Printf("Sandbox_GetData recv err=%q, exit", err)
			close(clientDone)
			return err
		}
		fmt.Printf("Sandbox_GetData rec'd request: %#v\n", req)
	}

}

func BuildAndListen(apiServerConfig config.ApiServerConfig, dbConfig config.DbConfig) error {
	bindAddr := apiServerConfig.BindAddr
	fmt.Printf("Picfinder GRPC Api Server startup, bindAddr=%q\n", bindAddr)

	db, err := dbutil.ConnectDatabase(dbConfig)
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
