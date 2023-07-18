package rocksdb

import "C"

func boolToUchar(t bool) C.uchar {
	if t {
		return C.uchar(1)
	} else {
		return C.uchar(0)
	}
}
