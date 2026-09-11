package _4_string

/*
题目描述 / Problem Description
给定非空字符串 s，判断它能否由某个非空子串重复多次构成。
Given a nonempty string s, determine whether it can
be constructed by repeating one of its nonempty substrings multiple times.

解题思路 / Solution Approach
文件提供三种方法：枚举可能的重复单元、在 (s+s)[1:2n-1] 中查找 s，
以及利用 KMP 最长相等前后缀判断字符串长度能否被重复周期整除。
The file provides enumeration, doubled-string matching,
and KMP solutions. KMP derives a possible period from the longest
equal prefix and suffix and checks whether it divides the string length.

关键逻辑：为什么这样做 / Why This Works
枚举法：重复至少两次，所以单元长度不超过 n/2；长度必须整除 n，否则最后剩下不足一个单元。
双倍字符串法：从 s+s 的内部起点 p（1<=p<n）取长度 n，就是把 s 循环移动 p 位。若仍等于 s，
按这个位移反复对应的字符都相等，字符串由长度 gcd(n,p) 的块重复组成；
去首尾正是排除 p=0、n 这两个任何字符串都能匹配的情况。
KMP 法：设最长相等真前后缀长度为 L，p=n-L。前后缀相等意味着 s[p:]=s[:n-p]，
即每个位置与前面相距 p 的位置相等，p 就是最小周期；只有 L>0 且 n%p==0，
才能把整串切成至少两个完整周期。例如 "ababa" 的 L=3、p=2，但 5%2!=0，最后半个周期不算完整重复。
Enumeration: at least two copies require block length <=n/2, and the length must divide n.
Doubling: a length-n match starting at internal offset p (1<=p<n) is a rotation by p.
Equality forces characters in each repeated-shift orbit to agree, yielding repetition
of blocks of length gcd(n,p). Removing both ends excludes trivial matches at offsets 0 and n.
KMP: for longest proper border length L, let p=n-L. Border equality means s[p:]=s[:n-p],
so characters p positions apart agree; p is the shortest period. L>0 and n%p==0 ensure
at least two complete copies. For "ababa", L=3 and p=2, but 5%2!=0 leaves an incomplete copy.

时间与空间复杂度 / Time and Space Complexity
n = len(s)。枚举法 repeatedSubstringPattern：保守时间上界 O(n²)，枚举 O(n) 个长度，
候选长度最多比较 O(n) 字节；辅助空间 O(1)，子串切片不复制内容。
双倍字符串法 repeatedSubstringPattern2：构造 O(n) 时间和空间，
总时间 O(n + 子串搜索成本)，不能仅因调用 Contains 就断言最坏 O(n)；
朴素搜索分析的保守上界为 O(n²)。KMP 法 repeatedSubstringPattern3：
时间 O(n)，辅助空间 O(n) 保存前缀表。
n = len(s). Enumeration has a conservative O(n²) time bound and O(1)
auxiliary space; substring slices do not copy bytes. The doubled-string
version uses O(n) construction time and space, plus substring-search time;
calling Contains alone does not prove worst-case linear time, and naive-search
analysis gives a conservative O(n²) bound. The KMP version takes O(n)
time and O(n) auxiliary space for the prefix table.

边界条件 / Edge Cases
题目保证 s 非空。枚举法对空串天然返回 false；双倍字符串法会切出非法下标 doubled[1:-1]，
KMP 法会读到越界的 next[n-1]，因此两者都先显式判空，使三种解法对空串结果一致。
The problem guarantees a nonempty s. Enumeration already returns false for an empty string,
while doubling would slice doubled[1:-1] and KMP would read next[n-1] out of range,
so both check for an empty string first and all three solutions now agree.
*/

import "strings"

/*
Check whether the string can be constructed by repeating one of its substrings.
判断字符串能否由它的某个非空子串重复多次构成。
*/

// Enumeration: try every possible pattern length. / 枚举法：尝试每种可能的重复单元长度。
func repeatedSubstringPattern(s string) bool {
	n := len(s)

	// Try every possible substring length.
	// 尝试每一种可能的子串长度。
	for length := 1; length <= n/2; length++ {
		// The substring length must divide the whole string length.
		// 子串长度必须能整除整个字符串长度。
		if n%length != 0 {
			continue
		}

		// pattern is the candidate repeated substring.
		// pattern 是候选的重复子串。
		pattern := s[:length]

		// Check whether every block equals pattern.
		// 检查每一段是否都等于 pattern。
		ok := true
		for start := length; start < n; start += length {
			if s[start:start+length] != pattern {
				ok = false
				break
			}
		}

		if ok {
			// Every block matched the candidate pattern.
			// 每一段都与候选 pattern 相同，说明字符串可由它重复构成。
			return true
		}
	}

	return false
}

/*
If s is made of a repeated substring, s appears inside (s+s) after removing the first and last characters.
如果 s 由重复子串组成，那么在 (s+s) 去掉首尾字符后，仍然能找到完整的 s。
*/
func repeatedSubstringPattern2(s string) bool {
	// An empty string has no nonempty unit to repeat, and doubled[1:len-1] would be an invalid slice.
	// 空串没有可重复的非空单元，而且 doubled[1:len-1] 的下标是非法的。
	if len(s) == 0 {
		return false
	}

	// Doubling contains every rotation of s.
	// s+s 包含 s 的所有循环位移结果。
	doubled := s + s
	// Remove both ends so the two trivial copies of s cannot be matched directly.
	// 去掉首尾字符，避免直接匹配原本位于两端的完整 s。
	middle := doubled[1 : len(doubled)-1]
	return strings.Contains(middle, s)
}

/*
KMP
The longer the equal prefix and suffix, the more the string overlaps with itself.
最长相等前后缀越长，说明字符串前后重叠越多。

Pattern length = n - next[n-1]. / 重复单元长度 = n - next[n-1]。
If n is divisible by that length, the pattern repeats exactly. / 如果 n 能整除该长度，就是重复子串。
*/
func repeatedSubstringPattern3(s string) bool {
	n := len(s)

	// An empty string has no last prefix-table entry to read, so next[n-1] would panic.
	// 空串没有前缀表的最后一项可读，直接访问 next[n-1] 会越界。
	if n == 0 {
		return false
	}

	// Build the prefix table for s.
	// 给 s 构造前缀表。
	next := make([]int, n)
	getNext(next, s)

	// longestPrefixSuffix is the longest prefix length that is also a suffix.
	// longestPrefixSuffix 是整个字符串的最长相等前后缀长度。
	// 之所以取最后一个，只因为最后一个位置代表的范围是整个字符串。
	longestPrefixSuffix := next[n-1]

	// If there is no repeated prefix/suffix, it cannot be built by repetition.
	// 如果没有相等前后缀，就不可能由重复子串组成。
	if longestPrefixSuffix == 0 {
		return false
	}

	// patternLength is the smallest possible repeated block length.
	// patternLength 是可能的最小重复单元长度。
	patternLength := n - longestPrefixSuffix

	// If n can be divided by patternLength, s is repeated by that block.
	// 如果总长度能被 patternLength 整除，说明可以完整重复。
	return n%patternLength == 0
}
