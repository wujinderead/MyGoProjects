package util

import (
	"fmt"
	"testing"
	"unsafe"
)

func Test_align1(t *testing.T) {
	func() {
		type ds struct {
			a [5]uint8
			b [2]uint8
		}
		s := ds{[5]uint8{1, 2, 3, 4, 5}, [2]uint8{1, 2}}
		s1 := &ds{[5]uint8{45, 6, 255, 254, 5}, [2]uint8{123, 2}}
		p := new(ds)
		p.a = [5]uint8{45, 6, 0, 254, 53}
		p.b = [2]uint8{4, 63}
		fmt.Println(unsafe.Sizeof(s), unsafe.Sizeof(s1), unsafe.Sizeof(p))
	}()
	func() {
		type ds struct {
			a [5]uint8
			b [2]uint
		}
		s := ds{[5]uint8{1, 2, 3, 4, 5}, [2]uint{1, 2}}
		s1 := &ds{[5]uint8{45, 6, 255, 254, 5}, [2]uint{123, 2}}
		p := new(ds)
		p.a = [5]uint8{45, 6, 0, 254, 53}
		p.b = [2]uint{4, 63}
		fmt.Println(unsafe.Sizeof(s), unsafe.Sizeof(s1), unsafe.Sizeof(p))
	}()
	func() {
		type ds struct {
			a [7]uint8
			b [2]uint
		}
		s := ds{[7]uint8{1, 2, 3, 4, 5}, [2]uint{1, 2}}
		s1 := &ds{[7]uint8{45, 6, 255, 254, 5}, [2]uint{123, 2}}
		p := new(ds)
		p.a = [7]uint8{45, 6, 0, 254, 53}
		p.b = [2]uint{4, 63}
		fmt.Println(unsafe.Sizeof(s), unsafe.Sizeof(s1), unsafe.Sizeof(p))
	}()
	func() {
		type ds struct {
			a [9]uint8
			b [2]uint
		}
		s := ds{[9]uint8{1, 2, 3, 4, 5}, [2]uint{1, 2}}
		s1 := &ds{[9]uint8{45, 6, 255, 254, 5}, [2]uint{123, 2}}
		p := new(ds)
		p.a = [9]uint8{45, 6, 0, 254, 53}
		p.b = [2]uint{4, 63}
		fmt.Println(unsafe.Sizeof(s), unsafe.Sizeof(s1), unsafe.Sizeof(p))
	}()
}

func Test_align2(t *testing.T) {
	func() {
		type ds struct {
			c string
			b [2]uint8
		}
		s := ds{"test", [2]uint8{1, 2}}
		fmt.Println(unsafe.Sizeof(s))
	}()
	func() {
		type ds struct {
			c string
			b [2]uint
		}
		s := ds{"test123", [2]uint{1, 2}}
		fmt.Println(unsafe.Sizeof(s))
	}()
	func() {
		type ds struct {
			c string
			b [2]uint
		}
		s := ds{"test", [2]uint{1, 2}}
		fmt.Println(unsafe.Sizeof(s))
	}()
	func() {
		type ds struct {
			c string
			b [2]uint8
		}
		s := ds{"test123", [2]uint8{1, 2}}
		fmt.Println(unsafe.Sizeof(s))
	}()
	func() {
		type ds struct {
			c *string
			b [2]uint
		}
		a := "test"
		s := ds{&a, [2]uint{1, 2}}
		fmt.Println(unsafe.Sizeof(s))
	}()
	func() {
		type ds struct {
			c *string
			b [2]uint8
		}
		a := "test123"
		s := ds{&a, [2]uint8{1, 2}}
		fmt.Println(unsafe.Sizeof(s))
	}()
}
