package _3_hash

/*
1002. 查找常用字符 / Find Common Characters

题目描述 / Problem Description
找出字符串数组 words 中每个字符串都出现的共用字符（含重复次数），按任意顺序返回。
words[i] 只含小写字母。
Return the characters that appear in every string of words, including duplicates, in any order.
Each words[i] contains only lowercase letters.

示例 / Example
["bella","label","roller"] → ["e","l","l"]

解法一：26 计数数组取最小值（推荐） / Method 1: 26-Slot Frequency Minima (Recommended)
用第一个字符串初始化 26 格频次数组。对每个后续字符串再统计一遍，逐格取 min。
最后把每个字母按其最小次数展开成结果。
Initialize a 26-slot array from the first string. Count each later string and take min per letter.
Expand each letter by its surviving count into the answer.

解法二：哈希表取最小值 / Method 2: Hash Map Minima
用 map 统计第一个字符串，再对其余字符串分别计数并取 min。字母范围仍是 26，只是容器换成 map。
Count the first string in a map, then take min against each later string. The alphabet is still 26; only the container changes.

关键逻辑：为什么这样做 / Why This Works
共用字符的可输出次数，等于它在每一串中出现次数的最小值：
任何一串更少，最终能同时从所有串里“抠”出的份数就更少。
取 min 后保留的正数，就是答案中该字母应重复出现的次数。
The shared count of a letter equals the minimum of its counts across all strings:
any scarcer string caps how many copies can be taken from every string at once.
Positive values after successive mins are exactly how many times that letter appears in the answer.

时间与空间复杂度 / Time and Space Complexity
n 为字符串个数，L 为所有字符总数。两种解法时间 O(L)。解法一辅助空间 O(1)；解法二 map 最多 26 个键，也记 O(1)。输出最多 O(minLen)。
For n strings and L total characters, both take O(L) time. Method 1 uses O(1) auxiliary space;
method 2 stores at most 26 keys, also O(1). Output is O(minLen).
*/

// 1. 26-slot frequency minima (recommended)
// 1. 26 计数数组取最小值（推荐）
// Seed from the first string, then take per-letter min against every later string; expand surviving counts into the answer.
// 用第一个字符串初始化频次，再对其余串逐字母取 min；把幸存次数展开成结果。
// Time: O(L) for L total characters, Space: O(1).
// 时间复杂度：O(L)（L 为全部字符数），空间复杂度：O(1)。
//
// 步骤与要点 / Steps and notes:
//  1. Shared count cannot exceed what this string can contribute.
//     公共次数不能超过当前字符串能提供的次数。
func commonChars(words []string) []string {
	if len(words) == 0 {
		return []string{}
	}

	hash := make([]int, 26)
	for _, c := range words[0] {
		hash[c-'a']++
	}
	for i := 1; i < len(words); i++ {
		other := make([]int, 26)
		for _, c := range words[i] {
			other[c-'a']++
		}
		for k := 0; k < 26; k++ {
			hash[k] = min(hash[k], other[k])
		}
	}

	result := []string{}
	for i := range 26 {
		for hash[i] > 0 {
			result = append(result, string(rune('a'+i)))
			hash[i]--
		}
	}
	return result
}

// 2. Hash map minima
// 2. 哈希表取最小值
// Same min-count idea with a map; absent letters become 0 via Go's zero value when tightening shared keys.
// 思路相同，容器换成 map；收紧共享键时，本串没有的字母会通过 Go 零值变成 0。
// Time: O(L), Space: O(1) for at most 26 keys.
// 时间复杂度：O(L)，空间复杂度：O(1)，最多 26 个键。
//
// 步骤与要点 / Steps and notes:
//  1. Only keys already in minCount can remain shared; missing keys read as 0.
//     只有已在 minCount 中的字母才可能继续共享；本串没有的键读出来是 0。
func commonCharsMap(words []string) []string {
	if len(words) == 0 {
		return []string{}
	}

	minCount := map[byte]int{}
	for i := 0; i < len(words[0]); i++ {
		minCount[words[0][i]]++
	}
	for i := 1; i < len(words); i++ {
		cur := map[byte]int{}
		for j := 0; j < len(words[i]); j++ {
			cur[words[i][j]]++
		}
		for c := range minCount {
			if cur[c] < minCount[c] {
				minCount[c] = cur[c]
			}
		}
	}

	result := []string{}
	for c, cnt := range minCount {
		for range cnt {
			result = append(result, string(c))
		}
	}
	return result
}
