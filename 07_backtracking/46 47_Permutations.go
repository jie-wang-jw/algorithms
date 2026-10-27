package _7_backtracking

import "sort"

/*
46. 全排列 / Permutations
47. 全排列 II / Permutations II

题目描述 / Problem Description
46：互不相同的数组，返回全部排列。
47：可含重复数字，返回不重复的全部排列。
46: return every permutation of a distinct-valued array.
47: the array may contain duplicates; return unique permutations only.

示例 / Examples
46: [1,2,3] → [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
47: [1,1,2] → [[1,1,2],[1,2,1],[2,1,1]]
*/

// 面试首选 / Interview first choice：46 全排列。
//
// 1. Permutations: try every unused index at each depth until the path holds n numbers.
// 1. 全排列：每一层尝试所有尚未使用的下标，直到路径长度为 n。
// Time: O(n*n!), Space: O(n) auxiliary plus O(n*n!) output.
// 时间复杂度：O(n*n!)，空间复杂度：辅助 O(n)，输出另占 O(n*n!)。
//
// 排列每层都从全部下标选数，used[i] 表示该下标是否已在当前路径中；与组合不同，不能限制递增起点。
// Permutations choose from all indices each level; used[i] marks membership in the current path, unlike increasing-start combinations.
// 选择后置 true，回退后恢复 false；题目保证值互异，满长时复制 path 保存。
// Mark on choice and clear on return; assume distinct values and copy complete paths into the result.
func permute(nums []int) [][]int {
	result := [][]int{}
	path := []int{}
	used := make([]bool, len(nums))

	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			saved := make([]int, len(path))
			copy(saved, path)
			result = append(result, saved)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backtrack()
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack()
	return result
}

// 面试首选 / Interview first choice：47 全排列 II。
//
// 2. Permutations II: sort, then skip an equal unused predecessor so sibling branches stay unique.
// 2. 全排列 II：先排序，再跳过“前一个相同值未使用”的情况，避免同一层重复。
// Time: O(n*n!), Space: O(n) auxiliary plus output.
// 时间复杂度：O(n*n!)，空间复杂度：辅助 O(n)，另加输出。
//
// 排序副本后，同值只允许按下标顺序进入路径：前一个相同值尚未使用时，当前值必须跳过。
// After sorting a copy, equal values must enter the path in index order; skip one if its equal predecessor is unused.
// !used[i-1] 表示同层会重复；若前者已在路径中，则此处可以继续选同值，保留含重复元素的排列。
// !used[i-1] prevents duplicate sibling choices; if the predecessor is already on the path, this equal value remains usable.
func permuteUnique(nums []int) [][]int {
	sorted := append([]int(nil), nums...)
	sort.Ints(sorted)

	result := [][]int{}
	path := []int{}
	used := make([]bool, len(sorted))

	var backtrack func()
	backtrack = func() {
		if len(path) == len(sorted) {
			saved := make([]int, len(path))
			copy(saved, path)
			result = append(result, saved)
			return
		}
		for i := 0; i < len(sorted); i++ {
			if used[i] {
				continue
			}
			if i > 0 && sorted[i] == sorted[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			path = append(path, sorted[i])
			backtrack()
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack()
	return result
}
