package set2

// Pkcs7Pad returns a padded version of src
// acording to PKCS#7 padding
func Pkcs7Pad(src []byte, blockSize int) []byte {
	n := len(src)
	padLen := blockSize - (n % blockSize)
	for i := 0; i < padLen; i++ {
		src = append(src, byte(padLen))
	}
	return src
}
