package rocksdb

/*
#include <stdlib.h>
*/
import "C"
import (
	"encoding/binary"
	"errors"
	"unsafe"
)

func boolToUchar(t bool) C.uchar {
	if t {
		return C.uchar(1)
	}
	return C.uchar(0)
}

func ucharToBool(t C.uchar) bool {
	return t != 0
}

func boolToInt(t bool) C.int {
	if t {
		return C.int(1)
	}
	return C.int(0)
}

func intToBool(t C.int) bool {
	return t != 0
}

func parseCerr(cErr *C.char) error {
	if len(C.GoString(cErr)) > 0 {
		return errors.New(C.GoString(cErr))
	}
	return nil
}

// result need to free
func StringsToCchar(strs []string) (**C.char, *C.size_t) {
	var t uint64
	var length = len(strs)
	cArray := C.malloc(C.size_t(length) * C.size_t(unsafe.Sizeof(uintptr(0))))
	bArray := C.malloc(C.size_t(length) * C.size_t(unsafe.Sizeof(t)))

	a := (*[1 << 30]*C.char)(cArray)
	b := (*[1 << 30]C.size_t)(bArray)
	for i, s := range strs {
		a[i] = C.CString(s)
		b[i] = C.size_t(len(s))
	}

	return (**C.char)(cArray), (*C.size_t)(bArray)
}

func CcharToStrings(size C.size_t, strs **C.char) []string {
	length := int(size)
	tmpslice := unsafe.Slice(strs, length)
	gostrings := make([]string, length)
	for i, s := range tmpslice {
		gostrings[i] = C.GoString(s)
	}
	return gostrings
}

func CSize(str string) C.size_t {
	return C.size_t(len(str))
}

const timestampSize = 8

func marshalTimestamp(ts uint64) []byte {
	b := make([]byte, timestampSize)
	binary.BigEndian.PutUint64(b, ts)
	return b
}

// charToByte converts a *C.char to a byte slice.
func charToByte(data *C.char, len C.size_t) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(data)), int(len))
}

// byteToChar returns *C.char from byte slice.
func byteToChar(b []byte) *C.char {
	var c *C.char
	if len(b) > 0 {
		c = (*C.char)(unsafe.Pointer(&b[0]))
	}
	return c
}
