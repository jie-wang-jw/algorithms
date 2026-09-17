package _4_string

/*
题目描述 / Problem Description
给定字符数组 s，原地反转数组中的字符。要求使用 O(1) 额外空间完成。
Given a character array s, reverse its characters in place using O(1) extra space.
*/

// 1. Two-pointer in-place byte swap (recommended)
// 1. 字节双指针原地交换（推荐）
// Swap from both ends until left meets right. Correct for printable ASCII (one byte per character).
// 从两端向中间交换，直到指针相遇。本题可打印 ASCII 下一字符一字节，结果正确。
// Time: O(n), Space: O(1).
// 时间复杂度：O(n)，空间复杂度：O(1)。
//
// 交换首尾字节后向内收缩，区间外已就位；奇数长度的中心无需交换。按题目单字节字符使用。
// Swap inward while outer bytes stay fixed; an odd center needs no swap. Use the problem's single-byte character contract.
func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// 2. Two-pointer rune swap
// 2. rune 双指针
// Decode to []rune first so multi-byte UTF-8 characters stay intact; cannot meet O(1) space.
// 先解码成 []rune，多字节 UTF-8 字符不会被拆坏；无法做到 O(1) 空间。
// Time: O(n), Space: O(u) for the rune slice plus O(n) output.
// 时间复杂度：O(n)，空间复杂度：rune 切片 O(u)，输出另占 O(n)。
//
// 先转为 rune 再交换，避免拆开 UTF-8 编码；反转单位是码点，不是可能由多个码点组成的可见字符。
// Swap runes to avoid splitting UTF-8 encodings; this reverses code points, not multi-code-point grapheme clusters.
func reverseStringUnicode(s string) string {
	runes := []rune(s)
	left, right := 0, len(runes)-1
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
	return string(runes)
}
