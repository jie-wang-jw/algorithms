package _6_binary_tree

/*
700. 二叉搜索树中的搜索 / Search in a Binary Search Tree

题目描述 / Problem Description
在二叉搜索树中找出值等于 val 的节点，并返回以该节点为根的子树；不存在则返回 nil。
BST 满足：左子树所有值 < 根 < 右子树所有值，左右子树也是 BST。
Find the node valued val in a BST and return that subtree, or nil if it is absent.
Every left subtree value is less than the root, every right subtree value is greater, and both children are BSTs.

示例 / Examples
      4
     / \
    2   7
   / \
  1   3

val=2 → 返回以 2 为根的子树：
    2
   / \
  1   3
val=5 → 沿 4→7 后落到空，返回 nil。
val=2 returns the subtree rooted at 2; val=5 walks 4→7 then falls off to nil.
复杂度记号：n 为节点数，h 为树高，w 为最大层宽；辅助空间不含返回结果。
Notation: n nodes, height h, maximum width w; auxiliary space excludes returned results.
*/

// 1. Recursion (recommended): search only the child that can contain val, and return that call's result.
// 1. 递归：推荐；只进入可能含 val 的一侧，并返回这次调用的结果。
// Time: O(h), Space: O(h).
// 时间复杂度：O(h)，空间复杂度：O(h)。
//
// BST 的整棵左树小于根、整棵右树大于根，故比较一次就可排除一侧；命中返回该节点，走到 nil 表示不存在。
// BST order excludes one whole side per comparison; return the matched node or nil if the search path ends.
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}

	if root.Val > val {
		return searchBST(root.Left, val)
	}

	return searchBST(root.Right, val)
}

// 2. Iteration: walk left or right in place; the ordered BST path needs no backtracking stack.
// 2. 迭代：按大小原地走向左或右，有序性已确定路径，不需要回溯栈。
// Time: O(h), Space: O(1).
// 时间复杂度：O(h)，空间复杂度：O(1)。
//
// BST 的整棵左树小于根、整棵右树大于根，故比较一次就可排除一侧；命中返回该节点，走到 nil 表示不存在。
// BST order excludes one whole side per comparison; return the matched node or nil if the search path ends.
func searchBSTIterative(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val > val {
			root = root.Left
		} else if root.Val < val {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}
