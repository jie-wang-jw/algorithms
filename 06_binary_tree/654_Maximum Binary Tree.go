package _6_binary_tree

/*
654. 最大二叉树 / Maximum Binary Tree

题目描述 / Problem Description
给定不含重复元素的整数数组 nums，按如下规则构造最大二叉树并返回根节点：
根是当前区间的最大元素；最大值左侧构造左子树，右侧构造右子树。
Given an integer array nums of unique values, construct the maximum binary tree:
the root is the largest value in the current range, the left subarray builds the left subtree,
and the right subarray builds the right subtree.

示例 / Example
nums = [3,2,1,6,0,5]
       6
      / \
     3   5
      \  /
      2 0
       \
        1
复杂度记号：n 为节点数，h 为树高，w 为最大层宽；辅助空间不含返回结果。
Notation: n nodes, height h, maximum width w; auxiliary space excludes returned results.
*/

// 1. Recursive slicing: scan the current slice for the maximum, then recurse on the left and right subarrays.
// 1. 递归切片：在当前切片中找最大值，再对左右两侧子数组递归构造。
// Time: O(n²) worst case, Space: O(h) plus O(n) for the output tree.
// 时间复杂度：最坏 O(n²)，空间复杂度：O(h)，输出树另占 O(n)。
//
// 区间最大值作根，最大值左侧和右侧分别递归建树，保持中序等于原数组；子范围必须排除根位置。
// Use the interval maximum as root and recurse on either side, preserving the original inorder; exclude the root index from child ranges.
// 按题意值互异；单调数组每次只缩小一个元素，重复扫描导致最坏 O(n²)。
// Assume distinct values; monotone input shrinks each interval by one and repeated scans cause worst-case O(n²).
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	maxIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}

	root := &TreeNode{Val: nums[maxIndex]}
	root.Left = constructMaximumBinaryTree(nums[:maxIndex])
	root.Right = constructMaximumBinaryTree(nums[maxIndex+1:])
	return root
}

// 2. Index ranges (recommended): split a half-open interval [left, right) around the maximum index.
// 2. 下标区间：推荐；在左闭右开区间 [left, right) 中找最大值下标并切分。
// Time: O(n²) worst case, Space: O(h) plus O(n) for the output tree.
// 时间复杂度：最坏 O(n²)，空间复杂度：O(h)，输出树另占 O(n)。
//
// 区间最大值作根，最大值左侧和右侧分别递归建树，保持中序等于原数组；子范围必须排除根位置。
// Use the interval maximum as root and recurse on either side, preserving the original inorder; exclude the root index from child ranges.
// 按题意值互异；单调数组每次只缩小一个元素，重复扫描导致最坏 O(n²)。
// Assume distinct values; monotone input shrinks each interval by one and repeated scans cause worst-case O(n²).
func constructMaximumBinaryTreeIndex(nums []int) *TreeNode {
	return constructMaximumBinaryTreeRange(nums, 0, len(nums))
}

// Half-open [left, right) describes the subarray that should become one maximum binary tree.
// 左闭右开区间 [left, right) 表示要构造成一棵最大二叉树的子数组。
// Time: O(n²) worst case, Space: O(h).
// 时间复杂度：最坏 O(n²)，空间复杂度：O(h)。
//
// 区间最大值作根，最大值左侧和右侧分别递归建树，保持中序等于原数组；子范围必须排除根位置。
// Use the interval maximum as root and recurse on either side, preserving the original inorder; exclude the root index from child ranges.
// 按题意值互异；单调数组每次只缩小一个元素，重复扫描导致最坏 O(n²)。
// Assume distinct values; monotone input shrinks each interval by one and repeated scans cause worst-case O(n²).
func constructMaximumBinaryTreeRange(nums []int, left, right int) *TreeNode {
	if left >= right {
		return nil
	}

	maxIndex := left
	for i := left + 1; i < right; i++ {
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}

	root := &TreeNode{Val: nums[maxIndex]}
	root.Left = constructMaximumBinaryTreeRange(nums, left, maxIndex)
	root.Right = constructMaximumBinaryTreeRange(nums, maxIndex+1, right)
	return root
}
