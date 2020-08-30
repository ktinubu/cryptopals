package set1

import (
	"math"
)

const (
	byteLen     = 8
	startKeyLen = 2
	endKeyLen   = 40
)

// CrackRepeatingKeyXOR decrypted a cyphetext repeated xor'ed by the same key
func CrackRepeatingKeyXOR(src []byte, freqTable []float64) (plaintext, key []byte) {
	keyLen := KeySize(src)
	slices := Transpose(src, keyLen)
	key = make([]byte, keyLen)
	for i := range slices {
		key[i], _ = CrackSingleByteXOR(slices[i], freqTable)
	}
	return RepeatedByteXOR(src, key), key
}

// KeySize returns the keysize that minimizes the hamming distance of
// between the keyLen size byte blocks of src
func KeySize(src []byte) int {
	currKeyLen := startKeyLen
	bestKeyLen := 0
	bestHDist := math.MaxFloat64
	// split up src into currKeyLen chunks and calculate average hamming
	// pairs: dist(chunk1, chunk,) dist(chunk2, chunk3)...
	n := len(src)
	for currKeyLen <= endKeyLen && (n/currKeyLen)/2 > 0 {
		distSum := 0.
		count := 0
		for i := 0; i < (n/currKeyLen)/2; i++ {
			leftIdx := i * 2 * currKeyLen
			rightIdx := leftIdx + 2*currKeyLen
			twoChunks := src[leftIdx:rightIdx]
			d := HammingDistance(twoChunks[:currKeyLen], twoChunks[currKeyLen:])
			distSum += d
			count++
		}
		currHDist := distSum / float64(count)
		if currHDist < bestHDist {
			bestHDist = currHDist
			bestKeyLen = currKeyLen
		}
		currKeyLen++
	}
	return bestKeyLen
}

// Transpose src such that each byte slice 'B' contains all src[i]
// where i % keyLen = B
func Transpose(src []byte, keyLen int) [][]byte {
	slices := make([][]byte, keyLen)
	for i := range src {
		slices[i%keyLen] = append(slices[i%keyLen], src[i])
	}
	return slices
}

// HammingDistance returns the normalized hamming distance
// betweeen 2 equal length byte slices
func HammingDistance(a, b []byte) float64 {
	if len(a) != len(b) {
		panic("slices are different lengths")
	}
	dist := 0
	for i := range a {
		dist += hammingDistanceByte(a[i], b[i])
	}
	return float64(dist) / float64(len(a))
}

func hammingDistanceByte(a, b byte) int {
	dist := 0
	xor := a ^ b
	mask := byte(1)
	for i := 0; i < byteLen; i++ {
		diff := int(mask & xor)
		dist += diff
		xor >>= 1
	}
	return dist
}
