package sandbox

import (
	// "github.com/dcrosby42/picfinder/commands"
	"fmt"
	"time"

	"github.com/dcrosby42/picfinder/dbutil"
	"github.com/dcrosby42/picfinder/fileinfo"
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
		Type:               "jpg",
		Kind:               "picture",
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
