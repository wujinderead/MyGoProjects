package cipher

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"crypto/rc4"
	"encoding/hex"
	"fmt"
	"testing"

	"golang.org/x/crypto/blowfish"
	"golang.org/x/crypto/tea"
	"golang.org/x/crypto/twofish"
	"golang.org/x/crypto/xtea"
)

// block encrypt only work for one block; for a series of blocks,
// should use group encryption like CBC, OFB, CFB, CTR ...
func TestBlock(t *testing.T) {
	aes256key := make([]byte, 32) // aes use 128, 192 or 256 bits key
	_, _ = rand.Read(aes256key)
	aesblock, _ := aes.NewCipher(aes256key)

	des3key := make([]byte, 24) // triple des use 192-bit key
	_, _ = rand.Read(des3key)
	des3block, _ := des.NewTripleDESCipher(des3key)

	teakey := make([]byte, tea.KeySize) // tea key 128-bit
	_, _ = rand.Read(teakey)
	teablock, _ := tea.NewCipher(teakey)

	xteakey := make([]byte, 16) // xtea key 128-bit
	_, _ = rand.Read(xteakey)
	xteablock, _ := xtea.NewCipher(xteakey)

	twofishkey := make([]byte, 24) // twofish key 128, 192 or 256 bits
	_, _ = rand.Read(twofishkey)
	twofishblock, _ := twofish.NewCipher(twofishkey)

	blowfishkey := make([]byte, 25) // blowfish key from 1 to 56 bytes
	_, _ = rand.Read(blowfishkey)
	blowfishblock, _ := blowfish.NewCipher(blowfishkey)

	data := []byte("rnzrohfkalcrryjvulkleanqvqhxctrrfbqvklqtbid")
	aesenc := make([]byte, aes.BlockSize)
	aesblock.Encrypt(aesenc, data)
	fmt.Println("aes enc:", hex.EncodeToString(aesenc))
	aesblock.Decrypt(aesenc, aesenc)
	fmt.Println("aes dec:", string(aesenc), bytes.Equal(aesenc, data[:aes.BlockSize]))
	fmt.Println()

	des3enc := make([]byte, des.BlockSize)
	des3block.Encrypt(des3enc, data)
	fmt.Println("des3 enc:", hex.EncodeToString(des3enc))
	des3block.Decrypt(des3enc, des3enc)
	fmt.Println("des3 dec:", string(des3enc), bytes.Equal(des3enc, data[:des.BlockSize]))
	fmt.Println()

	teaenc := make([]byte, tea.BlockSize)
	teablock.Encrypt(teaenc, data)
	fmt.Println("tea enc:", hex.EncodeToString(teaenc))
	teablock.Decrypt(teaenc, teaenc)
	fmt.Println("tea dec:", string(teaenc), bytes.Equal(teaenc, data[:tea.BlockSize]))
	fmt.Println()

	xteaenc := make([]byte, xtea.BlockSize)
	xteablock.Encrypt(xteaenc, data)
	fmt.Println("xtea enc:", hex.EncodeToString(xteaenc))
	xteablock.Decrypt(xteaenc, xteaenc)
	fmt.Println("xtea dec:", string(xteaenc), bytes.Equal(xteaenc, data[:xtea.BlockSize]))
	fmt.Println()

	twofishenc := make([]byte, twofish.BlockSize)
	twofishblock.Encrypt(twofishenc, data)
	fmt.Println("twofish enc:", hex.EncodeToString(twofishenc))
	twofishblock.Decrypt(twofishenc, twofishenc)
	fmt.Println("twofish dec:", string(twofishenc), bytes.Equal(twofishenc, data[:twofish.BlockSize]))
	fmt.Println()

	blowfishenc := make([]byte, blowfish.BlockSize)
	blowfishblock.Encrypt(blowfishenc, data)
	fmt.Println("blowfish enc:", hex.EncodeToString(blowfishenc))
	blowfishblock.Decrypt(blowfishenc, blowfishenc)
	fmt.Println("blowfish dec:", string(blowfishenc), bytes.Equal(blowfishenc, data[:blowfish.BlockSize]))
	fmt.Println()
}

// BlockMode represents a block cipher running in a block-based mode (CBC, ECB etc)
func TestBlockMode(t *testing.T) {
	aes256key := make([]byte, 32) // aes use 128, 192 or 256 bits key
	_, _ = rand.Read(aes256key)
	aesblock, _ := aes.NewCipher(aes256key)

	iv := make([]byte, aesblock.BlockSize()) // cbc iv size equals to block size
	_, _ = rand.Read(iv)
	fmt.Println("iv:", hex.EncodeToString(iv))
	cbcen := cipher.NewCBCEncrypter(aesblock, iv)
	cbcde := cipher.NewCBCDecrypter(aesblock, iv)

	// cbc encrypt
	enc := padToBlocksize(data1, cbcen.BlockSize()) // the input for BlockMode should pad to multiple of blocksize
	cbcen.CryptBlocks(enc, enc)
	fmt.Println("enc:", len(data1), len(enc), hex.EncodeToString(enc))

	// cbc decrypt
	cbcde.CryptBlocks(enc, enc)
	fmt.Println("dec:", bytes.Equal(enc[:len(data1)], data1), string(enc))

	// same original data, different encrypted data, that's because the iv is incrementing explicitly
	cbcen.CryptBlocks(enc, enc)
	fmt.Println("enc:", len(data1), len(enc), hex.EncodeToString(enc))
	cbcde.CryptBlocks(enc, enc)
	fmt.Println("dec:", bytes.Equal(enc[:len(data1)], data1), string(enc))
}

