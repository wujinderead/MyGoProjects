package util

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"hash/crc64"
	"strconv"
	"testing"
)

func TestCrc64(t *testing.T) {
	var hasher = crc64.New(crc64.MakeTable(crc64.ECMA))
	hasher.Write([]byte("赵客缦胡缨，吴钩霜雪明。"))
	hash_uint64 := hasher.Sum64()
	hash_bytes := hasher.Sum([]byte{})
	fmt.Println(hash_uint64)
	fmt.Println(strconv.FormatUint(hash_uint64, 16))
	fmt.Println("hash bytes: ", hex.EncodeToString(hash_bytes))
	fmt.Println(binary.BigEndian.Uint64(hash_bytes))
}

func TestCrc32(t *testing.T) {
	var hasher = crc32.New(crc32.MakeTable(crc32.IEEE))
	hasher.Write([]byte("赵客缦胡缨，吴钩霜雪明。"))
	hash_uint32 := hasher.Sum32()
	hash_bytes := hasher.Sum([]byte{})
	fmt.Println(hash_uint32)
	fmt.Println(strconv.FormatUint(uint64(hash_uint32), 16))
	fmt.Println("hash bytes: ", hex.EncodeToString(hash_bytes))
	fmt.Println(binary.BigEndian.Uint32(hash_bytes))
}
