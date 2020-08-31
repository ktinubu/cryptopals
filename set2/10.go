package set2

import (
	"crypto/aes"
	"crypto/cipher"
	set1 "cryptopals/set_1"
	"fmt"
)

// AesCbcEncrypter returns an AES block cipher
// in CBC mode that capable of encryption
func AesCbcEncrypter(key, iv []byte) set1.Encrypter {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	return (set1.Encrypter)(newECB(block, iv))
}

// AesCbcDecrypter returns an AES block cipher
// in CBC mode that capable of decryption
func AesCbcDecrypter(key, iv []byte) set1.Decrypter {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	return (set1.Decrypter)(newECB(block, iv))
}

type cbc struct {
	block      cipher.Block
	blockSize  int
	iv         []byte
	prevCipher []byte
}

func newECB(b cipher.Block, iv []byte) *cbc {
	return &cbc{b, b.BlockSize(), iv, iv}
}

func (b *cbc) EncryptBlocks(dst, src []byte) {
	if len(src)%b.blockSize != 0 {
		panic("Need a multiple of the blocksize")
	}
	if len(b.iv) != b.blockSize {
		panic("IV must be same length as block size")
	}
	if len(src) != len(dst) {
		panic("src and dst are not same length")
	}
	for len(src) > 0 {
		blockInput := set1.FixedXOR(src[:b.blockSize], b.prevCipher)
		b.block.Encrypt(dst, blockInput)
		b.prevCipher = dst[:b.blockSize]
		src = src[b.blockSize:]
		dst = dst[b.blockSize:]
	}
}

func (b *cbc) DecryptBlocks(dst, src []byte) {
	b.prevCipher = b.iv
	if len(src)%b.blockSize != 0 {
		panic("Need a multiple of the blocksize")
	}
	if len(src) != len(dst) {
		panic("src and dst are not same length")
	}
	if len(b.iv) != b.blockSize {
		panic("IV must be same length as block size")
	}
	for len(src) > 0 {
		b.block.Decrypt(dst, src)
		numCompied := copy(dst, set1.FixedXOR(dst[:b.blockSize], b.prevCipher))
		if numCompied != b.blockSize {
			panic(fmt.Sprintf("copied %d bytes where block size is %d", numCompied, b.blockSize))
		}
		b.prevCipher = src[:b.blockSize]
		src = src[b.blockSize:]
		dst = dst[b.blockSize:]
	}
}
