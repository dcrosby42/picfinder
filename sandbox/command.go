package sandbox

import (
	// "github.com/dcrosby42/picfinder/commands"
	"fmt"
	"os"
	"time"

	"github.com/dcrosby42/picfinder/dbutil"
	"github.com/dcrosby42/picfinder/fileinfo"
	"github.com/dcrosby42/picfinder/scan"
	"github.com/jmoiron/sqlx"
	"github.com/urfave/cli"
)

func Command() cli.Command {
	return cli.Command{
		Name:  "sandbox",
		Usage: "Holder for one-offs and experiments",
		Subcommands: []cli.Command{
			sandbox_insert_command(),
			sandbox_retrieve_command(),
			sandbox_scan_command(),
			sandbox_ext_command(),
		},
	}
}

func sandbox_insert_command() cli.Command {
	return cli.Command{
		Name:  "insert",
		Usage: "Try to insert FileInfo record",
		Action: func(c *cli.Context) error {
			db, err := dbutil.ConnectDatabase()
			if err != nil {
				return cli.NewExitError(err.Error(), -37)
			}
			err = insertFileInfo(db)
			if err != nil {
				return cli.NewExitError(err.Error(), -1)
			}
			return nil
		},
	}
}
func sandbox_retrieve_command() cli.Command {
	return cli.Command{
		Name:  "retrieve",
		Usage: "Try to query FileInfo record",
		Action: func(c *cli.Context) error {
			db, err := dbutil.ConnectDatabase()
			if err != nil {
				return cli.NewExitError(err.Error(), -37)
			}
			err = queryFileInfo(db)
			if err != nil {
				return cli.NewExitError(err.Error(), -1)
			}
			return nil
		},
	}
}

func insertFileInfo(db *sqlx.DB) error {
	f := fileinfo.FileInfo{
		// Id:                 1,
		Host:               "XandersBeatBox",
		Path:               []byte("/Users/crosby/something.txt"),
		PathHash:           1234,
		Size:               1024,
		ContentHash:        []byte("fake content hash"),
		ContentHashLower32: 5678,
		Type:               fileinfo.JpegType,
		Kind:               fileinfo.PictureKind,
		ScannedAtUnix:      time.Now().Unix(),
		FileModifiedAtUnix: time.Now().Add(-1 * time.Hour).Unix(),
	}
	_, err := fileinfo.Insert(db, &f)
	if err != nil {
		return err
	}
	fmt.Printf("!!!! insert FileInfo id=%d\n", f.Id)

	return nil
}
func queryFileInfo(db *sqlx.DB) error {
	f2, err := fileinfo.Get(db, 2)
	if err != nil {
		return err
	}
	fmt.Printf("!!!! f2: %#v\n", f2)

	return nil
}

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
