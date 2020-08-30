package set1_test

import (
	"bytes"
	set1 "cryptopals/set_1"
	"cryptopals/utils"
	"testing"
)

type messagePair struct {
	cyphertext []byte
	plaintext  []byte
}

var singleByteXORCipherInputs = []messagePair{
	{
		[]byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"),
		[]byte("Cooking MC's like a pound of bacon"),
	},
}

func TestSingleByteXORCipher(t *testing.T) {
	englishFreq := set1.CharacterFrquency(utils.GetData("data/shakespeare.txt"))
	for i, tt := range singleByteXORCipherInputs {
		decoded := set1.DecodeHex(tt.cyphertext)
		bestKey, _ := set1.CrackSingleByteXOR(decoded, englishFreq)
		bestDecrypted := set1.SingleByteXOR(decoded, bestKey)
		if !bytes.Equal(set1.SingleByteXOR(bestDecrypted, bestKey), decoded) {
			t.Errorf("#%d: Xor'ing twice \n  got: %v\n want: %v", i, bestDecrypted, tt.plaintext)
		}
		if !bytes.Equal(bestDecrypted, tt.plaintext) {
			t.Errorf("#%d: decrypted vs plaintext\n  got: %v\n want: %v", i, bestDecrypted, tt.plaintext)
		}
	}
}
