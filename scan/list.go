package scan

import (
	"io/ioutil"
	"os"
)

type DirList []os.FileInfo
type FileList []os.FileInfo

func ListFilesAndDirs(dirname string) (DirList, FileList, error) {
	infos, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, nil, err
	}
	dirs := make(DirList, 0)
	files := make(FileList, 0)
	for _, info := range infos {
		if info.IsDir() {
			dirs = append(dirs, info)
		} else {
			files = append(files, info)
		}
	}
	return dirs, files, nil
}
