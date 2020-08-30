package set1

import (
	"crypto/aes"
	"crypto/cipher"
)

// AesECB returns a block
func AesECB(key []byte) BlockCipher {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	return (BlockCipher)(newECB(block))
}

// BlockCipher represents a block cipher
// running in a block based mode
type BlockCipher interface {
	EncryptBlocks(dst, src []byte)
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
