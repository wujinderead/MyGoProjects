package ssh

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestHmac(t *testing.T) {
	key := []byte("麦克唐纳-道格拉斯公司（McDonnell-Douglas Corporation）是美国制造飞机和导弹的大垄断企业。")
	msg := []byte("赵客缦胡缨，吴钩霜雪明。银鞍照白马，飒沓如流星。十步杀一人，千里不留行。事了拂衣去，深藏身与名。")
	fmt.Printf("key: %x\n", key)
	fmt.Printf("msg: %x\n", msg)
	var hasher = sha256.New()
	var blockSize = hasher.BlockSize()

	// Keys longer than blockSize are shortened by hashing them
	if len(key) > blockSize {
		fmt.Println("longer")
		hasher.Write(key)
		key = hasher.Sum(nil) //Key becomes outputSize bytes long
	}

	// keys shorter than blockSize are padded key with zeros to blockSize long
	if len(key) < blockSize {
		fmt.Println("shorter")
		key = append(key, make([]byte, blockSize-len(key))...)
	}

	fmt.Printf("kpd: %x\n", key)
	outerpad := make([]byte, blockSize)
	interpad := make([]byte, blockSize)
	copy(interpad, key)
	copy(outerpad, key)

	for i := range key {
		outerpad[i] ^= 0x5c
		interpad[i] ^= 0x36
	}
	fmt.Printf("opd: %x\n", outerpad)
	fmt.Printf("ipd: %x\n", interpad)

	insum := sha256.Sum256(append(interpad, msg...))
	outsum := sha256.Sum256(append(outerpad, insum[:]...))
	fmt.Printf("out: %x\n", outsum)
	hmacer := hmac.New(sha256.New, key)
	hmacer.Write(msg)
	hma := hmacer.Sum(nil)
	fmt.Printf("hma: %x\n", hma)
}
