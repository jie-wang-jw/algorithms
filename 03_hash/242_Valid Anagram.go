package _3_hash

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
