package main

import "fmt"

// https://leetcode.com/problems/find-all-possible-recipes-from-given-supplies/

// You have information about n different recipes. You are given a string array
// recipes and a 2D string array ingredients. The iᵗʰ recipe has the name recipes[i],
// and you can create it if you have all the needed ingredients from ingredients[i].
// Ingredients to a recipe may need to be created from other recipes, i.e., ingredients[i]
// may contain a string that is in recipes.
// You are also given a string array supplies containing all the ingredients that you initially have,
// and you have an infinite supply of all of them.
// Return a list of all the recipes that you can create. You may return the answer in any order.
// Note that two recipes may contain each other in their ingredients.
// Example 1:
//   Input: recipes = ["bread"], ingredients = [["yeast","flour"]], supplies = ["yeast","flour","corn"]
//   Output: ["bread"]
//   Explanation:
//     We can create "bread" since we have the ingredients "yeast" and "flour".
// Example 2:
//   Input: recipes = ["bread","sandwich"], ingredients = [["yeast","flour"],["bread","meat"]], supplies = ["yeast","flour","meat"]
//   Output: ["bread","sandwich"]
//   Explanation:
//     We can create "bread" since we have the ingredients "yeast" and "flour".
//     We can create "sandwich" since we have the ingredient "meat" and can create the ingredient "bread".
// Example 3:
//   Input: recipes = ["bread","sandwich","burger"], ingredients = [["yeast","flour"],["bread","meat"],["sandwich","meat","bread"]], supplies = ["yeast","flour", "meat"]
//   Output: ["bread","sandwich","burger"]
//   Explanation:
//   We can create "bread" since we have the ingredients "yeast" and "flour".
//   We can create "sandwich" since we have the ingredient "meat" and can create the ingredient "bread".
//   We can create "burger" since we have the ingredient "meat" and can create the ingredients "bread" and "sandwich".
// Constraints:
//   n == recipes.length == ingredients.length
//   1 <= n <= 100
//   1 <= ingredients[i].length, supplies.length <= 100
//   1 <= recipes[i].length, ingredients[i][j].length, supplies[k].length <= 10
//   recipes[i], ingredients[i][j], and supplies[k] consist only of lowercase English letters.
//   All the values of recipes and supplies combined are unique.
//   Each ingredients[i] does not contain any duplicate values.

// topological sort
func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	rmap := make(map[string]int)
	graph := make(map[string][]string)
	for _, v := range recipes {
		rmap[v] = 0
	}
	var ans []string
	for i := range ingredients {
		for j := range ingredients[i] {
			graph[ingredients[i][j]] = append(graph[ingredients[i][j]], recipes[i])
			rmap[recipes[i]] = rmap[recipes[i]] + 1
		}
	}
	for len(supplies) > 0 {
		cur := supplies[len(supplies)-1]
		supplies = supplies[:len(supplies)-1]
		for _, v := range graph[cur] {
			x := rmap[v]
			if x == 0 {
				continue
			}
			x--
			if x == 0 {
				supplies = append(supplies, v)
				ans = append(ans, v)
			}
			rmap[v] = x
		}
	}
	return ans
}

func main() {
	for _, v := range []struct {
		r      []string
		i      [][]string
		s, ans []string
	}{
		{[]string{"bread"}, [][]string{{"yeast", "flour"}}, []string{"yeast", "flour", "corn"}, []string{"bread"}},
		{[]string{"bread", "sandwich"}, [][]string{{"yeast", "flour"}, {"bread", "meat"}}, []string{"yeast", "flour", "meat"}, []string{"bread", "sandwich"}},
		{[]string{"bread", "sandwich", "burger"}, [][]string{{"yeast", "flour"}, {"bread", "meat"}, {"sandwich", "meat", "bread"}}, []string{"yeast", "flour", "meat"}, []string{"bread", "sandwich", "burger"}},
	} {
		fmt.Println(findAllRecipes(v.r, v.i, v.s), v.ans)
	}
}
