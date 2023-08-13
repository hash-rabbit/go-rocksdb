package rocksdb

import "testing"

func TestGet(t *testing.T) {
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
		if val, _ := db.Get(CreateReadOptions(), v[0]); val != v[1] {
			t.Errorf("get key:%s value:%s real:%s", v[0], val, v[1])
		}
	}
	t.Log("test ok")
}

func TestKeyexist(t *testing.T) {
	op := CreateOptions()
	op.SetCreateIfMissing(true)

	db := Open(op, "./test")
	defer db.Close()

	db.Put(CreateWriteOption(), "123", "456")
	t.Log(db.KeyMayExist(CreateReadOptions(), "123", ""))

	db.Delete(CreateWriteOption(), "123")
	t.Log(db.KeyMayExist(CreateReadOptions(), "123", ""))
}

func TestGetTs(t *testing.T) {
	op := CreateOptions()
	op.SetCreateIfMissing(true)

	db := Open(op, "./test")
	defer db.Close()

	db.PutWithTs(CreateWriteOption(), "test_10001", "value_10000", "1")

	ro := CreateReadOptions()
	ro.SetTimestamp("1")
	value, ts, err := db.GetWithTs(ro, "test_10001")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("key:123 value:%s ts:%s", value, ts)
}
