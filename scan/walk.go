package scan

import "os"

func WalkFiles(dirname string, callback func(dirname string, info os.FileInfo) error) error {
	dirs, files, err := ListFilesAndDirs(dirname)
	if err != nil {
		return err
	}
	for _, info := range files {
		err := callback(dirname, info)
		if err != nil {
			return err
		}
	}
	for _, info := range dirs {
		err := WalkFiles(dirname+"/"+info.Name(), callback)
		if err != nil {
			return err
		}
	}
	return nil
}
