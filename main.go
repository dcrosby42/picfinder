package main

import (
	"os"

	"github.com/dcrosby42/picfinder/api_client"
	"github.com/dcrosby42/picfinder/api_server"
	"github.com/dcrosby42/picfinder/commands"
	"github.com/dcrosby42/picfinder/config"
	"github.com/dcrosby42/picfinder/dbutil"
	"github.com/dcrosby42/picfinder/sandbox"
	"github.com/urfave/cli"
	// _ "github.com/golang/protobuf/proto"
	// _ "github.com/golang/protobuf/protoc-gen-go"
	// _ "google.golang.org/grpc"
)

func main() {
	app := cli.NewApp()

	app.Flags = config.GlobalFlags()
	app.Commands = []cli.Command{
		commands.ScanCommand(),
		api_server.Command(),
		api_client.PingCommand(),
		dbutil.Command(),
		sandbox.Command(),
		config.DumpConfigCommand(),
	}

	app.Run(os.Args)
}
