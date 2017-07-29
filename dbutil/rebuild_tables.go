package dbutil

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func RebuildTables(db *sqlx.DB) error {
	return execStatements(db, []string{
		`DROP TABLE IF EXISTS file_info`,
		`CREATE TABLE file_info (
			id bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
			host varchar(255),
			path blob,
			path_hash int(11),
			content_hash blob,
			content_hash_lower_32 int(11),
			size int(11),
			type varchar(255),
			kind varchar(255),
			scanned_at int(11),
			file_modified_at int(11)
		)`,

		// `DROP INDEX file_info_path_hash`,
		`CREATE INDEX file_info_path_hash ON file_info (path_hash)`,

		// `DROP INDEX file_info_hashlow32`,
		`CREATE INDEX file_info_hashlow32 ON file_info (content_hash_lower_32)`,
	})

}

func execStatements(db *sqlx.DB, stmts []string) error {
	for _, s := range stmts {
		_, err := db.Exec(s)
		if err != nil {
			return fmt.Errorf("FAILED to exec %q, err=%s", s, err)
		}
	}
	return nil
}
