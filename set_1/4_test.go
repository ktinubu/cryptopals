package set1_test

import (
	"bytes"
	set1 "cryptopals/set_1"
	"log"
	"math"
	"testing"
)

func TestDetectSingleCharacterXOR(t *testing.T) {
	englishFreq := set1.CharacterFrquency(getData("data/shakespeare.txt"))
	scanner, file := scanner("data/4.txt")
	defer file.Close()
	lineNum := 0
	bestChiSqr := math.MaxFloat64
	plaintext := []byte{}
	for scanner.Scan() {
		decoded := set1.DecodeHex([]byte(scanner.Text()))
		key, currChiSqr := set1.CrackSingleByteXOR(decoded, englishFreq)
		if currChiSqr < bestChiSqr {
			bestChiSqr = currChiSqr
			plaintext = set1.SingleByteXOR(decoded, key)
		}
		lineNum++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	answerLine := 327
	if lineNum != answerLine {
		t.Errorf("got: %v\n want: %v", lineNum, answerLine)
	}
	answerText := []byte("Now that the party is jumping\n")
	if !bytes.Equal(plaintext, answerText) {
		t.Errorf("\n  got: %v\n want: %v", plaintext, answerText)
	}
}
