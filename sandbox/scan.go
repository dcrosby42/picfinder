package sandbox

import (
	"fmt"

	"github.com/dcrosby42/picfinder/fileinfo"
	"github.com/dcrosby42/picfinder/scan"
	"github.com/urfave/cli"
)

func sandbox_scan_command() cli.Command {
	return cli.Command{
		Name:  "scan",
		Usage: "scan dirs",
		Action: func(c *cli.Context) error {

			dirname := "/Users/crosby/Pictures"
			host := "XandersBeatBox"

			// result, err := scan.CountAll(dirname)
			// if err != nil {
			// 	return cli.NewExitError(err.Error(), -1)
			// }
			// fmt.Printf("Counted all in %q:\n", dirname)
			// fmt.Printf("Files: %d\n", result.Files)
			// fmt.Printf("Dirs: %d\n", result.Dirs)
			// fmt.Printf("Elapsed: %s\n", result.Elapsed)

			// fcount := 0
			// started := time.Now()
			// err := scan.WalkFiles(dirname, func(dname string, info os.FileInfo) error {
			// 	fcount += 1
			// 	filepath := dname + "/" + info.Name()
			// 	h := scan.HashStringMurmer32(filepath)
			// 	fmt.Printf("%q, murmer32=%d\n", filepath, h)
			// 	return nil
			// })
			// elapsed := time.Now().Sub(started)
			// fmt.Printf("Elapsed: %s\n", elapsed)
			// if err != nil {
			// 	return cli.NewExitError(err.Error(), -1)
			// }
			// fmt.Printf("Walked %d files\n", fcount)

			infoC := make(chan fileinfo.FileInfo, 100)
			go scan.GenerateMediaFiles(host, dirname, infoC)

			for info := range infoC {
				fmt.Printf("%#v\n", info)
			}

			return nil
		},
	}
}
