package scan

import (
	"fmt"
	"time"
)

type CountResult struct {
	Files   int
	Dirs    int
	Elapsed time.Duration
	Started time.Time
}

func CountAll(dirname string) (*CountResult, error) {
	res := &CountResult{}
	res.Started = time.Now()
	err := recurseCount(dirname, res)
	res.Elapsed = time.Now().Sub(res.Started)
	if err != nil {
		return res, err
	}
	return res, nil
}

func recurseCount(dirname string, res *CountResult) error {
	res.Dirs += 1
	dirs, files, err := ListFilesAndDirs(dirname)
	if err != nil {
		return fmt.Errorf("Error recursing dirname %q err=%s", dirname, err)
	}
	res.Files += len(files)

	for _, dir := range dirs {
		err := recurseCount(dirname+"/"+dir.Name(), res)
		if err != nil {
			return err
		}
	}
	return nil
}
