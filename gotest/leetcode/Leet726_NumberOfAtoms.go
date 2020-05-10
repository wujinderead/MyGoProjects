package main

import (
	"sort"
	"strconv"
	"strings"
	"fmt"
)

// https://leetcode.com/problems/number-of-atoms/

// Given a chemical formula (given as a string), return the count of each atom.
// An atomic element always starts with an uppercase character, then zero or more
// lowercase letters, representing the name. 1 or more digits representing the count
// of that element may follow if the count is greater than 1. If the count is 1,
// no digits will follow. For example, H2O and H2O2 are possible, but H1O2 is impossible.
// Two formulas concatenated together produce another formula. For example, H2O2He3Mg4
// is also a formula. A formula placed in parentheses, and a count (optionally added) is
// also a formula. For example, (H2O2) and (H2O2)3 are formulas.
// Given a formula, output the count of all elements as a string in the following
// form: the first name (in sorted order), followed by its count (if that count is
// more than 1), followed by the second name (in sorted order), followed by its count
// (if that count is more than 1), and so on.
// Example 1:
//   Input: formula = "H2O"
//   Output: "H2O"
//   Explanation: The count of elements are {'H': 2, 'O': 1}.
// Example 2:
//   Input: formula = "Mg(OH)2"
//   Output: "H2MgO2"
//   Explanation: The count of elements are {'H': 2, 'Mg': 1, 'O': 2}.
// Example 3:
//   Input: formula = "K4(ON(SO3)2)2"
//   Output: "K4N2O14S4"
//   Explanation: The count of elements are {'K': 4, 'N': 2, 'O': 14, 'S': 4}.
// Note:
//   All atom names consist of lowercase letters, except for the first character which is uppercase.
//   The length of formula will be in the range [1, 1000].
//   formula will only consist of letters, digits, and round parentheses, and is a
//     valid formula as defined in the problem.

func countOfAtoms(s string) string {
    buf := make([]pair, 0)
    stack := make([]int, 0)
    i := 0
    for i<len(s) {
    	if s[i]>='A' && s[i]<='Z' {
    		j := i+1
    		for j<len(s) && s[j]>='a' && s[j]<='z' {
    			j++
			}
			atom := s[i: j]
			count := 0
			for j<len(s) && s[j]>='0' && s[j]<='9' {
				count = count*10+int(s[j]-'0')
				j++
			}
			if count==0 {
				count = 1
			}
			buf = append(buf, pair{atom, count})
			i = j
			continue
		}
		if s[i] == '(' {
			stack = append(stack, len(buf))
			i++
			continue
		}
		if s[i] == ')' {
			count := 0
			j := i+1
			for j<len(s) && s[j]>='0' && s[j]<='9' {
				count = count*10+int(s[j]-'0')
				j++
			}
			if count==0 {
				count = 1
			}
			i = j
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if count>0 {
				for j:=top; j<len(buf); j++ {
					buf[j].count *= count
				}
			}
		}
	}
	countmap := make(map[string]int)
	for i := range buf {
		if v, ok := countmap[buf[i].atom]; ok {
			countmap[buf[i].atom] = v+buf[i].count
		} else {
			countmap[buf[i].atom] = buf[i].count
		}
	}
	keys := make([]string, 0, len(countmap))
	for k := range countmap {
		keys = append(keys, k)
	}
	sort.Sort(sort.StringSlice(keys))
	var b strings.Builder
	for i := range keys {
		b.WriteString(keys[i])
		if countmap[keys[i]]>1 {
			b.WriteString(strconv.Itoa(countmap[keys[i]]))
		}
	}
	return b.String()
}

type pair struct {
	atom string
	count int
}

func main() {
	fmt.Println(countOfAtoms("H2O"))
	fmt.Println(countOfAtoms("Mg(OH)2"))
	fmt.Println(countOfAtoms("K4(ON(SO3)2)2"))
	fmt.Println(countOfAtoms("K4(CN(SO3)2Be(C3CaC)3E)2F"))
}