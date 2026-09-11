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

// 1. strings.Fields 辅助：先切词再反转切片
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
// 2. 整体反转再逐词反转：不依赖库函数切词
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

// 3. 从右向左扫描单词：单词顺序自然倒转
func reverseWordsBackward(s string) string {
	// Builder appends the result without allocating a new string per word.
	// Builder 逐段追加结果，避免每个单词都创建一个新字符串。
	var result strings.Builder

	// right scans from the end of the string toward the beginning.
	// right 从字符串末尾向开头扫描。
	for right := len(s) - 1; right >= 0; {
		// Skip the spaces that separate this word from the one already handled.
		// 跳过当前单词与已处理部分之间的空格。
		for right >= 0 && s[right] == ' ' {
			right--
		}

		// Only spaces were left, so there is no further word to append.
		// 剩下的全是空格，说明已经没有单词可以追加。
		if right < 0 {
			break
		}

		// left walks past the word until it reaches a space or the string start.
		// left 一直向左走过整个单词，直到遇到空格或到达字符串开头。
		left := right
		for left >= 0 && s[left] != ' ' {
			left--
		}

		// Add a separator only when the result already holds a word, so there is no leading space.
		// 只有结果中已经有单词时才补分隔空格，因此不会产生前导空格。
		if result.Len() > 0 {
			result.WriteByte(' ')
		}

		// left stopped one position before the word, so the word is s[left+1:right+1].
		// left 停在单词前一个位置，因此单词是 s[left+1:right+1]。
		// Copying it forward keeps its letters in the original order; only the word order is reversed.
		// 按原方向复制可以保持单词内部字母不变，只有单词之间的顺序被倒转。
		result.WriteString(s[left+1 : right+1])

		// Continue from the character before this word.
		// 从这个单词前面的字符继续扫描。
		right = left - 1
	}

	return result.String()
}
