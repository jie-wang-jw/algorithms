package _3_hash

import "sort"

/*
题目描述 / Problem Description
给定两个只包含小写英文字母的字符串 s 和 t，判断 t 是否是 s 的字母异位词，
即两者是否包含完全相同且数量相同的字符。
Given two lowercase English strings s and t, determine whether t is an anagram of s,
meaning both contain exactly the same characters with the same frequencies.
*/

// 1. Fixed 26-slot frequency array: count s, then decrement with t.
// 1. 固定 26 计数数组：分两次遍历。
// Time: O(n+m), Space: O(1).
// 时间复杂度：O(n+m)，空间复杂度：O(1)。
//
// 按小写 a-z 计数，s 加一、t 减一；最终全部为零等价于每个字母次数相同。
// Count lowercase a-z positively for s and negatively for t; all zeros mean identical letter multiplicities.
func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	record := [26]int{}

	for _, r := range s {
		record[r-'a']++
	}
	for _, r := range t {
		record[r-'a']--
	}

	return record == [26]int{}
}

// 2. Fixed 26-slot frequency array: increment s and decrement t in one pass.
// 2. 固定 26 计数数组：一次遍历同时加减。
// Time: O(n+m), Space: O(1).
// 时间复杂度：O(n+m)，空间复杂度：O(1)。
//
// 先确认长度相同，才能在同一循环安全读取 s[i]、t[i]；26 个计数的最终差值都为零才是异位词。
// Equal lengths permit paired indexing; anagrams require every one of the 26 lowercase-letter count differences to be zero.
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	record := [26]int{}

	for i := 0; i < len(s); i++ {
		record[s[i]-'a']++
		record[t[i]-'a']--
	}

	for _, v := range record {
		if v != 0 {
			return false
		}
	}

	return true
}

// 3. Sort then compare: the character set is not limited to 26 letters.
// 3. 排序后比较：字符集不受 26 个字母限制。
// Time: O(n log n), Space: O(n).
// 时间复杂度：O(n log n)，空间复杂度：O(n)。
//
// 排序后相同字节会相邻；两份排序结果相等恰好表示各字节出现次数一致。按题目字母输入使用。
// Sorted byte slices are equal exactly when byte multiplicities agree; use the problem's letter input contract.
func isAnagramSorted(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	a, b := []byte(s), []byte(t)

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	return string(a) == string(b)
}

// 4. Rune frequency map: count by Unicode code point.
// 4. rune 频次表：按 Unicode 码点统计。
// Time: O(n+m) expected, Space: O(u) for distinct code points.
// 时间复杂度：期望 O(n+m)，空间复杂度：O(u)，u 为不同码点数。
//
// 按 rune 统计 Unicode 码点次数，不按 UTF-8 字节；只比较码点多重集合，不做大小写折叠或 Unicode 规范化。
// Count Unicode code points via rune, not UTF-8 bytes; compare multisets without case folding or Unicode normalization.
func isAnagramUnicode(s, t string) bool {
	counts := make(map[rune]int)

	for _, r := range s {
		counts[r]++
	}
	for _, r := range t {
		counts[r]--
	}

	for _, count := range counts {
		if count != 0 {
			return false
		}
	}

	return true
}
