package rocksdb

// TODO fix error and vaild use

/*
#include <stdlib.h>
#include "rocksdb/c.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

type Iterator struct {
	it *C.rocksdb_iterator_t
}

func CreateIterator(db *DB, ro *ReadOptions) *Iterator {
	return &Iterator{
		it: C.rocksdb_create_iterator(db.db, ro.opt),
	}
}

func CreateIteratorCf(db *DB, ro *ReadOptions, cf *ColumnFamily) *Iterator {
	return &Iterator{
		it: C.rocksdb_create_iterator_cf(db.db, ro.opt, cf.cf),
	}
}

// not use C.rocksdb_create_iterators
func CreateIterators(db *DB, ro *ReadOptions, cfs []*ColumnFamily) []*Iterator {
	its := make([]*Iterator, 0)
	for _, cf := range cfs {
		its = append(its, &Iterator{
			it: C.rocksdb_create_iterator_cf(db.db, ro.opt, cf.cf),
		})
	}
	return its
}

func (i *Iterator) Destory() {
	C.rocksdb_iter_destroy(i.it)
	i.it = nil
}

func (i *Iterator) Valid() bool {
	return ucharToBool(C.rocksdb_iter_valid(i.it))
}

func (i *Iterator) GetErr() error {
	var cErr *C.char
	C.rocksdb_iter_get_error(i.it, &cErr)
	if cErr != nil {
		return errors.New(C.GoString(cErr))
	}
	return nil
}

func (i *Iterator) SeekToFirst() {
	C.rocksdb_iter_seek_to_first(i.it)
}

func (i *Iterator) SeekToLast() {
	C.rocksdb_iter_seek_to_last(i.it)
}

func (i *Iterator) Seek(key string) {
	k := C.CString(key)
	defer C.free(unsafe.Pointer(k))

	C.rocksdb_iter_seek(i.it, k, CSize(key))
}

func (i *Iterator) SeekForPrev(key string) {
	k := C.CString(key)
	defer C.free(unsafe.Pointer(k))

	C.rocksdb_iter_seek_for_prev(i.it, k, CSize(key))
}

func (i *Iterator) Next() {
	C.rocksdb_iter_next(i.it)
}

func (i *Iterator) Prev() {
	C.rocksdb_iter_prev(i.it)
}

func (i *Iterator) Get() (string, string, string) {
	var len C.size_t
	key := C.GoString(C.rocksdb_iter_key(i.it, &len))
	value := C.GoString(C.rocksdb_iter_value(i.it, &len))
	ts := C.GoString(C.rocksdb_iter_timestamp(i.it, &len))
	return key, value, ts
}

// extern ROCKSDB_LIBRARY_API void rocksdb_iter_get_error(
//     const rocksdb_iterator_t*, char** errptr);
