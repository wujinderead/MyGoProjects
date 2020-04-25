package main

import "fmt"

// https://leetcode.com/problems/parsing-a-boolean-expression/

// Return the result of evaluating a given boolean expression, represented as a string.
// An expression can either be:
//   "t", evaluating to True;
//   "f", evaluating to False;
//   "!(expr)", evaluating to the logical NOT of the inner expression expr;
//   "&(expr1,expr2,...)", evaluating to the logical AND of 2 or more inner expressions expr1, expr2, ...;
//   "|(expr1,expr2,...)", evaluating to the logical OR of 2 or more inner expressions expr1, expr2, ...
// Example 1:
//   Input: expression = "!(f)"
//   Output: true
// Example 2:
//   Input: expression = "|(f,t)"
//   Output: true
// Example 3:
//   Input: expression = "&(t,f)"
//   Output: false
// Example 4:
//   Input: expression = "|(&(t,f,t),!(t))"
//   Output: false
// Constraints:
//   1 <= expression.length <= 20000
//   expression[i] consists of characters in {'(', ')', '&', '|', '!', 't', 'f', ','}.
//   expression is a valid expression representing a boolean, as given in the description.

func parseBoolExpr(expression string) bool {
    stack := make([]byte, len(expression))
    top := 0
    for i := range expression {
		if expression[i] == '(' || expression[i] == ',' {
			continue
		}
		if expression[i] == ')' {
			j := top-1
			for stack[j]=='t' || stack[j]=='f' {
				j--
			}
			//fmt.Println(string(stack[j: top]))
			var re bool
			if stack[j]=='!' {
				re = stack[j+1] == 'f'
			} else if stack[j] == '&' {
				re = true
				for k:=j+1; k<top; k++ {
					re = re && (stack[k]=='t')
				}
			} else if stack[j] == '|' {
				re = false
				for k:=j+1; k<top; k++ {
					re = re || (stack[k]=='t')
				}
			}
			if re {
				stack[j] = 't'
			} else {
				stack[j] = 'f'
			}
			top = j+1
			continue
		}
		// else for 't', 'f', '&', '|' or '!', just push to stack
		stack[top] = expression[i]
		top++
	}
    return stack[0] == 't'
}

func main() {
    fmt.Println(parseBoolExpr("!(f)"))
	fmt.Println(parseBoolExpr("|(f,t)"))
	fmt.Println(parseBoolExpr("&(t,f)"))
	fmt.Println(parseBoolExpr("|(&(t,f,t),!(t))"))
	fmt.Println(parseBoolExpr("|(&(t,f,t),!(f))"))
}