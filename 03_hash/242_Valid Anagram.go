package _3_hash

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
*/

/*
I use a frequency array to count characters.
我使用固定长度的频次数组统计字符出现次数。
For each position, I increment the count for s and decrement the count for t.
遇到 s 中的字符就加 1，遇到 t 中的字符就减 1。
If the two strings are anagrams, all counts should become zero.
如果两个字符串是字母异位词，最后每个字符的计数都应该回到 0。
*/

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
