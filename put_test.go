package rocksdb

import (
	"strconv"
	"testing"
)

func TestPut(t *testing.T) {
	op := CreateOptions()
	op.SetCreateIfMissing(true)

	db := Open(op, "./test")
	defer db.Close()

	ts := [][]string{
		{"123", "456"},
		{"abc", "def"},
		{"qwer", "tyui"},
		{"1_3", "4_6"},
		{"z 4", "jjj"},
	}

	for _, v := range ts {
		db.Put(CreateWriteOption(), v[0], v[1])
	}

	for _, v := range ts {
		if val, _ := db.Get(CreateReadOptions(), v[0]); val != v[1] {
			t.Errorf("get value:%s real:%s", val, v[1])
		}
	}
	t.Log("test ok")
}

func BenchmarkPut(b *testing.B) {
	op := CreateOptions()
	op.SetCreateIfMissing(true)

	db := Open(op, "./test")
	defer db.Close()

	wo := CreateWriteOption()
	defer wo.Destroy()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := "key" + strconv.Itoa(i)
		value := "value_" + key
		db.Put(wo, key, value)
	}
}
