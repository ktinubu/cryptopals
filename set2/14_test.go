package set2_test

import (
	"cryptopals/set2"
	"fmt"
	"testing"
)

// first decrypt, try empty string
func TestEcbRandomPrependMystery(t *testing.T) {
	if string(set2.DecryptEcbRandomMystery(set2.EcbRandomPrependMystery))[:6] != "Rollin" {
		fmt.Print(string(set2.DecryptEcbRandomMystery(set2.EcbRandomPrependMystery)))
		t.Error()
	}
}
