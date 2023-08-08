package main

import (
	"fmt"

	"github.com/hash-rabbit/go-rocksdb"
)

func main() {
	opt := rocksdb.CreateOptions()
	opt.SetCreateIfMissing(true)
	db := rocksdb.Open(opt, "../test")
	db.Put(rocksdb.CreateWriteOption(), "123", "789")
	value, _ := db.Get(rocksdb.CreateReadOptions(), "123")
	fmt.Print(value)
}
