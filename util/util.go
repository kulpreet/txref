/* Copyright (c) 2018 Kulpreet Singh
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */
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

func Encode(hrp string, magic int, height int, position int, vout int, nonStandard bool) (txref string, error error) {
	var data []int = make([]int, 13, 13)

	/* ensure we stay in boundaries */
    if (!nonStandard && (height > 0x1FFFFF || position > 0x1FFF || vout > 0x1FFF || magic > 0x1F)) ||
		(nonStandard && (height > 0x3FFFFFF || position > 0x3FFFF || vout > 0x1FFF || magic > 0x1F)) {
        return txref, errors.New("Height, position, vout and magic are not within boundaries")
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

		data[8] |= ((vout & 0x1F))
		data[9] |= ((vout & 0x3E0) >> 5)
		data[10] |= ((vout & 0x1C00) >> 10)
		
    } else {
		// use extended blockheight (up to 0x3FFFFFF)
		// use extended txpos (up to 0x3FFFF)
		data[5] |= ((height & 0xF80000) >> 19)
		data[6] |= ((height & 0x3000000) >> 24)
		
		data[6] |= ((position & 0x7) << 2)
		data[7] |= ((position & 0xF8) >> 3)
		data[8] |= ((position & 0x1F00) >> 8)
		data[9] |= ((position & 0x3E000) >> 13)

		data[10] |= ((vout & 0x1F))
		data[11] |= ((vout & 0x3E0) >> 5)
		data[12] |= ((vout & 0x1C00) >> 10)
    }
	
	if len(hrp) == 0 {
		hrp = Txref_bech32_hrp_mainnet
	}

	if !nonStandard {
		data = data[0:11]
	}

	encoded, err := bech32.Encode(hrp, data)
	if err != nil {
		return "", err
	}

	txref = insertSeparators(encoded, hrp, nonStandard)
	return txref, nil
}

func Decode(txref string) (hrp string, magic int, height int, position int, vout int, err error) {

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
	if len(data) != 11 && len(data) != 13 {
		err = errors.New("Decoded data is not 8 or 10 characters long")
		return
	}

	magic = int(data[0])

    /* set the block height */
    height = int(data[1]) >> 1;
    height |= int(data[2]) << 4;
    height |= int(data[3]) << 9;
    height |= int(data[4]) << 14;
    if (len(data) == 11) {
		height |= int((data[5] & 0x03)) << 19;

		/* set the tx position */
		position = int((data[5] & 0x1C)) >> 2;
		position |= int(data[6]) << 3;
		position |= int(data[7]) << 8;

		vout = int(data[8])
		vout |= int(data[9] << 5)
		vout |= int(data[10] << 10)
		
    } else {
		/* use extended blockheight / txpos (test networks) */
		height |= int(data[5]) << 19;
		height |= int((data[6] & 0x03)) << 24;

		/* set the tx position */
		position = int((data[6] & 0x1C)) >> 2;
		position |= int(data[7]) << 3;
		position |= int(data[8]) << 8;
		position |= int(data[9]) << 13;

		vout = int(data[10])
		vout |= int(data[11] << 5)
		vout |= int(data[12] << 10)
    }

	return hrp, magic, height, position, vout, nil
}

func insertSeparators(encoded string, hrp string, nonStandard bool) (txref string) {
	var hrplen = len(hrp)
	txref = encoded[0 : hrplen+1] + ":" +
		encoded[hrplen+1 : hrplen+1+4] + "-" +
		encoded[hrplen+1+4 : hrplen+1+4+4] + "-" +
		encoded[hrplen+1+4+4 : hrplen+1+4+4+4] + "-"
	if !nonStandard {
		txref += encoded[hrplen+1+4+4+4 : hrplen+1+4+4+4+5]
	} else {
		txref += encoded[hrplen+1+4+4+4 : hrplen+1+4+4+4+4] + "-"
		txref += encoded[hrplen+1+4+4+4+4 : hrplen+1+4+4+4+4+3]
	}
	return
}
