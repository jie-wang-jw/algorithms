package _5_stack_queue

/*
20. 有效的括号 / Valid Parentheses

题目描述 / Problem Description
给定一个只包含圆括号、方括号和花括号的字符串 s，判断所有括号是否类型匹配、
闭合顺序正确，并且每个右括号都有对应的左括号。
Given a string s containing only parentheses, square brackets,
and braces, determine whether every bracket has the correct type,
closing order, and matching partner.

解法一：栈保存期待的右括号 / Method 1: Stack of Expected Closing Brackets
使用栈保存每个左括号所期待的右括号。遇到右括号时，它必须等于栈顶；
如果栈为空、类型不匹配，或遍历结束后栈不为空，则字符串无效。
这样右括号只需与栈顶做一次字节相等比较，不必再查表还原括号类型。
Use a stack to store the closing bracket expected for each opening bracket.
A closing bracket must match the top; an empty stack, mismatch,
or leftover stack makes the string invalid.
A closing bracket then needs only one byte comparison against the top, with no lookup to recover its type.

关键逻辑：为什么这样做 / Why This Works
最近打开且尚未闭合的括号必须最先闭合，否则就会交叉嵌套，所以使用后进先出的栈。例如 "([)]" 在读到 ')' 时，
最近未闭合的是 '['，期待 ']'；即使两种括号数量都配对，也必须判错。结束后栈不空表示有左括号未闭合；
中途栈空却遇到右括号表示没有可配对的左括号。
The most recently opened unmatched bracket must close first; otherwise pairs cross.
That requires LIFO order. In "([)]", ')' encounters an unmatched '[' expecting ']';
balanced counts alone are insufficient. A nonempty final stack means unclosed openings;
a closing bracket with an empty stack has no matching opening.

解法二：哈希表映射，压入左括号 / Method 2: Hash Map, Pushing Opening Brackets
isValidWithMap 把 map[byte]byte 作为“右括号 -> 它期待的左括号”的映射，栈里保存左括号原样。
判定某个字符是不是右括号，就是查这张表：查不到说明它是左括号，直接入栈；
查到 expected 则要求栈顶恰好等于 expected。
两种解法的判定完全等价，只是把“开括号时换成右括号”改成“闭括号时换回左括号”。
表只有三条固定条目，可以写成局部 map；换成 switch 或数组同样成立，这里保留 map 以对照代码随想录的写法。
isValidWithMap uses a map[byte]byte from each closing bracket to the opening bracket it expects,
and the stack holds opening brackets unchanged.
A map lookup classifies the character: a miss means an opening bracket, which is pushed as is;
a hit yields expected, and the stack top must equal expected.
The two methods are equivalent; the translation just moves from push time to pop time.
The table has three fixed entries, so a local map is enough; a switch or array works identically,
and the map is kept here to match the 代码随想录 presentation.

时间与空间复杂度 / Time and Space Complexity
n = len(s)。两种解法都是最坏时间 O(n)，每个字符检查一次；奇数长度可 O(1) 提前返回。
辅助空间 O(n)，最坏情况下栈保存线性数量的括号。
解法二每个字符多一次哈希查找，按平均 O(1) 计入，并额外占用 O(1) 的三条映射表空间。
n = len(s). Both methods take worst-case O(n) time, examining each character once;
odd lengths return in O(1). Auxiliary space is O(n) for a stack holding a linear number of brackets
in the worst case. Method 2 adds one hash lookup per character, counted at average O(1),
plus O(1) space for the three-entry table.
*/

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

// 1. 栈保存期待的右括号：直接比较
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

// 2. 哈希表映射：压入左括号本身
func isValidWithMap(s string) bool {
	// A valid bracket string must have an even number of characters.
	// 有效括号字符串的字符数量一定是偶数。
	if len(s)%2 != 0 {
		return false
	}

	// pairs maps each closing bracket to the opening bracket it must match.
	// pairs 把每个右括号映射到它必须匹配的左括号。
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	// stack stores the opening brackets that are still unmatched.
	// stack 保存尚未闭合的左括号本身，而不是期待的右括号。
	stack := make([]byte, 0, len(s)/2)

	// Examine every bracket in the string.
	// 依次检查字符串中的每个括号。
	for i := 0; i < len(s); i++ {
		current := s[i]

		// A lookup miss means the character is an opening bracket.
		// 查不到映射，说明当前字符是左括号。
		expected, isClosing := pairs[current]
		if !isClosing {
			// Opening brackets are pushed unchanged; no translation happens here.
			// 左括号原样入栈，此处不做任何转换。
			stack = append(stack, current)
			continue
		}

		// A closing bracket cannot be matched if the stack is empty.
		// 如果栈为空，当前右括号就没有对应的左括号。
		if len(stack) == 0 {
			return false
		}

		// The innermost unmatched opening bracket must be exactly the expected one.
		// 最内层尚未闭合的左括号必须正是当前右括号期待的那一个。
		if stack[len(stack)-1] != expected {
			return false
		}

		// The pair is matched, so this opening bracket is no longer pending.
		// 配对成功，该左括号不再处于未闭合状态，弹出栈顶。
		stack = stack[:len(stack)-1]
	}

	// Any leftover opening bracket was never closed.
	// 栈中剩余的左括号都没有闭合，字符串无效。
	return len(stack) == 0
}
