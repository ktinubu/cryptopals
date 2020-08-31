package set2_test

import (
	"cryptopals/set2"
	set1 "cryptopals/set_1"
	"testing"
)

// Check for panic
func TestEcbCbcRandomPad(t *testing.T) {
	for i := 0; i < 20; i++ {
		set2.DetectRandomPadOracle(set1.AesBlockLen)
	}
}
