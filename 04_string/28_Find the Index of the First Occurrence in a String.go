package _4_string

/*
题目描述 / Problem Description
给定字符串 haystack 和 needle，返回 needle 在 haystack 中第一次出现的起始下标；
如果不存在则返回 -1。
Given strings haystack and needle, return the starting index of the first
occurrence of needle in haystack, or -1 if it does not occur.
*/

// 1. Brute-force sliding compare: try every starting index in haystack.
// 1. 暴力滑动比较：枚举每个起点。
// Time: O((n-m+1)m) worst case when m<=n, Space: O(1).
// 时间复杂度：m<=n 时最坏 O((n-m+1)m)，空间复杂度：O(1)。
//
// 只枚举还能容纳整个 needle 的起点 i<=n-m；从左往右逐个比较，首次完整匹配就是最早出现位置。
// Try only starts i<=n-m that fit needle; left-to-right comparison returns the earliest complete match.
// 空 needle 约定返回 0；needle 更长时直接返回 -1。
// An empty needle returns 0 by convention; a needle longer than haystack cannot match.
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	if len(needle) > len(haystack) {
		return -1
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		j := 0

		for j < len(needle) && haystack[i+j] == needle[j] {
			j++
		}

		if j == len(needle) {
			return i
		}
	}

	return -1
}

// 2. KMP prefix table: reuse matched prefix information so the text pointer never retreats.
// 2. KMP 前缀表：主串指针不回退。
// Time: O(n+m), Space: O(m) for the prefix table.
// 时间复杂度：O(n+m)，空间复杂度：O(m)，由 next 数组产生。
//
// j 表示 haystack[i] 之前已匹配的长度，这段等于 needle[:j]；失配时其末尾 next[j-1] 个字符仍等于模式前缀。
// j matched bytes before haystack[i] equal needle[:j]; on mismatch, the final next[j-1] bytes still equal a pattern prefix.
// 因此令 j=next[j-1] 保留这段匹配，继续用同一个 haystack[i] 尝试；仍失配就继续回退，i 无须回头。
// Set j=next[j-1] to retain those matches and retry the same haystack[i]; repeat fallback without retreating i.
// 相等时 j++；j==m 表示以 i 结尾的完整匹配，起点为 i-m+1。空模式返回 0，遍历完未命中返回 -1。
// Increment j on equality; j==m completes a match ending at i, starting at i-m+1. Empty patterns return 0; exhaustion returns -1.
func strStr2(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	next := make([]int, len(needle))
	getNext(next, needle)

	j := 0

	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}

		if haystack[i] == needle[j] {
			j++
		}

		if j == len(needle) {
			return i - len(needle) + 1
		}
	}

	return -1
}

// getNext builds the shared KMP prefix table. / 构造共用的 KMP 前缀表。
// 调用方提供 len(next)>=len(s)，长度和下标按字节计；空串不写表。
// The caller supplies len(next)>=len(s); indices count bytes and empty input writes nothing.
// Time O(m), extra space O(1) beyond next. / 时间 O(m)，除 next 外空间 O(1)。
// j 每轮最多增加 1，回退严格减小，所以回退总次数 O(m)。
// j increases at most once per round and strictly decreases on fallback, bounding total fallbacks by O(m).
//
// next[i] 是 s[:i+1] 的最长相等真前后缀长度；“真”排除整段本身，例如 "aaa" 的答案是 2。
// next[i] is the longest proper border of s[:i+1]; exclude the whole substring, so "aaa" has length 2.
// 处理 i 前，j=next[i-1] 描述旧串 s[:i]；比较 s[i] 与 s[j]，是在尝试把旧前后缀同时延长一位。
// Before i, j=next[i-1] describes s[:i]; comparing s[i] with s[j] attempts to extend that border by one byte.
// 失配时，长度 j 的已匹配部分末下标为 j-1；next[j-1] 给出它的最长较短前后缀，因此令 j=next[j-1]。
// On mismatch, the matched part of length j ends at j-1; next[j-1] gives its longest shorter border, the next candidate.
// 旧后缀等于 s[:j]，所以这个较短后缀仍等于整串开头；更短候选也必在此回退链上，不会漏解。
// The old suffix equals s[:j], so its shorter border is still a whole-string prefix; all shorter candidates lie on this chain.
// j>0 时才读 next[j-1]，每次回退仍尝试同一个 s[i]；相等则 j++，最终把 j 写到 next[i]。
// Read next[j-1] only for j>0 and retry the same s[i]; increment j on equality, then store it in next[i].
// 例如给 "aabaa" 追加 'a'：j=2 期待 'b'，失败后退到 next[1]=1；此时期待 'a'，成功得到 next[5]=2。
// Appending 'a' to "aabaa": j=2 expects 'b'; fall back to next[1]=1, match 'a', and obtain next[5]=2.
func getNext(next []int, s string) {
	if len(s) == 0 {
		return
	}

	j := 0
	next[0] = 0
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}

		if s[i] == s[j] {
			j++
		}

		next[i] = j
	}
}
