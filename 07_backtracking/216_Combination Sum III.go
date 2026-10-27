package _7_backtracking

/*
216. 组合总和 III / Combination Sum III

题目描述 / Problem Description
从 1..9 中选出 k 个不同数字，使它们的和为 n，返回所有组合。数字不可重复使用。
Find every combination of k distinct integers from 1..9 that sum to n. Each number may be used at most once.

示例 / Example
k=3, n=7 → [[1,2,4]]
k=3, n=9 → [[1,2,6],[1,3,5],[2,3,4]]
*/

// 面试首选 / Interview first choice：216 组合总和 III。
//
// 1. Backtracking with pruning: choose increasing digits from 1..9 until k numbers sum to n.
// 1. 回溯加剪枝：在 1..9 中递增地选出 k 个数，使它们的和为 n。
// Time: O(k*C(9,k)), Space: O(k) auxiliary plus output.
// 时间复杂度：O(k*C(9,k))，空间复杂度：辅助 O(k)，另加输出。
//
// start 保证从 1..9 递增选数，避免重复使用或重复排列；remaining 是还缺的和，need 是还缺的个数。
// start enforces increasing choices from 1..9; remaining is the unmet sum and need the unfilled count.
// 选 i 后必须还留 need-1 个更大数，所以 i<=9-need+1；正数使 i>remaining 后可直接停止。
// Leave need-1 larger choices after i, giving i<=9-need+1; positivity allows stopping when i>remaining.
//
// 边界：只选 1..9，每个数字一次；只有选满 k 个且 remaining==0 才是答案，按题意 2<=k<=9。
func combinationSum3(k int, n int) [][]int {
	result := [][]int{}
	path := []int{}

	var backtrack func(int, int)
	backtrack = func(start, remaining int) {
		if remaining == 0 && len(path) == k {
			saved := make([]int, k)
			copy(saved, path)
			result = append(result, saved)
			return
		}
		if remaining <= 0 || len(path) == k {
			return
		}
		need := k - len(path)
		for i := start; i <= 9-need+1; i++ {
			if i > remaining {
				break
			}
			path = append(path, i)
			backtrack(i+1, remaining-i)
			path = path[:len(path)-1]
		}
	}

	backtrack(1, n)
	return result
}
