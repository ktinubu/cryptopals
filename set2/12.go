package set2

import (
	set1 "cryptopals/set_1"
	"cryptopals/utils"
)

const (
	maxTries = 100000
)

var mysteryKey []byte = RandBytes(set1.AesBlockLen)
var oracleAppend []byte

func EcbMysteryAppend(src []byte) []byte {
	if len(oracleAppend) == 0 {
		oracleAppend = set1.DecodeB64(utils.GetDataTrimNewLine("data/12.txt"))
	}
	// src = append([]byte{byte(0)}, src...)
	src = append(src, oracleAppend...)
	src = Pkcs7Pad(src, set1.AesBlockLen)
	// fmt.Println(src)
	// fmt.Println("s")
	ecbEnc := set1.AesEcbEncrypter(mysteryKey)
	dst := make([]byte, len(src))
	ecbEnc.EncryptBlocks(dst, src)
	return dst
}

func DecryptMysteryAppend(oracleFunc func([]byte) []byte) []byte {
	blockLen, _ := DiscoverBlockSize(oracleFunc)
	isECB, _ := CheckDuplicates(EcbMysteryAppend(make([]byte, blockLen*3)), blockLen)
	if !isECB {
		panic("Not Ecb")
	}
	// 0 indexed
	currBlock := 0
	plaintext := []byte{}

	queryLen := blockLen - 1
	blockState := make([]byte, queryLen)
	currResponse := oracleFunc(blockState)

	for true {
		query := make([]byte, queryLen)
		currResponse = oracleFunc(query)
		boundaryBlock := currResponse[currBlock*blockLen : (currBlock*blockLen)+blockLen]
		possibleResponses := possiblePlaintexts(blockState)
		oracleResponses := make(map[string]byte)
		for byteVal, p := range possibleResponses {
			response := string(oracleFunc(p)[:len(blockState)+1])
			oracleResponses[response] = byte(byteVal)
		}
		b, ok := oracleResponses[string(boundaryBlock)]
		// response missing due to change of padding value
		if !ok {
			// Get rid of \x01 padding byte
			return plaintext[:len(plaintext)-1]
		}
		plaintext = append(plaintext, b)
		blockState = append(blockState, b)
		blockState = blockState[1:]
		queryLen--
		if queryLen == -1 {
			queryLen = blockLen - 1
			currBlock++
		}
	}
	return plaintext
}

// returns all possible oracle responses for (src || b) where b
// is any byte value
func possiblePlaintexts(src []byte) [][]byte {
	possiblePlaintexts := [][]byte{}
	for i := 0; i < set1.ByteSize; i++ {
		possible := append(duplicate(src), byte(i))
		possiblePlaintexts = append(possiblePlaintexts, possible)
	}
	return possiblePlaintexts
}

func duplicate(src []byte) []byte {
	dst := make([]byte, len(src))
	if n := copy(dst, src); n != len(src) {
		panic("incomplete copy")
	}
	return dst
}

func DiscoverBlockSize(oracleFunc func([]byte) []byte) (blockLen, plaintextLen int) {
	tries := 0
	input := []byte{0}
	lastResponseSize := len(oracleFunc(input))
	for true {
		if tries > maxTries {
			panic("max tries reached")
		}
		currResponseSize := len(oracleFunc(input))
		sizeDiff := currResponseSize - lastResponseSize
		lastResponseSize = currResponseSize
		if sizeDiff > 1 {
			return sizeDiff, currResponseSize - len(input) - sizeDiff
		}
		input = append(input, byte(0))
	}
	return -1, -1
}
