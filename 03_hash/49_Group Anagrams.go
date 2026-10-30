package _3_hash

/*
49. 字母异位词分组 / Group Anagrams

将只含小写英文字母的字符串按字母异位词分组，结果顺序不限。
Group lowercase English strings by anagram; the result may be in any order.

["eat","tea","tan","ate","nat","bat"] → [["eat","tea","ate"],["tan","nat"],["bat"]]
*/

// 面试首选 / Interview first choice：49 字母异位词分组。
//
// 状态：count[i] 是当前字符串中字母 'a'+i 的次数；groups[count] 保存计数相同的字符串。
// State: count[i] counts letter 'a'+i in the current string; groups[count] stores strings with that signature.
// 正确性：两个串互为异位词，当且仅当每种字母的次数相同；[26]int 按数组值比较，可直接作 map 的键。
// Correctness: two strings are anagrams exactly when all letter counts agree; [26]int keys compare by value.
// 边界：每个字符串重新计数；空串使用全零键，重复字符串保留，空输入返回空结果；仅适用于小写 a-z。
// Boundaries: reset counts per string; empty strings share the zero key, duplicates remain, and empty input yields no groups. Lowercase a-z only.
// 时间：期望 O(n+L)，n 为字符串数，L 为总字符数；内层合计扫描 L 个字符，每个键固定 26 项。
// Time: expected O(n+L) for n strings and L total characters; all inner loops scan L characters, with fixed-size keys.
// 空间：O(n)，包含分组表和结果中的字符串引用；不复制字符串的底层字符，map 遍历不保证组的顺序。
// Space: O(n) for grouping and result string references; character data is not copied and map iteration does not fix group order.
func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)

	for _, s := range strs {
		var count [26]int
		for _, ch := range s {
			count[ch-'a']++
		}
		groups[count] = append(groups[count], s)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}
