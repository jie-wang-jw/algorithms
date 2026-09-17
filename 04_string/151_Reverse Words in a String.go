package _4_string

/*
题目描述 / Problem Description
给定字符串 s，反转其中单词的顺序。单词之间可能有多个空格，
结果中单词之间只保留一个空格，且首尾不能有空格。
Given a string s, reverse the order of its words.
The input may contain extra spaces; the result must use one space
between words and have no leading or trailing spaces.
*/

import "strings"

// 1. strings.Fields helper
// 1. strings.Fields 辅助
// Fields drops extra whitespace; reverse the word slice, then Join with a single space.
// Fields 去掉多余空白；反转单词切片后用单个空格 Join。
// Time: O(n), Space: O(w) for the word slice plus O(n) output.
// 时间复杂度：O(n)，空间复杂度：单词切片 O(w)，输出 O(n)。
//
// Fields 提取非空单词并丢弃 Unicode 空白；只反转单词数组，不反转单词内部，最后用一个空格连接。
// Fields drops Unicode whitespace; reverse the word order, preserve each word, and join with one space.
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

// 2. Reverse whole string, then reverse each word
// 2. 整体反转再逐词反转
// Normalize spaces, reverse all bytes (flips word order and letters), then reverse each word to restore letters.
// 先清空格，再整体反转（颠倒单词序与字母序），最后逐词反转以恢复字母。
// Time: O(n), Space: O(n) for the []byte copy.
// 时间复杂度：O(n)，空间复杂度：O(n)，由 []byte(s) 副本产生。
//
// 先压缩空格，再整段反转：单词顺序和内部字节顺序都会反转；逐词再反转一次，就只剩词序反转。
// Normalize spaces, reverse the whole buffer, then reverse each word to restore its bytes while keeping reversed word order.
// i==len(b) 是最后一个词的结束哨兵，先判定它可避免越界读 b[i]；此版本只把 ASCII 空格当分隔符。
// i==len(b) closes the final word and is checked before b[i]; only ASCII spaces are separators here.
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

// Collapse extra spaces in place
// 原地清空格
// fast scans; slow writes. Insert one separator only before non-first words; return b[:slow].
// fast 扫描、slow 写入；仅在非首单词前补一个分隔空格；返回 b[:slow]。
// Time: O(n), Space: O(1).
// 时间复杂度：O(n)，空间复杂度：O(1)。
//
// b[:slow] 是已写好的结果；只在第二个及以后单词之前写一个空格，因此同时去掉首尾和多余空格。
// b[:slow] is the compacted output; write one space only before later words, removing leading, trailing and repeated spaces.
// slow 不超过读位置 fast，原地写入不会覆盖未读内容；返回缩短的切片。
// slow never overtakes fast, so writes preserve unread bytes; return the shortened slice.
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
//
// 反转闭区间 [left,right]，交换两端后向内收缩；空区间或单字节不动，非空区间下标须有效。
// Reverse inclusive [left,right] by swapping inward; empty/single-byte ranges are unchanged and nonempty bounds must be valid.
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
// 从右向左跳过空格并定位整词，按原字节顺序追加 s[left+1:right+1]；只在已有结果后追加分隔空格。
// Scan backward to locate whole words but append their bytes in original order; add one separator only after prior output.
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
