package set2

import (
	set1 "cryptopals/set_1"
	"strings"
)

var cbcKey = RandomKey()
var cbcIV []byte = RandBytes(set1.AesBlockLen)

func CbcSanitizeEncrypt(srcString string) []byte {
	srcString = strings.ReplaceAll(srcString, ";", "")
	srcString = strings.ReplaceAll(srcString, "=", "")
	src := []byte(srcString)
	src = append([]byte("comment1=cooking%20MCs;userdata="), src...)
	src = append(src, []byte(";comment2=%20like%20a%20pound%20of%20bacon")...)
	src = Pkcs7Pad(src, set1.AesBlockLen)
	enc := AesCbcEncrypter(cbcKey, cbcIV)
	dst := make([]byte, len(src))
	enc.EncryptBlocks(dst, src)
	return dst
}

func CbcOutputHacked(ciphertext []byte) (string, bool) {
	dec := AesCbcDecrypter(cbcKey, cbcIV)
	dst := make([]byte, len(ciphertext))
	dec.DecryptBlocks(dst, ciphertext)
	output := string(dst)
	splits := strings.Split(output, ";")

	for _, s := range splits {
		tokens := strings.SplitN(s, "=", -1)
		k, v := tokens[0], tokens[1]
		if k == "admin" && v == "true" {
			return output, true
		}
	}
	return output, false
}

func SmallEnc() []byte {
	enc := AesCbcEncrypter(cbcKey, cbcIV)
	src := make([]byte, set1.AesBlockLen*3)
	dst := make([]byte, len(src))
	enc.EncryptBlocks(dst, src)
	return dst
}

func SmallDec(src []byte) []byte {
	dec := AesCbcDecrypter(cbcKey, cbcIV)
	dst := make([]byte, len(src))
	dec.DecryptBlocks(dst, src)
	return dst
}
