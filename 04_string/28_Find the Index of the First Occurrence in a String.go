package _4_string

/*
题目描述 / Problem Description
给定字符串 haystack 和 needle，返回 needle 在 haystack 中第一次出现的起始下标；
如果不存在则返回 -1。
Given strings haystack and needle, return the starting index of the first
occurrence of needle in haystack, or -1 if it does not occur.

解题思路 / Solution Approach
文件提供暴力匹配和 KMP。暴力法尝试每个可能起点；KMP 使用前缀表，
在失配时复用已经匹配的信息，避免回退主串指针。
The file provides brute-force and KMP solutions. Brute force tries every starting position,
while KMP uses a prefix table to reuse matched information without moving the text pointer backward.

关键逻辑：为什么这样做 / Why This Works
KMP 比较前，j 表示已经匹配的长度，主串当前位置之前的 j 个字符等于 needle[:j]。
若下一字符失配，next[j-1] 给出这段已匹配内容的最长相等真前后缀长度 L：
末尾 L 个字符既然等于开头 L 个字符，就可以直接当作新一轮已经匹配的前缀，无需重新扫描主串。
例如已匹配 "abab"，下一字符不是期待的 'a'，可先保留末尾 "ab"，令 j=2，再用同一个主串字符比较 needle[2]；
仍失败就继续沿更短前后缀回退。建表 getNext 使用同样逻辑：寻找能接上新字符的最长旧前后缀。
匹配完成时 i 是末尾下标，起点为 i-m+1。
Before a KMP comparison, j characters immediately before the current text position equal needle[:j].
On mismatch, L=next[j-1] is the longest equal proper prefix/suffix of that matched part.
Its last L characters already equal the first L pattern characters, so reuse them as a matched prefix without rescanning the text.
After matching "abab", a mismatch against the expected 'a' first retains suffix "ab", sets j=2,
and compares the same text character with needle[2]. Further failures follow shorter borders.
getNext applies this same rule to find the longest old border extendable by the new character.
A full match ends at i, so its start is i-m+1.

时间与空间复杂度 / Time and Space Complexity
n = len(haystack)，m = len(needle)。strStr 暴力法：m<=n 时最坏 O((n-m+1)m)，
辅助空间 O(1)。strStr2 KMP：建表 O(m)、匹配 O(n)，总 O(n+m)，辅助空间 O(m)
保存 next。getNext 本身 O(m) 时间，除传入的 next 外 O(1) 空间。
For text length n and pattern length m, strStr takes worst-case O((n-m+1)m)
time when m<=n and O(1) auxiliary space. strStr2 takes O(n+m) time and O(m)
auxiliary space for next. getNext takes O(m) time and O(1) extra space beyond the supplied table.
*/

func strStr(haystack string, needle string) int {
	// If needle is empty, return 0 by convention.
	// 按题目约定，needle 为空时返回 0。
	if len(needle) == 0 {
		return 0
	}

	// If needle is longer than haystack, it cannot match.
	// needle 比 haystack 更长时，不可能匹配成功。
	if len(needle) > len(haystack) {
		return -1
	}

	// i is the starting position where we try to match needle.
	// i 表示本轮尝试匹配 needle 的起始位置。
	//
	// 为什么是 i <= len(haystack)-len(needle)?
	// Because we need enough remaining characters for needle.
	// 因为从 i 开始必须剩下足够多的字符，才能容纳整个 needle。
	for i := 0; i <= len(haystack)-len(needle); i++ {
		// j is the current index inside needle.
		// j 是当前正在比较的 needle 下标。
		j := 0

		// Compare haystack[i+j] with needle[j] one by one.
		// 从起点 i 开始，逐个比较 haystack[i+j] 和 needle[j]。
		for j < len(needle) && haystack[i+j] == needle[j] {
			j++
		}

		// If j reaches len(needle), the whole needle matched.
		// j 到达 len(needle)，表示 needle 的所有字符都匹配成功。
		if j == len(needle) {
			return i
		}
	}

	// Tried all starting positions and found no match.
	// 所有可能起点都尝试过，仍然没有找到匹配。
	return -1
}

