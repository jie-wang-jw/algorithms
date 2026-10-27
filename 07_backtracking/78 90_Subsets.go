package _7_backtracking

import "sort"

/*
78. 子集 / Subsets
90. 子集 II / Subsets II

题目描述 / Problem Description
78：返回数组所有子集。元素互不相同。
90：数组可含重复元素，解集不能含重复子集。
78: return every subset of a distinct-valued array.
90: the array may contain duplicates; the answer must not contain duplicate subsets.

示例 / Examples
78: [1,2,3] → [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
90: [1,2,2] → [[],[1],[1,2],[1,2,2],[2],[2,2]]
*/

// 面试首选 / Interview first choice：78 子集。
//
// 1. Subsets: collect the current path at every node, then try each later element as the next choice.
// 1. 子集：每个节点都收集当前路径，再从后面的元素里选下一个。
// Time: O(n*2^n), Space: O(n) auxiliary plus O(n*2^n) output.
// 时间复杂度：O(n*2^n)，空间复杂度：辅助 O(n)，输出另占 O(n*2^n)。
//
// 每个递归节点的 path 都是一个子集，包括初始空集，进入递归就保存副本；起点 i+1 防止重复选下标。
// Every path, including the initial empty one, is a subset; copy it on entry and advance to i+1 to avoid index reuse.
//
// 边界：题目保证元素互异；空数组也有一个子集，即空集。
func subsets(nums []int) [][]int {
	result := [][]int{}
	path := []int{}

	var backtrack func(int)
	backtrack = func(start int) {
		saved := make([]int, len(path))
		copy(saved, path)
		result = append(result, saved)

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return result
}

// 面试首选 / Interview first choice：90 子集 II。
//
// 2. Subsets II: sort, then skip equal values among siblings so duplicate subsets are not rebuilt.
// 2. 子集 II：先排序，再跳过同一层中相同的值，避免重复子集。
// Time: O(n*2^n), Space: O(n) auxiliary plus output.
// 时间复杂度：O(n*2^n)，空间复杂度：辅助 O(n)，另加输出。
//
// 先排序副本，让相同值相邻；仅跳过同层 i>start 的重复值，跨层仍可选同值以形成 [2,2] 等子集。
// Sort a copy and skip duplicates only among siblings with i>start; deeper equal choices preserve subsets such as [2,2].
// 每个节点都保存 path 副本，不只是递归叶子。
// Save a copy of every path, not only leaf paths.
func subsetsWithDup(nums []int) [][]int {
	sorted := append([]int(nil), nums...)
	sort.Ints(sorted)

	result := [][]int{}
	path := []int{}

	var backtrack func(int)
	backtrack = func(start int) {
		saved := make([]int, len(path))
		copy(saved, path)
		result = append(result, saved)

		for i := start; i < len(sorted); i++ {
			if i > start && sorted[i] == sorted[i-1] {
				continue
			}
			path = append(path, sorted[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return result
}
