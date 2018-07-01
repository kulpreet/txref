# Confirmed Transaction Reference Codes

Porting Jonas Schnelli's reference C implementation to Go

## Usage

```
import (
    "fmt"
    
	"github.com/kulpreet/txref"
)

/* ENCODE */
//buffer should be at least strlen(hrp) + 19 + 1(0 byte)
var encoded_txref = [32]byte
const hrp = "tx"; //mainnet

// now encode for height 100 and pos 100
encoded_txref, err = txref.Encode(hrp, TXREF_MAGIC_BTC_MAINNET, 100, 100)
fmt.Printf("Encoded txref is %s\n", encoded_txref)

/* DECODE */
const pos, height, buflen int = 10, 10, len(encoded_txref)
hrpbuf [buflen]string
magic string
hrpbuf, magic, err = txref.Decode(encoded_txref, height, pos)
```

## Install

`go get -u github.com/kulpreet/txref`
