package main

import "fmt"

func main() {
	testSwitch()
}

func testSwitch() {
	for _, v := range []int{1, 2, 3, 4, 5, 6} {
		fmt.Println("v =", v)
		switch v {
		case 1:
			fmt.Println(1)
		case 2:
		case 3:
			fmt.Println(2)
		case 4:
			fmt.Println(4)
			fallthrough
		case 5:
			fmt.Println(5)
			fallthrough
		default:
			fmt.Println(6)
		}
		fmt.Println()
	}
}
