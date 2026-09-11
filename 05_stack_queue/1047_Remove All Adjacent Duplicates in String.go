package _5_stack_queue

/*
1047. 删除字符串中的所有相邻重复项 / Remove All Adjacent Duplicates in String

题目描述 / Problem Description
给定一个只包含小写英文字母的字符串 s，不断删除两个相邻且相同的字符，
直到无法继续删除，并返回最终字符串。
Given a string s containing only lowercase English letters,
repeatedly remove adjacent equal character pairs until no
more removals are possible, and return the final string.

解法一：字节切片模拟栈（推荐） / Method 1: Byte Slice as a Stack (Recommended)
使用字节切片模拟栈。当前字符与栈顶相同时弹出栈顶，否则将当前字符入栈；
栈顶始终是最近保留的字符，因此可以自然处理连锁删除。
Use a byte slice as a stack. Pop when the current character matches the top;
otherwise push it. The top is always the most recently retained character,
so chain reactions are handled naturally.

关键逻辑：为什么这样做 / Why This Works
每轮开始时，栈是已扫描前缀消除相邻重复后的结果，内部已没有可删除的相邻对。
新字符只可能与栈顶形成新的一对；相同就弹栈且不压入当前字符，等于同时删掉两者。
例如 "abbaca"：a -> ab -> a -> 空 -> c -> ca，删掉 bb 后暴露的 a 会与下一个 a 抵消，因此不用回退原字符串下标。
At each step the stack is the fully reduced scanned prefix, containing no adjacent duplicate pair.
A new character can create a pair only with its top. On equality, pop and do not push the current character,
removing both. For "abbaca": a -> ab -> a -> empty -> c -> ca. Removing bb exposes a for
cancellation with the next a, without rewinding the input index.

解法二：快慢指针原地模拟栈 / Method 2: Fast and Slow Pointers as an In-Place Stack
removeDuplicatesTwoPointers 先把 s 复制成 []byte，然后把这份副本的前部当作栈。
buffer[:slow] 始终是已读前缀消除后的结果，slow 既是写入位置，也就是栈长度。
buffer[slow-1] 相当于栈顶：相等时 slow-- 抵消一对，否则写入 buffer[slow] 并递增。
覆盖写是安全的，因为每轮都有 slow<=fast，写指针永远不会越过读指针去破坏未读字符。
这与解法一是同一条消除规则，区别只是复用字节副本的前部，而不另外 append 一个栈。
Copy s into a []byte, then treat the front of that copy as the stack.
buffer[:slow] is always the reduced prefix of the scanned input, and slow is both the write position
and the stack size. buffer[slow-1] acts as the stack top: pop with slow-- on equality,
otherwise write buffer[slow] and advance.
Overwriting is safe because slow <= fast holds every round, so the write pointer never passes the read
pointer and never clobbers unread input.
This is the same cancellation rule as Method 1, reusing the front of the byte copy instead of a separate stack.

时间与空间复杂度 / Time and Space Complexity
n = len(s)。两种解法都是时间 O(n)：每个字符最多入栈、出栈各一次，最后转换字符串也是 O(n) 上界。
解法一 removeDuplicates：辅助空间 O(n) 保存栈；结果最多 O(n)，总空间仍为 O(n)。
解法二 removeDuplicatesTwoPointers：辅助空间同样是 O(n)。
Go 的字符串不可变，[]byte(s) 一定会复制一份，所以这里不能误写为 O(1) 原地算法；
它相对解法一只省下一次单独的栈分配，并且不修改调用方的数据。
n = len(s). Both methods take O(n) time: each character is pushed and popped at most once,
and the final string conversion is also at most O(n).
Method 1 removeDuplicates: auxiliary space O(n) for the stack, output at most O(n), total O(n).
Method 2 removeDuplicatesTwoPointers: auxiliary space is also O(n).
Go strings are immutable, so []byte(s) always copies; this must not be reported as an O(1) in-place algorithm.
Compared with Method 1 it merely avoids one separate stack allocation, and it leaves the caller's data unchanged.
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

// 1. 字节切片模拟栈：推荐
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

// 2. 快慢指针原地模拟栈：复用字节副本的前部
func removeDuplicatesTwoPointers(s string) string {
	// Go strings are immutable, so this conversion copies; the caller's data is never modified.
	// Go 的字符串不可变，这次转换一定会复制，因此不会修改调用方的数据。
	buffer := []byte(s)

	// slow is both the write position and the stack size; buffer[:slow] is the reduced prefix.
	// slow 既是写入位置，也是栈长度；buffer[:slow] 就是已读前缀消除后的结果。
	slow := 0

	// fast reads the original characters in order.
	// fast 按顺序读取原始字符。
	for fast := range buffer {
		// buffer[slow-1] is the stack top, valid only when the stack is nonempty.
		// buffer[slow-1] 相当于栈顶，只有栈非空时才可以读取。
		if slow > 0 && buffer[slow-1] == buffer[fast] {
			// The pair cancels: drop the top and do not write the current character.
			// 两个字符抵消：弹出栈顶，并且不写入当前字符。
			slow--
			continue
		}

		// No pair is formed, so keep the current character on the stack.
		// 没有形成重复对，把当前字符保留在栈上。
		//
		// slow <= fast always holds, so this write never clobbers an unread character.
		// slow <= fast 始终成立，所以这次写入不会覆盖尚未读取的字符。
		buffer[slow] = buffer[fast]
		slow++
	}

	// Only the first slow bytes survived; the rest of buffer is stale data.
	// 只有前 slow 个字节是结果，buffer 之后的内容是残留数据。
	return string(buffer[:slow])
}
