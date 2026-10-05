package _4_string

/*
题目描述 / Problem Description
给定字符串 s 和整数 k，从字符串开头起每计数 2k 个字符，就反转其中前 k 个字符。
剩余少于 k 个时全部反转，介于 k 和 2k 之间时只反转前 k 个。
Given a string s and integer k, reverse the first k characters for every
block of 2k characters. Reverse all remaining characters if fewer than k remain,
or only the first k otherwise.
*/

// 面试首选 / Interview first choice：541 反转字符串 II。
//
// 1. Walk every 2k-byte block and reverse the first min(k, remaining) characters with two pointers.
// 1. 按 2k 分组，用双指针反转每组前 min(k, 剩余长度) 个字符。
// Time: O(n), Space: O(n) for the []byte copy.
// 时间复杂度：O(n)，空间复杂度：O(n)，由 []byte 副本产生。
//
// 每次跳过 2*k 个字节，只反转这一块的前 min(k,剩余长度) 个；right 截到末尾处理不足 k 的情况。
// Advance by 2*k bytes and reverse the first min(k,remaining) bytes; clamp right for the final short block.
// k<=0 原样返回；其余按题目范围使用单字节字符和正 k。
// Return unchanged for k<=0; otherwise use the problem's single-byte characters and bounded positive k.
//
// 状态：start 是当前 2k 块的起点，left/right 是其中尚未反转的区间；两端交换后区间外已就位。
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

// 2. Alternate reversed and unchanged k-byte groups while building the result.
// 2. 每 k 个字节一组，交替倒序、正序写入结果：反一组，留一组。
// Time: O(n), Space: O(n) for the result buffer and returned string.
// 时间复杂度：O(n)，空间复杂度：O(n)，包含结果缓冲区和返回字符串。
//
// 每两组对应一个 2k 区间；reversed 表示当前组是否需要倒序，处理完一组后切换。
// Two groups form one 2k block; reversed controls the current group's order and toggles afterward.
// end 为组的右开边界，截到串尾即可处理不足 k 的最后一组；k<=0 原样返回，按题目约束处理单字节字符。
// Clamp the exclusive end to handle a short final group; return unchanged for k<=0 and use single-byte characters as required by the problem.
// 例如 k=3：abc|def|gh → cba|def|hg。
// Example with k=3: abc|def|gh becomes cba|def|hg.
func reverseStrAlternating(s string, k int) string {
	if k <= 0 {
		return s
	}

	result := make([]byte, 0, len(s))
	reversed := true

	for start := 0; start < len(s); {
		end := start + min(k, len(s)-start)

		if reversed {
			for i := end - 1; i >= start; i-- {
				result = append(result, s[i])
			}
		} else {
			result = append(result, s[start:end]...)
		}

		reversed = !reversed
		start = end
	}

	return string(result)
}
