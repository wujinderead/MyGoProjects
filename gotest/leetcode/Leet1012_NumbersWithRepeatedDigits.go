package leetcode

import "fmt"

func numDupDigitsAtMostN(N int) int {
	if N < 10 {
		return 0
	}
	digit := 0
	// get all digits, e.g., N=23456
	digits := make([]int, 10)
	for i := N; i > 0; i = i / 10 {
		digits[digit] = i % 10
		digit++
	}

	// calculate non repeat numbers to exclude
	cur := 9
	all := 1
	non := 1
	for i := digit - 1; i > 1; i-- {
		non *= cur
		all += non
		cur--
	}

	// plus non repeat in 10000-19999
	countNonRepeat := all*9 + (digits[digit-1]-1)*non*cur

	// count 20000-23456
	has := make([]int, 10)
	cur = digits[digit-1]
	has[cur] = 1 // most significant digit present, 2
	for i := digit - 2; i >= 0; i-- {
		cur = digits[i]     // current digit = 3
		curHas := digit - i // 23 has present, curHas=2
		non = 10 - curHas   // 10-curHas=8 digits can use
		all = 1
		for k := 0; k < i; k++ { // i=3 positions remain, 20xxx, can has 8*7*6 non repeat
			all *= non
			non--
		}
		if i == 0 {
			cur += 1 // if last digit, 0<=j<=cur; other wise, 0<=j<cur
		}
		for j := 0; j < cur; j++ { // from 0 to 2, test 20xxx, 21xxx, 22xxx
			if has[j] != 1 { // j=0,1 hasn't present, i.e., 20xxx and 21xxx
				countNonRepeat += all
			} // else j=2, 22xxx, has repeat, ignore
		}
		if i == 0 || has[cur] == 1 { // e.g. 22456, no need to continue, 22000-22456 all repeat
			break
		}
		has[cur] = 1 // e.g. 23456, move to next digit
	}
	return N - countNonRepeat
}

func main() {
	for _, i := range []int{5, 7, 9, 10, 11,
		43, 44, 48, 99, 100, 102, 1008, 19999, 20166, 20266, 23456} {
		fmt.Println(i, "=", numDupDigitsAtMostN(i))
	}
}
