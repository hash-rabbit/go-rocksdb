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

func (db *DB) Put(wo *WriteOption, key, value string) error {
	var cErr *C.char
	C.rocksdb_put(db.db, wo.opt, C.CString(key), C.size_t(len(key)), C.CString(value), C.size_t(len(value)), &cErr)
	if cErr != nil {
		defer C.free(unsafe.Pointer(cErr))
		return errors.New(C.GoString(cErr))
	}
	return nil
}
