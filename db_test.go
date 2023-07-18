package rocksdb

import (
	"fmt"
	"testing"
)

func TestStart(t *testing.T) {
	opt := NewOptions()
	opt.SetCreateIfMissing(true)
	db := Open(opt, "./test")
	db.Put(CreateWriteOption(), "123", "456")
	fmt.Print(db.Get(CreateReadOption(), "123"))
}
