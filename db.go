package rocksdb

/*
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <assert.h>

#include "rocksdb/c.h"
*/
import "C"
import "unsafe"

type DB struct {
	db *C.rocksdb_t
}

// Open open a rocksdb database
// if err: Invalid argument: Column families not opened:xxx please use:OpenColumnFamilys
func Open(option *Options, path string) (*DB, error) {
	var cErr *C.char

	name := C.CString(path)
	defer C.free(unsafe.Pointer(name))

	db := C.rocksdb_open(option.opt, name, &cErr)

	return &DB{
		db: db,
	}, parseCerr(cErr)
}

func OpenWithTTl(option *Options, path string, ttl int) *DB {
	var cErr *C.char
	name := C.CString(path)
	defer C.free(unsafe.Pointer(name))

	db := &DB{
		db: C.rocksdb_open_with_ttl(option.opt, name, C.int(ttl), &cErr),
	}

	return db
}

func OpenForReadOnly(option *Options, path string, errIfWALExist bool) *DB {
	var cErr *C.char
	name := C.CString(path)
	defer C.free(unsafe.Pointer(name))

	return &DB{
		db: C.rocksdb_open_for_read_only(option.opt, name, boolToUchar(errIfWALExist), &cErr),
	}
}

func OpenAsSecondary(option *Options, path, secondPath string) *DB {
	var cErr *C.char
	name, sname := C.CString(path), C.CString(secondPath)
	defer C.free(unsafe.Pointer(name))
	defer C.free(unsafe.Pointer(sname))

	return &DB{
		db: C.rocksdb_open_as_secondary(option.opt, name, sname, &cErr),
	}
}

func (db *DB) Close() {
	C.rocksdb_close(db.db)
}
