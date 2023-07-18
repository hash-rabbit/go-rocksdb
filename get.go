package rocksdb

/*
#include <stdlib.h>
#include "rocksdb/c.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

func (db *DB) Get(ro *ReadOption, key string) (string, error) {
	var cErr *C.char
	var cValLen C.size_t
	value := C.rocksdb_get(db.db, ro.opt, C.CString(key), C.size_t(len(key)), &cValLen, &cErr)
	if cErr != nil {
		defer C.free(unsafe.Pointer(cErr))
		return "", errors.New(C.GoString(cErr))
	}
	return C.GoString(value), nil
}
