package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/sender-with-largest-word-count/

// You have a chat log of n messages. You are given two string arrays messages and senders where
// messages[i] is a message sent by senders[i].
// A message is list of words that are separated by a single space with no leading or trailing spaces.
// The word count of a sender is the total number of words sent by the sender. Note that a sender may
// send more than one message.
// Return the sender with the largest word count. If there is more than one sender with the largest
// word count, return the one with the lexicographically largest name.
// Note:
//   Uppercase letters come before lowercase letters in lexicographical order.
//   "Alice" and "alice" are distinct.
// Example 1:
//   Input: messages = ["Hello userTwooo","Hi userThree","Wonderful day Alice",
//   "Nice day userThree"], senders = ["Alice","userTwo","userThree","Alice"]
//   Output: "Alice"
//   Explanation: Alice sends a total of 2 + 3 = 5 words.
//     userTwo sends a total of 2 words.
//     userThree sends a total of 3 words.
//     Since Alice has the largest word count, we return "Alice".
// Example 2:
//   Input: messages = ["How is leetcode for everyone","Leetcode is useful for
//   practice"], senders = ["Bob","Charlie"]
//   Output: "Charlie"
//   Explanation: Bob sends a total of 5 words.
//     Charlie sends a total of 5 words.
//     Since there is a tie for the largest word count, we return the sender with
//     the lexicographically larger name, Charlie.
// Constraints:
//   n == messages.length == senders.length
//   1 <= n <= 10⁴
//   1 <= messages[i].length <= 100
//   1 <= senders[i].length <= 10
//   messages[i] consists of uppercase and lowercase English letters and ' '.
//   All the words in messages[i] are separated by a single space.
//   messages[i] does not have leading or trailing spaces.
//   senders[i] consists of uppercase and lowercase English letters only.

func largestWordCount(messages []string, senders []string) string {
	mapp := make(map[string]int)
	for i := range messages {
		c := strings.Count(messages[i], " ") + 1
		mapp[senders[i]] = mapp[senders[i]] + c
	}
	max := 0
	maxk := ""
	for k, v := range mapp {
		if v > max || (v == max && strings.Compare(k, maxk) > 0) {
			max = v
			maxk = k
		}
	}
	return maxk
}

func main() {
	for _, v := range []struct {
		messages []string
		senders  []string
		ans      string
	}{
		{
			[]string{"Hello userTwooo", "Hi userThree", "Wonderful day Alice", "Nice day userThree"},
			[]string{"Alice", "userTwo", "userThree", "Alice"},
			"Alice",
		},
		{
			[]string{"How is leetcode for everyone", "Leetcode is useful for practice"},
			[]string{"Bob", "Charlie"},
			"Charlie",
		},
		{
			[]string{"b I j", "OK N x J jt b iO N Y", "Q h y CV UE Q A", "Qo Qy w Aw c", "oh", "OA kC G V GlX", "AD Z A YH Tyl", "MA", "sVD", "a BB o g o A hf H", "qu", "P nAx", "d e As Gd oD C RWb", "kS tI Lt U eq k M A", "cS e R h f gl", "AX dn b w nx", "nX T P B", "F", "Gk eGO", "l y Ue nC D", "o UV W P j p e Ov g", "aI Xr Fs NVz", "H f l", "B AY vs S", "rZ Ku S S pQ", "f N q cP lX o x", "W X X Za t", "Vp a xR X J G h A Vo"},
			[]string{"kXMEHbzSid", "LxSLj", "HvI", "rIffGg", "rIffGg", "RHiE", "HvI", "QWsD", "v", "QWsD", "VUCp", "vsp", "ArRIVvhn", "VUCp", "RHiE", "rIffGg", "FzxQzXec", "FzxQzXec", "VUCp", "VUCp", "vsp", "v", "rDkxpR", "rKsKmX", "rKsKmX", "HvI", "LxSLj", "grfeiaY"},
			"HvI",
		},
	} {
		fmt.Println(largestWordCount(v.messages, v.senders), v.ans)
	}
}
