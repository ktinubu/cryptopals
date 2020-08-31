package set2_test

import (
	"cryptopals/set2"
	"fmt"
	"testing"
)

// first decrypt, try empty string
func TestEcbMysteryAppend(t *testing.T) {
	if string(set2.DecryptMysteryAppend(set2.EcbMysteryAppend))[:6] != "Rollin" {
		fmt.Print(string(set2.DecryptMysteryAppend(set2.EcbMysteryAppend)))
		t.Error()
	}
}
