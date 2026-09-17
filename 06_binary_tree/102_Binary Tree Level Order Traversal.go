package _6_binary_tree

/*
102. 二叉树的层序遍历 / Binary Tree Level Order Traversal

题目描述 / Problem Description
给定二叉树的根节点 root，返回节点值的层序遍历结果，即从上到下、从左到右逐层访问所有节点。
Given the root of a binary tree, return its level-order traversal:
visit all nodes level by level, from left to right.
复杂度记号：n 为节点数，h 为树高，w 为最大层宽；辅助空间不含返回结果。
Notation: n nodes, height h, maximum width w; auxiliary space excludes returned results.
*/

/*
题目描述 / Problem Description

给定二叉树的根节点 root，返回节点值的层序遍历结果，
即从上到下、从左到右逐层访问节点。

Given the root of a binary tree, return its level-order traversal,
visiting nodes level by level from left to right.
*/

// 1. Level-by-level queue (recommended): save levelSize before enqueuing children so adjacent levels stay separated.
// 1. 队列层序：推荐；入队孩子之前先固定 levelSize，用来分隔相邻两层。
// Time: O(n), Space: O(w) for the queue.
// 时间复杂度：O(n)，空间复杂度：O(w)，由队列产生。
//
// 每轮先固定 levelSize=len(queue)，只弹出这些节点；新入队的孩子属于下一层，不能计入本轮。
// Freeze levelSize before each round; newly enqueued children belong to the next level, not the current one.
// 按从左到右的孩子顺序入队，确保同层输出顺序；空树返回空层列表。
// Enqueue children left to right to preserve order within each level; empty input produces no levels.
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)

		level := make([]int, 0, levelSize)

		for range levelSize {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, level)
	}

	return result
}

// 2. Recursion grouped by depth: carry depth as the result index and append a new subarray on first arrival.
// 2. 递归按深度分组：深度就是结果下标，第一次到达该层时先追加一个空子数组。
// Time: O(n), Space: O(h) for the recursion stack.
// 时间复杂度：O(n)，空间复杂度：O(h)，由递归栈产生。
//
// depth 是从根起的零基层号；首次到达新深度才新增结果层，再把节点值加入 result[depth]。
// depth is the zero-based level; create a result row on first reaching that depth, then append to result[depth].
// DFS 不按层访问，但按从左到右递归保证每层加入顺序正确。
// DFS does not visit by level, but left-to-right recursion preserves order within each result row.
func levelOrderRecursive(root *TreeNode) [][]int {
	result := [][]int{}

	var traverse func(*TreeNode, int)
	traverse = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		if depth == len(result) {
			result = append(result, []int{})
		}

		result[depth] = append(result[depth], node.Val)

		traverse(node.Left, depth+1)
		traverse(node.Right, depth+1)
	}

	traverse(root, 0)

	return result
}
