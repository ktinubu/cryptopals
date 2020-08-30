package set1_test

import (
	"bytes"
	set1 "cryptopals/set_1"
	"testing"
)

type hammingStruct struct {
	a        []byte
	b        []byte
	distance float64
}

var hammingInputs = []hammingStruct{
	{
		[]byte("aello"),
		[]byte("hello"),
		2. / 5.,
	},
	{
		[]byte("this is a test"),
		[]byte("wokka wokka!!!"),
		37. / 14.,
	},
}

type keySizeStruct struct {
	message []byte
	keySize int
}

// byte values encoded in b64
var keySizeInputs = []keySizeStruct{
	{
		[]byte("a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4631374860123a2b4826374860123a2120374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123a2b4820374860123"),
		8,
	},
}

func TestHammingDistance(t *testing.T) {
	for i, tt := range hammingInputs {
		dist := set1.HammingDistance(tt.a, tt.b)
		if dist != tt.distance {
			t.Errorf("#%d:\n  got: %v\n want: %v", i, dist, tt.distance)
		}
	}
}

func TestKeySize(t *testing.T) {
	for i, tt := range keySizeInputs {
		keyLen := set1.KeySize(set1.DecodeHex(tt.message))
		if keyLen != tt.keySize {
			t.Errorf("#%d:\n  got: %v\n want: %v", i, keyLen, tt.keySize)
		}
	}
}

func TestTranspose(t *testing.T) {
	input := []byte("1234567123456712345671234567123456712345671234567")
	t7 := set1.Transpose(input, 7)
	expceted := [][]byte{[]byte("11111111"), []byte("22222222"), []byte("33333333"), []byte("44444444"), []byte("55555555"), []byte("66666666"), []byte("77777777")}
	for i, _ := range expceted {
		if bytes.Equal(t7[i], expceted[i]) {
			t.Errorf("got: %v\n want: %v", t7, expceted[i])
		}
	}
}

func TestCrackRepeatingKeyXOR(t *testing.T) {
	englishFreq := set1.CharacterFrquency(getData("data/shakespeare.txt"))
	input := set1.DecodeB64(getDataTrimNewLine("data/6.txt"))
	_, key := set1.CrackRepeatingKeyXOR(input, englishFreq)
	if !bytes.Equal(key, []byte("Terminator X: Bring the noise")) {
		t.Errorf("got: %v\n want: %v", key, "Terminator X: Bring the noise")
	}
}
