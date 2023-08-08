package rocksdb

/*
#include "rocksdb/c.h"
*/
import "C"

func (db *DB) Write(wo *WriteOptions, wb *WriteBatch) error {
	var cErr *C.char
	C.rocksdb_write(db.db, wo.opt, wb.wb, &cErr)
	return parseCerr(cErr)
}
