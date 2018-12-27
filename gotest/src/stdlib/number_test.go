package stdlib

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"unsafe"
	"testing"
)

func TestCongruent(t *testing.T) {
	BI := new(big.Int)
	a1, _ := new(big.Int).SetString("6803298487826435051217540", 10)
	b1, _ := new(big.Int).SetString("411340519227716149383203", 10)
	a2, _ := new(big.Int).SetString("411340519227716149383203", 10)
	b2, _ := new(big.Int).SetString("21666555693714761309610", 10)
	c, _ := new(big.Int).SetString("2244035177043369699245575130906674863160948472041", 10)
	d, _ := new(big.Int).SetString("8912332268928859588025535178967163570016480830", 10)
	congruent := BI.Mul(a1, a2)
	congruent = BI.Div(congruent, b1)
	congruent = BI.Div(congruent, b2)
	fmt.Println(congruent.String())
	a1 = BI.Mul(a1, a1)
	b1 = BI.Mul(b1, b1)
	a2 = BI.Mul(a2, a2)
	b2 = BI.Mul(b2, b2)
	c = BI.Mul(c, c)
	d = BI.Mul(d, d)
	left := BI.Mul(c, b1)
	left = BI.Mul(left, b2)
	right1 := BI.Mul(a1, b2)
	right2 := BI.Mul(a2, b1)
	right := BI.Add(right1, right2)
	right = BI.Mul(right, d)
	fmt.Println(len(left.String()), len(left.Bytes()))
	fmt.Println(left.String())
	fmt.Println(right.String())
	fmt.Println(left.Cmp(right))
}

func TestCongruent1(t *testing.T) {
	BI := new(big.Int)
	a1, _ := new(big.Int).SetString("80155", 10)
	b1, _ := new(big.Int).SetString("20748", 10)
	a2, _ := new(big.Int).SetString("41496", 10)
	b2, _ := new(big.Int).SetString("3485", 10)
	c, _ := new(big.Int).SetString("90514617", 10)
	d, _ := new(big.Int).SetString("72306780", 10)
	congruent := BI.Mul(a1, a2)
	congruent = BI.Div(congruent, b1)
	congruent = BI.Div(congruent, b2)
	fmt.Println(congruent.String())
	a1 = BI.Mul(a1, a1)
	b1 = BI.Mul(b1, b1)
	a2 = BI.Mul(a2, a2)
	b2 = BI.Mul(b2, b2)
	c = BI.Mul(c, c)
	d = BI.Mul(d, d)
	left := BI.Mul(c, b1)
	left = BI.Mul(left, b2)
	right1 := BI.Mul(a1, b2)
	right2 := BI.Mul(a2, b1)
	right := BI.Add(right1, right2)
	right = BI.Mul(right, d)
	fmt.Println(len(left.String()), len(left.Bytes()))
	fmt.Println(left.String())
	fmt.Println(right.String())
	fmt.Println(left.Cmp(right))
	fmt.Println(left.Cmp(new(big.Int).SetInt64(int64(^uint64(0) >> 1))))
}

func TestMaxMin(t *testing.T) {
	uint8_max := ^uint8(0)
	uint8_min := uint8(0)
	int8_max := int8(^uint8(0) >> 1)
	int8_min := ^int8_max
	fmt.Println(uint8_max, reflect.TypeOf(uint8_max))
	fmt.Println(uint8_min, reflect.TypeOf(uint8_min))
	fmt.Println(int8_max, reflect.TypeOf(int8_max))
	fmt.Println(int8_min, reflect.TypeOf(int8_min))

	uint16_max := ^uint16(0)
	uint16_min := uint16(0)
	int16_max := int16(^uint16(0) >> 1)
	int16_min := ^int16_max
	fmt.Println(uint16_max, reflect.TypeOf(uint16_max))
	fmt.Println(uint16_min, reflect.TypeOf(uint16_min))
	fmt.Println(int16_max, reflect.TypeOf(int16_max))
	fmt.Println(int16_min, reflect.TypeOf(int16_min))

	uint32_max := ^uint32(0)
	uint32_min := uint32(0)
	int32_max := int32(^uint32(0) >> 1)
	int32_min := ^int32_max
	fmt.Println(uint32_max, reflect.TypeOf(uint32_max))
	fmt.Println(uint32_min, reflect.TypeOf(uint32_min))
	fmt.Println(int32_max, reflect.TypeOf(int32_max))
	fmt.Println(int32_min, reflect.TypeOf(int32_min))

	uint64_max := ^uint64(0)
	uint64_min := uint64(0)
	int64_max := int64(^uint64(0) >> 1)
	int64_min := ^int64_max
	fmt.Println(uint64_max, reflect.TypeOf(uint64_max))
	fmt.Println(uint64_min, reflect.TypeOf(uint64_min))
	fmt.Println(int64_max, reflect.TypeOf(int64_max))
	fmt.Println(int64_min, reflect.TypeOf(int64_min))

	pos := 9223372036854774801
	neg := -9223372036854774801
	fmt.Println(strconv.FormatUint(uint64(pos), 2))
	fmt.Println(strconv.FormatUint(uint64(neg), 2))

	var base = uintptr(unsafe.Pointer(&pos))
	fmt.Println("size: ", int(unsafe.Sizeof(pos)))
	for i := 0; i < int(unsafe.Sizeof(pos)); i++ {
		abyte := *(*byte)(unsafe.Pointer(base + uintptr(i)))
		fmt.Printf("%x\n", abyte)
	}

	var int32_e = int32(-2147481647)
	fmt.Printf("%x\n", int32_e)
	base = uintptr(unsafe.Pointer(&int32_e))
	fmt.Println("size: ", int(unsafe.Sizeof(int32_e)))
	for i := 0; i < int(unsafe.Sizeof(int32_e)); i++ {
		abyte := *(*byte)(unsafe.Pointer(base + uintptr(i)))
		fmt.Printf("%x\n", abyte)
	}

	bytea := make([]byte, 4)
	fmt.Printf("\n%x\n", uint32(int32_e))
	binary.LittleEndian.PutUint32(bytea, uint32(int32_e))
	for _, b := range bytea {
		fmt.Printf("%02x\n", b)
	}
	bytea1 := make([]byte, 4)
	fmt.Printf("\n%x\n", uint32(int32_e))
	binary.BigEndian.PutUint32(bytea1, uint32(int32_e))
	for _, b := range bytea1 {
		fmt.Printf("%02x\n", b)
	}
}

func TestBigLittle(t *testing.T) {
	bytea := []byte{0x28, 0xa7}
	a := binary.BigEndian.Uint16(bytea)
	b := binary.LittleEndian.Uint16(bytea)
	fmt.Printf("%v, %x\n", a, a)
	fmt.Printf("%v, %x\n", b, b)
	a1 := *(*int16)(unsafe.Pointer(&a))
	b1 := *(*int16)(unsafe.Pointer(&b))
	fmt.Printf("%v\n", a1)
	fmt.Printf("%v\n", b1)
}
