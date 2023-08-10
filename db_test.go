package rocksdb

import "testing"

func TestOpen(t *testing.T) {
	op := CreateOptions()
	op.SetCreateIfMissing(true)

	db := Open(op, "./test")
	defer db.Close()
}
