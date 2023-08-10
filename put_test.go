package rocksdb

import "testing"

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
