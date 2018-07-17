# Confirmed Transaction Reference Codes

[![Build Status](https://travis-ci.org/kulpreet/txref.svg?branch=master)](https://travis-ci.org/kulpreet/txref)

Porting Jonas Schnelli's reference C implementation to Go

## Usage

```golang
import (
    "fmt"
    
	"github.com/kulpreet/txref"
)

/* ENCODE */
var encoded_txref = string
const hrp string = "tx" //mainnet

// now encode for height 100 and pos 100
encoded_txref, err := txref.Encode(hrp, txref.Txref_magic_btc_mainnet, 100, 100, false)
fmt.Printf("Encoded txref is %s\n", encoded_txref)

/* DECODE */
decoded, magic, height, position, err := txref.Decode(encoded_txref)
```

### CLI Examples
```bash
./txref --encode --blockheight 1354001 --position 83 --vout 0 --magic=6 --nonstd

./txref --encode --blockheight 1354001 --position 83 --vout 0 --magic=3

./txref  --decode --txref txtest1:xz35-jzv2-qqqq-qg2u-sft
```

## Install

`go get -u github.com/kulpreet/txref`

## Run tests

`go test ./...`
