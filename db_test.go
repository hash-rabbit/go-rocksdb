package rocksdb

import "testing"

func TestOpen(t *testing.T) {
	op := CreateOptions()
	op.SetCreateIfMissing(true)

	db, err := Open(op, "./test")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer db.Close()
}
