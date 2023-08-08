package rocksdb

/*
#include "rocksdb/c.h"
*/
import "C"

type Options struct {
	opt *C.rocksdb_options_t
}

func CreateOptions() *Options {
	return &Options{
		opt: C.rocksdb_options_create(),
	}
}

func (o *Options) Destroy() {
	C.rocksdb_options_destroy(o.opt)
	o.opt = nil
}

func (o *Options) CreateCopy() *Options {
	return &Options{
		opt: C.rocksdb_options_create_copy(o.opt),
	}
}

func (o *Options) IncreaseParallelism(num int) {
	C.rocksdb_options_increase_parallelism(o.opt, C.int(num))
}

func (o *Options) OptimizeForPointLookup(block_cache_size_mb uint64) {
	C.rocksdb_options_optimize_for_point_lookup(o.opt, C.uint64_t(block_cache_size_mb))
}

func (o *Options) OptimizeLevelStyleCompaction(memtable_memory_budget uint64) {
	C.rocksdb_options_optimize_level_style_compaction(o.opt, C.uint64_t(memtable_memory_budget))
}

func (o *Options) SetCreateIfMissing(create bool) {
	C.rocksdb_options_set_create_if_missing(o.opt, boolToUchar(create))
}
