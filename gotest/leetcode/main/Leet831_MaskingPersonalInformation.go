package main

import "fmt"

func maskPII(S string) string {
	first := S[0]
	if (first >= 'a' && first <= 'z') || (first >= 'A' && first <= 'Z') { // email
		var tmp = make([]byte, len(S))
		i_of_at := 0
		for i := 0; i < len(S); i++ {
			if S[i] >= 'A' && S[i] <= 'Z' {
				tmp[i] = 'a' + (S[i] - 'A')
			} else {
				tmp[i] = S[i]
			}
			if S[i] == '@' {
				i_of_at = i
			}
		}
		var ans = make([]byte, 0)
		ans = append(ans, tmp[0], '*', '*', '*', '*', '*', tmp[i_of_at-1])
		for i := i_of_at; i < len(tmp); i++ {
			ans = append(ans, tmp[i])
		}
		return string(ans)
	} else { // phone
		var tmp = make([]byte, 0)
		for i := 0; i < len(S); i++ {
			if S[i] >= '0' && S[i] <= '9' {
				tmp = append(tmp, S[i])
			}
		}
		var ans = make([]byte, 0)
		if len(tmp) > 10 {
			ans = append(ans, '+')
			for i := 0; i < len(tmp)-10; i++ {
				ans = append(ans, '*')
			}
			ans = append(ans, '-')
		}
		ans = append(ans, '*', '*', '*', '-', '*', '*', '*', '-',
			tmp[len(tmp)-4], tmp[len(tmp)-3], tmp[len(tmp)-2], tmp[len(tmp)-1])
		return string(ans)
	}
}

func main() {
	fmt.Println(maskPII("Lgq@LeetCode.com"))
	fmt.Println(maskPII("ab@qq.Com"))
	fmt.Println(maskPII("1(234)567-890"))
	fmt.Println(maskPII("86-(10)12345678"))
}
