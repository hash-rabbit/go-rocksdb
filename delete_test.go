package rocksdb

import "testing"

func TestDelete(t *testing.T) {
	op := CreateOptions()
	op.SetCreateIfMissing(true)

	db, _ := Open(op, "./test")
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
			t.Errorf("get key:%s value:%s real:%s", v[0], val, v[1])
		}
	}

	db.Delete(CreateWriteOption(), "123")
	db.Delete(CreateWriteOption(), "abc")

	for _, v := range ts {
		val, err := db.Get(CreateReadOptions(), v[0])
		if err != nil {
			t.Errorf("get error:%s", err)
		} else {
			t.Logf("key:%s value:%s", v[0], val)
		}
	}
}
