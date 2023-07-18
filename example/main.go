package main

import (
	"fmt"

	"github.com/hash-rabbit/go-rocksdb"
)

func main() {
	opt := rocksdb.NewOptions()
	opt.SetCreateIfMissing(true)
	db := rocksdb.Open(opt, "../test")
	db.Put(rocksdb.CreateWriteOption(), "123", "789")
	value, _ := db.Get(rocksdb.CreateReadOption(), "123")
	fmt.Print(value)
}
