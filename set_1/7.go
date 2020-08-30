package set1

import (
	"crypto/aes"
	"crypto/cipher"
)

// AesECB returns a one block cipher capable of
// encyptyon and another block cipher capable
// of decryption
func AesECB(key []byte) (Encrypter, Decrypter) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	return (Encrypter)(newECB(block)), (Decrypter)(newECB(block))
}

// Encrypter represents a block cipher
// running in a block based mode that
// implements encryption
type Encrypter interface {
	EncryptBlocks(dst, src []byte)
}

// Decrypter represents a block cipher
// running in a block based mode that
// implements decryption
type Decrypter interface {
	DecryptBlocks(dst, src []byte)
}

type ecb struct {
	block     cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{b, b.BlockSize()}
}

func (b *ecb) EncryptBlocks(dst, src []byte) {
	if len(src)%b.blockSize != 0 {
		panic("Need a multiple of the blocksize")
	}
	if len(src) != len(dst) {
		panic("src and dst are not same length")
	}
	for len(src) > 0 {
		b.block.Encrypt(dst, src)
		src = src[b.blockSize:]
		dst = dst[b.blockSize:]
	}
}

func (b *ecb) DecryptBlocks(dst, src []byte) {
	if len(src)%b.blockSize != 0 {
		panic("Need a multiple of the blocksize")
	}
	if len(src) != len(dst) {
		panic("src and dst are not same length")
	}
	for len(src) > 0 {
		b.block.Decrypt(dst, src)
		src = src[b.blockSize:]
		dst = dst[b.blockSize:]
	}
}
