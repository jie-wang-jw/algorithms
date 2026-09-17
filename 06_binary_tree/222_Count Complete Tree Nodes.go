package _6_binary_tree

/*
222. 完全二叉树的节点个数 / Count Complete Tree Nodes

题目描述 / Problem Description
给定一棵完全二叉树的根节点 root，返回节点总数。
完全二叉树：除最后一层外，每层都填满；最后一层节点从左到右连续排列，中间不能有空位。
Given the root of a complete binary tree, return its number of nodes.
Every level except possibly the last is full, and the last level is filled continuously from left to right.

示例 / Example
       1
      / \
     2   3
    / \ /
   4  5 6
答案为 6；最后一层可以缺少右侧节点，但不能跳过左侧位置。
The answer is 6. Missing positions may occur on the right of the last level, not before existing nodes.
复杂度记号：n 为节点数，h 为树高，w 为最大层宽；辅助空间不含返回结果。
Notation: n nodes, height h, maximum width w; auxiliary space excludes returned results.
*/

// 1. Ordinary recursion: add both disjoint subtree counts, then include the current root.
// 1. 普通递归：左右子树节点数相加，再加上当前根。
// Time: O(n), Space: O(h); completeness makes h = O(log n).
// 时间复杂度：O(n)，空间复杂度：O(h)，本题完全二叉树使 h=O(log n)。
//
// 每层返回左树节点数+右树节点数+当前 1 个节点；适用于任意二叉树，空树为 0。
// Return left count plus right count plus one; works for any binary tree, with nil contributing zero.
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftCount := countNodes(root.Left)
	rightCount := countNodes(root.Right)

	return leftCount + rightCount + 1
}

// 2. Breadth-first traversal: increment the count on every dequeue; no levelSize is needed.
// 2. 层序遍历：每出队一个节点就计数，本题不需要分层。
// Time: O(n), Space: O(w) for the queue.
// 时间复杂度：O(n)，空间复杂度：O(w)，由队列产生。
//
// 每个非空节点入队一次、出队时计数一次；不需要分层，也不依赖完全二叉树性质。
// Enqueue each nonnil node once and count it on removal; neither level grouping nor completeness is required.
func countNodesIterative(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	count := 0

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		count++

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return count
}

// 3. Perfect-subtree shortcut (recommended): equal left/right boundary heights yield 2^h-1 without visiting the interior.
// 3. 满子树公式：推荐；最左、最右路径等高即可直接返回 2^h-1，不必再进入内部。
// Time: O(log² n), Space: O(log n) for the recursion stack.
// 时间复杂度：O(log² n)，空间复杂度：O(log n)，由递归栈产生。
//
// 只有在完全二叉树前提下，最左和最右链高度相同才能推出整棵子树是满的，节点数为 2^h-1。
// Only for a complete tree do equal outer-spine heights prove a perfect subtree with 2^h-1 nodes.
// 否则递归两边；每层至少一边是满子树而可直接计数，因此总时间 O(log² n)，普通二叉树不能套用。
// Otherwise recurse; at least one child subtree is perfect at each level, yielding O(log² n), not a rule for arbitrary trees.
func countNodesOptimized(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight, rightHeight := 0, 0

	for node := root; node != nil; node = node.Left {
		leftHeight++
	}

	for node := root; node != nil; node = node.Right {
		rightHeight++
	}

	if leftHeight == rightHeight {
		return (1 << leftHeight) - 1
	}

	return countNodesOptimized(root.Left) +
		countNodesOptimized(root.Right) + 1
}
