package _6_binary_tree

/*
111. 二叉树的最小深度 / Minimum Depth of Binary Tree

题目描述 / Problem Description
给定二叉树的根节点 root，返回从根节点到最近叶子节点的最短路径上的节点数量。
叶子节点是左右孩子都为空的节点；空树的最小深度为 0。
Given the root of a binary tree, return the number of nodes on the shortest path from the root to a leaf.
A leaf has neither a left nor a right child. An empty tree has minimum depth 0.

示例 / Examples
       3                 1
      / \                 \
     9  20                 2
       /  \                 \
      15   7                 3
左图最小深度为 2，路径为 3 -> 9；右图最小深度为 3，路径为 1 -> 2 -> 3。
The left tree has minimum depth 2 via 3 -> 9; the right tree has minimum depth 3 via 1 -> 2 -> 3.
复杂度记号：n 为节点数，h 为树高，w 为最大层宽；辅助空间不含返回结果。
Notation: n nodes, height h, maximum width w; auxiliary space excludes returned results.
*/

// 1. Recursion (recommended): skip a missing child so min cannot treat a nil pointer as a finished leaf path.
// 1. 递归：推荐；先排除缺失的一边，避免 min 把空指针当成已经到达叶子。
// Time: O(n), Space: O(h) for the recursion stack.
// 时间复杂度：O(n)，空间复杂度：O(h)，由递归栈产生。
//
// 最短路径必须到叶子；只有一边非空时必须沿非空边走，不能拿空边的 0 与另一边取 min。
// A valid path ends at a leaf; with one missing child, follow the existing child instead of minimizing against zero.
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return minDepth(root.Right) + 1
	}

	if root.Right == nil {
		return minDepth(root.Left) + 1
	}

	leftDepth := minDepth(root.Left)
	rightDepth := minDepth(root.Right)

	return min(leftDepth, rightDepth) + 1
}

// 2. Breadth-first search: return the current depth at the first node whose both children are nil.
// 2. 层序遍历：第一次遇到左右孩子都为空的节点，当前层数就是最小深度。
// Time: O(n), Space: O(w) for the queue.
// 时间复杂度：O(n)，空间复杂度：O(w)，由队列产生。
//
// BFS 按深度递增处理，第一次遇到左右孩子都空的叶子，其深度就是最短深度。
// BFS processes increasing depths; the first node with no children is a leaf at minimum depth.
func minDepthIterative(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	depth := 1

	for len(queue) > 0 {
		levelSize := len(queue)

		for range levelSize {
			node := queue[0]
			queue = queue[1:]

			if node.Left == nil && node.Right == nil {
				return depth
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		depth++
	}

	return 0
}
