package scan

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/dcrosby42/picfinder/fileinfo"
)

func GeneratePictureFiles(host string, dirname string, infoChan chan fileinfo.FileInfo) {
	GenerateFilteredFiles(host, dirname, IsPicture, infoChan)
}
func GenerateMediaFiles(host string, dirname string, infoChan chan fileinfo.FileInfo) {
	GenerateFilteredFiles(host, dirname, IsMedia, infoChan)
}
func GenerateAllFiles(host string, dirname string, infoChan chan fileinfo.FileInfo) {
	GenerateFilteredFiles(host, dirname, AllowAny, infoChan)
}

func GenerateFilteredFiles(host string, dirname string, shouldKeep func(string, os.FileInfo) bool, infoChan chan fileinfo.FileInfo) {
	WalkFiles(dirname, func(dname string, info os.FileInfo) error {
		ftype, fkind := InferFileTypeAndKind(info.Name())
		if shouldKeep(dname, info) {
			fpath := dname + "/" + info.Name()
			finfo := fileinfo.FileInfo{
				Host:               host,
				Path:               []byte(fpath),
				PathHash:           HashStringMurmer32(fpath),
				Size:               info.Size(),
				Type:               ftype,
				Kind:               fkind,
				ScannedAtUnix:      time.Now().Unix(),
				FileModifiedAtUnix: info.ModTime().Unix(),
			}
			infoChan <- finfo
		}
		return nil
	})
	close(infoChan)
}

func LowercaseFileExt(fname string) string {
	ext := filepath.Ext(fname)
	if ext == "" {
		return ""
	}
	return strings.ToLower(ext[1:])
}

func InferFileKind(fname string) fileinfo.FileKind {
	return fileinfo.FileKindForExt(LowercaseFileExt(fname))
}

func InferFileTypeAndKind(fname string) (fileinfo.FileType, fileinfo.FileKind) {
	ext := LowercaseFileExt(fname)
	return fileinfo.FileTypeForExt(ext), fileinfo.FileKindForExt(ext)
}

func IsPicture(dirname string, info os.FileInfo) bool {
	fkind := InferFileKind(info.Name())
	return fkind == fileinfo.PictureKind
}
func IsPictureOrMovie(dirname string, info os.FileInfo) bool {
	fkind := InferFileKind(info.Name())
	return fkind == fileinfo.PictureKind || fkind == fileinfo.MovieKind
}
func IsMedia(dirname string, info os.FileInfo) bool {
	fkind := InferFileKind(info.Name())
	return fkind == fileinfo.PictureKind || fkind == fileinfo.MovieKind || fkind == fileinfo.SoundKind
}

func AllowAny(dirname string, info os.FileInfo) bool {
	return true
}
