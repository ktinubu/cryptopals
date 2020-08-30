package set1_test

import (
	set1 "cryptopals/set_1"
	"testing"
)

func TestDetectAesECB(t *testing.T) {
	scanner, file := scanner("data/8.txt")
	defer file.Close()
	count := 0
	for scanner.Scan() {
		txt := scanner.Bytes()
		if set1.DetectAesECB(txt) {
			count++
		}
	}
	if count != 1 {
		t.Errorf("got %d want %d", count, 1)
	}
}
