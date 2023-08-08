package rocksdb

/*
#include <stdlib.h>
#include "rocksdb/c.h"
*/
import "C"
import "unsafe"

func (db *DB) Delete(wo *WriteOptions, key string) error {
	var cErr *C.char
	k := C.CString(key)
	defer C.free(unsafe.Pointer(k))

	C.rocksdb_delete(db.db, wo.opt, k, C.size_t(len(key)), &cErr)

	return parseCerr(cErr)
}

func (db *DB) DeleteCf(wo *WriteOptions, cf *ColumnFamily, key string) error {
	var cErr *C.char
	k := C.CString(key)
	defer C.free(unsafe.Pointer(k))

	C.rocksdb_delete_cf(db.db, wo.opt, cf.cf, k, C.size_t(len(key)), &cErr)

	return parseCerr(cErr)
}

func (db *DB) DeleteRangeCf(wo *WriteOptions, cf *ColumnFamily, start_key, end_key string) error {
	var cErr *C.char
	sk, ek := C.CString(start_key), C.CString(end_key)
	defer C.free(unsafe.Pointer(sk))
	defer C.free(unsafe.Pointer(ek))

	C.rocksdb_delete_range_cf(db.db, wo.opt, cf.cf, sk, C.size_t(len(start_key)), ek, C.size_t(len(end_key)), &cErr)

	return parseCerr(cErr)
}
