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

手写法 reverseWords2：先清理空格，再反转整个字节切片，最后反转每个单词，使单词内部恢复原顺序。
The manual reverseWords2 normalizes spaces, reverses the entire byte slice, and reverses each word to restore its internal order.

关键逻辑：为什么这样做 / Why This Works
reverseWords2 的逻辑是两次反转作用在不同范围：整体反转同时颠倒单词顺序和每个单词的字母顺序；
再单独反转每个单词，只恢复字母顺序，保留已经倒过来的单词顺序。例如 "the sky" -> "yks eht" -> "sky the"。
清空格时仅在一个非首单词开始前补一个空格，所以不会有首尾空格，单词间恰好一个。
扫描单词结束处用 i==len(b) 充当末尾分隔符，并把它放在 || 左侧，短路后不会访问越界的 b[i]。
reverseWords2 reverses two different scopes: reversing the whole string reverses both word order and letter order;
reversing each word restores only its letters, preserving reversed word order. Example: "the sky" -> "yks eht" -> "sky the".
Space cleanup inserts one separator only before non-first words, preventing leading/trailing spaces and
repeated separators. i==len(b) acts as the last word's delimiter; placing it first in || short-circuits the out-of-bounds b[i] access.

时间与空间复杂度 / Time and Space Complexity
n 为字符串字节数，w 为单词数。reverseWords：时间 O(n)，Fields 扫描、
反转单词、Join 总计线性；辅助空间 O(w) 保存单词切片，输出 O(n)。reverseWords2：
清理空格和两次反转总时间 O(n)，Go 的 []byte(s) 副本占 O(n) 辅助空间；
仅反转步骤为 O(1) 空间。两者包含输出均为 O(n)。
For n bytes and w words, reverseWords takes O(n) time, O(w) auxiliary
space for word slices, and O(n) output space. reverseWords2 takes O(n)
time for normalization and reversals; its []byte(s) copy uses O(n)
auxiliary space, though the reversal steps alone use O(1). Total space including output is O(n) for both.

补充解法：从右向左扫描单词 / Alternative: Scan Words Right to Left
reverseWordsBackward 跳过尾部空格，再向左找到一个完整单词，按原字母顺序追加到结果。
先遇到原字符串靠后的单词，因此单词顺序自然倒转，不需要先反转字母再恢复。
只在已有结果后添加分隔空格；本题分隔符为普通空格。时间 O(n)，结果缓冲区 O(n)，其他状态 O(1)。
Skip spaces from the right, locate a complete word, and append its letters in their original order.
Later words are discovered first, reversing word order without reversing letters.
Insert a separator only after an existing word. Uses ordinary spaces as specified.
Time O(n), result buffer O(n), other state O(1).
*/

import "strings"

// 1. strings.Fields helper
// 1. strings.Fields 辅助
// Fields drops extra whitespace; reverse the word slice, then Join with a single space.
// Fields 去掉多余空白；反转单词切片后用单个空格 Join。
// Time: O(n), Space: O(w) for the word slice plus O(n) output.
// 时间复杂度：O(n)，空间复杂度：单词切片 O(w)，输出 O(n)。
//
// 步骤与要点 / Steps and notes:
//  1. Prefer Fields over Split: Split keeps empty items for repeated spaces.
//     用 Fields 而不是 Split：Split 会保留连续空格产生的空串。
func reverseWords(s string) string {
	words := strings.Fields(s)
	left, right := 0, len(words)-1
	for left < right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}
	return strings.Join(words, " ")
}

/*
1. Remove extra spaces. / 去掉多余空格。
2. Reverse the whole string. / 整体反转整个字符串。
3. Reverse every individual word. / 再反转每一个单词。

Example / 示例:
"the sky" -> "yks eht" -> "sky the"
*/
// 2. Reverse whole string, then reverse each word
// 2. 整体反转再逐词反转
// Normalize spaces, reverse all bytes (flips word order and letters), then reverse each word to restore letters.
// 先清空格，再整体反转（颠倒单词序与字母序），最后逐词反转以恢复字母。
// Time: O(n), Space: O(n) for the []byte copy.
// 时间复杂度：O(n)，空间复杂度：O(n)，由 []byte(s) 副本产生。
//
// 步骤与要点 / Steps and notes:
// 1. i == len(b) also ends a word: the last word has no trailing space.
//    i == len(b) 也要收尾：最后一个单词后面没有空格。把 i==len(b) 放在 || 左侧可短路，避免越界读 b[i]。
func reverseWords2(s string) string {
	b := []byte(s)
	b = removeExtraSpaces(b)
	reverse(b, 0, len(b)-1)

	start := 0
	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {
			reverse(b, start, i-1)
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
// Collapse extra spaces in place
// 原地清空格
// fast scans; slow writes. Insert one separator only before non-first words; return b[:slow].
// fast 扫描、slow 写入；仅在非首单词前补一个分隔空格；返回 b[:slow]。
// Time: O(n), Space: O(1).
// 时间复杂度：O(n)，空间复杂度：O(1)。
//
// 步骤与要点 / Steps and notes:
// 1. One separator before every word except the first.
//    除第一个单词外，每个单词前补一个空格。
func removeExtraSpaces(b []byte) []byte {
	slow := 0
	for fast := 0; fast < len(b); fast++ {
		if b[fast] != ' ' {
			if slow != 0 {
				b[slow] = ' '
				slow++
			}
			for fast < len(b) && b[fast] != ' ' {
				b[slow] = b[fast]
				slow++
				fast++
			}
		}
	}
	return b[:slow]
}

// Reverse the inclusive range [left, right] in place.
// 原地反转闭区间 [left, right]。
// Time: O(right-left+1), Space: O(1).
// 时间复杂度：O(right-left+1)，空间复杂度：O(1)。
func reverse(b []byte, left, right int) {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

// 3. Scan words from right to left
// 3. 从右向左扫描单词
// Discover later words first and append them, so word order reverses without reversing letters.
// 先发现靠后的单词再追加，单词顺序自然倒转，字母顺序保持不变。
// Time: O(n), Space: O(n) for the result buffer.
// 时间复杂度：O(n)，空间复杂度：O(n)，由结果缓冲区产生。
//
// 步骤与要点 / Steps and notes:
//  1. Separator only when the result already holds a word (no leading space).
//     只有结果里已有单词才补分隔空格，避免前导空格。
//  2. left stopped one before the word, so the slice is s[left+1:right+1].
//     left 停在单词前一位，因此单词是 s[left+1:right+1]。
func reverseWordsBackward(s string) string {
	var result strings.Builder
	for right := len(s) - 1; right >= 0; {
		for right >= 0 && s[right] == ' ' {
			right--
		}
		if right < 0 {
			break
		}

		left := right
		for left >= 0 && s[left] != ' ' {
			left--
		}
		if result.Len() > 0 {
			result.WriteByte(' ')
		}
		result.WriteString(s[left+1 : right+1])
		right = left - 1
	}
	return result.String()
}
