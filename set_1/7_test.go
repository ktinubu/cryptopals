package set1_test

import (
	"bytes"
	set1 "cryptopals/set_1"
	"cryptopals/utils"
	"testing"
)

func TestAesECB(t *testing.T) {
	in := set1.DecodeB64([]byte(utils.GetDataTrimNewLine("data/7.txt")))
	_, decrypter := set1.AesECB([]byte("YELLOW SUBMARINE"))
	plaintext := make([]byte, len(in))
	decrypter.DecryptBlocks(plaintext, in)
	if !bytes.Equal(plaintext[:10], []byte("I'm back a")) {
		t.Errorf("got: %v\n want: %v", plaintext[:10], []byte("I'm back a"))

	}
}
