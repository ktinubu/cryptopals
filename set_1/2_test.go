package set1_test

import (
	"bytes"
	set1 "cryptopals/set_1"
	"testing"
)

type encoding int

const (
	utf8Enc encoding = iota
	hexEnc
	b64Enc
)

type HexXorTest struct {
	enc encoding
	a   []byte
	b   []byte
	out []byte
}

var testInputs = []HexXorTest{
	{
		utf8Enc,
		[]byte("a"),
		[]byte("B"),
		[]byte("#"),
	},
	{
		hexEnc,
		[]byte("49636520637265616d2069732067726561742e206265737420666f6f6f64"),
		[]byte("627572676572732061726520616c736f206772656174207365636f6e6421"),
		[]byte("2b161747060016410c520c53410b010a41135c4503115307450500010b45"),
	},
	{
		b64Enc,
		[]byte("WG9yaW5nIGlzIHByZXR0eSB0aXJpbmchISEhISEh"),
		[]byte("WXVwIEknbSBwcmV0dHkgZXhhdXN0ZWQgdGJxaC4u"),
		[]byte("ARoCSSdATUkDUhUGEQ1UHFgVHAEdCwMBVUNQSQ8P"),
	},
}

func noOpEncode(src []byte) []byte {
	return src
}
func noOpDecode(src []byte) []byte {
	return src
}
func TestXor(t *testing.T) {
	var encodeFunc func(src []byte) []byte = nil
	var decodeFunc func(src []byte) []byte = nil
	for i, tt := range testInputs {
		switch tt.enc {
		case utf8Enc:
			encodeFunc = noOpEncode
			decodeFunc = noOpDecode
		case hexEnc:
			encodeFunc = set1.EncodeHex
			decodeFunc = set1.DecodeHex
		case b64Enc:
			encodeFunc = set1.EncodeB64
			decodeFunc = set1.DecodeB64
		}
		decA := decodeFunc(tt.a)
		decB := decodeFunc(tt.b)

		result := encodeFunc(set1.FixedXOR(decA, decB))
		if !bytes.Equal(result, tt.out) {
			t.Errorf("#%d:\n  got: %v\n want: %v", i, result, tt.out)
		}
	}
}
