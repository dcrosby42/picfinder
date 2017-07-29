package fileinfo

type FileInfo struct {
	Id                 uint64 `db:"id"`
	Host               string `db:"host"`
	Path               []byte `db:"path"`
	PathHash           uint32 `db:"path_hash"`
	Size               uint32 `db:"size"`
	ContentHash        []byte `db:"content_hash"`
	ContentHashLower32 uint32 `db:"content_hash_lower_32"`
	Type               string `db:"type"`
	Kind               string `db:"kind"`
	ScannedAtUnix      int64  `db:"scanned_at"`
	FileModifiedAtUnix int64  `db:"file_modified_at"`
}
