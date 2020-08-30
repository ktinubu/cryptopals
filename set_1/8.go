package set1

const (
	// AesBlockLen is the lngth of AES block in bytes
	AesBlockLen = 16
)

// DetectAesECB return true if it detects a repeated
// block of cyphertext
func DetectAesECB(src []byte) bool {
	n := len(src)
	if n%AesBlockLen != 0 {
		panic("cyphertext is not multiple of block length")
	}
	seen := make(map[string]bool)
	for i := 0; i < n/16; i++ {
		leftIdx := i * 16
		b := string(src[leftIdx : leftIdx+16])
		if len(b) != AesBlockLen {
			panic("block string is not block length")
		}
		if seen[b] {
			return true
		}
		seen[b] = true
	}
	return false
}
