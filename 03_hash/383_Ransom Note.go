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

时间与空间复杂度 / Time and Space Complexity
n = len(ransomNote)，m = len(magazine)。暴力法时间 O(n*m)，每次删除还可能移动后缀。
哈希法时间 O(n+m)。两种辅助空间都是 O(1) 量级；暴力法的副本是 O(n) 字节。
For lengths n and m, brute force takes O(n*m) time because each erase may shift a suffix.
The frequency array takes O(n+m) time. Auxiliary space is O(1) for the array and O(n) for the brute-force copy.
*/

// 1. Brute-force erase: delete each magazine character from a copy of ransomNote when it matches.
// 1. 暴力删除：在 ransomNote 的副本里删掉 magazine 中匹配到的字符。
// Time: O(n*m), Space: O(n) for the copy.
// 时间复杂度：O(n*m)，空间复杂度：O(n)，由副本产生。
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

// 2. 26-slot frequency array (recommended): count magazine, then decrement with ransomNote.
// 2. 26 计数数组：推荐；先统计 magazine，再用 ransomNote 逐个减一。
// Time: O(n+m), Space: O(1).
// 时间复杂度：O(n+m)，空间复杂度：O(1)。
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
