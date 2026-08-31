package _4_string

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
