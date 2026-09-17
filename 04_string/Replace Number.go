package _4_string

/*
题目描述 / Problem Description
给定一个只包含小写字母和数字字符的字符串，将其中每个数字字符替换为字符串 "number"，
并返回替换后的结果。
Given a string containing only lowercase letters and digit characters,
replace every digit character with the word "number" and return the resulting string.
*/

import (
	"strings"
)

// 1. Scan left to right and build the result with strings.Builder.
// 1. 从左向右用 Builder 构造结果。
// Time: O(n), Space: O(n) for the Builder buffer.
// 时间复杂度：O(n)，空间复杂度：O(n)，由 Builder 缓冲区产生。
//
// 仅将 ASCII '0'..'9' 替换为 "number"，其他字节原样写入；Builder 避免每次拼接都复制已有结果。
// Replace only ASCII digits with "number" and preserve other bytes; Builder avoids copying prior output on each append.
func replaceNumber(s string) string {
	var builder strings.Builder

	for i := 0; i < len(s); i++ {
		ch := s[i]

		if ch >= '0' && ch <= '9' {
			builder.WriteString("number")
		} else {
			builder.WriteByte(ch)
		}
	}

	return builder.String()
}

// 2. Expand first, then fill from the back: write "number" backward so it reads forward.
// 2. 先扩容再从后向前填充：逆序写入 "number"，读出来才是正序。
// Time: O(n), Space: O(n) for the expanded buffer.
// 时间复杂度：O(n)，空间复杂度：O(n)，由扩容缓冲区产生。
//
// 每个数字从 1 字节变为 6 字节，故新长度为 len(s)+5*digits；从后向前写使 write>=read，不覆盖未读内容。
// Each digit expands by five bytes, giving len(s)+5*digits; backward filling keeps write>=read and preserves unread input.
// 写指针向左走，所以 "number" 也必须倒序写入，最终字符串中的顺序才正确。
// Because write moves left, emit "number" backward so it appears in the correct final order.
func replaceNumberBackward(s string) string {
	digits := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			digits++
		}
	}

	buffer := make([]byte, len(s)+5*digits)
	copy(buffer, s)

	write := len(buffer) - 1

	for read := len(s) - 1; read >= 0; read-- {
		ch := buffer[read]

		if ch >= '0' && ch <= '9' {
			for j := len("number") - 1; j >= 0; j-- {
				buffer[write] = "number"[j]
				write--
			}
		} else {
			buffer[write] = ch
			write--
		}

	}

	return string(buffer)
}
