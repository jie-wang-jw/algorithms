package _7_backtracking

/*
131. 分割回文串 / Palindrome Partitioning
93. 复原 IP 地址 / Restore IP Addresses

题目描述 / Problem Description
131：把 s 分割成若干回文子串，返回所有分割方案。
93：把只含数字的 s 分成恰好 4 段合法 IP，返回所有 IP 字符串。
131: split s into palindromic parts and return every partition.
93: split a digit string into exactly four valid IP segments and return those addresses.

示例 / Examples
131: s="aab" → [["a","a","b"],["aa","b"]]
93: s="25525511135" → ["255.255.11.135","255.255.111.35"]
*/

// 面试首选 / Interview first choice：131 分割回文串。
//
// 1. Palindrome partitioning: cut a palindromic prefix at every start index, then partition the suffix.
// 1. 分割回文串：在每个起点切开一段回文前缀，再分割剩下的后缀。
// Time: O(n*2^n), Space: O(n) auxiliary plus output.
// 时间复杂度：O(n*2^n)，空间复杂度：辅助 O(n)，另加输出。
//
// start 是下一段起点；枚举段尾 i，仅当 s[start:i+1] 回文才继续划分剩余串，故每条叶路径恰好覆盖整串。
// start begins the next piece; recurse past each palindromic s[start:i+1], so a completed path covers the whole string.
// start==len(s) 保存 path 副本；返回时撤销本段，避免不同划分相互污染。按题目单字节字符使用。
// At exhaustion save a copy of path; pop the piece on return to isolate partitions. Assume single-byte characters.
func partition(s string) [][]string {
	result := [][]string{}
	path := []string{}

	var backtrack func(int)
	backtrack = func(start int) {
		if start == len(s) {
			saved := make([]string, len(path))
			copy(saved, path)
			result = append(result, saved)
			return
		}
		for i := start; i < len(s); i++ {
			if !isPalindromeRange(s, start, i) {
				continue
			}
			path = append(path, s[start:i+1])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return result
}

// Two-pointer palindrome check on s[left..right].
// 用双指针判断 s[left..right] 是否回文。
// Time: O(n), Space: O(1).
// 时间复杂度：O(n)，空间复杂度：O(1)。
//
// 比较闭区间 [left,right] 的对称字节，一处不同即可失败；调用方保证非空区间的下标有效。
// Compare mirrored bytes in inclusive [left,right]; any mismatch fails. Nonempty bounds must be valid.
func isPalindromeRange(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 面试首选 / Interview first choice：93 复原 IP 地址。
//
// 2. Restore IP addresses: cut 1..3 digits per segment with range and leading-zero checks, exactly four times.
// 2. 复原 IP：每段切 1..3 位数字，检查范围和前导零，恰好切成四段。
// Time: O(1) for length <= 12, Space: O(1) auxiliary plus output.
// 时间复杂度：长度不超过 12，视为 O(1)，空间复杂度：辅助 O(1)，另加输出。
//
// 每层选一个 1..3 位的段，恰好四段且耗尽输入才输出；剩余字符数必须介于剩余段数和其三倍之间。
// Choose one 1..3-digit segment per level; emit only four segments consuming all input. Remaining digits must fit the remaining segment count.
// 段只允许单独的 "0"，多位不能前导零且数值<=255；输入按题意仅含十进制数字。
// Allow "0" but no multi-digit leading zero, with value<=255; assume decimal-digit input.
func restoreIpAddresses(s string) []string {
	result := []string{}
	path := []string{}

	var backtrack func(int)
	backtrack = func(start int) {
		if len(path) == 4 {
			if start == len(s) {
				result = append(result, path[0]+"."+path[1]+"."+path[2]+"."+path[3])
			}
			return
		}
		if start >= len(s) {
			return
		}

		remainParts := 4 - len(path)
		remainChars := len(s) - start
		if remainChars < remainParts || remainChars > remainParts*3 {
			return
		}

		maxLen := 3
		if maxLen > remainChars {
			maxLen = remainChars
		}
		for length := 1; length <= maxLen; length++ {
			segment := s[start : start+length]
			if !validIPSegment(segment) {
				continue
			}
			path = append(path, segment)
			backtrack(start + length)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return result
}

// A valid IP segment is 0, or a number 1..255 with no leading zero.
// 合法 IP 段是单独的 0，或 1..255 且没有前导零。
// Time: O(1), Space: O(1).
// 时间复杂度：O(1)，空间复杂度：O(1)。
//
// 检查长度 1..3、无多位前导零、数值<=255；此辅助函数依赖调用方传入纯数字段，不验证其他字符。
// Check length 1..3, no multi-digit leading zero, and value<=255; the caller guarantees digit-only input.
func validIPSegment(segment string) bool {
	if len(segment) == 0 || len(segment) > 3 {
		return false
	}
	if len(segment) > 1 && segment[0] == '0' {
		return false
	}
	value := 0
	for i := 0; i < len(segment); i++ {
		value = value*10 + int(segment[i]-'0')
	}
	return value <= 255
}
