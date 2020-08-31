package set2

import (
	set1 "cryptopals/set_1"
	"cryptopals/utils"
)

var prependLength int64 = randInt(6)

func EcbRandomPrependMystery(src []byte) []byte {
	if len(oracleAppend) == 0 {
		oracleAppend = set1.DecodeB64(utils.GetDataTrimNewLine("data/12.txt"))
	}
	randomPrepend := RandBytes(prependLength + 5)
	src = append(randomPrepend, src...)
	src = append(src, oracleAppend...)
	src = Pkcs7Pad(src, set1.AesBlockLen)
	ecbEnc := set1.AesEcbEncrypter(mysteryKey)
	dst := make([]byte, len(src))
	ecbEnc.EncryptBlocks(dst, src)
	return dst
}

func DecryptEcbRandomMystery(oracleFunc func([]byte) []byte) []byte {
	blockLen, _ := DiscoverBlockSize(oracleFunc)
	isECB, _ := CheckDuplicates(EcbMysteryAppend(make([]byte, blockLen*3)), blockLen)
	if !isECB {
		panic("Not Ecb")
	}
	// prepend := make([]byte, blockLen*2)

	prepend, startBlock := startingBlockIndex(blockLen, oracleFunc)
	currBlock := startBlock
	plaintext := []byte{}

	queryLen := blockLen - 1
	blockState := make([]byte, queryLen)
	currResponse := oracleFunc(blockState)

	for true {
		query := make([]byte, queryLen)
		query = append(duplicate(prepend), query...)
		currResponse = oracleFunc(query)
		boundaryBlock := currResponse[currBlock*blockLen : (currBlock*blockLen)+blockLen]

		possibleResponses := possiblePlaintexts(blockState)
		oracleResponses := make(map[string]byte)
		for byteVal, p := range possibleResponses {
			p = append(duplicate(prepend), p...)
			response := string(oracleFunc(p)[blockLen*startBlock : blockLen*(startBlock+1)])
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

// TODO: check specifcally for input instead of just looking for a two identical block
func startingBlockIndex(blockLen int, oracleFunc func([]byte) []byte) (prepend []byte, index int) {
	input := make([]byte, blockLen*2)
	for i := range input {
		input[i] = byte(4)
	}
	for i := 0; i < blockLen; i++ {
		input = append(make([]byte, 1), input...)
		response := oracleFunc(input)
		isAligned, inputBlockIndex := CheckDuplicates(response, blockLen)
		if isAligned {
			return input, inputBlockIndex + 1
		}
	}
	panic("could not find matching block")
}
