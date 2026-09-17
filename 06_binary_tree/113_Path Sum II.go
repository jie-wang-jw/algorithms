package _6_binary_tree

/*
113. 路径总和 II / Path Sum II

题目描述 / Problem Description
给定二叉树根节点 root 和整数 targetSum，返回所有从根到叶子、节点值之和等于 targetSum 的路径。
每条路径保存节点值，整体返回 [][]int，顺序不限。空树返回空结果。
Given a binary tree root and an integer targetSum, return all root-to-leaf paths whose values sum to targetSum.
Return node values as [][]int in any order. An empty tree returns no paths.

示例 / Example
       5
      / \
     4   8
    /   / \
   11  13  4
  / \     / \
 7   2   5   1
targetSum=22，结果为 [[5,4,11,2],[5,8,4,5]]。
For targetSum=22, return [[5,4,11,2],[5,8,4,5]].

与 112、257 的联系 / Relation to Problems 112 and 257
112 只判断存在性，找到一条即可结束；113 要收集全部匹配路径，不能提前结束整个搜索。
257 收集所有根到叶子的路径；113 增加路径和筛选，保存整数切片而非字符串。
112 checks existence and may stop at one match; 113 must collect all matches.
257 collects every root-to-leaf path; 113 filters by sum and stores integer slices rather than strings.
复杂度记号：n 为节点数，h 为树高，w 为最大层宽；辅助空间不含返回结果。
Notation: n nodes, height h, maximum width w; auxiliary space excludes returned results.
*/

// 1. Recursive backtracking (recommended): append on entry, copy a matching leaf path, then undo before returning.
// 1. 递归回溯：推荐；进入时加入路径，叶子且剩余为 0 时复制答案，返回前撤销。
// Time: O(n+S), Space: O(h) auxiliary plus O(S) output, where S is the total number of integers in all returned paths.
// 时间复杂度：O(n+S)，空间复杂度：辅助 O(h)，输出另占 O(S)，S 为所有返回路径中整数的总数。
//
// path 是当前根到节点路径，remaining 是尚缺的和；入节点时追加，退出时弹出，避免兄弟分支互相污染。
// path holds the current root-to-node route and remaining the unmet sum; append on entry and pop on exit to isolate branches.
// 仅在叶子且 remaining==0 时保存 path 的副本；不复制会让后续回溯覆盖已保存答案。
// At matching leaves save a copy of path; otherwise later backtracking can overwrite stored answers.
func pathSum(root *TreeNode, targetSum int) [][]int {
	result := [][]int{}
	path := []int{}

	var traverse func(*TreeNode, int)
	traverse = func(node *TreeNode, remaining int) {
		if node == nil {
			return
		}

		path = append(path, node.Val)
		remaining -= node.Val

		if node.Left == nil && node.Right == nil {
			if remaining == 0 {
				saved := make([]int, len(path))
				copy(saved, path)
				result = append(result, saved)
			}
		} else {
			traverse(node.Left, remaining)
			traverse(node.Right, remaining)
		}

		path = path[:len(path)-1]
	}

	traverse(root, targetSum)
	return result
}

// 2. Iterative DFS with independent paths: copy the parent prefix for each child so no explicit undo is needed.
// 2. 栈迭代保存独立路径：每个孩子复制父路径前缀，因此不必手动回溯。
// Time: O(nh) upper bound, Space: O(h²) auxiliary plus O(S) output.
// 时间复杂度：上界 O(nh)，空间复杂度：辅助上界 O(h²)，输出另占 O(S)。
//
// nodes/sums/paths 的同一位置描述同一状态；为每个孩子复制父路径再追加，避免不同分支共享可写底层数组。
// Aligned nodes/sums/paths describe one state; copy each parent path before appending a child to prevent writable array sharing.
// 只收集和达标的叶子路径；每条路径的复制成本计入时间和空间，不能仅按节点数声称 O(n)。
// Collect only matching leaf paths; path-copy costs count toward time and space, not just the number of nodes.
func pathSumIterative(root *TreeNode, targetSum int) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	nodes := []*TreeNode{root}
	sums := []int{root.Val}
	paths := [][]int{{root.Val}}

	for len(nodes) > 0 {
		last := len(nodes) - 1
		node, sum, path := nodes[last], sums[last], paths[last]
		nodes = nodes[:last]
		sums = sums[:last]
		paths = paths[:last]

		if node.Left == nil && node.Right == nil {
			if sum == targetSum {
				result = append(result, path)
			}
			continue
		}

		for _, child := range []*TreeNode{node.Right, node.Left} {
			if child == nil {
				continue
			}

			nextPath := make([]int, len(path)+1)
			copy(nextPath, path)
			nextPath[len(path)] = child.Val

			nodes = append(nodes, child)
			sums = append(sums, sum+child.Val)
			paths = append(paths, nextPath)
		}
	}

	return result
}

// 3. BFS with independent paths: three synchronized queues hold node, sum, and a private path copy.
// 3. 队列 BFS 保存独立路径：三个同步队列分别保存节点、累计和以及各自的完整路径。
// Time: O(nh) upper bound, Space: O(wh) auxiliary plus O(S) output.
// 时间复杂度：上界 O(nh)，空间复杂度：辅助上界 O(wh)，输出另占 O(S)。
//
// nodes/sums/paths 的同一位置描述同一状态；为每个孩子复制父路径再追加，避免不同分支共享可写底层数组。
// Aligned nodes/sums/paths describe one state; copy each parent path before appending a child to prevent writable array sharing.
// 只收集和达标的叶子路径；每条路径的复制成本计入时间和空间，不能仅按节点数声称 O(n)。
// Collect only matching leaf paths; path-copy costs count toward time and space, not just the number of nodes.
func pathSumBFS(root *TreeNode, targetSum int) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	nodes := []*TreeNode{root}
	sums := []int{root.Val}
	paths := [][]int{{root.Val}}

	for len(nodes) > 0 {
		node, sum, path := nodes[0], sums[0], paths[0]
		nodes = nodes[1:]
		sums = sums[1:]
		paths = paths[1:]

		if node.Left == nil && node.Right == nil {
			if sum == targetSum {
				result = append(result, path)
			}
			continue
		}

		for _, child := range []*TreeNode{node.Left, node.Right} {
			if child == nil {
				continue
			}

			nextPath := make([]int, len(path)+1)
			copy(nextPath, path)
			nextPath[len(path)] = child.Val

			nodes = append(nodes, child)
			sums = append(sums, sum+child.Val)
			paths = append(paths, nextPath)
		}
	}

	return result
}
