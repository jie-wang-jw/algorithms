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

补充解法：扩容后从后向前写 / Alternative: Expand and Fill Backward
replaceNumberBackward 先数 d 个数字，结果长度为 n+5d，因为每个数字由 1 字节变成 6 字节。
复制原文到扩容缓冲区，read 从原文末尾读，write 从新末尾写；数字写入 number 的逆向字节。
写指针始终不在读指针左侧，因此不会覆盖尚未读取的原文。时间 O(n)，Go 中缓冲区 O(n)。
“反向填充用 O(1) 状态”不代表整个接收 string 的函数只用 O(1) 空间。
Count d digits and allocate n+5d bytes. Read backward from the original end and write backward from the new end.
Write "number" backward so the final word reads forward. The writer never overwrites unread bytes.
Time O(n), buffer space O(n) in Go; O(1) pointer state does not make the whole string function constant-space.
*/

import (
	"strings"
)

// 1. 从左向右用 Builder 构造结果
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

// 2. 先扩容再从后向前填充
func replaceNumberBackward(s string) string {
	// Count the digits first so the final length is known before any writing.
	// 先统计数字个数，这样在写入之前就能知道结果的最终长度。
	digits := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			digits++
		}
	}

	// Each digit grows from 1 byte to the 6 bytes of "number", so the result needs 5 extra bytes per digit.
	// 每个数字由 1 个字节变成 "number" 的 6 个字节，因此每个数字多占 5 个字节。
	buffer := make([]byte, len(s)+5*digits)
	copy(buffer, s)

	// read consumes the original text from its end; write fills the enlarged buffer from its end.
	// read 从原文末尾向前读取，write 从扩容后的末尾向前写入。
	write := len(buffer) - 1

	for read := len(s) - 1; read >= 0; read-- {
		ch := buffer[read]

		if ch >= '0' && ch <= '9' {
			// Write "number" backward so that it reads forward in the finished buffer.
			// 逆序写入 "number" 的字节，最终在缓冲区中读出来才是正序。
			for j := len("number") - 1; j >= 0; j-- {
				buffer[write] = "number"[j]
				write--
			}
		} else {
			// Letters are copied unchanged to the current write position.
			// 字母原样复制到当前写入位置。
			buffer[write] = ch
			write--
		}

		/*
			write stays at or ahead of read: the gap equals 5 times the digits still unread,
			so filling backward never overwrites original bytes that have not been read yet.
			write 始终不在 read 左侧：两者的距离等于尚未读取部分中数字的个数乘以 5，
			因此从后向前填充不会覆盖还没读过的原文字节。
		*/
	}

	return string(buffer)
}
