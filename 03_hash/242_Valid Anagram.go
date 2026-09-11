package _3_hash

import "sort"

/*
题目描述 / Problem Description
给定两个只包含小写英文字母的字符串 s 和 t，判断 t 是否是 s 的字母异位词，
即两者是否包含完全相同且数量相同的字符。
Given two lowercase English strings s and t, determine whether t is an anagram of s,
meaning both contain exactly the same characters with the same frequencies.

解题思路 / Solution Approach
使用长度为 26 的频次数组。遍历 s 时增加计数，遍历 t 时减少计数；
最终所有计数均为 0 时，两者互为字母异位词。
Use a frequency array of length 26. Increment counts for s and decrement them for t;
the strings are anagrams exactly when every final count is zero.

时间与空间复杂度 / Time and Space Complexity
n = len(s)，m = len(t)。isAnagram1 和 isAnagram2 均为 O(n+m) 时间上界，
长度不同会 O(1) 提前返回。辅助空间 O(1)，字符范围固定为 26 个小写字母。
For lengths n and m, both isAnagram1 and isAnagram2 have an O(n+m) time bound,
with O(1) early return for unequal lengths. Auxiliary space O(1) uses a fixed 26-letter frequency array.

补充解法 / Additional Approaches
isAnagramSorted：复制为字节数组后分别排序，比较结果；排序把字节多重集合变成唯一的标准顺序。
它比较的是字节而不是码点，对单字节字符（含本题的小写英文）等价于比较字符多重集合。
时间 O(n log n)，辅助空间 O(n)；本题固定 26 计数更优。
Sort byte copies and compare their canonical order: O(n log n) time, O(n) auxiliary space.
It compares bytes rather than code points, which is equivalent to comparing characters for
single-byte input such as this problem's lowercase English. The 26-counter method is still better here.
isAnagramUnicode：用 rune 频次表，支持超出 26 个字母的字符，不做 Unicode 规范化。
按 Unicode 码点计数，不按 UTF-8 字节；组合字符和预组字符不会自动视为相同。
时间平均 O(n+m)（输入字节数），辅助空间 O(u)（不同码点数）。
Count Unicode code points using a rune map, without normalization.
Combining sequences and precomposed characters are not automatically equivalent.
Expected time O(n+m) in input bytes, auxiliary space O(u) for distinct code points.
*/

/*
I use a frequency array to count characters.
我使用固定长度的频次数组统计字符出现次数。
For each position, I increment the count for s and decrement the count for t.
遇到 s 中的字符就加 1，遇到 t 中的字符就减 1。
If the two strings are anagrams, all counts should become zero.
如果两个字符串是字母异位词，最后每个字符的计数都应该回到 0。
*/

// 1. 固定 26 计数数组：分两次遍历
func isAnagram1(s string, t string) bool {
	/*
		Different lengths cannot contain exactly the same characters.
		长度不同，就不可能包含数量完全相同的字符。
	*/
	if len(s) != len(t) {
		return false
	}

	// record[0] counts 'a', record[1] counts 'b', and so on.
	// record[0] 统计 'a'，record[1] 统计 'b'，以此类推。
	record := [26]int{}

	for _, r := range s {
		// r-'a' converts a lowercase letter into an index from 0 to 25.
		// r-'a' 把小写字母转换成 0 到 25 的数组下标。
		record[r-'a']++
	}
	for _, r := range t {
		record[r-'a']--
	}

	// Go arrays of the same type are directly comparable.
	// Go 中相同类型的数组可以直接比较；全为 0 才说明字符数量一致。
	return record == [26]int{}
}

// 2. 固定 26 计数数组：一次遍历同时加减
func isAnagram2(s string, t string) bool {
	// This version updates counts for s and t in the same loop.
	// 这个版本在同一个循环中同时更新 s 和 t 的计数。
	if len(s) != len(t) {
		return false
	}

	record := [26]int{}

	for i := 0; i < len(s); i++ {
		// Add the character from s and cancel it with the character from t.
		// s 中的字符加 1，t 中的字符减 1，相同字符最终会互相抵消。
		record[s[i]-'a']++
		record[t[i]-'a']--
	}

	for _, v := range record {
		// Any nonzero count means one string has extra occurrences of a letter.
		// 任意计数不为 0，都表示某个字母在两个字符串中的数量不同。
		if v != 0 {
			return false
		}
	}

	return true
}

// 3. 排序后比较：字符集不受 26 个字母限制
func isAnagramSorted(s, t string) bool {
	// Different lengths cannot hold the same multiset of characters.
	// 长度不同，就不可能是同一个字符多重集合。
	if len(s) != len(t) {
		return false
	}

	// Strings are immutable in Go, so sort byte copies instead.
	// Go 的 string 不能原地修改，因此排序的是字节副本。
	a, b := []byte(s), []byte(t)

	// Sorting turns each multiset into its one canonical ordering.
	// 排序把每个字符多重集合变成唯一的标准顺序。
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	// Equal canonical forms mean equal character counts.
	// 标准顺序相同，说明每个字符的数量都相同。
	return string(a) == string(b)
}

// 4. rune 频次表：按 Unicode 码点统计
func isAnagramUnicode(s, t string) bool {
	// Key: one Unicode code point; value: its signed count difference.
	// key 是一个 Unicode 码点，value 是它在两个字符串中的计数差。
	counts := make(map[rune]int)

	// Ranging over a string decodes UTF-8, so r is a code point rather than a byte.
	// range 遍历字符串时会解码 UTF-8，因此 r 是码点而不是单个字节。
	for _, r := range s {
		counts[r]++
	}
	for _, r := range t {
		counts[r]--
	}

	for _, count := range counts {
		// A nonzero difference means one side has extra occurrences of that code point.
		// 计数差不为 0，说明该码点在某一侧出现得更多。
		if count != 0 {
			return false
		}
	}

	// Comparing code points does not apply Unicode normalization.
	// 这里只比较码点，不做 Unicode 规范化，组合字符与预组字符不会自动视为相同。
	return true
}
