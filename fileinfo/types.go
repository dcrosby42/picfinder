package fileinfo

import (
	picfinder_grpc "github.com/dcrosby42/picfinder/grpc"
)

type FileInfo struct {
	Id                 int64    `db:"id"`
	Host               string   `db:"host"`
	Path               []byte   `db:"path"`
	PathHash           uint32   `db:"path_hash"`
	Size               int64    `db:"size"`
	ContentHash        []byte   `db:"content_hash"`
	ContentHashLower32 uint32   `db:"content_hash_lower_32"`
	Type               FileType `db:"type"`
	Kind               FileKind `db:"kind"`
	ScannedAtUnix      int64    `db:"scanned_at"`
	FileModifiedAtUnix int64    `db:"file_modified_at"`
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
