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

func (db *DB) CreateColumnFamily(cf_options *Options, cf_name string) (*ColumnFamily, error) {
	var cErr *C.char

	name := C.CString(cf_name)
	defer C.free(unsafe.Pointer(name))

	return &ColumnFamily{
		cf: C.rocksdb_create_column_family(db.db, cf_options.opt, name, &cErr),
	}, parseCerr(cErr)
}

func (db *DB) CreateColumnFamilys(cf_options *Options, cf_names []string) []*ColumnFamily {
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

func OpenColumnFamilys(opt *Options, path string, cf_names []string, cf_opts []*Options) (*DB, []*ColumnFamily, error) {
	var cErr *C.char

	name := C.CString(path)
	defer C.free(unsafe.Pointer(name))

	cfNames := make([]*C.char, 0)
	for _, v := range cf_names {
		vname := C.CString(v)
		defer C.free(unsafe.Pointer(vname))
		cfNames = append(cfNames, vname)
	}

	cfOps := make([]*C.rocksdb_options_t, 0)
	for _, v := range cf_opts {
		cfOps = append(cfOps, v.opt)
	}

	cfhs := make([]*C.rocksdb_column_family_handle_t, len(cf_names))

	db := C.rocksdb_open_column_families(opt.opt, name, C.int(len(cf_names)), &cfNames[0], &cfOps[0], &cfhs[0], &cErr)

	cfs := make([]*ColumnFamily, 0)
	for _, v := range cfhs {
		cfs = append(cfs, &ColumnFamily{
			cf: v,
		})
	}

	return &DB{db: db}, cfs, parseCerr(cErr)
}

func ListColumnFamilys(dbpath string, opt *Options) ([]string, error) {
	var cErr *C.char

	var valen C.size_t

	dn := C.CString(dbpath)
	defer C.free(unsafe.Pointer(dn))

	names := C.rocksdb_list_column_families(opt.opt, dn, &valen, &cErr)
	defer C.rocksdb_list_column_families_destroy(names, valen)

	return CcharToStrings(valen, names), parseCerr(cErr)
}

func (db *DB) DropColumnFamily(cf *ColumnFamily) error {
	var cErr *C.char
	C.rocksdb_drop_column_family(db.db, cf.cf, &cErr)
	return parseCerr(cErr)
}
