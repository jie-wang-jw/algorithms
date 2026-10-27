package _7_backtracking

/*
17. 电话号码的字母组合 / Letter Combinations of a Phone Number

题目描述 / Problem Description
给定仅含数字 2-9 的字符串 digits，返回它能表示的所有字母组合。答案顺序不限。空串返回空结果。
Given a string of digits from 2-9, return every letter combination those digits can represent.
Any order is acceptable. An empty input yields an empty result.

示例 / Example
digits="23" → ["ad","ae","af","bd","be","bf","cd","ce","cf"]
*/

// 面试首选 / Interview first choice：17 电话号码的字母组合。
//
// 1. Backtracking: each digit expands one depth of letters, then the next digit continues the prefix.
// 1. 回溯：每个数字对应一层可选字母，选完再处理下一个数字。
// Time: O(n*4^n), Space: O(n) auxiliary plus O(n*4^n) output.
// 时间复杂度：O(n*4^n)，空间复杂度：辅助 O(n)，输出另占 O(n*4^n)。
//
// index 是正在处理的数字位置，每层从该数字的字母集合选一个；走完所有位置时把 path 转为独立字符串。
// index selects a digit; choose one mapped letter per level and copy path into a string when all positions are filled.
// 输入仅含 '2'..'9'，空输入返回空列表；退出分支时弹出本层字母。
// Assume digits '2'..'9'; empty input yields no combinations, and each branch pops its appended letter on return.
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	letters := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	result := []string{}
	path := make([]byte, 0, len(digits))

	var backtrack func(int)
	backtrack = func(index int) {
		if index == len(digits) {
			result = append(result, string(path))
			return
		}
		choices := letters[digits[index]-'0']
		for i := 0; i < len(choices); i++ {
			path = append(path, choices[i])
			backtrack(index + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return result
}
