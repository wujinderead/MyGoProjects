package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	var dict = make(map[string]int)
	for _, str := range cpdomains {
		count, _ := strconv.Atoi(strings.Split(str, " ")[0])
		domains := strings.Split(strings.Split(str, " ")[1], ".")
		var cur = domains[len(domains)-1]
		dict[cur] = dict[cur] + count
		for i := len(domains) - 2; i >= 0; i-- {
			cur = domains[i] + "." + cur
			dict[cur] = dict[cur] + count
		}
	}
	var ret = make([]string, 0)
	for k, v := range dict {
		ret = append(ret, fmt.Sprintf("%d %s", v, k))
	}
	return ret
}

func main() {
	fmt.Println(subdomainVisits(
		[]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}))
}
