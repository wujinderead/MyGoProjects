package leetcode

import "fmt"

func pushDominoes(dominoes string) string {
	var ret = []byte(dominoes)
	i := 0
	R := -1
	L := -1
	for i < len(dominoes) {
		ret[i] = dominoes[i]
		if dominoes[i] == 'R' {
			if R > L {
				for j := R; j < i; j++ {
					ret[j] = 'R'
				}
			}
			R = i
		}
		if dominoes[i] == 'L' {
			if L > R || R == -1 {
				for j := L + 1; j < i; j++ {
					ret[j] = 'L'
				}
			} else if R > L {
				r := R + 1
				l := i - 1
				for r < l {
					ret[r] = 'R'
					ret[l] = 'L'
					r++
					l--
				}
			}
			L = i
		}
		i++
	}
	if R > L {
		for i := R; i < len(dominoes); i++ {
			ret[i] = 'R'
		}
	}
	return string(ret)
}

func main() {
	in := []string{".L.L.R.R...L.L.R..L..R..R..", ".L.R...LR..L..", "RR.L"}
	for _, str := range in {
		fmt.Println(str)
		fmt.Println(pushDominoes(str))
	}
}
