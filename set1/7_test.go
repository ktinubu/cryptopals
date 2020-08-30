package set1_test

import (
	"bytes"
	set1 "cryptopals/set_1"
	"testing"
)

func TestAesECB(t *testing.T) {
	in := set1.DecodeB64([]byte(getDataTrimNewLine("data/7.txt")))
	bc := set1.AesECB([]byte("YELLOW SUBMARINE"))
	plaintext := make([]byte, len(in))
	bc.DecryptBlocks(plaintext, in)
	if !bytes.Equal(plaintext[:10], []byte("I'm back a")) {
		t.Errorf("got: %v\n want: %v", plaintext[:10], []byte("I'm back a"))

	}
}
