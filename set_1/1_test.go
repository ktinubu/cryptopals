package set1_test

import (
	"bytes"
	set1 "cryptopals/set_1"
	"testing"
)

type encTest struct {
	hex  []byte
	b64  []byte
	utf8 []byte
}

var encTests = []encTest{
	{
		[]byte("68656c6c6f"),
		[]byte("aGVsbG8="),
		[]byte("hello"),
	},
	{
		[]byte("676f6f646279652e"),
		[]byte("Z29vZGJ5ZS4="),
		[]byte("goodbye."),
	},
	{
		[]byte("49206d6973732070697a7a6120402050697a7a6120487574"),
		[]byte("SSBtaXNzIHBpenphIEAgUGl6emEgSHV0"),
		[]byte("I miss pizza @ Pizza Hut"),
	},
}

func TestFromUtf8(t *testing.T) {
	for i, tt := range encTests {
		src := tt.utf8
		result := set1.EncodeHex(src)
		if !bytes.Equal(result, tt.hex) {
			t.Errorf("\"%s\" -> hex #%d:\n  got: %v\n want: %v", tt.utf8, i, result, tt.hex)
		}
		result = set1.EncodeB64(src)
		if !bytes.Equal(result, tt.b64) {
			t.Errorf("\"%s\" -> b64 #%d:\n  got: %v\n want: %v", tt.utf8, i, result, tt.b64)
		}
	}
}

func TestFromHex(t *testing.T) {
	for i, tt := range encTests {
		src := set1.DecodeHex(tt.hex)
		result := set1.EncodeB64(src)
		if !bytes.Equal(result, tt.b64) {
			t.Errorf("\"%s\" -> b64 #%d:\n  got: %v\n want: %v", tt.utf8, i, result, tt.b64)
		}
	}
}

func TestFrom64(t *testing.T) {
	for i, tt := range encTests {
		src := set1.DecodeB64(tt.b64)
		result := set1.EncodeHex(src)
		if !bytes.Equal(result, tt.hex) {
			t.Errorf("\"%s\" -> hex #%d:\n  got: %v\n want: %v", tt.utf8, i, result, tt.hex)
		}
	}
}
