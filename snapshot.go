package rocksdb

/*
#include "rocksdb/c.h"
*/
import "C"

type Snapshot struct {
	s  *C.rocksdb_snapshot_t
	db *DB
}

func CreateSnapshot(db *DB) *Snapshot {
	return &Snapshot{
		s:  C.rocksdb_create_snapshot(db.db),
		db: db,
	}
}

func (s *Snapshot) Release() {
	C.rocksdb_release_snapshot(s.db.db, s.s)
}
