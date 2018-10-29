package main

import (
	"fmt"
)

func champagneTower(poured int, query_row int, query_glass int) float64 {
	var amount = make([][]float64, query_row+1)
	for i := 0; i <= query_row; i++ {
		amount[i] = make([]float64, i+1)
	}
	var stop = true
	amount[0][0] = float64(poured)
	for i := 0; i < query_row; i++ {
		for j := 0; j <= i; j++ {
			if amount[i][j] > 1 {
				stop = false
				amount[i+1][j] = amount[i+1][j] + (amount[i][j]-1)/2
				amount[i+1][j+1] = amount[i+1][j+1] + (amount[i][j]-1)/2
			}
		}
		if stop {
			break
		}
	}
	if amount[query_row][query_glass] > 1 {
		return 1
	} else {
		return amount[query_row][query_glass]
	}
}

func main() {
	fmt.Println(champagneTower(0, 0, 0))
	fmt.Println(champagneTower(1, 0, 0))
	fmt.Println(champagneTower(2, 0, 0))
	fmt.Println(champagneTower(2, 1, 0))
	fmt.Println(champagneTower(2, 1, 1))
}
