package rocksdb

/*
#include <stdlib.h>
#include "rocksdb/c.h"
*/
import "C"
import (
	"unsafe"
)

func (db *DB) Get(ro *ReadOptions, key string) (string, error) {
	var cErr *C.char
	var cValLen C.size_t

	k := C.CString(key)
	defer C.free(unsafe.Pointer(k))

	value := C.rocksdb_get(db.db, ro.opt, k, CSize(key), &cValLen, &cErr)
	return C.GoStringN(value, C.int(cValLen)), parseCerr(cErr)
}

func (db *DB) GetWithTs(ro *ReadOptions, key string) (string, string, error) {
	var cErr *C.char
	var valen C.size_t

	k := C.CString(key)
	defer C.free(unsafe.Pointer(k))

	var ts *C.char
	var tslen C.size_t

	value := C.rocksdb_get_with_ts(db.db, ro.opt, k, CSize(key), &valen, &ts, &tslen, &cErr)
	return C.GoStringN(value, C.int(valen)), C.GoStringN(ts, C.int(tslen)), parseCerr(cErr)
}

func (db *DB) GetCf(ro *ReadOptions, cf *ColumnFamily, key string) (string, error) {
	var cErr *C.char
	var valen C.size_t

	k := C.CString(key)
	defer C.free(unsafe.Pointer(k))

	value := C.rocksdb_get_cf(db.db, ro.opt, cf.cf, k, CSize(key), &valen, &cErr)
	return C.GoStringN(value, C.int(valen)), parseCerr(cErr)
}

func (db *DB) GetCfWithTs(ro *ReadOptions, cf *ColumnFamily, key string) (string, string, error) {
	var cErr *C.char
	var valen C.size_t

	var ts *C.char
	var tslen C.size_t

	k := C.CString(key)
	defer C.free(unsafe.Pointer(k))

	value := C.rocksdb_get_cf_with_ts(db.db, ro.opt, cf.cf, k, CSize(key), &valen, &ts, &tslen, &cErr)
	return C.GoStringN(value, C.int(valen)), C.GoStringN(ts, C.int(tslen)), parseCerr(cErr)
}

func (db *DB) MultiGet(ro *ReadOptions, keys []string) ([]string, error) {
	var cErr *C.char

	ks, sizes := StringsToCchar(keys)
	defer C.free(unsafe.Pointer(ks))
	defer C.free(unsafe.Pointer(sizes))

	var values *C.char
	var valen C.size_t

	C.rocksdb_multi_get(db.db, ro.opt, C.size_t(len(keys)), ks, sizes, &values, &valen, &cErr)
	return CcharToStrings(valen, &values), parseCerr(cErr)
}

func (db *DB) MultiGetWithTs(ro *ReadOptions, keys []string) ([]string, error) {
	var cErr *C.char

	ks, sizes := StringsToCchar(keys)
	defer C.free(unsafe.Pointer(ks))
	defer C.free(unsafe.Pointer(sizes))

	var values *C.char
	var valen C.size_t

	var ts *C.char
	var tslen C.size_t

	C.rocksdb_multi_get_with_ts(db.db, ro.opt, C.size_t(len(keys)), ks, sizes, &values, &valen, &ts, &tslen, &cErr)
	return CcharToStrings(valen, &values), parseCerr(cErr)
}

func (db *DB) MultiGetCf(ro *ReadOptions, cf *ColumnFamily, keys []string) ([]string, error) {
	var cErr *C.char

	ks, sizes := StringsToCchar(keys)
	defer C.free(unsafe.Pointer(ks))
	defer C.free(unsafe.Pointer(sizes))

	var values **C.char
	var valen C.size_t

	C.rocksdb_multi_get_cf(db.db, ro.opt, &cf.cf, C.size_t(len(keys)), ks, sizes, values, &valen, &cErr)
	return CcharToStrings(valen, values), parseCerr(cErr)
}

func (db *DB) rocksdb_multi_get_cf_with_ts(ro *ReadOptions, cf *ColumnFamily, keys []string) ([]string, error) {
	var cErr *C.char

	ks, sizes := StringsToCchar(keys)
	defer C.free(unsafe.Pointer(ks))
	defer C.free(unsafe.Pointer(sizes))

	var values *C.char
	var valen C.size_t

	var ts *C.char
	var tslen C.size_t

	C.rocksdb_multi_get_cf_with_ts(db.db, ro.opt, &cf.cf, C.size_t(len(keys)), ks, sizes, &values, &valen, &ts, &tslen, &cErr)
	return CcharToStrings(valen, &values), parseCerr(cErr)
}

func (db *DB) KeyMayExist(ro *ReadOptions, key, timeStamp string) bool {
	k := C.CString(key)
	defer C.free(unsafe.Pointer(k))

	var values *C.char
	var valen C.size_t

	ts := C.CString(timeStamp)
	defer C.free(unsafe.Pointer(ts))

	valFind := boolToUchar(false)

	return ucharToBool(C.rocksdb_key_may_exist(db.db, ro.opt, k, C.size_t(len(key)), &values, &valen, ts, C.size_t(len(timeStamp)), &valFind))
}

func (db *DB) KeyMayExistCf(ro *ReadOptions, cf *ColumnFamily, key, timeStamp string) bool {
	k := C.CString(key)
	defer C.free(unsafe.Pointer(k))

	var value *C.char
	var valen C.size_t

	ts := C.CString(timeStamp)
	defer C.free(unsafe.Pointer(ts))

	valFind := boolToUchar(false)

	return ucharToBool(C.rocksdb_key_may_exist_cf(db.db, ro.opt, cf.cf, k, C.size_t(len(key)), &value, &valen, ts, C.size_t(len(timeStamp)), &valFind))
}
