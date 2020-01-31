package main

import "fmt"

// https://leetcode.com/problems/smallest-sufficient-team/

// In a project, you have a list of required skills req_skills, and a list of people.
// The i-th person people[i] contains a list of skills that person has.
// Consider a sufficient team: a set of people such that for every required skill in req_skills,
// there is at least one person in the team who has that skill. We can represent these teams
// by the index of each person: for example, team = [0, 1, 3] represents the people with skills
// people[0], people[1], and people[3].
// Return any sufficient team of the smallest possible size, represented by the index of each person.
// You may return the answer in any order. It is guaranteed an answer exists.
// Example 1:
//   Input: req_skills = ["java","nodejs","reactjs"], people = [["java"],["nodejs"],["nodejs","reactjs"]]
//   Output: [0,2]
// Example 2:
//   Input: req_skills = ["algorithms","math","java","reactjs","csharp","aws"], people = [
//     ["algorithms","math","java"],["algorithms","math","reactjs"],
//     ["java","csharp","aws"], ["reactjs","csharp"],["csharp","math"],["aws","java"]]
//   Output: [1,2]
// Constraints:
//   1 <= req_skills.length <= 16
//   1 <= people.length <= 60
//   1 <= people[i].length, req_skills[i].length, people[i][j].length <= 16
//   Elements of req_skills and people[i] are (respectively) distinct.
//   req_skills[i][j], people[i][j][k] are lowercase English letters.
//   Every skill in people[i] is a skill in req_skills.
//   It is guaranteed a sufficient team exists.

func smallestSufficientTeam(req_skills []string, people [][]string) []int {
	// use an integer to represent the skill set for a person
	pskill := make([]int, len(people))
	for i := range people {
		pskill[i] = getSkillSet(req_skills, people[i])
	}
	fmt.Println(pskill)
	// all 1 to represent all skills, e.g., 0b111=7 to represent all 3 skills
	allskill := (1 << uint(len(req_skills))) - 1

	// let sst(i, j) be the smallest team for people[0...i] to achieve skill set j
	// then there is two option:
	// include people[i], we need people[0...i-1] to represent j-(j&pskill[i])
	// not include people[i], we need people[0...i-1] to represent j
	// thus, sst(i, j) = min(sst(i-1, j-(j&pskill[i]))+1, sst(i-1, j))
	// base case: sst(x, 0)=0; sst(0, j)=1 if (pskill[0]|j)==pskill[0], else sst(0, j)=+INF
	// time O(len(people)*2^(len(req_skills)), space O(2^len(req_skills)) since update line by line

	// first int to record sst, second int to store the team members
	prev, cur := make([][2]int, allskill+1), make([][2]int, allskill+1)

	max := len(people) // all people can cover all skills
	for j := 1; j <= allskill; j++ {
		if pskill[0]|j == pskill[0] {
			prev[j][0] = 1
			prev[j][1] = 1 // 0b0...01 means first person
		} else {
			prev[j][0] = max
			prev[j][1] = (1 << uint(len(people))) - 1 // need all people
		}
	}
	for i := 1; i < len(people); i++ { // ith person
		for j := 1; j <= allskill; j++ {
			cur[j][0] = prev[j][0]
			cur[j][1] = prev[j][1]
			if prev[j-(j&pskill[i])][0]+1 < cur[j][0] {
				cur[j][0] = prev[j-(j&pskill[i])][0] + 1
				cur[j][1] = prev[j-(j&pskill[i])][1] | (1 << uint(i)) // add i-th person
			}
		}
		prev, cur = cur, prev
	}
	sst := make([]int, prev[allskill][0])
	ind := 0
	for i := 0; i < len(people); i++ {
		if (1<<uint(i))&prev[allskill][1] > 0 {
			sst[ind] = i
			ind++
		}
	}
	return sst
}

func getSkillSet(skills []string, myskills []string) int {
	sk := 0
	for _, v := range myskills {
		for i := range skills {
			if skills[i] == v {
				sk |= 1 << uint(i)
			}
		}
	}
	return sk
}

func main() {
	fmt.Println(smallestSufficientTeam([]string{"java", "nodejs", "reactjs"}, [][]string{
		{"java"}, {"nodejs"}, {"nodejs", "reactjs"},
	}))
	fmt.Println(smallestSufficientTeam([]string{"algorithms", "math", "java", "reactjs", "csharp", "aws"},
		[][]string{
			{"algorithms", "math", "java"}, {"algorithms", "math", "reactjs"},
			{"java", "csharp", "aws"}, {"reactjs", "csharp"}, {"csharp", "math"}, {"aws", "java"},
		}))
}
