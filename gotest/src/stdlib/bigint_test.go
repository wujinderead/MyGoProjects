package stdlib

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"testing"
)

func TestBigInt(t *testing.T) {
	x, _ := new(big.Int).SetString("3976292833650478493946274414963605588507971231667974845467637", 10)
	fmt.Println(x)
	fmt.Println(x.Text(16))
	fmt.Println(x.Text(2))
	fmt.Println(hex.EncodeToString(x.Bytes()))  // x.Bytes() return bytes in big-endian form
	fmt.Println()

	// the big.Int is underlying stored in []big.Word, i.e. []uint, in little endian form
	arr := x.Bits()
	for i := len(arr)-1; i>=0; i-- {  // reversing the order will be compatible with big-endian
		fmt.Print(uint64ToStr(uint(arr[i])), " ")
	}
	fmt.Println()
	fmt.Println(arr)
	fmt.Println()

	// negate x
	x.Neg(x)
	fmt.Println(x)
	fmt.Println(x.Text(16))
	fmt.Println(x.Text(2))
	fmt.Println(hex.EncodeToString(x.Bytes()))  // x.Bytes() return the abs, so negative is the same as positive
	fmt.Println()

	// x.Bits() return the underlying array, so modifying it can modify the big.Int
	arr[0] = 0x23456789
	arr[1] = 0xfdcaed
	fmt.Println(arr)
	fmt.Println(x)
}

func uint64ToStr(uinter uint) string {
	str := strconv.FormatUint(uint64(uinter), 2)
	return strings.Repeat("0", 64-len(str)) + str
}