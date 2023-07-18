package rocksdb

/*
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <assert.h>

#include "rocksdb/c.h"
*/
import "C"

type DB struct {
	db *C.rocksdb_t
}

func Open(option *Options, path string) *DB {
	var cErr *C.char

	db := &DB{
		db: C.rocksdb_open(option.opt, C.CString(path), &cErr),
	}

	return db
}

func OpenWithTTl(option *Options, path string, ttl int) *DB {
	var cErr *C.char

	db := &DB{
		db: C.rocksdb_open_with_ttl(option.opt, C.CString(path), C.int(ttl), &cErr),
	}

	return db
}

func (db *DB) Close() {
	C.rocksdb_close(db.db)
}
