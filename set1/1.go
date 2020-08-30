// Package set1 contain solutions for set 1 of the cryptopals cryptography challenges
package set1

import (
	"encoding/base64"
	"encoding/hex"
)

func EncodeHex(src []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return dst
}

func DecodeHex(hexStr []byte) []byte {
	dst := make([]byte, hex.DecodedLen(len(hexStr)))
	n, err := hex.Decode(dst, hexStr)
	if err != nil {
		panic(err)
	}
	return dst[:n]
}

func EncodeB64(src []byte) []byte {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(dst, src)
	return dst
}

func DecodeB64(b64Str []byte) []byte {
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(b64Str)))
	n, err := base64.StdEncoding.Decode(dst, b64Str)
	if err != nil {
		panic(err)
	}
	return dst[:n]
}
