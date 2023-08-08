package rocksdb

/*
#include "rocksdb/c.h"
*/
import "C"

type WriteOptions struct {
	opt *C.rocksdb_writeoptions_t
}

func CreateWriteOption() *WriteOptions {
	return &WriteOptions{
		opt: C.rocksdb_writeoptions_create(),
	}
}

func (wo *WriteOptions) Destroy() {
	C.rocksdb_writeoptions_destroy(wo.opt)
	wo.opt = nil
}

func (wo *WriteOptions) SetSync(v bool) {
	C.rocksdb_writeoptions_set_sync(wo.opt, boolToUchar(v))
}

func (wo *WriteOptions) GetSync() bool {
	return ucharToBool(C.rocksdb_writeoptions_get_sync(wo.opt))
}

func (wo *WriteOptions) DisableWAL(v bool) {
	C.rocksdb_writeoptions_disable_WAL(wo.opt, boolToInt(v))
}

func (wo *WriteOptions) GetDisableWAL() bool {
	return ucharToBool(C.rocksdb_writeoptions_get_disable_WAL(wo.opt))
}

func (wo *WriteOptions) SetIgnoreMissingColumnFamilies(v bool) {
	C.rocksdb_writeoptions_set_ignore_missing_column_families(wo.opt, boolToUchar(v))
}

func (wo *WriteOptions) GetIgnoreMissingColumnFamilies() bool {
	return ucharToBool(C.rocksdb_writeoptions_get_ignore_missing_column_families(wo.opt))
}

func (wo *WriteOptions) SetNoSlowdown(v bool) {
	C.rocksdb_writeoptions_set_no_slowdown(wo.opt, boolToUchar(v))
}

func (wo *WriteOptions) GetNoSlowdown() bool {
	return ucharToBool(C.rocksdb_writeoptions_get_no_slowdown(wo.opt))
}

func (wo *WriteOptions) SetLowPri(v bool) {
	C.rocksdb_writeoptions_set_low_pri(wo.opt, boolToUchar(v))
}

func (wo *WriteOptions) GetLowPri() bool {
	return ucharToBool(C.rocksdb_writeoptions_get_low_pri(wo.opt))
}

func (wo *WriteOptions) SetMemtableInsertHintPerBatch(v bool) {
	C.rocksdb_writeoptions_set_memtable_insert_hint_per_batch(wo.opt, boolToUchar(v))
}

func (wo *WriteOptions) GetMemtableInsertHintPerBatch() bool {
	return ucharToBool(C.rocksdb_writeoptions_get_memtable_insert_hint_per_batch(wo.opt))
}
