package _5_stack_queue

/*
150. 逆波兰表达式求值 / Evaluate Reverse Polish Notation

题目描述 / Problem Description
给定一个字符串数组 tokens，它表示一个合法的逆波兰表达式，
请计算并返回表达式的整数结果。
Given a string array tokens representing a valid Reverse
Polish Notation expression, evaluate it and return the integer result.

表达式支持四种运算符：
The expression supports four operators:
+  -  *  /
需要注意：
Important details:
- 每个操作数可以是整数，也可以是另一个表达式的计算结果。
  Each operand can be an integer or the result of another expression.
- 整数除法向零截断。
  Integer division truncates toward zero.
- 表达式中不会出现除以零。
  The expression does not contain division by zero.
- 所有输入都构成合法的逆波兰表达式。
  Every input is a valid Reverse Polish Notation expression.
*/

import "strconv"

// 1. Explicit stack, forward scan (recommended): push numbers, and each operator pops the right operand then the left.
// 1. 显式栈从前向后求值：推荐；数字入栈，遇到运算符先弹出右操作数再弹出左操作数。
// Time: O(n), Space: O(n) for the operand stack.
// 时间复杂度：O(n)，空间复杂度：O(n)，由操作数栈产生。
//
// 栈保存尚未被上层运算消费的子表达式值；运算符取栈顶两项，较早入栈的是左操作数。
// The stack holds completed subexpression values; an operator consumes two, with the earlier entry as its left operand.
// 顺序必须是 left-right、left/right；Go 整数除法向零截断。按题意保证表达式合法、除数非零且结果不溢出。
// Preserve left/right order for subtraction and division; Go truncates division toward zero. Assume valid, nonoverflowing expressions and nonzero divisors.
func evalRPN(tokens []string) int {
	stack := make([]int, 0, (len(tokens)+1)/2)

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			n := len(stack)
			left := stack[n-2]
			right := stack[n-1]

			stack = stack[:n-2]

			var result int

			switch token {
			case "+":
				result = left + right
			case "-":
				result = left - right
			case "*":
				result = left * right
			default:
				result = left / right
			}

			stack = append(stack, result)

		default:
			number, _ := strconv.Atoi(token)
			stack = append(stack, number)
		}
	}

	return stack[0]
}

// 2. Recursive reverse parsing: use the call stack instead of an explicit stack, parsing the right subexpression first.
// 2. 从后向前递归求值：用调用栈代替显式栈，先递归右子表达式再递归左子表达式。
// Time: O(n), Space: O(d) for the recursion stack, where d is the nesting depth, worst case O(n).
// 时间复杂度：O(n)，空间复杂度：O(d)，由递归调用栈产生，d 为表达式嵌套深度，最坏 O(n)。
//
// 后缀表达式从尾部读时先遇运算符，再遇右子表达式，最后才是左子表达式，因此先 parse right 再 parse left。
// Reading postfix backward encounters the operator, then its right expression, then its left; parse in that order.
// 共享 index 每读一项减一，每次递归返回一棵子表达式的值；沿用合法表达式、非零除数及不溢出的约束。
// The shared index consumes each token once; each call returns a subexpression value under the valid-expression contract.
func evalRPNRecursive(tokens []string) int {
	index := len(tokens) - 1

	var parse func() int

	parse = func() int {
		token := tokens[index]
		index--

		switch token {
		case "+", "-", "*", "/":
			right := parse()

			left := parse()

			switch token {
			case "+":
				return left + right
			case "-":
				return left - right
			case "*":
				return left * right
			default:
				return left / right
			}

		default:
			value, _ := strconv.Atoi(token)
			return value
		}
	}

	return parse()
}
