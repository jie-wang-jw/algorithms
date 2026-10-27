package _7_backtracking

import "sort"

/*
39. 组合总和 / Combination Sum
40. 组合总和 II / Combination Sum II

题目描述 / Problem Description
39：候选无重复，每个数字可无限次使用，求所有和为 target 的组合。
40：候选可有重复，每个下标只能用一次，解集不能包含重复组合。
39: candidates are unique and may be reused; return combinations that sum to target.
40: candidates may contain duplicates and each index is used once; the answer must not contain duplicate combinations.

示例 / Examples
39: candidates=[2,3,6,7], target=7 → [[2,2,3],[7]]
40: candidates=[10,1,2,7,6,1,5], target=8 → [[1,1,6],[1,2,5],[1,7],[2,6]]
     排序后 / after sort: [1,1,2,5,6,7,10]
*/

// 面试首选 / Interview first choice：39 组合总和。
//
// 1. Combination sum: reuse the current candidate by recursing from the same index.
// 1. 组合总和：递归仍从当前下标开始，同一个数字可以重复选取。
// n candidates, C recursive calls, S output entries: time O(n*C+S), auxiliary space O(t+1), output O(S).
// n 为候选数，C 为递归调用数，S 为输出元素总数：时间 O(n*C+S)，辅助空间 O(t+1)，输出 O(S)。
// For nonempty positive candidates, t=floor(target/minCandidate); at most n choices per level give C<=1+n+...+n^t.
// 候选非空且为正时，最大深度 t=target/最小候选（向下取整）；每层至多 n 个分支，C<=1+n+...+n^t。
//
// 候选数须互异且为正数；递归仍传 i 表示当前数可复用，起点不减则同一组合不会因选择顺序不同而重复。
// Candidates must be distinct and positive; recurse with i to allow reuse, while nondecreasing indices suppress order duplicates.
// 数组未排序，所以某个值>remaining 只能 continue，不能 break；正数确保 remaining 每次递减并最终终止。
// The input is unsorted, so skip an oversized value without breaking; positivity ensures the remaining target strictly decreases.
func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	path := []int{}

	var backtrack func(int, int)
	backtrack = func(start, remaining int) {
		if remaining == 0 {
			saved := make([]int, len(path))
			copy(saved, path)
			result = append(result, saved)
			return
		}
		for i := start; i < len(candidates); i++ {
			if candidates[i] > remaining {
				continue
			}
			path = append(path, candidates[i])
			backtrack(i, remaining-candidates[i])
			path = path[:len(path)-1]
		}
	}

	backtrack(0, target)
	return result
}

// 面试首选 / Interview first choice：40 组合总和 II。
//
// 2. Combination sum II: sort, use each index once, and skip equal values among siblings.
// 2. 组合总和 II：先排序，每个下标只用一次，并跳过同一层中相同的值。
// Time: O(n*2^n) including copying solutions; auxiliary space O(n), output separate.
// 时间：含结果复制的保守上界 O(n*2^n)；辅助空间 O(n)，输出另计。
//
// 排序副本后，同层 i>start 且值相同就跳过；跨层可选相同值，才能保留由不同重复元素组成的合法答案。
// Sort a copy and skip equal values only when i>start; deeper levels may use equal values from distinct elements.
// 递归传 i+1 表示每个下标最多用一次；正数且已排序，所以当前值>remaining 时后续也不可能成功。
// Pass i+1 to use each index once; sorted positive values permit breaking when a value exceeds remaining.
func combinationSum2(candidates []int, target int) [][]int {
	sorted := append([]int(nil), candidates...)
	sort.Ints(sorted)

	result := [][]int{}
	path := []int{}

	var backtrack func(int, int)
	backtrack = func(start, remaining int) {
		if remaining == 0 {
			saved := make([]int, len(path))
			copy(saved, path)
			result = append(result, saved)
			return
		}
		for i := start; i < len(sorted); i++ {
			if i > start && sorted[i] == sorted[i-1] {
				continue
			}
			if sorted[i] > remaining {
				break
			}
			path = append(path, sorted[i])
			backtrack(i+1, remaining-sorted[i])
			path = path[:len(path)-1]
		}
	}

	backtrack(0, target)
	return result
}
