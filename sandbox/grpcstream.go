package sandbox

import (
	"fmt"
	"io"
	"sync"

	context "golang.org/x/net/context"

	"github.com/dcrosby42/picfinder/api_client"
	picfinder_grpc "github.com/dcrosby42/picfinder/grpc"
)

func TestGrpcStream(host, port string) {
	err := api_client.DoPingServer(host, port)
	if err != nil {
		fmt.Printf("TestGrpcStream ping err=%s\n", err)
		return
	}

	client, closeConn, err := api_client.NewClient_HostPort(host, port)
	if err != nil {
		fmt.Printf("TestGrpcStream connect err=%s\n", err)
		return
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		doGetData(client, "A")
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		doGetData(client, "B")
		wg.Done()
	}()
	wg.Wait()
	defer closeConn()
}

func doGetData(client picfinder_grpc.PicfinderClient, name string) {
	stream, err := client.Sandbox_GetData(context.Background())
	if err != nil {
		fmt.Printf("(%s) TestGrpcStream rpc err=%s\n", name, err)
		return
	}

	waitc := make(chan bool)
	count := 0
	go func() {
		fmt.Printf("(%s) TestGrpcStream spinning off a reader\n", name)
		defer close(waitc)
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				fmt.Printf("(%s) TestGrpcStream recv EOF\n", name)
				return
			}
			if err != nil {
				fmt.Printf("(%s) TestGrpcStream recv err=%s\n", name, err)
				return
			}
			fmt.Printf("(%s) TestGrpcStream recv: %s\n", name, string(in.Data))
			count++
			if count >= 5 {
				fmt.Printf("(%s) TestGrpcStream count to 10, returning\n", name)
				return
			}
		}
	}()

	req := &picfinder_grpc.Sandbox_GetDataRequest{}
	err = stream.Send(req)
	if err != nil {
		fmt.Printf("TestGrpcStream send err=%s\n", err)
		return
	}

	<-waitc
	stream.CloseSend()

}
