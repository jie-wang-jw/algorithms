package _7_backtracking

/*
77. 组合 / Combinations

题目描述 / Problem Description
给定 n 和 k，返回 [1,n] 中所有长度为 k 的组合。组合内顺序无关，答案顺序不限。
Given n and k, return every combination of k numbers chosen from 1..n.
Order inside a combination does not matter; the list of combinations may appear in any order.

示例 / Example
n=4, k=2 → [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
*/

// 1. Backtracking: choose the next number from start, recurse with i+1, and copy a completed k-length path.
// 1. 回溯：从 start 起选下一个数，递归时用 i+1，凑满 k 个就复制路径。
// Time: O(Σ(j=0..k) C(n,j) + k*C(n,k)); auxiliary space O(k), output O(k*C(n,k)).
// 时间 O(Σ(j=0..k) C(n,j)+k*C(n,k))，辅助空间 O(k)，输出 O(k*C(n,k))；未剪枝会访问凑不满 k 个数的分支。
//
// 从 start 开始递增选择，递归传 i+1，保证每个数最多选一次且同一集合只出现一种顺序；满 k 项保存副本。
// Choose increasing values and recurse from i+1, using each value once and each set in one order; copy paths of size k.
func combine(n int, k int) [][]int {
	result := [][]int{}
	path := []int{}

	var backtrack func(int)
	backtrack = func(start int) {
		if len(path) == k {
			saved := make([]int, k)
			copy(saved, path)
			result = append(result, saved)
			return
		}
		for i := start; i <= n; i++ {
			path = append(path, i)
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(1)
	return result
}

// 面试首选 / Interview first choice：77 组合。
//
// 2. Pruning (recommended): stop when the remaining numbers cannot fill the leftover slots.
// 2. 剪枝：推荐；剩下的数字已经凑不满还缺的个数时，不必再枚举。
// Time: O(k*C(n,k)), Space: O(k) auxiliary plus O(k*C(n,k)) output.
// 时间复杂度：O(k*C(n,k))，空间复杂度：辅助 O(k)，输出另占 O(k*C(n,k))。
//
// need=k-len(path)；若选 i，还需 i 之后有 need-1 个数，所以 n-i>=need-1，即 i<=n-need+1。
// With need=k-len(path), choosing i requires n-i>=need-1, hence i<=n-need+1.
// 剪去的分支连个数都凑不够，不会丢答案；递归传 i+1 保证组合不重复。
// Pruned branches cannot fill the remaining slots; advancing to i+1 prevents duplicate combinations.
//
// 状态与边界：path 是已选的递增序列，start 是下一候选下界；按题意 1<=k<=n，凑满 k 个才复制答案。
func combinePruned(n int, k int) [][]int {
	result := [][]int{}
	path := []int{}

	var backtrack func(int)
	backtrack = func(start int) {
		if len(path) == k {
			saved := make([]int, k)
			copy(saved, path)
			result = append(result, saved)
			return
		}
		need := k - len(path)
		for i := start; i <= n-need+1; i++ {
			path = append(path, i)
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(1)
	return result
}
