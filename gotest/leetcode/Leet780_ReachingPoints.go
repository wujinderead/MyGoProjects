package main

import "fmt"

// https://leetcode.com/problems/reaching-points

// A move consists of taking a point (x, y) and transforming it to either
// (x, x+y) or (x+y, y). Given a starting point (sx, sy) and a target point
// (tx, ty), return True if and only if a sequence of moves exists to transform
// the point (sx, sy) to (tx, ty). Otherwise, return False.
// Example1:
//   Input: sx = 1, sy = 1, tx = 3, ty = 5
//   Output: True
//   Explanation:
//     One series of moves that transforms the starting point to the target is:
//     (1, 1) -> (1, 2)
//     (1, 2) -> (3, 2)
//     (3, 2) -> (3, 5)
// Example2:
//   Input: sx = 1, sy = 1, tx = 2, ty = 2
//   Output: False
// Example3:
//   Input: sx = 1, sy = 1, tx = 1, ty = 1
//   Output: True
// Note:
//   sx, sy, tx, ty will all be integers in the range [1, 10^9].

// for (a, b) is from (a, b-a) if b>a; from (a-b, b) if a>b.
// if a>>b, (a, b) is from (a-kb, b), k increase until a-kb<=b
func reachingPoints(sx int, sy int, tx int, ty int) bool {
	for {
		if sx == tx && sy == ty {
			return true
		}
		if (tx == ty && sx != tx) || sx > tx || sy > ty {
			return false
		}
		if tx > ty && ty == sy {
			if (tx-sx)%ty == 0 {
				return true
			}
			return false
		}
		if tx > ty {
			if tx%ty == 0 {
				tx = ty
			} else {
				tx = tx % ty
			}
			continue
		}
		if tx < ty && tx == sx {
			if (ty-sy)%tx == 0 {
				return true
			}
			return false
		}
		if tx < ty {
			if ty%tx == 0 {
				ty = tx
			} else {
				ty = ty % tx
			}
			continue
		}
	}
}

func main() {
	fmt.Println(reachingPoints(1, 1, 3, 5))
	fmt.Println(reachingPoints(1, 1, 2, 2))
	fmt.Println(reachingPoints(1, 1, 1, 1))
}
