package _3_hash

/*
383. 赎金信 / Ransom Note

题目描述 / Problem Description
判断 ransomNote 能否由 magazine 中的字符构成。每个 magazine 字符只能用一次。两串均只含小写字母。
Return whether ransomNote can be built from magazine characters. Each magazine character may be used once.
Both strings contain only lowercase letters.

示例 / Example
canConstruct("a", "b") = false
canConstruct("aa", "aab") = true
*/

// 1. Brute-force erase
// 1. 暴力删除
// For each magazine letter, erase at most one matching byte from a mutable copy of ransomNote; empty copy means success.
// 对 magazine 每个字母，在 ransomNote 的可变副本里最多删掉一个相同字节；副本删空即成功。
// Time: O(n*m), Space: O(n) for the copy.
// 时间复杂度：O(n*m)，空间复杂度：O(n)，由副本产生。
//
// note 保存尚未满足的字母；每个 magazine 字节最多删去一项，删除后 break 防止把同一个来源字母重复使用。
// note holds unmet letters; each magazine byte removes at most one item, with break preventing reuse of that source byte.
func canConstructBruteForce(ransomNote string, magazine string) bool {
	note := []byte(ransomNote)
	for i := 0; i < len(magazine); i++ {
		for j := 0; j < len(note); j++ {
			if magazine[i] == note[j] {
				note = append(note[:j], note[j+1:]...)
				break
			}
		}
	}
	return len(note) == 0
}

// 2. 26-slot frequency array (recommended)
// 2. 26 计数数组（推荐）
// Count magazine letters into a 26-slot array, then spend them for ransomNote; a negative count means shortfall.
// 先把 magazine 字母记进 26 格数组，再按 ransomNote 消耗；减成负数说明不够。
// Time: O(n+m), Space: O(1).
// 时间复杂度：O(n+m)，空间复杂度：O(1)。
//
// record 表示 magazine 的剩余字母数量，每取一个就减一；出现负数说明该字母供应不足，可立即失败。
// record tracks remaining magazine letters; consuming one decrements its count, and a negative count proves insufficiency.
// 仅适用于题目中的小写 a-z；短于 ransomNote 的来源串不可能满足需求。
// Use lowercase a-z only; a source shorter than ransomNote cannot supply all required letters.
func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	record := make([]int, 26)
	for _, v := range magazine {
		record[v-'a']++
	}
	for _, v := range ransomNote {
		record[v-'a']--
		if record[v-'a'] < 0 {
			return false
		}
	}
	return true
}
