package _5_stack_queue

/*
Use a stack to track the expected closing brackets.
使用栈记录接下来期待出现的右括号。

When an opening bracket appears, push its matching closing bracket.
遇到左括号时，将与它匹配的右括号压入栈中。

When a closing bracket appears, it must match the top of the stack.
遇到右括号时，它必须与栈顶保存的期待字符相同。

After processing the whole string, the stack must be empty.
遍历完整个字符串后，栈必须为空。

I use a stack because brackets must be closed in reverse order.
When I see an opening bracket, I push its expected closing bracket.
When I see a closing bracket, it must match the top of the stack.
If the stack is empty or the characters do not match, I return false.
After processing the string, the stack must be empty.
The time complexity is \(O(n)\), and the space complexity is \(O(n)\).

我使用栈，因为括号必须按照与出现顺序相反的顺序闭合。遇到左括号时，
我把对应的右括号压入栈；遇到右括号时，它必须与栈顶期待的字符相同。
如果栈为空或者字符不匹配，就返回 false。遍历结束后，栈也必须为空。
时间复杂度是 \(O(n)\)，空间复杂度是 \(O(n)\)。
*/

func isValid(s string) bool {
	// A valid bracket string must have an even number of characters.
	// 有效括号字符串的字符数量一定是偶数。
	if len(s)%2 != 0 {
		return false
	}

	// stack stores the closing brackets expected later.
	// stack 保存之后期待遇到的右括号。
	stack := make([]byte, 0, len(s)/2)

	// Examine every bracket in the string.
	// 依次检查字符串中的每个括号。
	for i := 0; i < len(s); i++ {
		current := s[i]

		switch current {
		case '(':
			// After '(', the matching closing bracket must be ')'.
			// 遇到 '(' 后，之后必须使用 ')' 与它匹配。
			stack = append(stack, ')')

		case '[':
			// After '[', the matching closing bracket must be ']'.
			// 遇到 '[' 后，之后必须使用 ']' 与它匹配。
			stack = append(stack, ']')

		case '{':
			// After '{', the matching closing bracket must be '}'.
			// 遇到 '{' 后，之后必须使用 '}' 与它匹配。
			stack = append(stack, '}')

		default:
			// A closing bracket cannot be matched if the stack is empty.
			// 如果栈为空，当前右括号就没有对应的左括号。
			if len(stack) == 0 {
				return false
			}

			// The current closing bracket must match the expected bracket.
			// 当前右括号必须与栈顶期待的右括号相同。
			top := stack[len(stack)-1]
			if current != top {
				return false
			}

			// Remove the matched expected bracket from the stack.
			// 当前括号匹配成功，弹出栈顶元素。
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
