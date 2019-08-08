package stdlib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLeftShift(t *testing.T) {
	// uint left shift, right shift are logical shift
	// int left shift is logical shift (ditch most left bit, insert 0 at end)
	// int right shift is arithmetical shift (ditch most right bit, insert sign bit at head)

	// the shift size can override the int size, for instance,
	// if uint32(11) << 34, the result will be 0
	int32a := []int32{96754441, -206107026, 803362655}
	uint32a := []uint32{96754441, 0xf3b70e6e, 803362655}

	for _, val := range int32a {
		fmt.Println(reflect.TypeOf(val), val, "<<")
		for i := uint(0); i < 34; i++ {
			fmt.Println(val << i)
		}
		fmt.Println()
	}

	for _, val := range int32a {
		fmt.Println(reflect.TypeOf(val), val, ">>")
		for i := uint(0); i < 34; i++ {
			fmt.Println(val >> i)
		}
		fmt.Println()
	}

	for _, val := range uint32a {
		fmt.Println(reflect.TypeOf(val), val, "<<")
		for i := uint(0); i < 34; i++ {
			fmt.Println(val << i)
		}
		fmt.Println()
	}

	for _, val := range uint32a {
		fmt.Println(reflect.TypeOf(val), val, ">>")
		for i := uint(0); i < 34; i++ {
			fmt.Println(val >> i)
		}
		fmt.Println()
	}
}
