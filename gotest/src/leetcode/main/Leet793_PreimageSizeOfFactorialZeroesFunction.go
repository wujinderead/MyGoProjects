package main

import "fmt"

func preimageSizeFZF1(K int) int {
	N := K * 4
	//fmt.Println("f(",N,")=",F(N))
	start := N + (K-F(N))*5
	for F(start) > K {
		//fmt.Println("f(",start,")=",F(start))
		start = start - 5
	}
	if F(start) == K {
		return 5
	} else {
		return 0
	}
}

func preimageSizeFZF(K int) int {
	N := K * 4
	start := F(N)
	fmt.Println("f(", N, ")=", F(N))
	for start < K {
		N = N + 5
		start = F(N)
		fmt.Println("f(", N, ")=", F(N))
	}
	if start == K {
		return 5
	} else {
		return 0
	}
}

func F(N int) int {
	var count int = 0
	var pow int = 5
	for N/pow > 0 {
		count += N / pow
		pow = pow * 5
	}
	return count
}

func main() {
	var i int = 12345687
	fmt.Println(preimageSizeFZF(i))
	fmt.Println(preimageSizeFZF1(i))
}
