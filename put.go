package rocksdb

/*
#include <stdlib.h>
#include "rocksdb/c.h"
*/
import "C"
import (
	"unsafe"
)

func (db *DB) Put(wo *WriteOptions, key, value string) error {
	var cErr *C.char
	k, v := C.CString(key), C.CString(value)
	defer C.free(unsafe.Pointer(k))
	defer C.free(unsafe.Pointer(v))

	C.rocksdb_put(db.db, wo.opt, k, C.size_t(len(key)), v, C.size_t(len(value)), &cErr)

	return parseCerr(cErr)
}

func (db *DB) PutCf(wo *WriteOptions, cf *ColumnFamily, key, value string) error {
	var cErr *C.char
	k, v := C.CString(key), C.CString(value)
	defer C.free(unsafe.Pointer(k))
	defer C.free(unsafe.Pointer(v))

	C.rocksdb_put_cf(db.db, wo.opt, cf.cf, k, C.size_t(len(key)), v, C.size_t(len(value)), &cErr)

	return parseCerr(cErr)
}
