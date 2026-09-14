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

关键逻辑 / Key Logic
BST 的中序是非降序（本题值互不相同则为严格递增）。相邻中序值的差是所有节点对里可能最小的差，
不必比较不相邻节点：递增序列中不相邻差一定不小于夹在中间的相邻差。
BST inorder is sorted. The minimum difference must occur between neighboring inorder values;
a non-adjacent pair in a sorted sequence cannot be smaller than the adjacent pairs it spans.

解法一：中序递归记录前驱（推荐） / Method 1: Inorder Recursion with a Predecessor (Recommended)
中序访问时，用当前值减去前驱值更新答案，再把前驱改成当前节点。
During inorder, subtract the predecessor from the current value, update the answer, then move the predecessor forward.

解法二：中序迭代 / Method 2: Iterative Inorder
用栈模拟同样的中序前驱比较。
The same predecessor comparison, using an explicit inorder stack.

时间与空间复杂度 / Time and Space Complexity
n 为节点数，h 为树高。两版时间 O(n)，辅助空间 O(h)。
For n nodes and height h, both take O(n) time and O(h) auxiliary space.
*/

// 1. Inorder recursion with a predecessor (recommended): the minimum gap is between neighboring inorder values.
// 1. 中序递归记录前驱：推荐；最小差只可能出现在中序相邻的两个值之间。
// Time: O(n), Space: O(h).
// 时间复杂度：O(n)，空间复杂度：O(h)。
func getMinimumDifference(root *TreeNode) int {
	// Seed with a large sentinel; real nonnegative gaps will replace it.
	// 先用大哨兵占位，真实的非负差值会把它换掉。
	minDiff := int(^uint(0) >> 1)
	var prev *TreeNode

	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)

		// First visited node has no predecessor; later nodes update the minimum gap.
		// 第一个访问的节点没有前驱；之后每个节点都用与前驱的差更新答案。
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

	// Empty or single-node tree: no pair exists.
	// 空树或单节点：不存在可比较的一对。
	if prev == nil || minDiff == int(^uint(0)>>1) {
		return 0
	}
	return minDiff
}

// 2. Iterative inorder: walk left-root-right with a stack and compare each node with its predecessor.
// 2. 中序迭代：用栈走左、根、右，并把每个节点与前驱比较。
// Time: O(n), Space: O(h).
// 时间复杂度：O(n)，空间复杂度：O(h)。
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
