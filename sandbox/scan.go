package sandbox

import (
	"fmt"
	"os"
	"time"

	"github.com/dcrosby42/picfinder/scan"
	"github.com/urfave/cli"
)

func sandbox_scan_command() cli.Command {
	return cli.Command{
		Name:  "scan",
		Usage: "scan dirs",
		Action: func(c *cli.Context) error {

			dirname := "/Users/crosby/Pictures"

			// result, err := scan.CountAll(dirname)
			// if err != nil {
			// 	return cli.NewExitError(err.Error(), -1)
			// }
			// fmt.Printf("Counted all in %q:\n", dirname)
			// fmt.Printf("Files: %d\n", result.Files)
			// fmt.Printf("Dirs: %d\n", result.Dirs)
			// fmt.Printf("Elapsed: %s\n", result.Elapsed)

			fcount := 0
			started := time.Now()
			err := scan.WalkFiles(dirname, func(dname string, info os.FileInfo) error {
				fcount += 1
				filepath := dname + "/" + info.Name()
				// contentHash, herr := scan.HashFileContentSha256(filepath)
				// if herr != nil {
				// 	fmt.Printf("!!!! ERR filepath=%s err=%s\n", filepath, herr)
				// }
				// fmt.Printf("%q, %x\n", filepath, contentHash)
				fmt.Printf("%q\n", filepath)
				return nil
			})
			elapsed := time.Now().Sub(started)
			fmt.Printf("Elapsed: %s\n", elapsed)
			if err != nil {
				return cli.NewExitError(err.Error(), -1)
			}

			fmt.Printf("Walked %d files\n", fcount)
			// fmt.Printf("Files: %d\n", result.Files)
			// fmt.Printf("Dirs: %d\n", result.Dirs)
			// fmt.Printf("Elapsed: %s\n", result.Elapsed)
			return nil
		},
	}
}
