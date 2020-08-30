package set1

// RepeatedByteXOR repeatedly xor's src by
// the key and returns the result
func RepeatedByteXOR(src []byte, key []byte) []byte {
	dst := []byte{}
	keyLen := len(key)
	for i, c := range src {
		dst = append(dst, c^key[i%keyLen])
	}
	return dst
}
