package main

import "fmt"

// https://leetcode.com/problems/process-restricted-friend-requests/

// You are given an integer n indicating the number of people in a network. Each person is labeled
// from 0 to n - 1.
// You are also given a 0-indexed 2D integer array restrictions, where restrictions[i] = [xi, yi]
// means that person xi and person yi cannot become friends, either directly or indirectly through
// other people.
// Initially, no one is friends with each other. You are given a list of friend requests as a 0-indexed
// 2D integer array requests, where requests[j] = [uj, vj] is a friend request between person uj and
// person vj.
// A friend request is successful if uj and vj can be friends. Each friend request is processed in the
// given order (i.e., requests[j] occurs before requests[j + 1]), and upon a successful request,
// uj and vj become direct friends for all future friend requests.
// Return a boolean array result, where each result[j] is true if the jth friend request is successful
// or false if it is not.
// Note: If uj and vj are already direct friends, the request is still successful.
// Example 1:
//   Input: n = 3, restrictions = [[0,1]], requests = [[0,2],[2,1]]
//   Output: [true,false]
//   Explanation:
//     Request 0: Person 0 and person 2 can be friends, so they become direct friends .
//     Request 1: Person 2 and person 1 cannot be friends since person 0 and person 1 would be indirect friends (1--2--0).
// Example 2:
//   Input: n = 3, restrictions = [[0,1]], requests = [[1,2],[0,2]]
//   Output: [true,false]
//   Explanation:
//     Request 0: Person 1 and person 2 can be friends, so they become direct friends .
//     Request 1: Person 0 and person 2 cannot be friends since person 0 and person 1 would be indirect friends (0--2--1).
// Example 3:
//   Input: n = 5, restrictions = [[0,1],[1,2],[2,3]], requests = [[0,4],[1,2],[3,1],[3,4]]
//   Output: [true,false,true,false]
//   Explanation:
//     Request 0: Person 0 and person 4 can be friends, so they become direct friends.
//     Request 1: Person 1 and person 2 cannot be friends since they are directly restricted.
//     Request 2: Person 3 and person 1 can be friends, so they become direct friends .
//     Request 3: Person 3 and person 4 cannot be friends since person 0 and person 1 would be indirect friends (0--4--3--1).
// Constraints:
//   2 <= n <= 1000
//   0 <= restrictions.length <= 1000
//   restrictions[i].length == 2
//   0 <= xi, yi <= n - 1
//   xi != yi
//   1 <= requests.length <= 1000
//   requests[j].length == 2
//   0 <= uj, vj <= n - 1
//   uj != vj

// use union-find, time complexity: len(restrictions)*len(requests)*logn
// for example, we have a two sets like these, and 1 and 4 is the root for these sets,
// if we have a request to connect 2 and 5, we find its roots are 1 and 4.
// then we process all restrictions, if we have 3 and 6 restricted, and we found their roots are also 1 and 4.
// so we can't connect 2 and 5.
//   1   4
//   |   |
//   2   5
//   |   |
//   3   6
func friendRequests(n int, restrictions [][]int, requests [][]int) []bool {
	root := make([]int, n)
	for i := range root {
		root[i] = -1
	}
	ans := make([]bool, len(requests))
	for i, r := range requests {
		r0 := getRoot(root, r[0])
		r1 := getRoot(root, r[1])
		ans[i] = true // default to true
		if r0 == r1 { // already same set
			continue
		}
		for _, rest := range restrictions {
			rr0 := getRoot(root, rest[0])
			rr1 := getRoot(root, rest[1])
			if (r0 == rr0 && r1 == rr1) || (r1 == rr0 && r0 == rr1) {
				ans[i] = false
				break
			}
		}
		if ans[i] {
			root[r1] = r0
		}
	}
	return ans
}

func getRoot(root []int, i int) int {
	if root[i] != -1 {
		x := getRoot(root, root[i])
		root[i] = x
		return x
	}
	return i
}

