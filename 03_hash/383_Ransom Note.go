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

解法一：暴力删除 / Method 1: Brute-Force Erase
外层扫描 magazine，内层在 ransomNote 的副本里找相同字符并删掉。删空则成功。
Scan magazine; for each character, erase one matching byte from a copy of ransomNote. Success is an empty copy.

解法二：26 计数数组（推荐） / Method 2: 26-Slot Frequency Array (Recommended)
先统计 magazine 各字母次数，再遍历 ransomNote 逐个减一；减到负数说明不够。
若 ransomNote 更长可直接失败。
Count magazine letters, then decrement for each ransomNote letter; a negative count means magazine is short.
If ransomNote is longer than magazine, return false immediately.

关键逻辑：为什么这样做 / Why This Works
magazine 提供可用字母库存；ransomNote 的每个字母都要消耗一次库存。
计数减到负数，说明某个字母的需求超过了供给。长度检查是剪枝：
ransomNote 更长时，即使 magazine 字母种类齐全，总数也一定不够。
Magazine is the available letter inventory; each ransomNote letter spends one unit.
A negative count means demand exceeded supply for that letter. The length check is pruning:
if ransomNote is longer, total supply cannot cover demand even when letter kinds match.

时间与空间复杂度 / Time and Space Complexity
n = len(ransomNote)，m = len(magazine)。暴力法时间 O(n*m)，每次删除还可能移动后缀。
哈希法时间 O(n+m)。两种辅助空间都是 O(1) 量级；暴力法的副本是 O(n) 字节。
For lengths n and m, brute force takes O(n*m) time because each erase may shift a suffix.
The frequency array takes O(n+m) time. Auxiliary space is O(1) for the array and O(n) for the brute-force copy.
*/

// 1. Brute-force erase
// 1. 暴力删除
// For each magazine letter, erase at most one matching byte from a mutable copy of ransomNote; empty copy means success.
// 对 magazine 每个字母，在 ransomNote 的可变副本里最多删掉一个相同字节；副本删空即成功。
// Time: O(n*m), Space: O(n) for the copy.
// 时间复杂度：O(n*m)，空间复杂度：O(n)，由副本产生。
//
// 步骤与要点 / Steps and notes:
//  1. Join prefix and suffix to erase note[j]; one magazine letter spends at most one match.
//     拼接前后缀删掉 note[j]；一个 magazine 字母最多只消耗一次匹配。
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
// 步骤与要点 / Steps and notes:
//  1. Longer note needs more letters than magazine can possibly supply.
//     ransomNote 更长时，总需求一定超过 magazine 总供给，可直接失败。
//  2. Map 'a'..'z' onto slots 0..25.
//     用 v-'a' 把 'a'..'z' 映射到下标 0..25。
//  3. Negative means this letter was demanded more times than magazine provided.
//     变成负数，说明该字母的需求超过了 magazine 的供给。
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
