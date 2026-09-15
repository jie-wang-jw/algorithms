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

关键逻辑 / Key Logic
有序数组的中点作为根，左边构造左子树，右边构造右子树，自然满足 BST 且尽量平衡。
区间坚持左闭右开 [left,right)。空区间 left>=right 返回 nil。中点取 left+(right-left)/2，避免溢出，也让左右半段长度差至多为 1。
The midpoint becomes the root; the left subarray builds the left subtree and the right subarray the right, which is both a BST and balanced.
Use a half-open range [left,right). left>=right is empty. The midpoint left+(right-left)/2 keeps the split balanced.

解法一：下标递归（推荐） / Method 1: Index Recursion (Recommended)
文章强调不要每层复制新数组，直接用下标在原数组上切分。
The article warns against copying a new array each time; split the original array with indices.

解法二：递归切片 / Method 2: Recursive Slicing
用切片视图切出左右半段，逻辑与解法一相同。Go 切片不复制元素。
Use slice views for the two halves. The idea matches method 1; Go slices do not copy elements.

时间与空间复杂度 / Time and Space Complexity
n 为数组长度。两版时间 O(n)，每个元素恰好成为一次根。
辅助空间 O(log n) 来自平衡树的递归栈；输出树 O(n)。
For n values, both take O(n) time because each element becomes a root once.
Auxiliary space is O(log n) from the balanced recursion; the output tree uses O(n).
*/

// 1. Index recursion (recommended): the midpoint of [left, right) is the root of a balanced subtree.
// 1. 下标递归：推荐；左闭右开区间 [left, right) 的中点就是当前平衡子树的根。
// Time: O(n), Space: O(log n) plus O(n) for the output tree.
// 时间复杂度：O(n)，空间复杂度：O(log n)，输出树另占 O(n)。
func sortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBSTRange(nums, 0, len(nums))
}

// Half-open [left, right) becomes one balanced BST.
// 左闭右开区间 [left, right) 构造成一棵平衡二叉搜索树。
// Time: O(n), Space: O(log n).
// 时间复杂度：O(n)，空间复杂度：O(log n)。
//
// 步骤与要点 / Steps and notes:
//  1. Empty half-open range: no node.
//     空的左闭右开区间：不建节点。
//  2. Midpoint root keeps left and right halves size-balanced.
//     中点作根，左右半段长度差至多为 1。
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
// 步骤与要点 / Steps and notes:
//  1. Same midpoint idea; slice views avoid copying element storage.
//     同样取中点；切片视图不会复制底层元素。
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
