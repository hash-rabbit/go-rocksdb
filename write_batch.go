package rocksdb

/*
#include "rocksdb/c.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

type WriteBatch struct {
	wb *C.rocksdb_writebatch_t
}

func CreateWriteBatch() *WriteBatch {
	return &WriteBatch{
		wb: C.rocksdb_writebatch_create(),
	}

}

func CreateWriteBatchFrom(rep string) *WriteBatch {
	r := C.CString(rep)
	defer C.free(unsafe.Pointer(r))

	return &WriteBatch{
		wb: C.rocksdb_writebatch_create_from(r, C.size_t(len(rep))),
	}

}

func (w *WriteBatch) Destroy() {
	C.rocksdb_writebatch_destroy(w.wb)
}

func (w *WriteBatch) Clear() {
	C.rocksdb_writebatch_clear(w.wb)
}

func (w *WriteBatch) Count() int {
	return int(C.rocksdb_writebatch_count(w.wb))
}

func (w *WriteBatch) Put(key, value string) {
	k, v := C.CString(key), C.CString(value)
	defer C.free(unsafe.Pointer(k))
	defer C.free(unsafe.Pointer(v))

	C.rocksdb_writebatch_put(w.wb, k, C.size_t(len(key)), v, C.size_t(len(value)))
}

func (w *WriteBatch) PutCf(cf *ColumnFamily, key, value string) {
	k, v := C.CString(key), C.CString(value)
	defer C.free(unsafe.Pointer(k))
	defer C.free(unsafe.Pointer(v))

	C.rocksdb_writebatch_put_cf(w.wb, cf.cf, k, C.size_t(len(key)), v, C.size_t(len(value)))
}

func (w *WriteBatch) PutCfWithTs() {

}

func (w *WriteBatch) Putv() {

}

func (w *WriteBatch) PutvCf() {

}

func (w *WriteBatch) Merge() {

}

func (w *WriteBatch) MergeCf() {

}

func (w *WriteBatch) Mergev() {

}

func (w *WriteBatch) MergevCf() {

}

func (w *WriteBatch) Delete(key string) {
	k := C.CString(key)
	defer C.free(unsafe.Pointer(k))

	C.rocksdb_writebatch_delete(w.wb, k, C.size_t(len(key)))
}

func (w *WriteBatch) Singledelete() {

}

func (w *WriteBatch) DeleteCf() {

}

func (w *WriteBatch) DeleteCfWithTs() {

}

func (w *WriteBatch) SingledeleteCf() {

}

func (w *WriteBatch) SingledeleteCfWithTs() {

}

func (w *WriteBatch) Deletev() {

}

func (w *WriteBatch) DeletevCf() {

}

func (w *WriteBatch) DeleteRange() {

}

func (w *WriteBatch) DeleteRangeCf() {

}

func (w *WriteBatch) DeleteRangev() {

}

func (w *WriteBatch) DeleteRangevCf() {

}

func (w *WriteBatch) PutLogData() {

}

func (w *WriteBatch) Iterate() {

}

func (w *WriteBatch) Data() {

}

func (w *WriteBatch) SetSavePoint() {

}

func (w *WriteBatch) RollbackToSavePoint() {

}

func (w *WriteBatch) PopSavePoint() {

}

// extern ROCKSDB_LIBRARY_API void rocksdb_writebatch_destroy(
//     rocksdb_writebatch_t*);
// extern ROCKSDB_LIBRARY_API void rocksdb_writebatch_clear(rocksdb_writebatch_t*);
// extern ROCKSDB_LIBRARY_API int rocksdb_writebatch_count(rocksdb_writebatch_t*);
// extern ROCKSDB_LIBRARY_API void rocksdb_writebatch_put(rocksdb_writebatch_t*,
//                                                        const char* key,
//                                                        size_t klen,
//                                                        const char* val,
//                                                        size_t vlen);
// extern ROCKSDB_LIBRARY_API void rocksdb_writebatch_put_cf(
//     rocksdb_writebatch_t*, rocksdb_column_family_handle_t* column_family,
//     const char* key, size_t klen, const char* val, size_t vlen);
// extern ROCKSDB_LIBRARY_API void rocksdb_writebatch_put_cf_with_ts(
//     rocksdb_writebatch_t*, rocksdb_column_family_handle_t* column_family,
//     const char* key, size_t klen, const char* ts, size_t tslen, const char* val,
//     size_t vlen);
// extern ROCKSDB_LIBRARY_API void rocksdb_writebatch_putv(
//     rocksdb_writebatch_t* b, int num_keys, const char* const* keys_list,
//     const size_t* keys_list_sizes, int num_values,
//     const char* const* values_list, const size_t* values_list_sizes);
// extern ROCKSDB_LIBRARY_API void rocksdb_writebatch_putv_cf(
//     rocksdb_writebatch_t* b, rocksdb_column_family_handle_t* column_family,
//     int num_keys, const char* const* keys_list, const size_t* keys_list_sizes,
//     int num_values, const char* const* values_list,
//     const size_t* values_list_sizes);
// extern ROCKSDB_LIBRARY_API void rocksdb_writebatch_merge(rocksdb_writebatch_t*,
//                                                          const char* key,
//                                                          size_t klen,
//                                                          const char* val,
//                                                          size_t vlen);
// extern ROCKSDB_LIBRARY_API void rocksdb_writebatch_merge_cf(
//     rocksdb_writebatch_t*, rocksdb_column_family_handle_t* column_family,
//     const char* key, size_t klen, const char* val, size_t vlen);
// extern ROCKSDB_LIBRARY_API void rocksdb_writebatch_mergev(
//     rocksdb_writebatch_t* b, int num_keys, const char* const* keys_list,
//     const size_t* keys_list_sizes, int num_values,
//     const char* const* values_list, const size_t* values_list_sizes);
// extern ROCKSDB_LIBRARY_API void rocksdb_writebatch_mergev_cf(
//     rocksdb_writebatch_t* b, rocksdb_column_family_handle_t* column_family,
//     int num_keys, const char* const* keys_list, const size_t* keys_list_sizes,
//     int num_values, const char* const* values_list,
//     const size_t* values_list_sizes);
// extern ROCKSDB_LIBRARY_API void rocksdb_writebatch_delete(rocksdb_writebatch_t*,
//                                                           const char* key,
//                                                           size_t klen);
// extern ROCKSDB_LIBRARY_API void rocksdb_writebatch_singledelete(
//     rocksdb_writebatch_t* b, const char* key, size_t klen);
// extern ROCKSDB_LIBRARY_API void rocksdb_writebatch_delete_cf(
//     rocksdb_writebatch_t*, rocksdb_column_family_handle_t* column_family,
//     const char* key, size_t klen);
// extern ROCKSDB_LIBRARY_API void rocksdb_writebatch_delete_cf_with_ts(
//     rocksdb_writebatch_t*, rocksdb_column_family_handle_t* column_family,
//     const char* key, size_t klen, const char* ts, size_t tslen);
// extern ROCKSDB_LIBRARY_API void rocksdb_writebatch_singledelete_cf(
//     rocksdb_writebatch_t* b, rocksdb_column_family_handle_t* column_family,
//     const char* key, size_t klen);
// extern ROCKSDB_LIBRARY_API void rocksdb_writebatch_singledelete_cf_with_ts(
//     rocksdb_writebatch_t* b, rocksdb_column_family_handle_t* column_family,
//     const char* key, size_t klen, const char* ts, size_t tslen);
// extern ROCKSDB_LIBRARY_API void rocksdb_writebatch_deletev(
//     rocksdb_writebatch_t* b, int num_keys, const char* const* keys_list,
//     const size_t* keys_list_sizes);
// extern ROCKSDB_LIBRARY_API void rocksdb_writebatch_deletev_cf(
//     rocksdb_writebatch_t* b, rocksdb_column_family_handle_t* column_family,
//     int num_keys, const char* const* keys_list, const size_t* keys_list_sizes);
// extern ROCKSDB_LIBRARY_API void rocksdb_writebatch_delete_range(
//     rocksdb_writebatch_t* b, const char* start_key, size_t start_key_len,
//     const char* end_key, size_t end_key_len);
// extern ROCKSDB_LIBRARY_API void rocksdb_writebatch_delete_range_cf(
//     rocksdb_writebatch_t* b, rocksdb_column_family_handle_t* column_family,
//     const char* start_key, size_t start_key_len, const char* end_key,
//     size_t end_key_len);
// extern ROCKSDB_LIBRARY_API void rocksdb_writebatch_delete_rangev(
//     rocksdb_writebatch_t* b, int num_keys, const char* const* start_keys_list,
//     const size_t* start_keys_list_sizes, const char* const* end_keys_list,
//     const size_t* end_keys_list_sizes);
// extern ROCKSDB_LIBRARY_API void rocksdb_writebatch_delete_rangev_cf(
//     rocksdb_writebatch_t* b, rocksdb_column_family_handle_t* column_family,
//     int num_keys, const char* const* start_keys_list,
//     const size_t* start_keys_list_sizes, const char* const* end_keys_list,
//     const size_t* end_keys_list_sizes);
// extern ROCKSDB_LIBRARY_API void rocksdb_writebatch_put_log_data(
//     rocksdb_writebatch_t*, const char* blob, size_t len);
// extern ROCKSDB_LIBRARY_API void rocksdb_writebatch_iterate(
//     rocksdb_writebatch_t*, void* state,
//     void (*put)(void*, const char* k, size_t klen, const char* v, size_t vlen),
//     void (*deleted)(void*, const char* k, size_t klen));
// extern ROCKSDB_LIBRARY_API const char* rocksdb_writebatch_data(
//     rocksdb_writebatch_t*, size_t* size);
// extern ROCKSDB_LIBRARY_API void rocksdb_writebatch_set_save_point(
//     rocksdb_writebatch_t*);
// extern ROCKSDB_LIBRARY_API void rocksdb_writebatch_rollback_to_save_point(
//     rocksdb_writebatch_t*, char** errptr);
// extern ROCKSDB_LIBRARY_API void rocksdb_writebatch_pop_save_point(
//     rocksdb_writebatch_t*, char** errptr);
