package _4_string

/*
题目描述 / Problem Description
给定字符串 s，反转其中单词的顺序。单词之间可能有多个空格，
结果中单词之间只保留一个空格，且首尾不能有空格。
Given a string s, reverse the order of its words.
The input may contain extra spaces; the result must use one space
between words and have no leading or trailing spaces.

解题思路 / Solution Approach
使用 strings.Fields 提取所有非空单词并自动忽略多余空白，
再用双指针反转单词切片，最后以单个空格连接。
Use strings.Fields to extract words while discarding extra whitespace,
reverse the word slice with two pointers, and join it with single spaces.

时间与空间复杂度 / Time and Space Complexity
n 为字符串字节数，w 为单词数。reverseWords：时间 O(n)，Fields 扫描、
反转单词、Join 总计线性；辅助空间 O(w) 保存单词切片，输出 O(n)。reverseWords2：
清理空格和两次反转总时间 O(n)，Go 的 []byte(s) 副本占 O(n) 辅助空间；
仅反转步骤为 O(1) 空间。两者包含输出均为 O(n)。
For n bytes and w words, reverseWords takes O(n) time, O(w) auxiliary
space for word slices, and O(n) output space. reverseWords2 takes O(n)
time for normalization and reversals; its []byte(s) copy uses O(n)
auxiliary space, though the reversal steps alone use O(1). Total space including output is O(n) for both.
*/

import "strings"

func reverseWords(s string) string {
	// Split keeps empty items for repeated spaces, while Fields removes extra whitespace.
	// Split 会保留连续空格产生的空字符串；Fields 会自动去掉多余空白。
	//words := strings.Split(s, " ")
	words := strings.Fields(s)

	// Reverse the word slice with two pointers.
	// 使用左右双指针反转单词切片。
	left, right := 0, len(words)-1
	for left < right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}

	// Join the reversed words with exactly one space.
	// 用一个空格连接反转后的单词。
	return strings.Join(words, " ")
}

/*
1. Remove extra spaces. / 去掉多余空格。
2. Reverse the whole string. / 整体反转整个字符串。
3. Reverse every individual word. / 再反转每一个单词。

Example / 示例:
"the sky" -> "yks eht" -> "sky the"
*/
func reverseWords2(s string) string {
	// Go strings are immutable, so convert to []byte for in-place changes.
	// Go 的 string 不能原地修改，所以先转换成 []byte。
	b := []byte(s)

	// First normalize spaces, then reverse the entire byte slice.
	// 先清理多余空格，再整体反转字节切片。
	b = removeExtraSpaces(b)
	reverse(b, 0, len(b)-1)

	// Reverse each word to restore its character order.
	// 逐个反转单词，使单词内部字符恢复正常顺序。
	start := 0
	// The last word has no trailing space, so i == len(b) must also trigger processing.
	// 最后一个单词后面没有空格，因此 i == len(b) 时也要触发一次处理。
	for i := 0; i <= len(b); i++ {
		// A space or the end of the slice marks the end of one word.
		// 遇到空格或切片末尾，表示找到一个单词的结束位置。
		if i == len(b) || b[i] == ' ' {
			reverse(b, start, i-1)
			// The next word begins immediately after the current separator.
			// 下一个单词从当前空格后面开始。
			start = i + 1
		}
	}
	return string(b)
}

/*
1. Remove leading spaces. / 去掉开头空格。
2. Remove trailing spaces. / 去掉结尾空格。
3. Keep exactly one space between words. / 单词之间只保留一个空格。
*/
func removeExtraSpaces(b []byte) []byte {
	// slow is the next write position in the cleaned result.
	// slow 是清理后结果的下一个写入位置。
	slow := 0

	/*
		fast scans the original bytes; slow writes the cleaned result.
		fast 负责扫描原字符串；slow 负责写入清理后的结果。
	*/
	for fast := 0; fast < len(b); fast++ {
		// Ignore spaces until fast reaches the beginning of a word.
		// 忽略空格，直到 fast 找到一个单词的开头。
		if b[fast] != ' ' {
			// Add one separator before every word except the first word.
			// 除第一个单词外，每个单词前补一个空格作为分隔符。
			if slow != 0 {
				b[slow] = ' '
				slow++
			}
			// Copy the whole word to the current write position.
			// 把当前完整单词复制到 slow 指向的写入区域。
			for fast < len(b) && b[fast] != ' ' {
				b[slow] = b[fast]
				slow++
				fast++
			}
		}
	}
	// Only [0, slow) contains the cleaned result.
	// 只有 [0, slow) 区间属于清理后的有效结果。
	return b[:slow]
}

// Reverse the inclusive range [left, right] in place.
// 原地反转闭区间 [left, right]。
func reverse(b []byte, left, right int) {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}
