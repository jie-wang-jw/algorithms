package _4_string

/*
题目描述 / Problem Description
给定非空字符串 s，判断它能否由某个非空子串重复多次构成。
Given a nonempty string s, determine whether it can
be constructed by repeating one of its nonempty substrings multiple times.
*/

import "strings"

// 1. Enumeration: try every possible repeated-unit length that divides n.
// 1. 枚举法：尝试每种能整除 n 的重复单元长度。
// Time: O(n²) conservative, Space: O(1).
// 时间复杂度：保守上界 O(n²)，空间复杂度：O(1)。
//
// 至少重复两次，所以单元长度最多 n/2，且须整除 n；逐块比较是否都等于首块，有一块不同就排除此长度。
// At least two copies require a block length <=n/2 dividing n; reject a length as soon as any block differs from the first.
func repeatedSubstringPattern(s string) bool {
	n := len(s)

	for length := 1; length <= n/2; length++ {
		if n%length != 0 {
			continue
		}

		pattern := s[:length]

		ok := true
		for start := length; start < n; start += length {
			if s[start:start+length] != pattern {
				ok = false
				break
			}
		}

		if ok {
			return true
		}
	}

	return false
}

// 2. Doubled-string search: s appears inside (s+s) after the first and last characters are removed.
// 2. 双倍字符串查找：去掉 (s+s) 的首尾后仍能找到 s。
// Time: O(n²) conservative for naive search, Space: O(n) for the doubled string.
// 时间复杂度：朴素搜索保守上界 O(n²)，空间复杂度：O(n)，由双倍字符串产生。
//
// s+s 从位置 p 取长度 n，得到 s 循环移动 p 位的结果；去首尾排除 p=0、n 两个必然匹配的起点。
// A length-n slice at p in s+s rotates s by p; removing the ends excludes trivial matches at p=0 and p=n.
// 若某个 0<p<n 仍匹配，反复移动 p 位对应的字符都相等，故由长度 gcd(n,p) 的块重复组成。
// A match at 0<p<n makes all positions linked by repeated shifts equal, yielding repeated blocks of length gcd(n,p).
// 先判空，避免 doubled[1:len-1] 构成非法区间。
// Reject empty input before slicing, which would otherwise create invalid bounds.
func repeatedSubstringPattern2(s string) bool {
	if len(s) == 0 {
		return false
	}

	doubled := s + s
	middle := doubled[1 : len(doubled)-1]
	return strings.Contains(middle, s)
}

// 3. KMP longest border: the candidate period is n-L and must divide n.
// 3. KMP 最长相等前后缀：候选周期为 n-L，且必须整除 n。
// Time: O(n), Space: O(n) for the prefix table.
// 时间复杂度：O(n)，空间复杂度：O(n)，由前缀表产生。
//
// L=next[n-1] 对应整串最长相等真前后缀；令 p=n-L，则 s[:L]==s[p:]，即相距 p 的字节相等。
// L=next[n-1] is the whole string's longest proper border; p=n-L gives s[:L]==s[p:], so bytes p apart agree.
// 前后缀越长，错开距离越小，因此 p 是最小周期；L>0 确保 p<n，排除只把整串取一次。
// A longer border means a smaller shift, so p is the shortest period; L>0 ensures p<n rather than one whole-string copy.
// 还要 n%p==0 才能全由完整单元组成："abab" 成立；"ababa" 虽有周期 2，却剩半个单元，不成立。
// Require n%p==0 for complete copies: "abab" qualifies, while period-2 "ababa" ends with an incomplete copy.
func repeatedSubstringPattern3(s string) bool {
	n := len(s)

	if n == 0 {
		return false
	}

	next := make([]int, n)
	getNext(next, s)

	longestPrefixSuffix := next[n-1]

	if longestPrefixSuffix == 0 {
		return false
	}

	patternLength := n - longestPrefixSuffix

	return n%patternLength == 0
}
