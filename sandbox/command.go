package sandbox

import (
	// "github.com/dcrosby42/picfinder/commands"
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
			sandbox_fileinfo_command(),
		},
	}
}

func sandbox_fileinfo_command() cli.Command {
	return cli.Command{
		Name:  "insert",
		Usage: "Try to insert and retrieve FileInfo records",
		Action: func(c *cli.Context) error {
			db, err := dbutil.ConnectDatabase()
			if err != nil {
				return cli.NewExitError(err.Error(), -37)
			}
			err = insertAndRetrieveFileInfo(db)
			if err != nil {
				return cli.NewExitError(err.Error(), -1)
			}
			return nil
		},
	}
}

func insertAndRetrieveFileInfo(db *sqlx.DB) error {
	f := fileinfo.FileInfo{
		Id:                 1,
		Host:               "XandersBeatBox",
		Path:               []byte("/Users/crosby/something.txt"),
		PathHash:           1234,
		Size:               1024,
		ContentHash:        []byte("fake content hash"),
		ContentHashLower32: 5678,
		Type:               "jpg",
		Kind:               "picture",
		ScannedAt:          time.Now(),
		FileModifiedAt:     time.Now().Add(-1 * time.Hour),
	}
	_, err := db.NamedExec("INSERT INTO file_info (host,path,path_hash,content_hash,content_hash_lower_32,size,type,kind,scanned_at,file_modified_at) VALUES(:host,:path,:path_hash,:content_hash,:content_hash_lower_32,:size,:type,:kind,:scanned_at,:file_modified_at)", f)
	return err

	return nil
}
