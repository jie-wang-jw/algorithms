package _5_stack_queue

/*
1047. 删除字符串中的所有相邻重复项 / Remove All Adjacent Duplicates in String

题目描述 / Problem Description
给定一个只包含小写英文字母的字符串 s，不断删除两个相邻且相同的字符，
直到无法继续删除，并返回最终字符串。
Given a string s containing only lowercase English letters,
repeatedly remove adjacent equal character pairs until no
more removals are possible, and return the final string.
*/

// 1. Byte slice as a stack (recommended): pop when the current character equals the top, otherwise push it.
// 1. 字节切片模拟栈：推荐；当前字符与栈顶相同就弹栈，否则入栈，连锁删除自然完成。
// Time: O(n), Space: O(n) for the stack, and the result is also at most O(n).
// 时间复杂度：O(n)，空间复杂度：O(n)，由栈产生，结果最多也是 O(n)。
//
// 栈始终保存已扫描前缀消除相邻重复后的结果；新字符只能与栈顶产生新的相邻重复。
// The stack is the reduced scanned prefix; a new byte can form a new duplicate pair only with its top.
// 相同就弹出抵消，不同就入栈；弹出暴露的新栈顶会与后续字符继续比较，从而处理连锁消除。
// Pop equal pairs or push a different byte; later comparisons against the exposed top handle cascading removals.
func removeDuplicates(s string) string {
	stack := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		current := s[i]

		if len(stack) > 0 && stack[len(stack)-1] == current {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, current)
		}
	}

	return string(stack)
}

// 2. Fast and slow pointers over a byte copy: buffer[:slow] is the stack, so slow is both write position and stack size.
// 2. 快慢指针原地模拟栈：复用字节副本的前部，buffer[:slow] 就是栈，slow 既是写入位置也是栈长度。
// Time: O(n), Space: O(n), because []byte(s) copies the immutable string; this is not an O(1) in-place algorithm.
// 时间复杂度：O(n)，空间复杂度：O(n)，因为 []byte(s) 会复制不可变字符串，所以它不是 O(1) 原地算法。
//
// buffer[:slow] 充当栈，fast 是读位置；匹配栈顶就 slow--，否则写入后 slow++。
// buffer[:slow] acts as a stack and fast reads input; pop with slow-- or push by writing and incrementing slow.
// slow<=fast 保证不会覆盖未读字节；空间仍为 O(n)，因为字符串先复制成了 buffer。
// slow<=fast preserves unread bytes; space is still O(n) because the string is copied into buffer.
func removeDuplicatesTwoPointers(s string) string {
	buffer := []byte(s)

	slow := 0

	for fast := range buffer {
		if slow > 0 && buffer[slow-1] == buffer[fast] {
			slow--
			continue
		}

		buffer[slow] = buffer[fast]
		slow++
	}

	return string(buffer[:slow])
}
