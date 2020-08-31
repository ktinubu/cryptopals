package set2_test

import (
	"bytes"
	"cryptopals/set2"
	"testing"
)

func TestValidatePadding(t *testing.T) {
	valid, err := set2.ValidatePadding([]byte("ICE ICE BABY\x04\x04\x04\x04"))
	if err != nil {
		t.Error(err)
	}
	if bytes.Equal(valid, []byte("ICE ICE BABY")) {
		t.Errorf("\n  got: %v\n want: %v", []byte("ICE ICE BABY"), valid)
	}
	_, err = set2.ValidatePadding([]byte("ICE ICE BABY\x05\x05\x05\x05"))
	if err == nil {
		t.Error(err)
	}
	_, err = set2.ValidatePadding([]byte("ICE ICE BABY\x01\x02\x03\x04"))
	if err == nil {
		t.Error(err)
	}

}
