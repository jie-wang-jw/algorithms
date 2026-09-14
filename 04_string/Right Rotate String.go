package _4_string

/*
卡码 55. 右旋字符串 / Right-Rotate String

题目描述 / Problem Description
把字符串尾部的 k 个字符移到前面。例如 s="abcdefg"、k=2 → "fgabcde"。
Move the last k characters of s to the front. For s="abcdefg" and k=2 the result is "fgabcde".

题目要求尽量在本串上操作、不申请与串等长的额外工作区（Go 的 string 不可变，需要 []byte 副本）。
The article asks for in-string reversals without an extra n-length workspace.
Go strings are immutable, so a []byte copy is still required.

解法一：先整体反转再反转两段 / Method 1: Reverse All, Then Reverse Each Part
右旋 k 位等于把后段放到前段前面。整体反转先交换两段的位置，再分别反转两段以恢复各自内部顺序。
k 先对长度取模，避免 k 大于长度时下标越界。
Right rotation by k swaps the two blocks. Reversing the whole string swaps the blocks;
reversing each block restores letter order inside them. Reduce k modulo n so k larger than n stays in range.

解法二：先反转两段再整体反转 / Method 2: Reverse Each Part, Then Reverse All
先反转长度为 n-k 的前段和长度为 k 的后段，再整体反转，得到同样的右旋结果。
Reverse the prefix of length n-k and the suffix of length k, then reverse the whole string.

时间与空间复杂度 / Time and Space Complexity
n 为字节长度。两种解法各做三次线性反转，时间 O(n)。
Go 中 []byte(s) 副本占 O(n)；反转步骤本身是 O(1) 额外空间。
For n bytes both methods reverse three ranges, so time is O(n).
The []byte(s) copy uses O(n) space in Go; the reversals themselves use O(1) extra space.
*/

// 1. Reverse the whole string, then reverse the new prefix of length k and the suffix.
// 1. 先整体反转，再反转长度为 k 的前段和后面的后缀。
// Time: O(n), Space: O(n) for the []byte copy.
// 时间复杂度：O(n)，空间复杂度：O(n)，由 []byte(s) 副本产生。
func rightRotateString(s string, k int) string {
	b := []byte(s)
	n := len(b)
	if n == 0 {
		return s
	}
	k %= n
	reverse(b, 0, n-1)
	reverse(b, 0, k-1)
	reverse(b, k, n-1)
	return string(b)
}

// 2. Reverse the prefix of length n-k and the last k characters, then reverse the whole string.
// 2. 先反转长度为 n-k 的前段和末尾 k 个字符，再整体反转。
// Time: O(n), Space: O(n) for the []byte copy.
// 时间复杂度：O(n)，空间复杂度：O(n)，由 []byte(s) 副本产生。
func rightRotateStringPartsFirst(s string, k int) string {
	b := []byte(s)
	n := len(b)
	if n == 0 {
		return s
	}
	k %= n
	reverse(b, 0, n-k-1)
	reverse(b, n-k, n-1)
	reverse(b, 0, n-1)
	return string(b)
}
