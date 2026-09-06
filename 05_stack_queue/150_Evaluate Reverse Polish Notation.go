package _5_stack_queue

/*
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

解题思路 / Solution Approach
使用栈保存操作数和中间结果。遇到数字时入栈；遇到运算符时先弹出右操作数，再弹出左操作数，
计算 left operator right 后将结果压回栈。
Use a stack for operands and intermediate results. Push numbers; for an operator,
pop the right operand first and then the left operand, compute left operator right, and push the result back.

时间与空间复杂度 / Time and Space Complexity
n 为 token 数。题目整数长度有界时，时间 O(n)，每个 token 只处理一次；
若考虑任意长度数字解析，则计入全部 token 字符数。辅助空间 O(n)，保存操作数与中间结果；返回整数 O(1)。
n is the token count. Time O(n) under the problem's bounded integer-token lengths;
for arbitrary-length tokens, include the total character count for parsing.
Auxiliary space O(n) stores operands and intermediate results; the integer output uses O(1).
*/

import "strconv"

/*
Use a stack to evaluate the Reverse Polish Notation expression.
使用栈计算逆波兰表达式。

When a number appears, convert it to an integer and push it.
遇到数字时，将它转换成整数并压入栈。

When an operator appears, pop the right and left operands,
calculate the result, and push the result back.
遇到运算符时，依次取出右操作数和左操作数，
完成计算后再将结果压回栈。

The only value remaining in the stack is the final answer.
最后栈中唯一剩余的数字就是最终答案。
*/

func evalRPN(tokens []string) int {
	// At most about half of the tokens are operands.
	// 操作数最多约占 token 数量的一半。
	stack := make([]int, 0, (len(tokens)+1)/2)

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			// The stack top is the right operand.
			// 栈顶元素是右操作数。
			n := len(stack)
			left := stack[n-2]
			right := stack[n-1]

			// Remove both operands before pushing the result.
			// 先删除两个操作数，再压入计算结果。
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
				// Division is the only remaining valid operator.
				// 唯一剩余的合法运算符是除法。
				result = left / right
			}

			// Push the intermediate result only once.
			// 统一将中间结果压入栈中。
			stack = append(stack, result)

		default:
			// Every non-operator token is guaranteed to be a valid integer.
			// 题目保证所有非运算符 token 都是合法整数。
			number, _ := strconv.Atoi(token)
			stack = append(stack, number)
		}
	}

	// A valid expression leaves exactly one result.
	// 合法表达式最终只留下一个结果。
	return stack[0]
}
