package _6_binary_tree

/*
108. 将有序数组转换为二叉搜索树 / Convert Sorted Array to Binary Search Tree

题目描述 / Problem Description
将升序数组转换成高度平衡的二叉搜索树：每个节点左右子树高度差不超过 1。
Convert a sorted array into a height-balanced BST, where every node's child-subtree heights differ by at most one.

示例 / Examples
nums = [-10,-3,0,5,9] 的一种平衡结果 / one balanced result:
      0
     / \
   -3   9
   /   /
 -10  5
中点 0 作根；左半 [-10,-3] 与右半 [5,9] 再各自取中点，高度差不超过 1。
Midpoint 0 is the root; left half [-10,-3] and right half [5,9] each pick their midpoints so heights differ by at most one.
复杂度记号：n 为节点数，h 为树高，w 为最大层宽；辅助空间不含返回结果。
Notation: n nodes, height h, maximum width w; auxiliary space excludes returned results.
*/

// 1. Index recursion (recommended): the midpoint of [left, right) is the root of a balanced subtree.
// 1. 下标递归：推荐；左闭右开区间 [left, right) 的中点就是当前平衡子树的根。
// Time: O(n), Space: O(log n) plus O(n) for the output tree.
// 时间复杂度：O(n)，空间复杂度：O(log n)，输出树另占 O(n)。
//
// 递增数组选中点为根，左侧都较小、右侧都较大，满足 BST；递归均分使两侧高度差不超过 1。
// Choose the midpoint of a sorted array: smaller values go left, larger right; recursively balanced splits keep height difference at most one.
// 子范围排除已选中的 mid；空范围返回 nil，偶数长度选两个中点中的右者也合法。
// Exclude mid from child ranges; empty ranges return nil, and the upper midpoint is valid for even lengths.
func sortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBSTRange(nums, 0, len(nums))
}

// Half-open [left, right) becomes one balanced BST.
// 左闭右开区间 [left, right) 构造成一棵平衡二叉搜索树。
// Time: O(n), Space: O(log n).
// 时间复杂度：O(n)，空间复杂度：O(log n)。
//
// 递增数组选中点为根，左侧都较小、右侧都较大，满足 BST；递归均分使两侧高度差不超过 1。
// Choose the midpoint of a sorted array: smaller values go left, larger right; recursively balanced splits keep height difference at most one.
// 子范围排除已选中的 mid；空范围返回 nil，偶数长度选两个中点中的右者也合法。
// Exclude mid from child ranges; empty ranges return nil, and the upper midpoint is valid for even lengths.
func sortedArrayToBSTRange(nums []int, left, right int) *TreeNode {
	if left >= right {
		return nil
	}

	mid := left + (right-left)/2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBSTRange(nums, left, mid)
	root.Right = sortedArrayToBSTRange(nums, mid+1, right)
	return root
}

// 2. Recursive slicing: take the midpoint of the current slice, then build both halves from views.
// 2. 递归切片：取当前切片中点为根，再用左右两半的视图继续构造。
// Time: O(n), Space: O(log n) plus O(n) for the output tree.
// 时间复杂度：O(n)，空间复杂度：O(log n)，输出树另占 O(n)。
//
// 递增数组选中点为根，左侧都较小、右侧都较大，满足 BST；递归均分使两侧高度差不超过 1。
// Choose the midpoint of a sorted array: smaller values go left, larger right; recursively balanced splits keep height difference at most one.
// 子范围排除已选中的 mid；空范围返回 nil，偶数长度选两个中点中的右者也合法。
// Exclude mid from child ranges; empty ranges return nil, and the upper midpoint is valid for even lengths.
func sortedArrayToBSTSlice(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBSTSlice(nums[:mid])
	root.Right = sortedArrayToBSTSlice(nums[mid+1:])
	return root
}