/*
KMP avoids moving the text pointer backward.
KMP 的目的：匹配失败时，不让 haystack 的 i 回头，只让 needle 的 j 回退。

When a mismatch happens, it uses the prefix table to move the pattern pointer.
发生不匹配时，利用前缀表移动模式串指针。
*/

func strStr2(haystack string, needle string) int {
	// Empty pattern matches at index 0.
	// 空模式串默认匹配在下标 0。
	if len(needle) == 0 {
		return 0
	}

	// Build the prefix table for needle.
	// 为 needle 构造前缀表 next。
	next := make([]int, len(needle))
	getNext(next, needle)

	// j means how many characters in needle have been matched.
	// j 表示 needle 当前已经匹配了多少个字符。
	j := 0

	// i scans haystack from left to right and never moves backward.
	// i 从左到右扫描 haystack，并且不会回退。
	for i := 0; i < len(haystack); i++ {
		// Mismatch: fall back j using the prefix table.
		// 不匹配：根据 next 数组回退 j。
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}

		// Match: move j forward.
		// 匹配：j 往前走一步。
		if haystack[i] == needle[j] {
			j++
		}

		// If j reaches len(needle), the whole needle has been matched.
		// 如果 j 等于 needle 长度，说明整个 needle 匹配完成。
		if j == len(needle) {
			return i - len(needle) + 1
		}
	}

	// No match found.
	// 没有找到匹配。
	return -1
}

/*
1. Initialize 初始化
2. Mismatch: fall back 前后缀不相同：j 回退
3. Match: move forward 前后缀相同：j 前进
4. Update the prefix table 更新 next 数组
next: prefix table / 前缀表
next[i]: the longest equal proper prefix and suffix length in s[0:i+1]
next[i]：s[0:i+1] 这一段中，最长相等真前缀和真后缀的长度
i=0: "a"       next[0]=0
i=1: "aa"      前缀 "a" == 后缀 "a"        next[1]=1
i=2: "aab"     没有相等前后缀             next[2]=0
i=3: "aaba"    前缀 "a" == 后缀 "a"        next[3]=1
i=4: "aabaa"   前缀 "aa" == 后缀 "aa"      next[4]=2
i=5: "aabaaf"  没有相等前后缀             next[5]=0

i: the character currently being processed. / 当前正在处理的字符位置。
j: the current longest equal prefix-suffix length. / 当前最长相等前后缀的长度。
next[i]: the answer for substring s[0:i+1]. / 到 i 为止这一段的最长相等前后缀长度。
*/
func getNext(next []int, s string) {
	// j is the current matched prefix length.
	// j 表示当前已经匹配上的前后缀长度。
	j := 0
	next[0] = 0

	/*
			i is the position whose next value we are computing; j is the reusable prefix length.
			i 是当前要计算 next[i] 的位置；j 是目前可以复用的前缀长度。

		s[i] is the new character, and s[j] is the next character expected by the old prefix.
			s[i] 是新加入的字符，s[j] 是旧前缀接下来期待匹配的字符。
	*/
	for i := 1; i < len(s); i++ {
		/*
			If s[i] does not match s[j], the current prefix of length j cannot continue.
			如果 s[i] 和 s[j] 不相同，说明当前长度为 j 的前后缀无法继续扩展。

			Shorten j to the next reusable prefix and try again.
			把 j 缩短到更短的可复用前缀，再继续尝试。
		*/
		for j > 0 && s[i] != s[j] {
			// next[j-1] finds the longest border within the matched prefix; its suffix is reusable as a prefix.
			// next[j-1] 找已匹配前缀内部的最长相等前后缀；末尾这段仍能当作开头使用，所以无需从零重试。
			j = next[j-1]
		}

		// If the new characters match, the reusable prefix grows by one.
		// 如果新字符匹配，最长相等前后缀长度增加 1。
		if s[i] == s[j] {
			j++
		}

		// Store the longest equal prefix-suffix length for s[0:i+1].
		// 记录 s[0:i+1] 当前这段的最长相等前后缀长度。
		next[i] = j
	}
}
