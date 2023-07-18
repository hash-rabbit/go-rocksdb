package rocksdb

/*
#include "rocksdb/c.h"
*/
import "C"

type Options struct {
	opt *C.rocksdb_options_t
}

func NewOptions() *Options {
	return &Options{
		opt: C.rocksdb_options_create(),
	}
}

func (o *Options) SetCreateIfMissing(create bool) *Options {
	C.rocksdb_options_set_create_if_missing(o.opt, boolToUchar(create))
	return o
}

func (o *Options) Clear() {
	C.rocksdb_options_destroy(o.opt)
}

func (o *Options) Copy() *Options {
	return &Options{
		opt: C.rocksdb_options_create_copy(o.opt),
	}
}
