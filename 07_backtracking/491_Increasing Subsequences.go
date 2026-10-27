package _7_backtracking

/*
491. 递增子序列 / Increasing Subsequences

题目描述 / Problem Description
找出数组中所有长度至少为 2 的递增子序列。相等视为递增。不能对原数组排序。答案不能含重复序列。
Find every increasing subsequence of length at least 2. Equal values count as increasing.
Do not sort the array. The answer must not contain duplicate subsequences.

示例 / Example
[4,6,7,7] → [[4,6],[4,7],[4,6,7],[4,6,7,7],[6,7],[6,7,7],[7,7],[4,7,7]]
*/

// 面试首选 / Interview first choice：491 非递减子序列。
//
// 1. Per-level hash set: skip a value that is smaller than the path tail or already used among siblings.
// 1. 本层哈希集合去重：小于路径末尾、或本层已经用过的值都跳过。
// Time: O(n*2^n), Space: O(n) auxiliary plus output.
// 时间复杂度：O(n*2^n)，空间复杂度：辅助 O(n)，另加输出。
//
// 保留原数组顺序，不能排序；只选不小于 path 末项的后续元素，路径长度>=2 时即保存，不必等到叶子。
// Preserve input order and choose later values no smaller than the path tail; save every path of length>=2, not only leaves.
// used 每次递归重新创建，只对同一层相同值去重；若全局共用会误删 [7,7] 等合法结果。
// Create used per recursive call to deduplicate siblings only; a global set would lose valid results such as [7,7].
// 数组版用值+100 作下标，额外依赖题目值域 [-100,100]。
// The array variant indexes by value+100 and therefore requires values in [-100,100].
func findSubsequences(nums []int) [][]int {
	result := [][]int{}
	path := []int{}

	var backtrack func(int)
	backtrack = func(start int) {
		if len(path) > 1 {
			saved := make([]int, len(path))
			copy(saved, path)
			result = append(result, saved)
		}
		used := map[int]bool{}
		for i := start; i < len(nums); i++ {
			if (len(path) > 0 && nums[i] < path[len(path)-1]) || used[nums[i]] {
				continue
			}
			used[nums[i]] = true
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return result
}

// 2. Per-level array hash (recommended): values in [-100,100] become indices in used[201].
// 2. 本层数组哈希：推荐；[-100,100] 映射到 used[201] 的下标。
// Time: O(n*2^n), Space: O(n) auxiliary plus output.
// 时间复杂度：O(n*2^n)，空间复杂度：辅助 O(n)，另加输出。
//
// 保留原数组顺序，不能排序；只选不小于 path 末项的后续元素，路径长度>=2 时即保存，不必等到叶子。
// Preserve input order and choose later values no smaller than the path tail; save every path of length>=2, not only leaves.
// used 每次递归重新创建，只对同一层相同值去重；若全局共用会误删 [7,7] 等合法结果。
// Create used per recursive call to deduplicate siblings only; a global set would lose valid results such as [7,7].
// 数组版用值+100 作下标，额外依赖题目值域 [-100,100]。
// The array variant indexes by value+100 and therefore requires values in [-100,100].
func findSubsequencesArray(nums []int) [][]int {
	result := [][]int{}
	path := []int{}

	var backtrack func(int)
	backtrack = func(start int) {
		if len(path) > 1 {
			saved := make([]int, len(path))
			copy(saved, path)
			result = append(result, saved)
		}
		used := [201]int{}
		for i := start; i < len(nums); i++ {
			if (len(path) > 0 && nums[i] < path[len(path)-1]) || used[nums[i]+100] == 1 {
				continue
			}
			used[nums[i]+100] = 1
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return result
}
