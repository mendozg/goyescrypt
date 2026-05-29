// goyescrypt.go
package goyescrypt

/*
#include "yescrypt.c"
#cgo LDFLAGS: -Wl,--allow-multiple-definition
*/
import "C"
import "unsafe"

// Hash computes the yescrypt hash of the input and writes it to dst.
// dst must be at least 32 bytes long.
func Hash(src, dst []byte) {
	if len(dst) < 32 {
		panic("dst buffer too small, need at least 32 bytes")
	}
	C.yescrypt_hash(
		(*C.char)(unsafe.Pointer(&src[0])),
		(*C.char)(unsafe.Pointer(&dst[0])),
	)
}