# Install
go get github.com/mendozg/goyescrypt
# usage
```
//main.go
package main

import (
    "fmt"
    "log"
    "encoding/hex"
    "github.com/mining-pool/not-only-mining-pool/utils"
    "github.com/mendozg/goyescrypt"
)

func main() {
    err := yescrypthash()
   if err != nil {
        log.Fatal(err)
    }
}

func yescrypthash() error {
	out := make([]byte, 32)
	headerhex := "240a5a20727df120f94999fd6e1df9a0dee583541e829597090ccaa5573b33b89f19121dbab36a503dfea48d17a160d100a78187ee80cf8ffd027bed3e82d03aa11e2d59da1cbc5ba79d011d00000a78"
	headerBytes, err := hex.DecodeString(headerhex)
	if err != nil {
		log.Fatal(err)
	}

	goyescrypt.Hash(headerBytes, out)

	result := hex.EncodeToString(utils.ReverseBytes(out))
	if result != "00000000fb7bf19efdaf25b5e31bd9caf785cf418bdcad4f5f18cddd3f4aaeef" {
		fmt.Println("%s: invalid hash", result )
	} else {
		fmt.Println("%s: Good hash", result )
	}

    return nil
}
```

