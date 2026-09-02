package _4_string

/*
题目描述 / Problem Description
给定字符数组 s，原地反转数组中的字符。要求使用 O(1) 额外空间完成。
Given a character array s, reverse its characters in place using O(1) extra space.

解题思路 / Solution Approach
使用左右双指针，从数组两端向中间移动，每轮交换两个指针指向的字符，直到指针相遇或交错。
Use two pointers starting at both ends. Swap their characters and move inward until the pointers meet or cross.
*/

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
