package goyescrypt

/*
#include "yescrypt.c"
#cgo LDFLAGS: -Wl,--allow-multiple-definition
*/
import "C"

import (
    "unsafe"
)

func Hash(src []byte, dst []byte){
	C.yescrypt_hash((*C.char)(unsafe.Pointer(&src[0])), (*C.char)(unsafe.Pointer(&dst[0])))
}
