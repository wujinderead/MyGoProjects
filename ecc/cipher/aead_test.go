package cipher

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"testing"

	"golang.org/x/crypto/chacha20poly1305"
)

var data1 = []byte("明月出天山，苍茫云海间。长风几万里，吹度玉门关。汉下白登道，胡窥青海湾。由来征战地，不见有人还。戍客望边色，思归多苦颜。高楼当此夜，叹息未应闲。")
var data2 = []byte("前出塞九首·其六  挽弓当挽强，用箭当用长。射人先射马，擒贼先擒王。杀人亦有限，列国自有疆。苟能制侵陵，岂在多杀伤。")

// aead (authenticated encryption with associated data), encryption with authentication.
func TestAesGcm(t *testing.T) {
	blocksize := 16
	key := make([]byte, blocksize) // aes key size 128 192 256 bits
	_, _ = rand.Read(key)
	fmt.Println("aes128 key:", hex.EncodeToString(key))
	aes128block, _ := aes.NewCipher(key)
	enc, dec := make([]byte, blocksize), make([]byte, blocksize)

	// block encrypt only work for one block; for a series of blocks, should use group encryption like CBC, OFB, CFB, CTR
	aes128block.Encrypt(enc, []byte("计算过程中!"))
	fmt.Println("enc:", hex.EncodeToString(enc))

	// block decrypt
	aes128block.Decrypt(dec, enc)
	fmt.Println("dec:", string(dec))

	// wrong key to decrypt
	key[0]++
	wrongkeydecrypter, _ := aes.NewCipher(key)
	wrongkeydecrypter.Decrypt(dec, enc)
	fmt.Println("dec:", string(dec)) // wrong key can still decrypt, decrypter can't know the result is wrong

	// aead-aes-gcm-128, Galois Counter Mode (Galois message authentication code + counter mode)
	aesgcm128, _ := cipher.NewGCM(aes128block)
	fmt.Println("overhead:", aesgcm128.Overhead())
	fmt.Println("noncesize:", aesgcm128.NonceSize())

	nonce := make([]byte, aesgcm128.NonceSize())
	_, _ = rand.Read(nonce) // make a nonce

	// encrypt
	end := aesgcm128.Seal([]byte{}, nonce, data1, []byte("are you ok?"))
	fmt.Println("enc:", len(data1), len(end), hex.EncodeToString(end))

	// decrypt
	dec, err := aesgcm128.Open([]byte{}, nonce, end, []byte("are you ok?"))
	if err != nil {
		fmt.Println("aes gcm open err:", err)
		return
	}
	fmt.Println(bytes.Equal(data1, dec), string(dec))

	// authenticate failed
	_, err = aesgcm128.Open([]byte{}, nonce, end, []byte("are you ok!")) // wrong additional data
	fmt.Println("aes gcm open err:", err)

	// for aead, the same nonce generate same data, we must increment nonce manually
	incrementNonce(nonce)
	end = aesgcm128.Seal([]byte{}, nonce, data2, []byte("i am ok!"))
	fmt.Println("enc:", len(data2), len(end), hex.EncodeToString(end))

	// decrypt should use the same nonce
	dec, err = aesgcm128.Open([]byte{}, nonce, end, []byte("i am ok!"))
	if err != nil {
		fmt.Println("aes gcm open err:", err)
		return
	}
	fmt.Println(bytes.Equal(data2, dec), string(dec))
}

func TestChacha20Poly1305(t *testing.T) {
	key := make([]byte, chacha20poly1305.KeySize)
	_, _ = rand.Read(key)
	c2p1, _ := chacha20poly1305.New(key)
	fmt.Println("overhead:", c2p1.Overhead())
	fmt.Println("noncesize:", c2p1.NonceSize())

	nonce := make([]byte, c2p1.NonceSize())
	_, _ = rand.Read(nonce) // make a nonce

	// encrypt
	end := c2p1.Seal([]byte{}, nonce, data1, []byte("are you ok?"))
	fmt.Println("enc:", len(data1), len(end), hex.EncodeToString(end))

	// decrypt
	dec, err := c2p1.Open([]byte{}, nonce, end, []byte("are you ok?"))
	if err != nil {
		fmt.Println("aes gcm open err:", err)
		return
	}
	fmt.Println(bytes.Equal(data1, dec), string(dec))

	// authenticate failed
	_, err = c2p1.Open([]byte{}, nonce, end, []byte("are you ok!")) // wrong additional data
	fmt.Println("aes gcm open err:", err)

	// increment nonce for next data
	incrementNonce(nonce)
	end = c2p1.Seal([]byte{}, nonce, data2, []byte("i am ok!"))
	fmt.Println("enc:", len(data2), len(end), hex.EncodeToString(end))

	// decrypt should use the same nonce
	dec, err = c2p1.Open([]byte{}, nonce, end, []byte("i am ok!"))
	if err != nil {
		fmt.Println("aes gcm open err:", err)
		return
	}
	fmt.Println(bytes.Equal(data2, dec), string(dec))
}

func TestXChacha20Poly1305(t *testing.T) {
	key := make([]byte, chacha20poly1305.KeySize)
	_, _ = rand.Read(key)
	xc2p1, _ := chacha20poly1305.NewX(key)
	fmt.Println("overhead:", xc2p1.Overhead())
	fmt.Println("noncesize:", xc2p1.NonceSize())

	nonce := make([]byte, xc2p1.NonceSize())
	_, _ = rand.Read(nonce) // make a nonce

	// encrypt
	end := xc2p1.Seal([]byte{}, nonce, data1, []byte("are you ok?"))
	fmt.Println("enc:", len(data1), len(end), hex.EncodeToString(end))

	// decrypt
	dec, err := xc2p1.Open([]byte{}, nonce, end, []byte("are you ok?"))
	if err != nil {
		fmt.Println("aes gcm open err:", err)
		return
	}
	fmt.Println(bytes.Equal(data1, dec), string(dec))

	// authenticate failed
	_, err = xc2p1.Open([]byte{}, nonce, end, []byte("are you ok!")) // wrong additional data
	fmt.Println("aes gcm open err:", err)

	// increment nonce for next data
	incrementNonce(nonce)
	end = xc2p1.Seal([]byte{}, nonce, data2, []byte("i am ok!"))
	fmt.Println("enc:", len(data2), len(end), hex.EncodeToString(end))

	// decrypt should use the same nonce
	dec, err = xc2p1.Open([]byte{}, nonce, end, []byte("i am ok!"))
	if err != nil {
		fmt.Println("aes gcm open err:", err)
		return
	}
	fmt.Println(bytes.Equal(data2, dec), string(dec))
}

func incrementNonce(b []byte) {
	for i := range b {
		b[i]++
		if b[i] != 0 {
			return
		}
	}
}
