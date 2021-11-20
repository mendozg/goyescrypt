package yescrypt

/*
#include "yescrypt.c"
*/
import "C"

import (
    "errors"
    "fmt"
    "log"
    "unsafe"
    "encoding/hex"
    "github.com/mining-pool/not-only-mining-pool/utils"
)

func hash(src []byte, dst []byte){
	ret, err := C.yescrypt_hash((*C.char)(unsafe.Pointer(&src[0])), (*C.char)(unsafe.Pointer(&dst[0])))
}
