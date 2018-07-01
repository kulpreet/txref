package util

import (
	"fmt"
	"errors"
	"strings"

	"github.com/kulpreet/txref/bech32"
)

const Txref_magic_btc_mainnet int = 0x03
const Txref_bech32_hrp_mainnet string = "tx"

const Txref_magic_btc_testnet int = 0x06
const Txref_bech32_hrp_testnet string = "txtest"

const Txref_len_without_hrp = 15;

func Encode(hrp string, magic int, height int, position int, nonStandard bool) (txref string, error error) {
	var data []int = make([]int, 10, 10)

	/* ensure we stay in boundaries */
    if (!nonStandard && (height > 0x1FFFFF || position > 0x1FFF || magic > 0x1F)) ||
		(nonStandard && (height > 0x3FFFFFF || position > 0x3FFFF || magic > 0x1F)) {
        return txref, errors.New("Height, position and magic are not within boundaries")
	}
	
	// set magic
	data[0] = magic

	// be explicit about setting version to 0
	data[1] &^= (1 << 0)

	data[1] |= ((height & 0xF) << 1)
    data[2] |= ((height & 0x1F0) >> 4)
    data[3] |= ((height & 0x3E00) >> 9)
    data[4] |= ((height & 0x7C000) >> 14)
    if !nonStandard {
      data[5] |= ((height & 0x180000) >> 19)
      data[5] |= ((position & 0x7) << 2)
      data[6] |= ((position & 0xF8) >> 3)
      data[7] |= ((position & 0x1F00) >> 8)
    } else {
      // use extended blockheight (up to 0x3FFFFFF)
      // use extended txpos (up to 0x3FFFF)
      data[5] |= ((height & 0xF80000) >> 19)
      data[6] |= ((height & 0x3000000) >> 24)

      data[6] |= ((position & 0x7) << 2)
      data[7] |= ((position & 0xF8) >> 3)
      data[8] |= ((position & 0x1F00) >> 8)
      data[9] |= ((position & 0x3E000) >> 13)
    }
		
	if len(hrp) == 0 {
		hrp = Txref_bech32_hrp_mainnet
	}

	if !nonStandard {
		data = data[0:8]
	}

	encoded, err := bech32.Encode(hrp, data)
	if err != nil {
		return "", err
	}

	txref = insertSeparators(encoded, hrp, nonStandard)
	return txref, nil
}

func Decode(txref string) (hrp string, magic int, height int, position int, err error) {

    /* max TXREF_LEN_WITHOUT_HRP (+4 separators) chars are allowed for now */
    if (len(txref) < Txref_len_without_hrp+4) {
		var msg string = fmt.Sprintf("max Txref_len_without_hrp (+4 separators) chars are allowed for now, %s %d",
			txref, len(txref))
		err = errors.New(msg)
        return
    }

	txref = strings.Join(strings.Split(txref, ":"), "")
	txref = strings.Join(strings.Split(txref, "-"), "")

	hrp, data, err := bech32.Decode(txref)
	if err != nil {
		return
	}	
	if len(data) != 8 && len(data) != 10 {
		err = errors.New("Decoded data is not 8 or 10 characters long")
		return
	}

	magic = int(data[0])

    /* set the block height */
    height = int(data[1]) >> 1;
    height |= int(data[2]) << 4;
    height |= int(data[3]) << 9;
    height |= int(data[4]) << 14;
    if (len(data) == 8) {
		height |= int((data[5] & 0x03)) << 19;

		/* set the tx position */
		position = int((data[5] & 0x1C)) >> 2;
		position |= int(data[6]) << 3;
		position |= int(data[7]) << 8;
    } else {
		/* use extended blockheight / txpos (test networks) */
		height |= int(data[5]) << 19;
		height |= int((data[6] & 0x03)) << 24;

		/* set the tx position */
		position = int((data[6] & 0x1C)) >> 2;
		position |= int(data[7]) << 3;
		position |= int(data[8]) << 8;
		position |= int(data[9]) << 13;
    }

	return hrp, magic, height, position, nil
}

func insertSeparators(encoded string, hrp string, nonStandard bool) (txref string) {
	var hrplen = len(hrp)
	txref = encoded[0:hrplen+1] + ":" +
		encoded[hrplen+1:hrplen+1+4] + "-" +
		encoded[hrplen+1+4:hrplen+1+4+4] + "-" +
		encoded[hrplen+1+4+4:hrplen+1+4+4+4] + "-"
	if !nonStandard {
		txref += encoded[hrplen+1+4+4+4:hrplen+1+4+4+4+2]
	} else {
		txref += encoded[hrplen+1+4+4+4:hrplen+1+4+4+4+4]
	}
	return
}
