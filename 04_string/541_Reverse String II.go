package _4_string

/*
题目描述 / Problem Description
给定字符串 s 和整数 k，从字符串开头起每计数 2k 个字符，就反转其中前 k 个字符。
剩余少于 k 个时全部反转，介于 k 和 2k 之间时只反转前 k 个。
Given a string s and integer k, reverse the first k characters for every
block of 2k characters. Reverse all remaining characters if fewer than k remain,
or only the first k otherwise.
*/

// 1. Walk every 2k-byte block and reverse the first min(k, remaining) characters with two pointers.
// 1. 按 2k 分组，用双指针反转每组前 min(k, 剩余长度) 个字符。
// Time: O(n), Space: O(n) for the []byte copy.
// 时间复杂度：O(n)，空间复杂度：O(n)，由 []byte 副本产生。
//
// 每次跳过 2*k 个字节，只反转这一块的前 min(k,剩余长度) 个；right 截到末尾处理不足 k 的情况。
// Advance by 2*k bytes and reverse the first min(k,remaining) bytes; clamp right for the final short block.
// k<=0 原样返回；其余按题目范围使用单字节字符和正 k。
// Return unchanged for k<=0; otherwise use the problem's single-byte characters and bounded positive k.
func reverseStr(s string, k int) string {
	if k <= 0 {
		return s
	}

	bytes := []byte(s)

	for start := 0; start < len(bytes); start += 2 * k {
		left := start
		right := start + k - 1

		if right >= len(bytes) {
			right = len(bytes) - 1
		}

		for left < right {
			bytes[left], bytes[right] = bytes[right], bytes[left]
			left++
			right--
		}
	}

	return string(bytes)
}
