package _5_stack_queue

/*
Use a byte slice as a stack to store characters that have not been removed.
使用字节切片作为栈，保存尚未被删除的字符。

For each character, compare it with the top of the stack.
对于每个字符，将它与栈顶字符进行比较。

If they are equal, remove the stack top because the two characters cancel out.
如果两者相同，就弹出栈顶，因为这两个相邻字符需要一起删除。

If they are different, push the current character onto the stack.
如果两者不同，就将当前字符压入栈中。

The remaining stack forms the final string.
最后栈中剩余的字符组成最终字符串。
*/

func removeDuplicates(s string) string {
	// stack stores the characters that remain after processing.
	// stack 保存处理过程中尚未被删除的字符。
	stack := make([]byte, 0, len(s))

	// Process each character from left to right.
	// 从左到右处理字符串中的每个字符。
	for i := 0; i < len(s); i++ {
		// current is the character currently being processed.
		// current 是当前正在处理的字符。
		current := s[i]

		// If the stack is not empty and its top equals current,
		// the two adjacent duplicate characters should be removed.
		// 如果栈不为空，并且栈顶字符与 current 相同，
		// 说明出现了两个相邻重复字符，需要将它们删除。
		if len(stack) > 0 && stack[len(stack)-1] == current {
			// Pop the top character.
			// 弹出栈顶字符。
			stack = stack[:len(stack)-1]
		} else {
			// No duplicate pair is formed, so keep the current character.
			// 当前字符没有形成重复对，因此将它保留下来。
			stack = append(stack, current)
		}
	}

	// Convert the remaining characters back into a string.
	// 将栈中剩余的字符转换回字符串。
	return string(stack)
}
