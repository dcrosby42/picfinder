package fileinfo

import (
	"fmt"

	picfinder_grpc "github.com/dcrosby42/picfinder/grpc"
)

type FileInfo struct {
	Id                 int64    `db:"id"`
	Host               string   `db:"host"`
	Path               []byte   `db:"path"`
	PathHash           uint64   `db:"path_hash"`
	Size               int64    `db:"size"`
	ContentHash        []byte   `db:"content_hash"`
	ContentHashLower32 uint32   `db:"content_hash_lower_32"`
	Type               FileType `db:"type"`
	Kind               FileKind `db:"kind"`
	ScannedAtUnix      int64    `db:"scanned_at"`
	FileModifiedAtUnix int64    `db:"file_modified_at"`
}

func (me FileInfo) PathString() string {
	return string(me.Path) // TODO this may or may not be unicode safe?
}

func (me FileInfo) String() string {
	return fmt.Sprintf("FileInfo[host=%s path=%s kind=%s type=%s size=%d contentHashLower32=%d contentHash=%x", me.Host, me.PathString(), me.Kind, me.Type, me.Size, me.ContentHashLower32, me.ContentHash)
}

func FromGrpcFileInfo(ginfo *picfinder_grpc.FileInfo) FileInfo {
	return FileInfo{
		Id:                 ginfo.Id,
		Host:               ginfo.Host,
		Path:               ginfo.Path,
		PathHash:           ginfo.PathHash,
		Size:               ginfo.Size,
		ContentHash:        ginfo.ContentHash,
		ContentHashLower32: ginfo.ContentHashLower32,
		Type:               FileType(ginfo.Type),
		Kind:               FileKind(ginfo.Kind),
		ScannedAtUnix:      ginfo.ScannedAtUnix,
		FileModifiedAtUnix: ginfo.FileModifiedAtUnix,
	}
}

func ToGrpcFileInfo(info FileInfo) *picfinder_grpc.FileInfo {
	return &picfinder_grpc.FileInfo{
		Id:                 info.Id,
		Host:               info.Host,
		Path:               info.Path,
		PathHash:           info.PathHash,
		Size:               info.Size,
		ContentHash:        info.ContentHash,
		ContentHashLower32: info.ContentHashLower32,
		Type:               string(info.Type),
		Kind:               string(info.Kind),
		ScannedAtUnix:      info.ScannedAtUnix,
		FileModifiedAtUnix: info.FileModifiedAtUnix,
	}
}
