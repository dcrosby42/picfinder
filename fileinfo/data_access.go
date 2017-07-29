package fileinfo

import (
	"github.com/jmoiron/sqlx"
)

func Insert(db *sqlx.DB, finfo *FileInfo) (*FileInfo, error) {
	res, err := db.NamedExec(`INSERT INTO file_info (host,path,path_hash,content_hash,content_hash_lower_32,size,type,kind,scanned_at,file_modified_at) 
	                        VALUES(:host,:path,:path_hash,:content_hash,:content_hash_lower_32,:size,:type,:kind,:scanned_at,:file_modified_at)`, finfo)
	if err != nil {
		return nil, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	finfo.Id = lastId
	return finfo, err

}

func Get(db *sqlx.DB, id int64) (*FileInfo, error) {
	var finfo FileInfo
	err := db.Get(&finfo, "SELECT * FROM file_info WHERE id=?", 2)
	return &finfo, err
}
