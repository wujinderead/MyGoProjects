package main

import (
	"fmt"
)

// https://leetcode.com/problems/minimum-number-of-frogs-croaking/

// Given the string croakOfFrogs, which represents a combination of the string
// "croak" from different frogs, that is, multiple frogs can croak at the same time,
// so multiple “croak” are mixed. Return the minimum number of different frogs to
// finish all the croak in the given string.
// A valid "croak" means a frog is printing 5 letters ‘c’, ’r’, ’o’, ’a’, ’k’
// sequentially. The frogs have to print all five letters to finish a croak. If the
// given string is not a combination of valid "croak" return -1.
// Example 1:
//   Input: croakOfFrogs = "croakcroak"
//   Output: 1
//   Explanation: One frog yelling "croak" twice.
// Example 2:
//   Input: croakOfFrogs = "crcoakroak"
//   Output: 2
//   Explanation: The minimum number of frogs is two. 
//     The first frog could yell "crcoakroak".
//     The second frog could yell later "crcoakroak".
// Example 3:
//   Input: croakOfFrogs = "croakcrook"
//   Output: -1
//   Explanation: The given string is an invalid combination of "croak" from different frogs.
// Example 4:
//   Input: croakOfFrogs = "croakcroa"
//   Output: -1
// Constraints:
//   1 <= croakOfFrogs.length <= 10^5
//   All characters in the string are: 'c', 'r', 'o', 'a' or 'k'.

func minNumberOfFrogs(croakOfFrogs string) int {
	c, r, o, a, k := 0, 0, 0, 0, 0
	// first round to valid the sequence
    for _, ch := range croakOfFrogs {
    	if ch=='c' {
    		c++
		}
		if ch=='r' {
			r++
		}
		if ch=='o' {
			o++
		}
		if ch=='a' {
			a++
		}
		if ch=='k' {
			k++
		}
		if !(c>=r && r>=o && o>=a && a>=k) {   // should be in order
			return -1
		}
	}
	if !(c==r && r==o && o==a && a==k) {   // the sequence is invalid
		return -1
	}
	max := 0
	i := 0
	count := 0
	for i<len(croakOfFrogs) {
		if croakOfFrogs[i] == 'c' {
			count++
		}
		if croakOfFrogs[i] == 'k' {
			if count>max {
				max = count
			}
			count--
		}
		i++
	}
	return max
}

func main() {
	fmt.Println(minNumberOfFrogs("croakcroak"))
	fmt.Println(minNumberOfFrogs("crcoakroak"))
	fmt.Println(minNumberOfFrogs("croakcrook"))
	fmt.Println(minNumberOfFrogs("croakcroa"))
	fmt.Println(minNumberOfFrogs("cccrrroooakakakccrrooaakk"))
	fmt.Println(minNumberOfFrogs("crcroaaokk"))
	fmt.Println(minNumberOfFrogs("ccccccccccrrccccccrcccccccccccrcccccccccrcccccccccccrcccccrcccrrcccccccccccccrocrrcccccccccrccrocccccrccccrrcccccccrrrcrrcrccrcoccroccrccccccccorocrocccrrrrcrccrcrcrcrccrcroccccrccccroorcacrkcccrrroacccrrrraocccrrcrrccorooccrocacckcrcrrrrrrkrrccrcoacrcorcrooccacorcrccccoocroacroraoaarcoorrcrcccccocrrcoccarrorccccrcraoocrrrcoaoroccooccororrrccrcrocrrcorooocorarccoccocrrrocaccrooaaarrcrarooaarrarrororrcrcckracaccorarorocacrrarorrraoacrcokcarcoccoorcrrkaocorcrcrcrooorrcrroorkkaaarkraroraraarooccrkcrcraocooaoocraoorrrccoaraocoorrcokrararrkaakaooroorcororcaorckrrooooakcarokokcoarcccroaakkrrororacrkraooacrkaraoacaraorrorrakaokrokraccaockrookrokoororoooorroaoaokccraoraraokakrookkroakkaookkooraaocakrkokoraoarrakakkakaroaaocakkarkoocokokkrcorkkoorrkraoorkokkarkakokkkracocoaaaaakaraaooraokarrakkorokkoakokakakkcracarcaoaaoaoorcaakkraooaoakkrrroaoaoaarkkarkarkrooaookkroaaarkooakarakkooaokkoorkroaaaokoarkorraoraorcokokaakkaakrkaaokaaaroarkokokkokkkoakaaookkcakkrakooaooroaaaaooaooorkakrkkakkkkaokkooaakorkaroaorkkokaakaaaaaocrrkakrooaaroroakrakrkrakaoaaakokkaaoakrkkoakocaookkakooorkakoaaaaakkokakkorakaaaaoaarkokorkakokakckckookkraooaakokrrakkrkookkaaoakaaaokkaokkaaoakarkakaakkakorkaakkakkkakaaoaakkkaoaokkkakkkoaroookakaokaakkkkkkakoaooakcokkkrrokkkkaoakckakokkocaokaakakaaakakaakakkkkrakoaokkaakkkkkokkkkkkkkrkakkokkroaakkakaoakkoakkkkkkakakakkkaakkkkakkkrkoak"))
}