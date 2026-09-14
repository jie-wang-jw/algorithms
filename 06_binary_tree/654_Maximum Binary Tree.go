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

解法一：递归切片 / Method 1: Recursive Slicing
在当前切片中扫描最大值下标 k，用 nums[k] 建根。
左子树递归 nums[:k]，右子树递归 nums[k+1:]。空切片返回 nil。
Scan the current slice for the max index k and build the root from nums[k].
Recurse on nums[:k] and nums[k+1:]. Return nil for an empty slice.

解法二：下标区间（推荐） / Method 2: Index Ranges (Recommended)
与解法一相同，但用左闭右开区间 [left,right) 描述当前子数组，避免每层切出新切片。
left==right 表示空区间。最大值下标 maxIndex 把区间切成 [left,maxIndex) 与 [maxIndex+1,right)。
Same idea as method 1, using a half-open range [left,right) instead of new slices.
left==right is empty. maxIndex splits the range into [left,maxIndex) and [maxIndex+1,right).

文章指出：第一版用 if 拦住空区间，不让空节点进入递归；第二版允许空区间进入，靠 left>=right 终止。
The article notes that method 1 guards empty children with if, while method 2 lets empty ranges recurse and stop at left>=right.

时间与空间复杂度 / Time and Space Complexity
n 为数组长度，h 为结果树高度。两版最坏时间都是 O(n²)，因为每层都可能线性扫描当前区间。
辅助空间 O(h)；输出树另占 O(n)。Go 切片是视图，解法一不会像 C++ vector 那样复制元素。
For n values and height h, both methods take O(n²) worst-case time from linear max scans.
Auxiliary space is O(h); the output tree uses O(n). Go slices are views, so method 1 does not copy elements.
*/

// 1. Recursive slicing: scan the current slice for the maximum, then recurse on the left and right subarrays.
// 1. 递归切片：在当前切片中找最大值，再对左右两侧子数组递归构造。
// Time: O(n²) worst case, Space: O(h) plus O(n) for the output tree.
// 时间复杂度：最坏 O(n²)，空间复杂度：O(h)，输出树另占 O(n)。
func constructMaximumBinaryTree(nums []int) *TreeNode {
	// An empty range produces no node.
	// 空区间不能构造节点。
	if len(nums) == 0 {
		return nil
	}

	// The largest value in this slice becomes the current root.
	// 当前切片中的最大值就是这一层的根。
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
func constructMaximumBinaryTreeIndex(nums []int) *TreeNode {
	return constructMaximumBinaryTreeRange(nums, 0, len(nums))
}

// Half-open [left, right) describes the subarray that should become one maximum binary tree.
// 左闭右开区间 [left, right) 表示要构造成一棵最大二叉树的子数组。
// Time: O(n²) worst case, Space: O(h).
// 时间复杂度：最坏 O(n²)，空间复杂度：O(h)。
func constructMaximumBinaryTreeRange(nums []int, left, right int) *TreeNode {
	// Equal bounds mean this interval is empty.
	// 两端相等表示当前区间为空。
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
	// Left child uses values strictly left of the maximum.
	// 左子树使用最大值左侧的区间。
	root.Left = constructMaximumBinaryTreeRange(nums, left, maxIndex)
	// Right child uses values strictly right of the maximum.
	// 右子树使用最大值右侧的区间。
	root.Right = constructMaximumBinaryTreeRange(nums, maxIndex+1, right)
	return root
}
