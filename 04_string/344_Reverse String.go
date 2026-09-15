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

// 1. Two-pointer in-place byte swap (recommended)
// 1. 字节双指针原地交换（推荐）
// Swap from both ends until left meets right. Correct for printable ASCII (one byte per character).
// 从两端向中间交换，直到指针相遇。本题可打印 ASCII 下一字符一字节，结果正确。
// Time: O(n), Space: O(1).
// 时间复杂度：O(n)，空间复杂度：O(1)。
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
// 步骤与要点 / Steps and notes:
//  1. []rune(s) decodes UTF-8 so each element is one complete code point, not one byte.
//     []rune(s) 解码 UTF-8，每个元素是完整码点而不是单个字节。
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
