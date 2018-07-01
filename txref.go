package main

import (
	"fmt"
	"os"

	"github.com/kulpreet/txref/util"
	"github.com/jessevdk/go-flags"
)

// Flags.
var opts = struct {
	Encode  bool   `short:"e" long:"encode" description:"Encode a txid/vout into a txref"`
	Decode  bool   `short:"d" long:"decode" description:"Decode a txref into txid/vout"`
	Height  int   `short:"b" long:"blockheight" description:"Block height to use for encoding"`
	Position  int   `short:"p" long:"position" description:"Position to use for encoding"`
	Txref  string  `short:"r" long:"txref" description:"Txref to decode"`
	Magic  int  `short:"m" long:"magic" description:"Magic to use in the HRP"`
	NonStandard bool  `short:"n" long:"nonstd" description:"Do we use the non standard ranges"`
}{
	Encode:  false,
	Decode: false,
	Height: -1,
	Position: -1,
	Txref: "",
}

func init() {
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}
}

func main() {

	const hrp = "tx"

	if !opts.Encode && !opts.Decode {
		fmt.Printf("Please choose either encode or decode\n")
		os.Exit(1)
	} else if opts.Encode {
		if opts.Height == -1 || opts.Position == -1 {		
			fmt.Printf("Please provide height and position to encode\n")
			os.Exit(1)
		} else {
			fmt.Printf("Calling Encode with height %d and position %d\n", opts.Height, opts.Position)
			encodedTxref, err := util.Encode(hrp, opts.Magic, opts.Height, opts.Position, opts.NonStandard)
			if err != nil {
				fmt.Printf("err: %s\n", err)
				return
			}
			fmt.Printf("encodedTxref: %s\n", encodedTxref)
		}
	} else {
		if opts.Txref == "" { 
			fmt.Printf("Please provide a txref to decode\n")
			os.Exit(1)
		} else {
			fmt.Printf("Calling Decode with txref %s\n", opts.Txref)
			decodedTxref, magic, height, position, err := util.Decode(opts.Txref)
			if err != nil {
				fmt.Printf("err: %s\n", err)
				return
			}
			fmt.Printf("decodedTxref: %s, magic: %d, height: %d, position: %d\n",
				decodedTxref, magic, height, position)
		}
	}	
}
