package set2_test

import (
	"cryptopals/set2"
	set1 "cryptopals/set_1"
	"testing"
)

// first decrypt, try empty string
func TestCbcBitFlip(t *testing.T) {
	target := []byte("nnone;admin=true")
	plaintext := set2.RandBytes(16)
	txp := set1.FixedXOR(target, plaintext)
	encrypted := set2.CbcSanitizeEncrypt(string(txp))
	prevBlock := encrypted[set1.AesBlockLen : set1.AesBlockLen*2]
	exp := set1.FixedXOR(prevBlock, plaintext)
	copy(prevBlock, exp)
	data, bitFlipped := set2.CbcOutputHacked(encrypted)
	if !bitFlipped {
		t.Errorf("\n  got: %q\n want: %q", data, ";admin=true;")
	}
}
