package ugly_number

import (
	"testing"
	"fmt"
)

func TestUglies235(t *testing.T) {
	uglies235 := findNthUglyNumber235(50)
	uglies235Native := findNthUglyNumber235Native(50)
	fmt.Println(uglies235)
	fmt.Println(uglies235Native)
	for i := range uglies235 {
		if uglies235[i] != uglies235Native[i] {
			t.Errorf("%d-th ugly number, expect %d, actual: %d",
				i, uglies235Native[i], uglies235[i])
		}
	}
}

func TestUglies357(t *testing.T) {
	uglies357 := findNthUglyNumber357(50)
	uglies357Native := findNthUglyNumber357Native(50)
	fmt.Println(uglies357)
	fmt.Println(uglies357Native)
	for i := range uglies357 {
		if uglies357[i] != uglies357Native[i] {
			t.Errorf("%d-th ugly number, expect %d, actual: %d",
				i, uglies357Native[i], uglies357[i])
		}
	}
}