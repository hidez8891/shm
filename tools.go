package shm

import (
	"reflect"
	"unsafe"
)

func copySlice2Ptr(b []byte, p uintptr, off int64, size int32) int {
	h := reflect.SliceHeader{}
	h.Cap = int(size)
	h.Len = int(size)
	h.Data = p

	bb := *(*[]byte)(unsafe.Pointer(&h))
	return copy(bb[off:], b)
}

func copyPtr2Slice(p uintptr, b []byte, off int64, size int32) int {
	h := reflect.SliceHeader{}
	h.Cap = int(size)
	h.Len = int(size)
	h.Data = p

	bb := *(*[]byte)(unsafe.Pointer(&h))
	return copy(b, bb[off:size])
}
