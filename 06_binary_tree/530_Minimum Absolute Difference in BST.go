package _6_binary_tree

/*
530. 二叉搜索树的最小绝对差 / Minimum Absolute Difference in BST

题目描述 / Problem Description
给定二叉搜索树根节点，返回树中任意两不同节点值之差的绝对值的最小值。
节点值非负。空树与单节点在题目中不会作为有效输入来求差，测试对它们返回 0。
Given a BST root, return the minimum absolute difference between any two distinct node values.
Node values are nonnegative. Empty and single-node trees are not meaningful inputs here; tests treat them as 0.

示例 / Examples
      4
     / \
    2   6
   / \
  1   3
中序序列 1,2,3,4,6；相邻差为 1,1,1,2，最小绝对差是 1。
Inorder values 1,2,3,4,6; adjacent gaps are 1,1,1,2; the minimum is 1.
复杂度记号：n 为节点数，h 为树高，w 为最大层宽；辅助空间不含返回结果。
Notation: n nodes, height h, maximum width w; auxiliary space excludes returned results.
*/

// 1. Inorder recursion with a predecessor (recommended): the minimum gap is between neighboring inorder values.
// 1. 中序递归记录前驱：推荐；最小差只可能出现在中序相邻的两个值之间。
// Time: O(n), Space: O(h).
// 时间复杂度：O(n)，空间复杂度：O(h)。
//
// BST 中序有序，任意不相邻差值是若干相邻差值之和，因此最小差一定出现在中序相邻节点间。
// In sorted inorder, a nonadjacent difference sums adjacent differences, so the minimum occurs between neighbors.
// prev 保存前一个访问节点，先计算差再更新 prev；按题目值域使用，差值不溢出且小于哨兵最大 int。
// prev is the prior visited node; compare before replacing it. Assume problem bounds keep differences below the max-int sentinel.
func getMinimumDifference(root *TreeNode) int {
	minDiff := int(^uint(0) >> 1)
	var prev *TreeNode

	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)

		if prev != nil {
			diff := node.Val - prev.Val
			if diff < minDiff {
				minDiff = diff
			}
		}
		prev = node
		traverse(node.Right)
	}

	traverse(root)

	if prev == nil || minDiff == int(^uint(0)>>1) {
		return 0
	}
	return minDiff
}

// 2. Iterative inorder: walk left-root-right with a stack and compare each node with its predecessor.
// 2. 中序迭代：用栈走左、根、右，并把每个节点与前驱比较。
// Time: O(n), Space: O(h).
// 时间复杂度：O(n)，空间复杂度：O(h)。
//
// BST 中序有序，任意不相邻差值是若干相邻差值之和，因此最小差一定出现在中序相邻节点间。
// In sorted inorder, a nonadjacent difference sums adjacent differences, so the minimum occurs between neighbors.
// prev 保存前一个访问节点，先计算差再更新 prev；按题目值域使用，差值不溢出且小于哨兵最大 int。
// prev is the prior visited node; compare before replacing it. Assume problem bounds keep differences below the max-int sentinel.
func getMinimumDifferenceIterative(root *TreeNode) int {
	minDiff := int(^uint(0) >> 1)
	stack := []*TreeNode{}
	var prev *TreeNode
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if prev != nil {
			diff := cur.Val - prev.Val
			if diff < minDiff {
				minDiff = diff
			}
		}
		prev = cur
		cur = cur.Right
	}
	if prev == nil || minDiff == int(^uint(0)>>1) {
		return 0
	}
	return minDiff
}
