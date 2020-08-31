package set2_test

import (
	"bytes"
	"cryptopals/set2"
	"testing"
)

func TestPkcs7Pad(t *testing.T) {
	inBytes := [][]byte{
		[]byte("YELLOW SUBMARINE"),
		[]byte("hello1234goodbye"),
	}
	inLens := []int{
		20,
		16,
	}
	out := [][]byte{
		[]byte("YELLOW SUBMARINE\x04\x04\x04\x04"),
		[]byte("hello1234goodbye\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10"),
	}

	for i := range inBytes {
		result := set2.Pkcs7Pad(inBytes[i], inLens[i])
		if !bytes.Equal(result, out[i]) {
			t.Errorf("\n  got: %v\n want: %v", result, out[i])
		}
	}
}
