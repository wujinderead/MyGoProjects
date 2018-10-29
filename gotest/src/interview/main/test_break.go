package main

import "fmt"

func main() {
loop1:
	for i := 0; i < 5; i++ {
		//loop2:
		for j := 0; j < 5; j++ {
			//loop3:
			for k := 0; k < 5; k++ {
				if k%3 == 2 {
					break loop1
				}
				fmt.Printf("%d,%d,%d\n", i, j, k)
			}
		}
	}
}
