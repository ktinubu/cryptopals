package set1_test

import (
	"bytes"
	set1 "cryptopals/set_1"
	"testing"
)

type challenge5Input struct {
	in  string
	out string
	key []byte
}

var repeatingXorInputs = []challenge5Input{
	{
		"data/5.txt",
		"data/5_xor_hex.txt",
		[]byte("ICE"),
	},
}

func TestRepeatingXorInputs(t *testing.T) {
	for i, tt := range repeatingXorInputs {
		in := []byte(getData(tt.in))
		xored := set1.RepeatedByteXOR(in, tt.key)
		encodedXor := set1.EncodeHex(xored)
		out := []byte(getDataTrimNewLine(tt.out))
		if !bytes.Equal(encodedXor, out) {
			t.Errorf("#%d:\n  got: %q\n want: %q", i, encodedXor, out)
		}
		if !bytes.Equal(set1.RepeatedByteXOR(xored, tt.key), in) {
			t.Errorf("xor twice #%d:\n  got: %v\n want: %v", i, xored, out)
		}

	}
}
