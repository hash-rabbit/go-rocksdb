package rocksdb

/*
#include "rocksdb/c.h"
*/
import "C"

type WriteOption struct {
	opt *C.rocksdb_writeoptions_t
}

func CreateWriteOption() *WriteOption {
	return &WriteOption{
		opt: C.rocksdb_writeoptions_create(),
	}
}
