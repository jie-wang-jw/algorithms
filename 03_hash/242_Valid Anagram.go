package _3_hash

/*
I use a frequency array to count characters.
For each position, I increment the count for s and decrement the count for t.
If the two strings are anagrams, all counts should become zero.
*/

func isAnagram1(s string, t string) bool {
	/*最好先判断长度。
	如果题目保证都是小写字母，长度不同肯定不是 anagram：*/
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
