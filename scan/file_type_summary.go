package scan

import (
	"fmt"
	"time"

	"github.com/dcrosby42/picfinder/fileinfo"
)

type ftSummary struct {
	Count int
}

func PrintFileTypeSummary(host string, dirname string, scanAll bool) error {
	infoC := make(chan fileinfo.FileInfo, 100)

	if scanAll {
		fmt.Printf("Scanning ALL files in %q...\n", dirname)
		go GenerateAllFiles(host, dirname, infoC)
	} else {
		fmt.Printf("Scanning media files in %q...\n", dirname)
		go GenerateMediaFiles(host, dirname, infoC)
	}

	fcount := 0
	byType := make(map[fileinfo.FileType]*ftSummary)
	started := time.Now()
	for info := range infoC {
		fcount++
		summ, ok := byType[info.Type]
		if !ok {
			summ = &ftSummary{}
			byType[info.Type] = summ
		}
		summ.Count++
	}
	elapsed := time.Now().Sub(started)
	fmt.Printf("Elapsed: %s\n", elapsed)
	fmt.Printf("Total: %d\n", fcount)
	for ftype, summ := range byType {
		fmt.Printf("%s: %d\n", ftype, summ.Count)
	}

	return nil
}
