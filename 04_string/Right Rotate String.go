package _4_string

/*
卡码 55. 右旋字符串 / Right-Rotate String

题目描述 / Problem Description
把字符串尾部的 k 个字符移到前面。例如 s="abcdefg"、k=2 → "fgabcde"。
Move the last k characters of s to the front. For s="abcdefg" and k=2 the result is "fgabcde".

题目要求尽量在本串上操作、不申请与串等长的额外工作区（Go 的 string 不可变，需要 []byte 副本）。
The article asks for in-string reversals without an extra n-length workspace.
Go strings are immutable, so a []byte copy is still required.
*/

// 1. Reverse all, then reverse each part
// 1. 先整体反转，再反转两段
// Full reverse swaps the two blocks; then reverse each block to restore letter order inside them.
// 整体反转交换前后两段；再分别反转两段，恢复各自内部字母顺序。
// Time: O(n), Space: O(n) for the []byte copy.
// 时间复杂度：O(n)，空间复杂度：O(n)，由 []byte(s) 副本产生。
//
// 设 s=A+B，B 长 k；整体反转得 reverse(B)+reverse(A)，再分别反转两段即得 B+A。
// Split s=A+B with |B|=k; reverse all, then reverse each part to obtain B+A.
// 空串先返回以避免模零，k%=n 消去整圈；要求 k>=0，按字节旋转。
// Return early for empty input, reduce full turns with k%=n, and require k>=0; rotation is byte-based.
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

// 2. Reverse each part, then reverse all
// 2. 先反转两段，再整体反转
// Same transforms in reverse order: reverse prefix n-k and suffix k, then the whole string.
// 同一组变换的逆序组合：先反转前 n-k 与后 k，再整体反转。
// Time: O(n), Space: O(n) for the []byte copy.
// 时间复杂度：O(n)，空间复杂度：O(n)，由 []byte(s) 副本产生。
//
// 设 s=A+B、B 长 k；先分别反转 A、B，再整体反转，得到 B+A，段内顺序恢复。
// For s=A+B with |B|=k, reverse both parts then the whole buffer to obtain B+A with each part restored.
// 先判空再取 k%=n；要求 k>=0，按字节旋转。
// Guard empty input before k%=n; require nonnegative k and use byte-based rotation.
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
