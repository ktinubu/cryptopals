package set2

import (
	"errors"
)

// PKCS#7 padding validation. returns input with padding
// stripped
func ValidatePadding(src []byte) ([]byte, error) {
	numPad := int(src[len(src)-1:][0])
	if numPad >= len(src) {
		return nil, errors.New("invalid pad")
	}
	padding := src[len(src)-numPad:]
	stripped := src[len(src)-len(padding):]
	for _, b := range padding {
		if int(b) != numPad {
			return nil, errors.New("invalid pad")
		}
	}
	return stripped, nil
}
