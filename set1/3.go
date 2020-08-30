package set1

import (
	"math"
)

const (
	// ByteSize is the number of bits in byte
	ByteSize = 256
)

// CrackSingleByteXOR returns the sinlge byte key that results in the highest Chi Squared
// Statistic statistic when analyzing  the character frequencies of src after being xor'ed with
// compared with that of expFreq.
func CrackSingleByteXOR(src []byte, expFreq []float64) (key byte, chiSquared float64) {
	bestChiSqr := ChiSquared(expFreq, CharacterFrquency(src))
	currentKey := byte(0)
	bestKey := byte(0)
	for i := 0; i < ByteSize; i++ {
		decryptedFreq := CharacterFrquency(SingleByteXOR(src, currentKey))
		currChiSqr := ChiSquared(expFreq, decryptedFreq)
		if currChiSqr < bestChiSqr {
			bestChiSqr = currChiSqr
			bestKey = currentKey
		}
		currentKey++
	}
	return bestKey, bestChiSqr
}

// SingleByteXOR repeatedly xor's src by
// the key and returns the result
func SingleByteXOR(src []byte, key byte) []byte {
	dst := []byte{}
	for _, c := range src {
		dst = append(dst, c^key)
	}
	return dst
}

// ChiSquared Returns Chi Squared Statistic between expected and
// observed frequencies
func ChiSquared(exp, obs []float64) float64 {
	if len(exp) != len(obs) {
		panic("Tables for Chi Square are different lengths!")
	}
	sum := 0.
	for i := 0; i < len(exp); i++ {
		if exp[i] != 0 {
			sum += math.Pow(obs[i]-exp[i], 2) / exp[i]
		}
	}
	return sum
}

// CharacterFrquency returns a lenth 256 slice with the
// frequncy of each byte value given
func CharacterFrquency(data []byte) []float64 {
	freq := make([]float64, ByteSize)
	sum := 0.
	for _, char := range data {
		freq[char]++
		sum++
	}
	return freq
}
