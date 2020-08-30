package set1

// FixedXOR returns the xor of equal sized
// byte slices
func FixedXOR(a, b []byte) []byte {
	n := len(a)
	if n != len(b) {
		panic("slice's are different lengths")
	}
	dst := make([]byte, n)
	for i := 0; i < n; i++ {
		dst[i] = a[i] ^ b[i]
	}
	return dst
}
