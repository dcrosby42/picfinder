package sandbox

import (
	// "github.com/dcrosby42/picfinder/commands"
	"fmt"
	"os"
	"time"

	context "golang.org/x/net/context"

	"github.com/dcrosby42/picfinder/api_client"
	"github.com/dcrosby42/picfinder/fileinfo"
	picfinder_grpc "github.com/dcrosby42/picfinder/grpc"
	"github.com/dcrosby42/picfinder/scan"
	"github.com/urfave/cli"
)

func Command() cli.Command {
	return cli.Command{
		Name:  "sandbox",
		Usage: "Holder for one-offs and experiments",
		Subcommands: []cli.Command{
			// sandbox_insert_command(),
			// sandbox_retrieve_command(),
			sandbox_scan_command(),
			sandbox_ext_command(),
			sandbox_client_command(),
			sandbox_testgrpcstream_command(),
		},
	}
}

// func sandbox_insert_command() cli.Command {
// 	return cli.Command{
// 		Name:  "insert",
// 		Usage: "Try to insert FileInfo record",
// 		Action: func(c *cli.Context) error {
// 			db, err := dbutil.ConnectDatabase()
// 			if err != nil {
// 				return cli.NewExitError(err.Error(), -37)
// 			}
// 			err = insertFileInfo(db)
// 			if err != nil {
// 				return cli.NewExitError(err.Error(), -1)
// 			}
// 			return nil
// 		},
// 	}
// }
// func sandbox_retrieve_command() cli.Command {
// 	return cli.Command{
// 		Name:  "retrieve",
// 		Usage: "Try to query FileInfo record",
// 		Action: func(c *cli.Context) error {
// 			db, err := dbutil.ConnectDatabase()
// 			if err != nil {
// 				return cli.NewExitError(err.Error(), -37)
// 			}
// 			err = queryFileInfo(db)
// 			if err != nil {
// 				return cli.NewExitError(err.Error(), -1)
// 			}
// 			return nil
// 		},
// 	}
// }
//
// func insertFileInfo(db *sqlx.DB) error {
// 	f := fileinfo.FileInfo{
// 		// Id:                 1,
// 		Host:               "XandersBeatBox",
// 		Path:               []byte("/Users/crosby/something.txt"),
// 		PathHash:           1234,
// 		Size:               1024,
// 		ContentHash:        []byte("fake content hash"),
// 		ContentHashLower32: 5678,
// 		Type:               fileinfo.JpegType,
// 		Kind:               fileinfo.PictureKind,
// 		ScannedAtUnix:      time.Now().Unix(),
// 		FileModifiedAtUnix: time.Now().Add(-1 * time.Hour).Unix(),
// 	}
// 	_, err := fileinfo.Insert(db, f)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Printf("!!!! insert FileInfo id=%d\n", f.Id)
//
// 	return nil
// }
// func queryFileInfo(db *sqlx.DB) error {
// 	f2, err := fileinfo.Get(db, 2)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Printf("!!!! f2: %#v\n", f2)
//
// 	return nil
// }

func sandbox_ext_command() cli.Command {
	return cli.Command{
		Name:  "findext",
		Usage: "Scan a dir and summarize all file extensions therein",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "dir",
				Usage: "The dir to start scanning in",
				Value: "/Users/crosby/Pictures",
			},
		},
		Action: func(c *cli.Context) error {
			dirname := c.String("dir")
			err := summarizeFileExtensions(dirname)
			if err != nil {
				return cli.NewExitError(err.Error(), -1)
			}
			return nil
		},
	}
}

func summarizeFileExtensions(dirname string) error {
	extSet := make(map[string]int)
	scan.WalkFiles(dirname, func(dname string, info os.FileInfo) error {
		ext := scan.LowercaseFileExt(info.Name())
		if ext != "" {
			count, ok := extSet[ext]
			if !ok {
				extSet[ext] = 1
			} else {
				extSet[ext] = count + 1
			}
			// extSet[ext] = true
		}
		return nil
	})
	fmt.Printf("Found %d distinct extensions\n", len(extSet))
	byKind := make(map[fileinfo.FileKind]int)
	for ext, count := range extSet {
		kind := fileinfo.FileKindForExt(ext)
		fmt.Printf("  %s, %s(%s), %d\n", ext, kind, fileinfo.FileTypeForExt(ext), count)
		kc, ok := byKind[kind]
		if !ok {
			kc = 0
		}
		byKind[kind] = kc + count

	}
	fmt.Printf("By kind:\n")
	for kind, count := range byKind {
		fmt.Printf("Kind %s: %d\n", kind, count)
	}
	return nil
}

func sandbox_client_command() cli.Command {
	return cli.Command{
		Name:  "client",
		Usage: "Test the client conn to picfinder server",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "host",
				Usage: "The api server host",
				Value: "127.0.0.1",
			},
			cli.StringFlag{
				Name:  "port",
				Usage: "The api server port",
				Value: "13131",
			},
		},
		Action: func(c *cli.Context) error {
			host := c.String("host")
			port := c.String("port")

			client, closeConn, err := api_client.NewClient_HostPort(host, port)
			if err != nil {
				return cli.NewExitError(err.Error(), -1)
			}

			defer closeConn()

			request := &picfinder_grpc.AddFileRequest{}
			request.FileInfo = &picfinder_grpc.FileInfo{
				Host:               "XandersBeatBox",
				Path:               []byte("/Users/crosby/something.txt"),
				PathHash:           1234,
				Size:               1024,
				ContentHash:        []byte("fake content hash"),
				ContentHashLower32: 5678,
				Type:               string(fileinfo.JpegType),
				Kind:               string(fileinfo.PictureKind),
				ScannedAtUnix:      time.Now().Unix(),
				FileModifiedAtUnix: time.Now().Add(-1 * time.Hour).Unix(),
			}

			fmt.Printf("!!!! Sending AddFile request %#v\n", request)

			resp, err := client.AddFile(context.Background(), request)
			if err != nil {
				return cli.NewExitError(err.Error(), -1)
			}
			fmt.Printf("!!!! Response to AddFile(): %#v\n", resp)

			return nil
		},
	}
}

func sandbox_testgrpcstream_command() cli.Command {
	return cli.Command{
		Name:  "clientstr",
		Usage: "GRPC client stream test",
		Flags: api_client.RemoteServerFlags(),
		Action: func(c *cli.Context) error {
			TestGrpcStream(c.String("host"), c.String("port"))
			return nil
		},
	}
}
