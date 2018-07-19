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

func TestExtendedEncoding(t *testing.T) {

	for i, tc := range ExtendedTestVectors[21:22] {
		tc := tc // capture the test case variable
		t.Run("Running test vector", func(t *testing.T) {
			t.Logf("%d : %s\n", i, tc.EncodedTxref)
			if tc.EncFail != 3 {
				t.Logf("Encoding magic: %08b height: %08b, position %08b and vout %08b\n", tc.Magic, tc.Height, tc.Position, tc.Vout)
				t.Logf("Encoding magic: %d height: %d, position %d and vout %d\n", tc.Magic, tc.Height, tc.Position, tc.Vout)
				txref, err := Encode(tc.Hrp, tc.Magic, tc.Height, tc.Position, tc.Vout, tc.NonStd)
				if err != nil {
					if tc.EncFail == 0 {
						t.Error(err)
					} else {
						return
					}
				}
				if tc.EncodedTxref != txref && tc.EncFail == 0 {
					t.Errorf("%d: %d %d %d failed to encode to %s, instead encoding to %s",
						i, tc.Height, tc.Position, tc.Vout, tc.EncodedTxref, txref)
				} else {
					hrp, decodedMagic, decodedHeight, decodedPosition, decodedVout, err := Decode(txref)
					if err != nil {
						if tc.DecFail == 0 {
							t.Error(err)
						} else {
							return
						}
					}
					if hrp != tc.Hrp ||
						decodedMagic != tc.Magic ||
						decodedHeight != tc.Height ||
						decodedPosition != tc.Position ||
						decodedVout != tc.Vout {
						t.Errorf("%d: %d %d %d failed to decode to %s from %s\n" +
							"Decoded hrp: %s, magic: %d height: %d, position: %d, vout: %d",
							i, tc.Height, tc.Position, tc.Vout,
							tc.EncodedTxref, txref,
							hrp, decodedMagic, decodedHeight, decodedPosition, decodedVout,
						)
					}		
				}
			} else {
				hrpbuf, decodedMagic, decodedHeight, decodedPosition, decodedVout, err := Decode(tc.EncodedTxref)
				if err != nil {
					if tc.EncFail == 0 {
						t.Error(err)
					} else {
						return
					}
				}
				if hrpbuf != tc.Hrp ||
					decodedMagic != tc.Magic ||
					decodedHeight != tc.Height ||
					decodedPosition != tc.Position ||
					decodedVout != tc.Vout {
					if tc.EncFail == 0 {
						t.Errorf("%d: %d %d %d failed to decode to %s from %s", i, tc.Height, tc.Position, tc.Vout,
							tc.EncodedTxref, tc.EncodedTxref)
					} else {
						return
					}					
				}		
			}
		})		
	}	
}
