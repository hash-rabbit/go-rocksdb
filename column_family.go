package rocksdb

/*
#include "rocksdb/c.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

type ColumnFamily struct {
	cf *C.rocksdb_column_family_handle_t
}

func CreateColumnFamily(db *DB, cf_options *Options, cf_name string) *ColumnFamily {
	var cErr *C.char
	name := C.CString(cf_name)
	defer C.free(unsafe.Pointer(name))

	return &ColumnFamily{
		cf: C.rocksdb_create_column_family(db.db, cf_options.opt, name, &cErr),
	}
}

func CreateColumnFamilys(db *DB, cf_options *Options, cf_names []string) []*ColumnFamily {
	var cErr *C.char

	cfs := make([]*ColumnFamily, 0)
	for _, cf_name := range cf_names {
		name := C.CString(cf_name)
		defer C.free(unsafe.Pointer(name))

		cfs = append(cfs, &ColumnFamily{
			cf: C.rocksdb_create_column_family(db.db, cf_options.opt, name, &cErr),
		})
	}
	return cfs
}
