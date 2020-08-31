package set2

import (
	"crypto/rand"
	set1 "cryptopals/set_1"
	"math/big"
)

// EcbCbcRandomPad returns randomly padded src encrypted using AES in
// ECB or CBC mode with equal probability.
func EcbCbcRandomPad(src []byte) []byte {
	key := RandBytes(set1.AesBlockLen)
	return (EcbCbcRandom(randomPad(src), key))
}

// EcbCbcRandom returns src encrypted using AES in ECB or CBC mode with
// equal probability
func EcbCbcRandom(src, key []byte) []byte {
	randIV := RandBytes(set1.AesBlockLen)
	ecbEnc := set1.AesEcbEncrypter(key)
	cbcEnc := AesCbcEncrypter(key, randIV)
	ecbOrcbc := randInt(2)
	dst := make([]byte, len(src))
	if ecbOrcbc != 0 {
		ecbEnc.EncryptBlocks(dst, src)
		return dst
	}
	cbcEnc.EncryptBlocks(dst, src)
	return dst
}

func randomPad(src []byte) []byte {
	randomPrepend := RandBytes(randInt(6) + 5)
	randomAppend := RandBytes(randInt(6) + 5)
	src = append(randomPrepend, src...)
	src = append(src, randomAppend...)
	src = Pkcs7Pad(src, set1.AesBlockLen)
	return src
}

// DetectRandomPadOracle call returns true if EcbAesRandom output
// was encrypted in ECB mode
func DetectRandomPadOracle(blockLen int) bool {
	input := make([]byte, blockLen*3)
	isEcb, _ := CheckDuplicates(EcbCbcRandomPad(input), blockLen)
	return isEcb
}

// CheckDuplicates breaks src into blockLen size chunks and returns
// true if there are identical chuncks, returning the chunk index
// of the second instance
func CheckDuplicates(src []byte, blockLen int) (seenDup bool, chunkIndex int) {
	seen := make(map[string]bool)
	chunkIndex = 0
	for len(src) > 0 {
		blockString := string(src[:blockLen])
		if seen[blockString] {
			return true, chunkIndex
		}
		seen[blockString] = true
		src = src[blockLen:]
		chunkIndex++
	}
	return false, chunkIndex
}

func randInt(max int64) int64 {
	r, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		panic(err)
	}
	if !r.IsInt64() {
		panic("big int too large to be int64")
	}
	return r.Int64()
}

func RandBytes(len int64) []byte {
	b := make([]byte, len)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}
