package ugly_number

func findNthUglyNumber235(n int) int {
	uglies := make([]int, n+1)
	uglies[0] = 1
	i := 1
	i2, i3, i5 := 0, 0, 0
	for i<=n {
		curmin :=  min(min(uglies[i2]*2, uglies[i3]*3), uglies[i5]*5)
		uglies[i] = curmin
		i++
		for uglies[i2]*2<=curmin {
			i2++
		}
		for uglies[i3]*3<=curmin {
			i3++
		}
		for uglies[i5]*5<=curmin {
			i5++
		}
	}
	return uglies[n]
}

func findNthUglyNumber235Native(n int) int {
	count := 0
	i:=0
	for count<n {
		if isUgly235(i) {
			count++
		}
		i++
	}
	return i-1
}

func isUgly235(n int) bool {
	for n%2==0 {
		n/=2
	}
	for n%3==0 {
		n/=3
	}
	for n%5==0 {
		n/=5
	}
	return n==1
}

func min(a, b int) int {
	if a<b {
		return a
	} else {
		return b
	}
}