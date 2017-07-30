package api_client

import (
	"context"
	"fmt"

	picfinder_grpc "github.com/dcrosby42/picfinder/grpc"
	"github.com/urfave/cli"
)

func RemoteServerFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "server",
			Usage: "The api server host",
			Value: "127.0.0.1",
		},
		cli.StringFlag{
			Name:  "port",
			Usage: "The api server port",
			Value: "13131",
		},
	}
}

func PingCommand() cli.Command {
	return cli.Command{
		Name:  "ping",
		Usage: "Ping the Picfinder GRPC API Server",
		Flags: RemoteServerFlags(),
		Action: func(c *cli.Context) error {
			host := c.String("server")
			port := c.String("port")
			err := DoPingServer(host, port)
			if err != nil {
				return cli.NewExitError(err.Error(), -1)
			}
			return nil

		},
	}
}

func DoPingServer(host, port string) error {
	fmt.Printf("Pinging api server %s:%s\n", host, port)
	client, closeConn, err := NewClient_HostPort(host, port)
	if err != nil {
		return err
	}

	defer closeConn()

	request := &picfinder_grpc.PingRequest{}
	resp, err := client.Ping(context.Background(), request)
	if err != nil {
		return err
	}
	fmt.Printf("PingResponse Status=%d Message=%q\n", resp.Header.Status, resp.Header.Message)
	return nil
}
