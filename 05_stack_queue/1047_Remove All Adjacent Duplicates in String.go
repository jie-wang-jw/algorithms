package _5_stack_queue

/*
题目描述 / Problem Description
给定一个只包含小写英文字母的字符串 s，不断删除两个相邻且相同的字符，直到无法继续删除，并返回最终字符串。
Given a string s containing only lowercase English letters, repeatedly remove adjacent equal character pairs until no more removals are possible, and return the final string.

解题思路 / Solution Approach
使用字节切片模拟栈。当前字符与栈顶相同时弹出栈顶，否则将当前字符入栈；栈顶始终是最近保留的字符，因此可以自然处理连锁删除。
Use a byte slice as a stack. Pop when the current character matches the top; otherwise push it. The top is always the most recently retained character, so chain reactions are handled naturally.

时间与空间复杂度 / Time and Space Complexity
n = len(s)。时间 O(n)，每个字符最多入栈、出栈各一次，最后转换字符串也是 O(n) 上界。辅助空间 O(n) 保存栈；结果最多 O(n)，总空间仍为 O(n)。
n = len(s). Time O(n), pushing and popping each character at most once; final string conversion is also at most O(n). Auxiliary stack space and output space are each O(n), so total space is O(n).
*/

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