// A Stream represents a stream cipher.
// cfb: cipher feedback mode
func TestStreamCfb(t *testing.T) {
	des3key := make([]byte, 24) // triple des use 192-bit key
	_, _ = rand.Read(des3key)
	des3block, _ := des.NewTripleDESCipher(des3key)

	iv := make([]byte, des3block.BlockSize()) // cfb iv size must equal to block size, otherwise it panics
	_, _ = rand.Read(iv)
	fmt.Println("iv:", hex.EncodeToString(iv))
	cfben := cipher.NewCFBEncrypter(des3block, iv)
	cfbde := cipher.NewCFBDecrypter(des3block, iv)

	// cfb encrypt, stream don't need pad, the dst length is euqal to src
	enc := make([]byte, len(data2))
	cfben.XORKeyStream(enc, data2)
	fmt.Println("enc:", len(data2), len(enc), hex.EncodeToString(enc))

	// cfb decrypt
	cfbde.XORKeyStream(enc, enc)
	fmt.Println("dec:", bytes.Equal(enc, data2), string(enc))

	// encrypt same data, get different result cause iv is incrementing
	cfben.XORKeyStream(enc, data2)
	fmt.Println("enc:", len(data2), len(enc), hex.EncodeToString(enc))

	// cfb decrypt
	cfbde.XORKeyStream(enc, enc)
	fmt.Println("dec:", bytes.Equal(enc, data2), string(enc))
}

func TestStreamOfb(t *testing.T) {
	teakey := make([]byte, tea.KeySize) // tea key 128-bit
	_, _ = rand.Read(teakey)
	teablock, _ := tea.NewCipher(teakey)

	iv := make([]byte, teablock.BlockSize()) // ofb iv size must equal to block size, otherwise it panics
	_, _ = rand.Read(iv)
	fmt.Println("iv:", hex.EncodeToString(iv))
	ofben := cipher.NewOFB(teablock, iv)
	ofbde := cipher.NewOFB(teablock, iv)

	// ofb encrypt, stream don't need pad, the dst length is equal to src
	enc := make([]byte, len(data2))
	ofben.XORKeyStream(enc, data2)
	fmt.Println("enc:", len(data2), len(enc), hex.EncodeToString(enc))

	// ofb decrypt
	ofbde.XORKeyStream(enc, enc)
	fmt.Println("dec:", bytes.Equal(enc, data2), string(enc))

	// encrypt same data, get different result cause iv is incrementing
	ofben.XORKeyStream(enc, data2)
	fmt.Println("enc:", len(data2), len(enc), hex.EncodeToString(enc))

	// cfb decrypt
	ofbde.XORKeyStream(enc, enc)
	fmt.Println("dec:", bytes.Equal(enc, data2), string(enc))
}

func TestStreamCtr(t *testing.T) {
	xteakey := make([]byte, 16) // xtea key 128-bit
	_, _ = rand.Read(xteakey)
	xteablock, _ := xtea.NewCipher(xteakey)

	iv := make([]byte, xteablock.BlockSize()) // ctr iv size must equal to block size, otherwise it panics
	_, _ = rand.Read(iv)
	fmt.Println("iv:", hex.EncodeToString(iv))
	ctren := cipher.NewCTR(xteablock, iv)
	ctrde := cipher.NewCTR(xteablock, iv)

	// ctr encrypt, stream don't need pad, the dst length is equal to src
	enc := make([]byte, len(data2))
	ctren.XORKeyStream(enc, data2)
	fmt.Println("enc:", len(data2), len(enc), hex.EncodeToString(enc))

	// ctr decrypt
	ctrde.XORKeyStream(enc, enc)
	fmt.Println("dec:", bytes.Equal(enc, data2), string(enc))

	// encrypt same data, get different result cause iv is incrementing
	ctren.XORKeyStream(enc, data2)
	fmt.Println("enc:", len(data2), len(enc), hex.EncodeToString(enc))

	// ctr decrypt
	ctrde.XORKeyStream(enc, enc)
	fmt.Println("dec:", bytes.Equal(enc, data2), string(enc))
}

func TestStreamRc4(t *testing.T) {
	rc4key := make([]byte, 101) // rc4 key from 1 to 256 bytes
	_, _ = rand.Read(rc4key)
	rc4en, _ := rc4.NewCipher(rc4key) // *rc4.Cipher implements cipher.Stream interface
	rc4de, _ := rc4.NewCipher(rc4key)

	// rc4 encrypt, stream don't need pad, the dst length is equal to src
	enc := make([]byte, len(data2))
	rc4en.XORKeyStream(enc, data2)
	fmt.Println("enc:", len(data2), len(enc), hex.EncodeToString(enc))

	// ctr decrypt
	rc4de.XORKeyStream(enc, enc)
	fmt.Println("dec:", bytes.Equal(enc, data2), string(enc))

	// encrypt same data, get different result for rc4
	rc4en.XORKeyStream(enc, data2)
	fmt.Println("enc:", len(data2), len(enc), hex.EncodeToString(enc))

	// ctr decrypt
	rc4de.XORKeyStream(enc, enc)
	fmt.Println("dec:", bytes.Equal(enc, data2), string(enc))
}

func padToBlocksize(data []byte, blocksize int) []byte {
	newsize := ((len(data)-1)/blocksize + 1) * blocksize
	pad := make([]byte, newsize)
	copy(pad, data)
	return pad
}
