package rocksdb

/*
#include "rocksdb/c.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

type ReadOptions struct {
	opt *C.rocksdb_readoptions_t
}

// use Destory() to release
func CreateReadOptions() *ReadOptions {
	return &ReadOptions{
		opt: C.rocksdb_readoptions_create(),
	}
}

func (ro *ReadOptions) Destroy() {
	C.rocksdb_readoptions_destroy(ro.opt)
	ro.opt = nil
}

func (ro *ReadOptions) SetVerifyChecksums(v bool) {
	C.rocksdb_readoptions_set_verify_checksums(ro.opt, boolToUchar(v))
}

func (ro *ReadOptions) GetVerifyChecksums() bool {
	return ucharToBool(C.rocksdb_readoptions_get_verify_checksums(ro.opt))
}

func (ro *ReadOptions) SetFillCache(v bool) {
	C.rocksdb_readoptions_set_fill_cache(ro.opt, boolToUchar(v))
}

func (ro *ReadOptions) GetFillCache() bool {
	return ucharToBool(C.rocksdb_readoptions_get_fill_cache(ro.opt))
}

func (ro *ReadOptions) SetSnapshot(s *Snapshot) {
	C.rocksdb_readoptions_set_snapshot(ro.opt, s.s)
}

func (ro *ReadOptions) SetIterateUpperBound(key string) {
	k := C.CString(key)
	defer C.free(unsafe.Pointer(k))

	C.rocksdb_readoptions_set_iterate_upper_bound(ro.opt, k, C.size_t(len(key)))
}

func (ro *ReadOptions) SetIterateLowerBound(key string) {
	k := C.CString(key)
	defer C.free(unsafe.Pointer(k))

	C.rocksdb_readoptions_set_iterate_lower_bound(ro.opt, k, C.size_t(len(key)))
}

type TierType int

const (
	kReadAllTier    TierType = 0x0 // data in memtable, block cache, OS cache or storage
	kBlockCacheTier TierType = 0x1 // data in memtable or block cache
	kPersistedTier  TierType = 0x2 // persisted data.  When WAL is disabled, this option will skip data in memtable. Note thatthis ReadTier currently only supports Get and MultiGet and does not support iterators.
	kMemtableTier   TierType = 0x3 // data in memtable. used for memtable-only iterators.
)

func (ro *ReadOptions) SetReadTier(v TierType) {
	C.rocksdb_readoptions_set_read_tier(ro.opt, C.int(v))
}

func (ro *ReadOptions) GetReadTier() TierType {
	return TierType(C.rocksdb_readoptions_get_read_tier(ro.opt))
}

func (ro *ReadOptions) SetTailing(v bool) {
	C.rocksdb_readoptions_set_tailing(ro.opt, boolToUchar(v))
}

func (ro *ReadOptions) GetTailing() bool {
	return ucharToBool(C.rocksdb_readoptions_get_tailing(ro.opt))
}

// Deprecated: The functionality that this option controlled has been removed.
func (ro *ReadOptions) SetManaged(v bool) {
	C.rocksdb_readoptions_set_managed(ro.opt, boolToUchar(v))
}

func (ro *ReadOptions) SetReadaheadSize(size uint) {
	C.rocksdb_readoptions_set_readahead_size(ro.opt, C.size_t(size))
}

func (ro *ReadOptions) GetReadaheadSize() uint {
	return uint(C.rocksdb_readoptions_get_readahead_size(ro.opt))
}

func (ro *ReadOptions) SetPrefixSameAsStart(v bool) {
	C.rocksdb_readoptions_set_prefix_same_as_start(ro.opt, boolToUchar(v))
}

func (ro *ReadOptions) GetPrefixSameAsStart() bool {
	return ucharToBool(C.rocksdb_readoptions_get_prefix_same_as_start(ro.opt))
}

func (ro *ReadOptions) SetPinData(v bool) {
	C.rocksdb_readoptions_set_pin_data(ro.opt, boolToUchar(v))
}

func (ro *ReadOptions) GetPinData() bool {
	return ucharToBool(C.rocksdb_readoptions_get_pin_data(ro.opt))
}

func (ro *ReadOptions) SetTotalOrderSeek(v bool) {
	C.rocksdb_readoptions_set_total_order_seek(ro.opt, boolToUchar(v))
}

func (ro *ReadOptions) GetTotalOrderSeek() bool {
	return ucharToBool(C.rocksdb_readoptions_get_total_order_seek(ro.opt))
}

func (ro *ReadOptions) SetMaxSkippableInternalKeys(v uint64) {
	C.rocksdb_readoptions_set_max_skippable_internal_keys(ro.opt, C.uint64_t(v))
}

func (ro *ReadOptions) GetMaxSkippableInternalKeys() uint64 {
	return uint64(C.rocksdb_readoptions_get_max_skippable_internal_keys(ro.opt))
}

func (ro *ReadOptions) SetBackgroundPurgeOnIteratorCleanup(v bool) {
	C.rocksdb_readoptions_set_background_purge_on_iterator_cleanup(ro.opt, boolToUchar(v))
}

func (ro *ReadOptions) GetBackgroundPurgeOnIteratorCleanup() bool {
	return ucharToBool(C.rocksdb_readoptions_get_background_purge_on_iterator_cleanup(ro.opt))
}

func (ro *ReadOptions) SetIgnoreRangeDeletions(v bool) {
	C.rocksdb_readoptions_set_ignore_range_deletions(ro.opt, boolToUchar(v))
}

func (ro *ReadOptions) GetIgnoreRangeDeletions() bool {
	return ucharToBool(C.rocksdb_readoptions_get_ignore_range_deletions(ro.opt))
}

func (ro *ReadOptions) SetDeadline(v uint64) {
	C.rocksdb_readoptions_set_deadline(ro.opt, C.uint64_t(v))
}

func (ro *ReadOptions) GetDeadline() uint64 {
	return uint64(C.rocksdb_readoptions_get_deadline(ro.opt))
}

func (ro *ReadOptions) SetIoTimeout(v uint64) {
	C.rocksdb_readoptions_set_io_timeout(ro.opt, C.uint64_t(v))
}

func (ro *ReadOptions) GetIoTimeout() uint64 {
	return uint64(C.rocksdb_readoptions_get_io_timeout(ro.opt))
}

func (ro *ReadOptions) SetAsyncIo(v bool) {
	C.rocksdb_readoptions_set_async_io(ro.opt, boolToUchar(v))
}

func (ro *ReadOptions) GetAsyncIo() bool {
	return ucharToBool(C.rocksdb_readoptions_get_async_io(ro.opt))
}

func (ro *ReadOptions) SetTimestamp(ts string) {
	t := C.CString(ts)
	defer C.free(unsafe.Pointer(t))

	C.rocksdb_readoptions_set_timestamp(ro.opt, t, C.size_t(len(ts)))
}

func (ro *ReadOptions) SetIterStartTs(ts string) {
	t := C.CString(ts)
	defer C.free(unsafe.Pointer(t))

	C.rocksdb_readoptions_set_iter_start_ts(ro.opt, t, C.size_t(len(ts)))
}
