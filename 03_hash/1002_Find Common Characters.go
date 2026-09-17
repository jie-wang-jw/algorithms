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
*/

// 1. 26-slot frequency minima (recommended)
// 1. 26 计数数组取最小值（推荐）
// Seed from the first string, then take per-letter min against every later string; expand surviving counts into the answer.
// 用第一个字符串初始化频次，再对其余串逐字母取 min；把幸存次数展开成结果。
// Time: O(L) for L total characters, Space: O(1).
// 时间复杂度：O(L)（L 为全部字符数），空间复杂度：O(1)。
//
// 仅适用于小写 a-z。hash[c] 是已处理所有单词中 c 次数的最小值；新单词逐字母取 min，得到多重集合交集。
// For lowercase a-z, hash[c] is the minimum count across processed words; taking minima computes the multiset intersection.
// 输出 hash[c] 份字母而非一份，才能保留重复公共字符；空输入返回空结果。
// Emit hash[c] copies, not just one, to preserve repeated common characters; empty input yields no result.
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
		for k := range 26 {
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
// minCount 只需保存首个单词中的字节，因为公共字符不可能来自首词之外；每个词都将次数更新为较小值。
// Only bytes in the first word can be common; reduce each count to the minimum seen across words.
// 按次数输出保留重复项；map 遍历顺序不固定。按题目小写字母约束使用，不是 Unicode 字符统计。
// Emit each retained multiplicity; map order is unspecified. Use the lowercase-letter contract, not Unicode character counting.
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

	var result []string
	for c, cnt := range minCount {
		for range cnt {
			result = append(result, string(c))
		}
	}
	return result
}
