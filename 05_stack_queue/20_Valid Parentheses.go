package _5_stack_queue

/*
20. 有效的括号 / Valid Parentheses

题目描述 / Problem Description
给定一个只包含圆括号、方括号和花括号的字符串 s，判断所有括号是否类型匹配、
闭合顺序正确，并且每个右括号都有对应的左括号。
Given a string s containing only parentheses, square brackets,
and braces, determine whether every bracket has the correct type,
closing order, and matching partner.
*/

// 1. Stack of expected closing brackets: push the closer an opening bracket demands, then compare bytes directly.
// 1. 栈保存期待的右括号：直接比较；入栈时换成对应的右括号，之后只做一次字节相等判断。
// Time: O(n), Space: O(n) for the stack.
// 时间复杂度：O(n)，空间复杂度：O(n)，由栈产生。
//
// 栈存“等待的右括号”，最近打开的括号必须最先闭合，所以当前右括号只能匹配栈顶。
// Store expected closing brackets; the most recently opened pair must close first, so compare only with the top.
// 遇到空栈或类型不符立即失败，结束还剩期待项也失败；输入仅含六种括号，奇数长度可直接排除。
// Reject an empty/mismatched stack or leftover expectations; the input contains only six bracket characters, so odd lengths fail.
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := make([]byte, 0, len(s)/2)

	for i := 0; i < len(s); i++ {
		current := s[i]

		switch current {
		case '(':
			stack = append(stack, ')')

		case '[':
			stack = append(stack, ']')

		case '{':
			stack = append(stack, '}')

		default:
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			if current != top {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

// 2. Hash map lookup: push opening brackets unchanged, and map each closing bracket back to the opener it expects.
// 2. 哈希表映射：压入左括号本身，闭括号时用映射表换回它期待的左括号。
// Time: O(n) average, Space: O(n) for the stack plus O(1) for the three-entry table.
// 时间复杂度：平均 O(n)，空间复杂度：O(n)，由栈产生，外加三条映射表的 O(1)。
//
// 栈存未闭合的左括号，pairs 将右括号映射到应匹配的左括号；只能消除栈顶，不能在栈中任意寻找。
// Store unclosed opening brackets; pairs maps each closing bracket to the required top, not any earlier stack entry.
// 输入仅含六种括号；结束栈为空才说明每个左括号都按嵌套顺序闭合。
// For the six bracket characters, an empty final stack means every opening closed in nesting order.
func isValidWithMap(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := make([]byte, 0, len(s)/2)

	for i := 0; i < len(s); i++ {
		current := s[i]

		expected, isClosing := pairs[current]
		if !isClosing {
			stack = append(stack, current)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		if stack[len(stack)-1] != expected {
			return false
		}

		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}
