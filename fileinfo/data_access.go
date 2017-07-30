package fileinfo

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func Insert(db *sqlx.DB, finfo FileInfo) (FileInfo, error) {
	res, err := db.NamedExec(`INSERT INTO file_info (host,path,path_hash,content_hash,content_hash_lower_32,size,type,kind,scanned_at,file_modified_at) 
	                        VALUES(:host,:path,:path_hash,:content_hash,:content_hash_lower_32,:size,:type,:kind,:scanned_at,:file_modified_at)`, finfo)
	if err != nil {
		return finfo, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return finfo, err
	}
	finfo.Id = lastId
	return finfo, err

}

func Get(db *sqlx.DB, id int64) (FileInfo, error) {
	var finfo FileInfo
	err := db.Get(&finfo, "SELECT * FROM file_info WHERE id=?", 2)
	return finfo, err
}

type UpdateAction string

const UpdateAction_Insert = UpdateAction("insert")
const UpdateAction_Update = UpdateAction("update")
const UpdateAction_Skip = UpdateAction("skip")

func InsertOrUpdateByHostPath(db *sqlx.DB, finfo FileInfo) (FileInfo, UpdateAction, error) {
	updateAction := UpdateAction_Insert
	var existing FileInfo
	err := db.Get(&existing, "SELECT * FROM file_info WHERE host=? AND path_hash=?", finfo.Host, finfo.PathHash)
	if err != nil {
		if err == sql.ErrNoRows {
			// there's no matching host+path in the db, so let's just insert:
			finfo, err := Insert(db, finfo)
			return finfo, updateAction, err
		}
		fmt.Printf("!!!! InsertOrUpdateByHostPath Get existing %#v\n", err)
		return finfo, updateAction, err
	}

	// There's already a record for this host+path, so let's see if it should be updated:
	if finfo.ScannedAtUnix > existing.ScannedAtUnix {
		updateAction = UpdateAction_Update
		finfo.Id = existing.Id
		err := Update(db, finfo)
		if err != nil {
			return finfo, updateAction, err
		}
		return finfo, updateAction, nil
	} else {
		updateAction = UpdateAction_Skip
		return existing, updateAction, nil
	}

}

func Update(db *sqlx.DB, finfo FileInfo) error {
	_, err := db.NamedExec(`UPDATE file_info SET host=:host, path=:path, path_hash=:path_hash,content_hash=:content_hash, content_hash_lower_32=:content_hash_lower_32, size=:size, type=:type, kind=:kind, scanned_at=:scanned_at, file_modified_at=:file_modified_at WHERE id=:id`, finfo)
	return err
}

// 	var err error
//
// 	contentDupes, err := findContentDupes(db, finfo)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	var repeater *FileInfo
// 	if len(contentDupes) > 0 {
// 		repeats := []*FileInfo{}
// 		for _, dupe := range contentDupes {
// 			if dupe.PathHash == finfo.PathHash && dupe.Host == finfo.Host && bytes.Equal(dupe.Path, finfo.Path) {
// 				repeats = append(repeats, dupe)
// 			}
// 		}
// 		if len(repeats) > 0 {
// 			repeater = repeats[0]
// 			// FILE REPEAT: this exact file+content has been added before for this host, at least once
// 			// Update the scanned_at date, if the incoming one is newer
// 			if finfo.ScannedAtUnix > repeater.ScannedAtUnix {
// 				repeater.ScannedAtUnix = finfo.ScannedAtUnix
// 				Update(repeater)
// 			}
// 			if len(repeats) > 1 {
// 				// Cruft in the db... multiple repeats. Should happen.  But...
// 				doomedIds := make([]int64, 0)
// 				for _, doomed := range repeats[1:] {
// 					doomedIds = append(doomedIds, doomed.Id)
// 				}
// 				query, args, err := sqlx.In("DELETE FROM file_info WHERE id IN (?);", doomedIds)
// 				query = db.Rebind(query)
// 				_, err := db.Exec(query, args...)
// 				if err != nil {
// 					fmt.Printf("!!!! InsertDeduped cleaning up doomed ids %v got err=%s", err)
// 				}
// 			}
// 		}
// 	}
// 	if repeater != nil {
// 		return repeater, nil // "That'll happen."  "That will happen."
// 	}
// 	return nil, nil // FIXME
//
// 	// altered,err := findWithSameHostPathButDifferentContent(db, finfo)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// if len(altered) > 0 {
// 	// }
//
// }
//
// func findContentDupes(db *sqlx.DB, finfo *FileInfo) ([]*FileInfo, error) {
// 	matched := []FileInfo{}
// 	err = db.Select(&matched, "SELECT * FROM file_info WHERE content_hash_lower_32 = ? AND size = ?", finfo.ContentHashLower32, finfo.Size)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(matched > 0) {
// 		dupes := []FileInfo{}
// 		for _, match := range matched {
// 			if match.ContentHash == finfo.ContentHash {
// 				dupes = append(dupes, match)
// 			}
// 		}
// 		return dupes, nil
// 	}
