package rocksdb

/*
#include "rocksdb/c.h"
*/
import "C"

type ReadOption struct {
	opt *C.rocksdb_readoptions_t
}

func CreateReadOption() *ReadOption {
	return &ReadOption{
		opt: C.rocksdb_readoptions_create(),
	}
}
