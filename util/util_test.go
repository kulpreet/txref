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
package util_test

import (
	"testing"

	. "github.com/kulpreet/txref/util"
)

var testVectors = []struct {
	magic int
    hrp string
    encodedTxref string
    height int
    position int
    encFail int //0 == must not fail, 1 == can fail, 2 == can fail and continue with next test, 3 == skip
    decFail int
    nonStd bool
}{
	{
        Txref_magic_btc_mainnet,
        Txref_bech32_hrp_mainnet,
        "tx1:rqqq-qqqq-qmhu-qk",
        0,
        0,
        0,0,false,
    },
	{
        Txref_magic_btc_mainnet,
        Txref_bech32_hrp_mainnet,
        "tx1:rjk0-u5ng-4jsf-mc",
        466793,
        2205,
        0,0,false,
    },
    {
        Txref_magic_btc_mainnet,
        Txref_bech32_hrp_mainnet,
        "tx1:rjk0-u5n1-2jsi-mc", /* error correct test >2tsi< instead of >4jsf<*/
        466793,
        2205,
        1,0,false,
    },
    {
        Txref_magic_btc_mainnet,
        Txref_bech32_hrp_mainnet,
        "tx1:rqqq-qqqq-qmhu-qk",
        0,
        0,
        0,0,false,
    },
    {
        Txref_magic_btc_mainnet,
        Txref_bech32_hrp_mainnet,
        "tx1:rzqq-qqqq-uvlj-ez",
        1,
        0,
        0,0,false,
    },
    {
        Txref_magic_btc_mainnet,
        Txref_bech32_hrp_mainnet,
        "tx1:rjk0-u5ng-4jsf-mc", /* complete invalid */
        0,
        0,
        1,1,false, /* enc & dec must fail */
    },
    {
        Txref_magic_btc_mainnet,
        Txref_bech32_hrp_mainnet,
        "tx1:r7ll-lrar-a27h-kt",
        2097151, /* last valid block height with current enc/dec version is 0x1FFFFF*/
        1000,
        0,0,false,
    },
    {
        Txref_magic_btc_mainnet,
        Txref_bech32_hrp_mainnet,
        "", /* encoding must fail, no txref to chain against */
        2097152, /* invalid height */
        1000,
        2,1,false,
    },
    {
        Txref_magic_btc_mainnet,
        Txref_bech32_hrp_mainnet,
        "tx1:r7ll-llll-khym-tq",
        2097151, /* last valid block height with current enc/dec version is 0x1FFFFF*/
        8191, /* last valid tx pos is 0x1FFF */
        0,0,false,
    },
    {
        Txref_magic_btc_mainnet,
        Txref_bech32_hrp_mainnet,
        "", /* encoding must fail, no txref to chain against */
        2097151, /* last valid block height with current enc/dec version is 0x1FFFFF*/
        8192, /* invalid tx pos */
        2,1,false,
    },
    {
        Txref_magic_btc_mainnet,
        Txref_bech32_hrp_mainnet,
        "tx1:r7ll-lrqq-vq5e-gg",
        2097151, /* last valid block height with current enc/dec version is 0x1FFFFF*/
        0,
        0,0,false,
    },
    {
        Txref_magic_btc_mainnet,
        Txref_bech32_hrp_mainnet,
        "tx1:rqqq-qull-6v87-r7",
        0,
        8191, /* last valid tx pos is 0x1FFF */
        0,0,false,
    },
    {
        Txref_magic_btc_mainnet,
        Txref_bech32_hrp_mainnet,
        "tx1:rjk0-u5ng-gghq-fkg7", /* valid Bech32, but 10x5bit packages instead of 8 */
        0,
        0,
        3,2,false, /* ignore encoding */
    },
    {
        Txref_magic_btc_mainnet,
        Txref_bech32_hrp_mainnet,
        "tx1:rjk0-u5qd-s43z", /* valid Bech32, but 6x5bit packages instead of 8 */
        0,
        0,
        3,2,false, /* ignore encoding */
    },
    {
        0xB,
        Txref_bech32_hrp_mainnet,
        "tx1:t7ll-llll-gey7-ez",
        2097151,
        8191,
        0,0,false, /* ignore encoding */
    },
    {
        Txref_magic_btc_mainnet,
        Txref_bech32_hrp_mainnet,
        "tx1:rk63-uvxf-9pqc-sy",
        467883,
        2355,
        0,0,false, /* ignore encoding */
    },
    {
        Txref_magic_btc_testnet,
        Txref_bech32_hrp_testnet,
        "txtest1:xk63-uqvx-fqx8-xqr8",
        467883,
        2355,
        0,0,true, /* ignore encoding */
    },
    {
        Txref_magic_btc_testnet,
        Txref_bech32_hrp_testnet,
        "txtest1:xqqq-qqqq-qqkn-3gh9",
        0,
        0,
        0,0,true, /* ignore encoding */
    },
    {
        Txref_magic_btc_testnet,
        Txref_bech32_hrp_testnet,
        "txtest1:x7ll-llll-llj9-t9dk",
        0x3FFFFFF,
        0x3FFFF,
        0,0,true, /* ignore encoding */
    },

}  

func TestEncoding(t *testing.T) {

	for i, tc := range testVectors {
		tc := tc // capture the test case variable
		t.Run("Running test vector", func(t *testing.T) {
			t.Logf("%d : %s\n", i, tc.encodedTxref)
			if tc.encFail != 3 {
				t.Logf("Encoding magic: %08b height: %08b and position %08b\n", tc.magic, tc.height, tc.position)
				txref, err := Encode(tc.hrp, tc.magic, tc.height, tc.position, tc.nonStd)
				if err != nil {
					if tc.encFail == 0 {
						t.Error(err)
					} else {
						return
					}
				}
				if tc.encodedTxref != txref && tc.encFail == 0 {
					t.Errorf("%d: %d %d failed to encode to %s", i, tc.height, tc.position, tc.encodedTxref)
				} else {
					hrp, decodedMagic, decodedHeight, decodedPosition, err := Decode(txref)
					if err != nil {
						if tc.decFail == 0 {
							t.Error(err)
						} else {
							return
						}
					}
					if hrp != tc.hrp ||
						decodedMagic != tc.magic ||
						decodedHeight != tc.height ||
						decodedPosition != tc.position {
						t.Errorf("%d: %d %d failed to decode to %s from %s\n" +
							"Decoded hrp: %s, magic: %d height: %d, position: %d",
							i, tc.height, tc.position,
							tc.encodedTxref, txref,
							hrp, decodedMagic, decodedHeight, decodedPosition,
						)
					}		
				}
			} else {
				hrpbuf, decodedMagic, decodedHeight, decodedPosition, err := Decode(tc.encodedTxref)
				if err != nil {
					if tc.encFail == 0 {
						t.Error(err)
					} else {
						return
					}
				}
				if hrpbuf != tc.hrp ||
					decodedMagic != tc.magic ||
					decodedHeight != tc.height ||
					decodedPosition != tc.position {
					if tc.encFail == 0 {
						t.Errorf("%d: %d %d failed to decode to %s from %s", i, tc.height, tc.position,
							tc.encodedTxref, tc.encodedTxref)
					} else {
						return
					}					
				}		
			}
		})		
	}	
}
