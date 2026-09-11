package _4_string

/*
题目描述 / Problem Description
给定字符串 s 和整数 k，从字符串开头起每计数 2k 个字符，就反转其中前 k 个字符。
剩余少于 k 个时全部反转，介于 k 和 2k 之间时只反转前 k 个。
Given a string s and integer k, reverse the first k characters for every
block of 2k characters. Reverse all remaining characters if fewer than k remain,
or only the first k otherwise.

解题思路 / Solution Approach
先将字符串转换为可修改的字节切片，每次以 2k 为步长定位一个分组，
用双指针反转该组前 min(k, 剩余长度) 个字符。
Convert the string to a mutable byte slice, advance in steps of 2k,
and use two pointers to reverse the first min(k, remaining length) characters of each block.

时间与空间复杂度 / Time and Space Complexity
n = len(s)。时间 O(n)，每个字符最多参与一次分组反转。辅助空间 O(n)，
Go 字符串不可修改，需要 []byte 副本；返回字符串也占 O(n)。
n = len(s). Time O(n), with each character participating in
at most one block reversal. Auxiliary space O(n) is needed for
the mutable byte copy; the returned string also takes O(n).

边界条件 / Edge Cases
步长是 2k，因此 k<=0 时 start 不会前进，循环无法结束。题目保证 k>=1，
函数仍然先判断 k<=0 并原样返回，避免死循环。
The loop advances by 2k, so a nonpositive k never moves start and the loop cannot finish.
The constraints guarantee k >= 1, but the function still returns s unchanged for k <= 0.
*/

/*
I convert the string to a byte slice because strings are immutable in Go.
Go 的 string 不能原地修改，所以先转换成 []byte。

Then I process the string in chunks of 2k characters.
每次处理长度为 2k 的一组字符。

For each chunk, I reverse only the first k characters.
每组只反转前 k 个字符。

If fewer than k characters remain, I reverse all remaining characters.
如果剩余字符少于 k 个，就把剩余字符全部反转。
*/

func reverseStr(s string, k int) string {
	/*
		The step below is 2k, so a nonpositive k would never advance start and would loop forever.
		下面的步长是 2k，k 不是正数时 start 永远不会前进，会造成死循环。

		k >= 1 由题目保证，这里只是防御性返回原字符串。
		The constraints guarantee k >= 1; this is only a defensive guard.
	*/
	if k <= 0 {
		return s
	}

	// bytes allows in-place character swaps.
	// bytes 允许我们原地交换字符。
	bytes := []byte(s)

	// start jumps by 2k because only the first k characters of each 2k block change.
	// start 每次跳过 2k，因为每个 2k 分组只需要处理前 k 个字符。
	for start := 0; start < len(bytes); start += 2 * k {
		// The intended reverse range is [start, start+k-1].
		// 计划反转的闭区间是 [start, start+k-1]。
		left := start
		right := start + k - 1

		/*
			If fewer than k characters remain, clamp right to the last valid index.
			如果剩余字符不足 k 个，就把 right 限制在最后一个有效下标。

			例如：
			s = "abc", k = 5
			right = start + k - 1 = 4，但数组最后一个下标是 2。
		*/
		if right >= len(bytes) {
			right = len(bytes) - 1
		}

		// Reverse the selected range by swapping from both ends.
		// 从区间两端向中间交换，完成这一段的反转。
		for left < right {
			bytes[left], bytes[right] = bytes[right], bytes[left]
			left++
			right--
		}
	}

	// Convert the modified byte slice back to a string.
	// 把修改后的字节切片转换回字符串。
	return string(bytes)
}
