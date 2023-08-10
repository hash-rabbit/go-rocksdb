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

type ComparessionType int

const (
	rocksdb_no_compression ComparessionType = iota
	rocksdb_snappy_compression
	rocksdb_zlib_compression
	rocksdb_bz2_compression
	rocksdb_lz4_compression
	rocksdb_lz4hc_compression
	rocksdb_xpress_compression
	rocksdb_zstd_compression
)

func (o *Options) SetCompression(comtype ComparessionType) {
	C.rocksdb_options_set_compression(o.opt, C.int(comtype))
}

func (o *Options) GetCompression() ComparessionType {
	return ComparessionType(C.rocksdb_options_get_compression(o.opt))
}

func (o *Options) SetBottommostCompression(comtype ComparessionType) {
	C.rocksdb_options_set_bottommost_compression(o.opt, C.int(comtype))
}

func (o *Options) GetBottommostCompression() ComparessionType {
	return ComparessionType(C.rocksdb_options_get_bottommost_compression(o.opt))
}
