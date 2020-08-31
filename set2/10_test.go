package set2_test

import (
	"bytes"
	"cryptopals/set2"
	set1 "cryptopals/set_1"
	"cryptopals/utils"
	"fmt"
	"testing"
)

func TestAesCBC(t *testing.T) {
	inFile := "data/10.txt"
	key := []byte("YELLOW SUBMARINE")
	iv := make([]byte, 16)
	// iv[0] = byte(44)
	in := set1.DecodeB64(utils.GetDataTrimNewLine(inFile))
	dec := set2.AesCbcDecrypter(key, iv)
	plaintext := make([]byte, len(in))
	fmt.Println(string(plaintext))
	dec.DecryptBlocks(plaintext, in)
	startOfOut := "I'm back and I'm ringin' the bell"
	if string(plaintext[:len(startOfOut)]) != startOfOut {
		t.Errorf("\n  got: %v\n want: %v", plaintext[:len(startOfOut)], []byte(startOfOut))
	}
	reconPlaintext := make([]byte, len(in))
	enc := set2.AesCbcEncrypter(key, iv)
	enc.EncryptBlocks(reconPlaintext, plaintext)
	// fmt.Println(string(in))
	if !bytes.Equal(reconPlaintext, in) {
		t.Errorf("\n  got: %v\n want: %v", reconPlaintext, in)
	}
}
