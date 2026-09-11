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

解法一：显式栈从前向后求值（推荐） / Method 1: Explicit Stack, Forward Scan (Recommended)
使用栈保存操作数和中间结果。遇到数字时入栈；遇到运算符时先弹出右操作数，再弹出左操作数，
计算 left operator right 后将结果压回栈。
Use a stack for operands and intermediate results. Push numbers; for an operator,
pop the right operand first and then the left operand, compute left operator right, and push the result back.

关键逻辑：为什么这样做 / Why This Works
逆波兰写法是“左表达式 右表达式 运算符”，所以右表达式的值最后入栈，在栈顶；左值位于它下面。
遇到运算符，就是把这两个已经算好的子表达式合并成一个结果并压回栈。例如 ["5","2","-"] 必须算 5-2，而不是 2-5。
合法二元表达式的数字数比运算符数多 1，每个运算把栈中两个值合为一个，所以结束恰好剩一个值。
Postfix order is left-expression, right-expression, operator, so the right value is pushed later and sits above the left.
An operator combines these completed subexpressions into one stack value. Thus ["5","2","-"] means 5-2, not 2-5.
A valid binary expression has one more numeric token than operators; each operator combines two values into one,
leaving exactly one final value.

解法二：从后向前递归求值 / Method 2: Recursive Reverse Parsing
evalRPNRecursive 的游标 index 从末尾向前走，每次读取并消耗一个 token。
数字直接返回它的值；运算符先递归读取右表达式，再递归读取左表达式。
逆序读取“左、右、运算符”就是“运算符、右、左”，所以两次递归调用的顺序不能颠倒，
否则 ["5","2","-"] 会算成 2-5。
每次递归调用恰好消耗一个完整的子表达式，因此不必预先构造表达式树；
调用栈代替了解法一的显式栈，深度等于表达式的嵌套层数。
index 是闭包捕获的共享游标，两次递归调用之间必须按顺序执行，不能并发求值。
Read tokens backward with a cursor index that consumes exactly one token per call.
A number returns its value; an operator parses the right expression first, then the left.
Reversing left-right-operator gives operator-right-left, so the two recursive calls
cannot be swapped; otherwise ["5","2","-"] would evaluate to 2-5.
Each call consumes one complete subexpression, so no expression tree is needed;
the call stack replaces Method 1's explicit stack, with depth equal to the nesting depth.
index is a shared cursor captured by the closure, so the two calls must run in order, never concurrently.

时间与空间复杂度 / Time and Space Complexity
n 为 token 数，题目整数长度有界。
解法一：时间 O(n)，每个 token 只处理一次；辅助空间 O(n)，保存操作数与中间结果；返回整数 O(1)。
解法二：时间 O(n)，每个 token 也只被游标读取一次；辅助空间为调用栈 O(d)，
d 是表达式嵌套深度，最坏 O(n)（如 ["1","1","+","1","+",...]）。
若考虑任意长度数字解析，两者都要计入全部 token 字符数。默认优先解法一：递归深度不受输入控制。
n is the token count, with the problem's bounded integer-token lengths.
Method 1: time O(n), processing each token once; auxiliary space O(n) for operands
and intermediate results; the integer output uses O(1).
Method 2: time O(n), since the cursor also reads each token once; auxiliary space is the
call stack O(d) for nesting depth d, worst case O(n).
For arbitrary-length tokens, both include the total character count for parsing.
Prefer Method 1 by default, because its stack depth is heap-allocated rather than bounded by the goroutine stack.
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

// 1. 显式栈从前向后求值：推荐
func evalRPN(tokens []string) int {
	// Each binary operator reduces the stack size by one; one final value means operands=operators+1.
	// Thus (len(tokens)+1)/2 is the exact numeric-token count and an upper bound on stack size.
	// 每个二元运算使栈大小减一，最终剩一个值，因此数字数=运算符数+1。
	// 所以 (len(tokens)+1)/2 恰好是数字 token 数，也是栈大小的上界。
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

// 2. 从后向前递归求值：用调用栈代替显式栈
func evalRPNRecursive(tokens []string) int {
	// index is the shared cursor; it starts at the last token and only ever moves left.
	// index 是共享游标，从最后一个 token 开始，且只向左移动。
	index := len(tokens) - 1

	// parse consumes exactly one complete subexpression and returns its value.
	// parse 恰好消耗一个完整子表达式，并返回它的值。
	var parse func() int

	parse = func() int {
		// Read the token under the cursor, then move past it.
		// 读取游标处的 token，然后让游标越过它。
		token := tokens[index]
		index--

		switch token {
		case "+", "-", "*", "/":
			// Reverse postfix order is operator, right, left: the right operand comes first.
			// 逆序读取后缀表达式的顺序是“运算符、右、左”，所以先得到右操作数。
			right := parse()

			// Only after the whole right subexpression is consumed does the left one begin.
			// 只有整个右子表达式都被消耗完，左子表达式才开始。
			left := parse()

			switch token {
			case "+":
				return left + right
			case "-":
				// Operand order matters: this must be left-right, not right-left.
				// 操作数顺序有意义：必须是 left-right，不能写成 right-left。
				return left - right
			case "*":
				return left * right
			default:
				// Division is the only remaining valid operator and truncates toward zero.
				// 唯一剩余的合法运算符是除法，Go 的整数除法向零截断。
				return left / right
			}

		default:
			// Every non-operator token is guaranteed to be a valid integer.
			// 题目保证所有非运算符 token 都是合法整数。
			value, _ := strconv.Atoi(token)
			return value
		}
	}

	// The outermost call consumes the entire expression.
	// 最外层调用消耗整个表达式。
	return parse()
}
