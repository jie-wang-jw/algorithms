package _4_string

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
