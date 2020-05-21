package main

import "fmt"

// https://leetcode.com/problems/rearrange-words-in-a-sentence/

// Given a sentence text (A sentence is a string of space-separated words) in the
// following format:
//   First letter is in upper case.
//   Each word in text are separated by a single space.
// Your task is to rearrange the words in text such that all words are rearranged
// in an increasing order of their lengths. If two words have the same length, arrange
// them in their original order.
// Return the new text following the format shown above.
// Example 1:
//   Input: text = "Leetcode is cool"
//   Output: "Is cool leetcode"
//   Explanation: There are 3 words, "Leetcode" of length 8, "is" of length 2 and "cool" of length 4.
//     Output is ordered by length and the new first word starts with capital letter.
// Example 2:
//   Input: text = "Keep calm and code on"
//   Output: "On and keep calm code"
//   Explanation: Output is ordered as follows:
//     "On" 2 letters.
//     "and" 3 letters.
//     "keep" 4 letters in case of tie order by position in original text.
//     "calm" 4 letters.
//     "code" 4 letters.
// Example 3:
//   Input: text = "To be or not to be"
//   Output: "To be or to be not"
// Constraints:
//   text begins with a capital letter and then contains lowercase letters and
//     single space between words.
//   1 <= text.length <= 10^5

func arrangeWords(text string) string {
	buf := make([]byte, 0, len(text))
	mapp := make(map[int][]int)
	start := 0
	maxlen := 0
	for i:=0; i<=len(text); i++ {
		ch := byte(' ')
		if i<len(text) {
			ch = text[i]
		}
		if ch==' ' {
			if _, ok := mapp[i-start]; ok {
				mapp[i-start] = append(mapp[i-start], start)
			} else {
				mapp[i-start] = []int{start}
			}
			if maxlen<i-start {
				maxlen= i-start
			}
			start = i+1
		}
	}
	for i:=1; i<=maxlen; i++ {
		if v, ok := mapp[i]; ok {
			for _, j := range v {
				if text[j]>='A' && text[j]<='Z' {
					buf = append(buf, text[j]+'a'-'A')
					buf = append(buf, text[j+1: j+i]...)
					buf = append(buf, ' ')
					continue
				}
				buf = append(buf, text[j: j+i]...)
				buf = append(buf, ' ')
			}
		}
	}
	if buf[0]>='a' && buf[0]<='z' {
		buf[0] = buf[0]+'A'-'a'
	}
	return string(buf[:len(buf)-1])
}

func main() {
	fmt.Printf("%q\n", arrangeWords("Leetcode is cool"))
	fmt.Printf("%q\n", arrangeWords("Keep calm and code on"))
	fmt.Printf("%q\n", arrangeWords("To be or not to be"))
	fmt.Printf("%q\n", arrangeWords("Ilovettt"))
	fmt.Printf("%q\n", arrangeWords("I"))
	fmt.Printf("%q\n", arrangeWords("Ok i am not ok"))
}