package _4_string

/*
题目描述 / Problem Description
给定字符数组 s，原地反转数组中的字符。要求使用 O(1) 额外空间完成。
Given a character array s, reverse its characters in place using O(1) extra space.

解题思路 / Solution Approach
使用左右双指针，从数组两端向中间移动，每轮交换两个指针指向的字符，直到指针相遇或交错。
Use two pointers starting at both ends. Swap their characters and move inward until the pointers meet or cross.

时间与空间复杂度 / Time and Space Complexity
n = len(s)。时间 O(n)，交换约 n/2 对字符。辅助空间 O(1)，直接修改输入切片。
n = len(s). Time O(n), swapping about n/2 character pairs. Auxiliary space O(1), modifying the input slice directly.

补充解法：按码点反转 / Alternative: Reverse by Code Point
reverseString 交换的是字节。本题输入是可打印 ASCII，一个字符正好一个字节，所以结果正确；
但对多字节 UTF-8 字符，逐字节反转会把编码字节的顺序也打乱，产生无效编码。
reverseStringUnicode 先用 []rune(s) 解码成码点再反转，因此每个字符保持完整。
它不能做到 O(1) 空间，也不做 Unicode 规范化：由多个码点组合出的字符（如基字符加变音符）仍会被拆开。
n 为字节数，u 为码点数。时间 O(n)，rune 切片辅助空间 O(u)，输出另占 O(n)。
reverseString swaps bytes, which is correct here because the input is printable ASCII: one byte per character.
Reversing bytes of multi-byte UTF-8 characters would also reverse their encoding bytes and produce invalid text.
reverseStringUnicode decodes to code points with []rune(s) first, so each character stays intact.
It cannot run in O(1) space and applies no normalization, so a character built from several
code points, such as a base letter plus a combining accent, is still split apart.
For n bytes and u code points: time O(n), auxiliary space O(u) for the rune slice, plus O(n) output.
*/

// 1. 字节双指针原地交换：推荐
// Reverse the byte slice in place with two pointers.
// 使用左右双指针原地反转字节切片。
func reverseString(s []byte) {
	// left starts at the first character and right starts at the last character.
	// left 从第一个字符开始，right 从最后一个字符开始。
	left, right := 0, len(s)-1

	// Stop when the pointers meet or cross because every pair has been swapped.
	// 当两个指针相遇或交错时停止，因为所有成对字符都已交换。
	for left < right {
		// Swap the characters at the two ends of the current range.
		// 交换当前区间两端的字符。
		s[left], s[right] = s[right], s[left]
		// Move both pointers toward the center.
		// 两个指针同时向中间移动。
		left++
		right--
	}
}

// 2. rune 双指针：多字节字符不会被拆坏
func reverseStringUnicode(s string) string {
	// []rune(s) decodes UTF-8, so every element is one complete code point instead of one byte.
	// []rune(s) 会解码 UTF-8，因此每个元素都是一个完整码点，而不是一个字节。
	runes := []rune(s)

	// The two-pointer swap is identical; only the element type changed.
	// 双指针交换的写法完全相同，只是元素类型变成了 rune。
	left, right := 0, len(runes)-1

	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	// Converting back re-encodes each code point as valid UTF-8.
	// 转换回字符串时，每个码点会被重新编码成合法的 UTF-8。
	return string(runes)
}
