package rocksdb

import "testing"

func TestCf(t *testing.T) {
	op := CreateOptions()
	op.SetCreateIfMissing(true)

	db, err := Open(op, "./test")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer db.Close()

	cf, err := db.CreateColumnFamily(op, "test1")
	if err != nil {
		t.Fatal(err.Error())
	}

	err = db.DropColumnFamily(cf)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestListCf(t *testing.T) {
	op := CreateOptions()
	names, err := ListColumnFamilys("./test", op)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Logf("value:%+v", names)
}

func TestOpenCf(t *testing.T) {
	db, _, err := OpenColumnFamilys(CreateOptions(), "./test", []string{"default", "test1"}, []*Options{CreateOptions(), CreateOptions()})
	checkerror(err)
	db.Close()
}

func checkerror(e1 error) {
	if e1 != nil {
		panic(e1)
	}
}