func main() {
	for _, v := range []struct {
		n         int
		rest, req [][]int
		ans       []bool
	}{
		{3, [][]int{{0, 1}}, [][]int{{0, 2}, {2, 1}}, []bool{true, false}},
		{3, [][]int{{0, 1}}, [][]int{{1, 2}, {0, 2}}, []bool{true, false}},
		{5, [][]int{{0, 1}, {1, 2}, {2, 3}}, [][]int{{0, 4}, {1, 2}, {3, 1}, {3, 4}}, []bool{true, false, true, false}},
	} {
		fmt.Println(friendRequests(v.n, v.rest, v.req), v.ans)
	}
	n := 53
	rest := [][]int{{44, 5}, {51, 50}, {38, 9}, {38, 31}, {22, 16}, {39, 42}, {30, 22}, {19, 26}, {40, 48}, {45, 31}, {37, 22}, {16, 28}, {32, 18}, {31, 34}, {44, 23}, {5, 52}, {8, 26}, {13, 27}, {14, 50}, {24, 37}, {32, 41}, {47, 34}, {13, 5}, {36, 26}, {35, 4}, {43, 0}, {23, 13}, {20, 44}, {6, 23}, {9, 32}, {2, 18}, {1, 50}, {22, 17}, {27, 0}, {48, 34}, {20, 38}, {32, 48}, {3, 9}, {25, 44}, {47, 29}, {1, 17}, {29, 26}, {13, 21}, {10, 23}, {21, 12}, {41, 50}, {45, 24}, {46, 11}, {15, 22}, {49, 45}, {10, 8}, {1, 42}, {11, 15}, {10, 45}, {33, 43}, {14, 1}, {40, 19}, {35, 15}, {46, 49}, {14, 38}, {20, 21}, {33, 35}, {6, 29}, {20, 31}, {5, 10}, {26, 51}, {2, 46}, {37, 3}, {42, 30}, {10, 52}, {36, 41}, {50, 48}, {7, 41}, {9, 28}, {48, 51}, {43, 31}, {44, 17}, {43, 35}, {31, 10}, {43, 22}, {41, 19}, {29, 5}, {39, 46}, {15, 13}, {45, 1}, {10, 43}, {37, 41}, {25, 9}, {6, 14}, {52, 41}, {15, 12}, {52, 21}, {46, 12}, {40, 24}, {48, 33}, {1, 24}, {15, 7}, {37, 36}, {11, 38}, {21, 31}, {39, 21}, {40, 46}, {3, 52}, {36, 17}, {3, 4}, {50, 37}, {31, 48}, {34, 41}, {24, 51}, {44, 9}, {12, 24}, {46, 18}, {1, 0}, {50, 42}, {41, 12}, {44, 7}, {35, 36}, {20, 45}, {24, 52}, {24, 19}, {46, 50}, {36, 34}, {33, 7}, {2, 24}, {31, 40}, {20, 5}, {25, 37}, {52, 22}, {28, 50}, {52, 16}, {25, 36}, {44, 40}, {42, 2}, {51, 21}, {16, 4}, {7, 27}, {43, 48}, {6, 50}, {2, 10}, {2, 34}, {37, 16}, {7, 49}, {39, 9}, {0, 28}, {21, 33}, {5, 18}, {13, 19}, {33, 39}, {19, 25}, {13, 16}, {37, 33}, {14, 11}, {42, 26}, {17, 2}, {23, 40}, {18, 15}, {38, 16}, {47, 12}, {29, 3}, {46, 13}, {50, 8}, {15, 5}, {26, 33}, {16, 12}, {25, 42}, {48, 36}, {4, 26}, {20, 52}, {51, 32}, {13, 42}, {16, 17}, {16, 24}, {21, 5}, {40, 49}, {39, 37}, {25, 43}, {29, 19}, {42, 43}, {37, 28}, {12, 26}, {4, 52}, {8, 44}, {0, 19}, {51, 16}, {3, 46}, {35, 23}, {44, 19}, {15, 39}, {17, 19}, {38, 50}, {42, 19}, {44, 46}, {8, 38}, {10, 26}, {49, 27}, {47, 18}, {31, 0}, {37, 7}, {34, 16}, {21, 1}, {51, 22}, {48, 9}, {26, 22}, {42, 29}, {18, 6}, {36, 38}, {33, 30}, {5, 22}, {21, 7}, {26, 37}, {11, 23}, {13, 8}, {11, 24}, {40, 3}, {17, 6}, {12, 14}, {28, 2}, {33, 31}, {2, 16}, {5, 36}, {20, 18}, {14, 36}, {24, 39}, {36, 44}, {47, 39}, {41, 29}, {33, 1}, {7, 5}, {28, 33}, {16, 39}, {23, 52}, {19, 35}, {35, 6}, {29, 18}, {6, 31}, {51, 38}, {28, 1}, {37, 13}, {8, 25}, {19, 46}, {32, 29}, {15, 9}, {7, 3}, {0, 37}, {51, 12}, {19, 12}, {7, 43}, {52, 19}, {27, 30}, {12, 9}, {32, 39}, {49, 9}, {50, 32}, {12, 32}, {17, 38}, {39, 19}, {31, 23}, {32, 46}, {3, 27}, {29, 34}, {32, 6}, {25, 46}, {37, 49}, {5, 39}, {23, 41}, {2, 1}, {25, 10}, {44, 15}, {18, 39}, {42, 12}, {31, 5}, {27, 22}, {26, 17}, {26, 1}, {51, 14}, {20, 36}, {2, 49}, {6, 36}, {18, 50}, {48, 4}, {3, 10}, {40, 18}, {41, 21}, {7, 30}, {45, 0}, {33, 17}, {22, 32}, {31, 3}, {48, 41}, {1, 6}, {51, 23}, {15, 48}, {7, 1}, {5, 34}, {33, 2}, {45, 17}, {28, 14}, {31, 22}, {19, 21}, {40, 30}, {2, 11}, {35, 21}, {12, 30}, {8, 5}, {50, 19}, {36, 18}, {23, 29}, {39, 41}, {34, 15}, {35, 34}, {31, 35}, {6, 47}, {50, 7}, {45, 33}, {44, 10}, {11, 47}, {15, 50}, {29, 4}, {25, 48}, {22, 48}, {43, 39}, {35, 26}, {18, 22}, {18, 4}, {5, 50}, {51, 30}, {6, 13}, {42, 8}, {30, 17}, {43, 4}, {22, 38}, {25, 23}, {4, 15}, {27, 31}, {50, 21}, {11, 13}, {15, 45}, {10, 12}, {36, 46}, {35, 12}, {50, 35}, {30, 20}, {16, 35}, {33, 6}, {28, 17}, {20, 8}, {22, 9}, {41, 43}, {45, 13}, {49, 38}, {45, 37}, {41, 26}, {0, 10}, {28, 38}, {25, 50}, {13, 0}, {28, 52}, {6, 21}, {44, 34}, {18, 24}, {47, 17}, {27, 26}, {5, 47}, {8, 47}, {30, 37}, {3, 38}, {3, 49}, {34, 43}, {18, 35}, {14, 34}, {21, 47}, {38, 40}, {39, 44}, {32, 0}, {26, 16}, {50, 43}, {50, 16}, {29, 11}, {37, 42}, {27, 52}, {26, 13}, {11, 39}, {41, 9}, {38, 48}, {28, 29}, {30, 32}, {33, 34}, {40, 37}, {19, 16}, {11, 40}, {32, 19}, {16, 43}, {25, 49}, {5, 11}, {40, 25}, {10, 19}, {48, 6}, {29, 20}, {2, 19}, {27, 17}, {16, 47}, {20, 27}, {27, 43}, {7, 0}, {33, 47}, {25, 38}, {51, 49}, {24, 26}, {51, 1}, {17, 25}, {8, 46}, {47, 20}, {2, 26}, {36, 39}, {21, 29}, {38, 4}, {2, 50}, {24, 31}, {35, 13}, {5, 28}, {51, 27}, {42, 44}, {23, 15}, {8, 27}, {50, 26}, {29, 37}, {12, 34}, {52, 6}, {50, 40}, {51, 39}, {5, 1}, {24, 48}, {48, 49}, {20, 46}, {11, 25}, {10, 33}, {48, 5}, {15, 10}, {1, 52}, {14, 33}, {45, 34}, {6, 44}, {31, 7}, {17, 29}, {25, 7}, {9, 34}, {6, 45}, {52, 49}, {1, 3}, {38, 34}, {23, 36}, {45, 14}, {33, 3}, {14, 13}, {28, 51}, {32, 33}, {2, 21}, {18, 9}, {52, 25}, {12, 27}, {8, 21}, {6, 2}, {12, 17}, {33, 4}, {9, 5}, {37, 51}, {49, 26}, {38, 24}, {28, 43}, {32, 4}, {4, 37}, {22, 7}, {45, 21}, {45, 8}, {45, 36}, {49, 41}, {26, 6}, {33, 5}, {37, 38}, {13, 41}, {20, 26}, {35, 3}, {51, 46}, {49, 6}, {13, 29}, {15, 38}, {2, 45}, {8, 6}, {16, 0}, {17, 35}, {0, 17}, {50, 44}, {0, 6}, {40, 8}, {10, 11}, {3, 26}, {1, 12}, {37, 8}, {28, 6}, {4, 20}, {23, 22}, {44, 4}, {10, 27}, {24, 23}, {50, 0}, {14, 3}, {1, 46}, {18, 3}, {16, 40}, {23, 2}, {50, 52}, {25, 32}, {32, 16}, {14, 9}, {12, 23}, {13, 36}, {40, 10}, {27, 36}, {8, 2}, {45, 50}, {11, 44}, {8, 19}, {40, 29}, {27, 42}, {31, 8}, {0, 22}, {8, 28}, {48, 23}, {19, 14}, {0, 15}, {51, 52}, {11, 7}, {47, 15}, {47, 0}, {32, 2}, {12, 29}, {5, 14}, {32, 36}, {31, 17}, {40, 51}, {44, 43}, {44, 35}, {15, 26}, {39, 34}, {44, 3}, {42, 9}, {21, 36}, {22, 39}, {41, 44}, {50, 39}, {20, 7}, {41, 40}, {22, 11}, {33, 44}, {5, 17}, {52, 31}, {12, 45}, {49, 43}, {2, 48}, {43, 24}, {24, 25}, {40, 15}, {22, 20}, {19, 36}, {11, 51}, {46, 17}, {37, 5}, {28, 45}, {36, 40}, {0, 30}, {12, 7}, {23, 19}, {5, 38}, {21, 43}, {34, 40}, {43, 40}, {10, 41}, {14, 22}, {6, 46}, {48, 29}, {28, 22}, {46, 29}, {1, 16}, {18, 27}, {25, 0}, {47, 3}, {39, 49}, {4, 28}, {31, 25}, {44, 22}, {34, 13}, {38, 46}, {6, 15}, {19, 18}, {9, 7}, {49, 10}, {51, 34}, {50, 36}, {22, 12}, {6, 12}, {32, 15}, {20, 25}, {29, 31}, {36, 9}, {20, 33}, {36, 0}, {21, 3}, {7, 26}, {8, 36}, {33, 51}, {49, 19}, {36, 42}, {25, 15}, {43, 52}, {52, 15}, {11, 18}, {15, 30}, {40, 52}, {2, 38}, {19, 51}, {23, 47}, {20, 43}, {0, 4}, {28, 47}, {14, 26}, {1, 44}, {9, 6}, {46, 52}, {1, 9}, {17, 40}, {11, 21}, {31, 15}, {32, 24}, {32, 47}, {22, 25}, {14, 41}, {13, 12}, {40, 33}, {39, 40}, {46, 37}, {52, 29}, {15, 3}, {34, 4}, {49, 30}, {40, 47}, {26, 44}, {30, 39}, {3, 12}, {28, 42}, {39, 20}, {22, 33}, {14, 29}, {29, 10}, {43, 37}, {35, 41}}
	req := [][]int{{32, 31}, {22, 4}, {6, 31}, {52, 36}, {1, 25}, {14, 1}, {1, 47}, {33, 19}, {22, 32}, {52, 15}, {45, 11}, {35, 4}, {26, 17}, {46, 7}, {29, 3}, {52, 36}, {36, 0}, {30, 4}, {6, 5}, {21, 12}, {17, 4}, {42, 44}, {29, 4}, {50, 26}, {28, 10}, {4, 51}, {15, 3}, {13, 16}, {17, 11}, {46, 12}, {34, 42}, {18, 24}, {41, 9}, {32, 19}, {29, 20}, {7, 5}, {38, 48}, {47, 20}, {40, 21}, {14, 50}, {50, 47}, {27, 42}, {30, 26}, {21, 36}, {17, 4}, {50, 39}, {14, 4}, {21, 18}, {8, 41}, {0, 22}, {3, 36}, {11, 28}, {45, 18}, {23, 26}, {44, 12}, {33, 13}, {11, 21}, {18, 42}, {26, 9}, {4, 15}, {41, 6}, {52, 15}, {2, 13}, {46, 23}, {15, 20}, {4, 5}, {45, 48}, {37, 3}, {37, 48}, {43, 9}, {35, 49}, {51, 20}, {39, 3}, {39, 35}, {10, 17}, {1, 30}, {33, 2}, {15, 21}, {40, 5}, {19, 47}, {41, 15}, {14, 17}, {30, 9}, {4, 49}, {40, 24}, {30, 18}, {45, 3}, {8, 27}, {9, 16}, {2, 1}, {21, 17}, {52, 2}, {23, 15}, {11, 12}, {7, 3}, {25, 42}, {28, 43}, {36, 29}, {25, 9}, {42, 16}, {31, 23}, {51, 30}, {51, 46}, {11, 48}, {32, 6}, {48, 36}, {33, 39}, {32, 31}, {1, 15}, {10, 6}, {0, 37}, {45, 24}, {13, 42}, {36, 41}, {40, 35}, {47, 0}, {5, 45}, {6, 29}, {51, 9}, {44, 12}, {35, 49}, {23, 2}, {35, 27}, {20, 34}, {22, 16}, {45, 24}, {10, 19}, {31, 49}, {3, 19}, {37, 19}, {45, 0}, {18, 39}, {27, 30}, {5, 22}, {27, 11}, {16, 39}, {5, 25}, {13, 16}, {15, 22}, {49, 13}, {32, 39}, {3, 23}, {0, 3}, {22, 39}, {17, 40}, {4, 39}, {40, 37}, {19, 48}, {27, 40}, {41, 2}, {3, 8}, {1, 9}, {41, 22}, {24, 30}, {8, 36}, {13, 30}, {33, 41}, {26, 17}, {52, 13}, {29, 7}, {18, 17}, {28, 21}, {21, 5}, {48, 8}, {36, 51}, {45, 17}, {48, 36}, {23, 7}, {9, 16}, {44, 2}, {39, 7}, {39, 1}, {36, 17}, {44, 9}, {44, 2}, {16, 31}, {0, 6}, {39, 46}, {35, 48}, {12, 45}, {48, 18}, {19, 15}, {47, 25}, {50, 19}, {25, 3}, {44, 46}, {35, 37}, {3, 13}, {39, 40}, {22, 12}, {6, 13}, {11, 38}, {49, 26}, {44, 21}, {37, 17}, {47, 3}, {1, 24}, {1, 20}, {12, 9}, {9, 31}, {5, 39}, {22, 4}, {33, 16}, {52, 6}, {16, 43}, {4, 39}, {14, 39}, {4, 13}, {4, 47}, {45, 31}, {40, 46}, {5, 41}, {23, 45}, {51, 12}, {26, 9}, {29, 16}, {11, 36}, {29, 3}, {28, 36}, {1, 3}, {5, 38}, {21, 17}, {12, 26}, {16, 30}, {30, 46}, {26, 11}, {14, 13}, {41, 50}, {13, 12}, {22, 24}, {27, 4}, {34, 40}, {18, 27}, {49, 29}, {49, 6}, {51, 2}, {27, 40}, {41, 29}, {21, 12}, {20, 46}, {3, 38}, {29, 34}, {35, 45}, {24, 26}, {6, 43}, {24, 3}, {23, 16}, {14, 20}, {49, 11}, {20, 21}, {0, 24}, {11, 48}, {7, 32}, {5, 23}, {13, 41}, {44, 2}, {49, 26}, {43, 48}, {13, 29}, {32, 6}, {43, 3}, {38, 32}, {31, 36}, {37, 9}, {33, 3}, {6, 29}, {6, 24}, {28, 38}, {18, 31}, {32, 39}, {20, 24}, {46, 16}, {33, 9}, {27, 19}, {34, 3}, {37, 15}, {13, 16}, {19, 47}, {18, 0}, {46, 14}, {0, 46}, {42, 15}, {9, 31}, {22, 7}, {24, 3}, {20, 45}, {19, 12}, {23, 17}, {44, 18}, {14, 29}, {15, 7}, {30, 20}, {26, 40}, {4, 49}, {16, 28}, {9, 10}, {1, 48}, {47, 29}, {47, 51}, {20, 34}, {2, 50}, {41, 28}, {51, 25}, {29, 9}, {4, 11}, {20, 2}, {19, 16}, {27, 39}, {41, 44}, {31, 22}, {32, 16}, {20, 32}, {27, 40}, {40, 21}, {51, 49}, {36, 17}, {26, 11}, {21, 17}, {6, 34}, {4, 8}, {35, 24}, {3, 4}, {0, 2}, {32, 6}, {14, 37}, {31, 50}, {49, 23}, {34, 16}, {27, 19}, {40, 9}, {7, 18}, {18, 26}, {10, 52}, {31, 36}, {47, 39}, {3, 19}, {37, 41}, {15, 30}, {40, 0}, {28, 43}, {5, 34}, {15, 2}, {35, 5}, {20, 21}, {4, 37}, {43, 40}, {25, 33}, {37, 28}, {42, 8}, {42, 8}, {19, 43}, {46, 49}, {38, 24}, {24, 42}, {39, 37}, {23, 26}, {49, 41}, {36, 2}, {25, 9}, {50, 19}, {12, 8}, {25, 10}, {8, 25}, {51, 52}, {38, 48}, {50, 47}, {50, 8}, {29, 35}, {17, 9}, {48, 17}, {10, 50}, {12, 33}, {24, 51}, {28, 29}, {30, 47}, {43, 47}, {44, 0}, {2, 7}, {6, 45}, {23, 28}, {24, 50}, {35, 34}, {11, 33}, {34, 0}, {16, 48}, {11, 47}, {42, 33}, {1, 17}, {8, 24}, {25, 10}, {0, 22}, {45, 33}, {43, 39}, {40, 49}, {16, 6}, {19, 26}, {5, 18}, {37, 28}, {5, 27}, {23, 50}, {39, 3}, {42, 41}, {5, 41}, {43, 4}, {10, 12}, {15, 22}, {7, 1}, {35, 12}, {40, 33}, {33, 43}, {42, 48}, {11, 12}, {49, 43}, {40, 0}, {25, 0}, {20, 13}, {21, 43}, {50, 21}, {45, 37}, {23, 21}, {45, 13}, {25, 9}, {43, 39}, {6, 11}, {45, 4}, {36, 38}, {33, 5}, {44, 46}, {36, 18}, {29, 31}, {18, 9}, {37, 19}, {36, 44}, {22, 40}, {22, 9}, {3, 36}}
	real := []bool{true, true, false, true, true, false, true, true, false, false, true, false, false, true, false, true, false, false, true, false, false, false, false, false, true, false, false, false, false, false, true, false, false, false, false, false, false, false, true, false, false, false, true, false, false, false, false, false, true, false, false, false, false, true, true, false, false, true, true, false, false, false, true, false, true, false, true, false, false, true, true, true, true, false, false, false, false, false, false, false, false, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false, false, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, true, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}
	ans := friendRequests(n, rest, req)
	for i := range real {
		if real[i] != ans[i] {
			fmt.Println(i, real[i])
		}
	}
}
