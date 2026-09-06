package _4_string

/*
题目描述 / Problem Description
给定一个只包含小写字母和数字字符的字符串，将其中每个数字字符替换为字符串 "number"，
并返回替换后的结果。
Given a string containing only lowercase letters and digit characters,
replace every digit character with the word "number" and return the resulting string.

解题思路 / Solution Approach
从左到右遍历字符串，使用 strings.Builder 构造结果。
遇到 '0' 到 '9' 时写入 "number"，否则保留原字符。
Scan the string from left to right and build the result
with strings.Builder. Write "number" for digits from '0' to '9';
otherwise keep the original character.

时间与空间复杂度 / Time and Space Complexity
n 为输入长度，d 为数字字符数量，输出长度为 n+5d，最多 6n。时间 O(n)，
每个字符最多写入 6 个字节。Builder 的结果缓冲区占 O(n)，包含输出总空间 O(n)，
除此以外仅用 O(1) 状态。
For input length n and d digit characters, output length is n+5d,
at most 6n. Time O(n), writing at most six bytes per input character.
The Builder result buffer uses O(n) space; other state is O(1), so total space including output is O(n).
*/

import (
	"strings"
)

func replaceNumber(s string) string {
	// Builder appends output efficiently without repeatedly creating new strings.
	// Builder 可以高效追加内容，避免反复创建新的字符串。
	var builder strings.Builder

	// The problem contains only lowercase ASCII letters and digits, so byte iteration is enough.
	// 题目只有小写 ASCII 字母和数字，因此按 byte 遍历即可。
	for i := 0; i < len(s); i++ {
		ch := s[i]

		// ASCII digits are in the continuous range from '0' to '9'.
		// ASCII 数字字符位于连续区间 '0' 到 '9' 中。
		if ch >= '0' && ch <= '9' {
			// Replace one digit character with the whole word "number".
			// 遇到一个数字字符，就写入完整单词 "number"。
			builder.WriteString("number")
		} else {
			// Keep lowercase letters unchanged.
			// 小写字母保持不变，直接写入结果。
			builder.WriteByte(ch)
		}
	}

	// Build and return the final string.
	// 生成并返回最终字符串。
	return builder.String()
}
