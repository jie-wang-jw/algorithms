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

关键逻辑 / Key Logic
普通二叉树搜索最坏要看整棵树；BST 每次比较后只进入一侧，路径唯一。
当前值 > val → 只可能在左；当前值 < val → 只可能在右；相等则找到。
递归时必须 return 递归调用的结果，否则找到的子树根会在返回途中丢失。
A plain binary-tree search may scan the whole tree; a BST enters only one child after each comparison.
Larger current value → go left; smaller → go right; equal → found.
Recursive calls must return their results, or a found subtree root is discarded on the way up.

解法一：递归（推荐） / Method 1: Recursion (Recommended)
空节点或当前值等于 val 时返回当前节点；否则只搜一侧并返回该侧结果。
Return the current node when it is nil or equals val; otherwise search only one side and return that result.

解法二：迭代 / Method 2: Iteration
不需要栈：按大小原地走向左或右，相等返回，走到空则失败。
No stack: walk left or right in place, return on a match, fail when the walk hits nil.

时间与空间复杂度 / Time and Space Complexity
h 为树高。两版时间 O(h)：平衡约 O(log n)，链状最坏 O(n)。
递归辅助空间 O(h)；迭代 O(1)。
For height h, both take O(h) time: about O(log n) when balanced, O(n) when skewed.
Recursion uses O(h) space; iteration uses O(1).
*/

// 1. Recursion (recommended): search only the child that can contain val, and return that call's result.
// 1. 递归：推荐；只进入可能含 val 的一侧，并返回这次调用的结果。
// Time: O(h), Space: O(h).
// 时间复杂度：O(h)，空间复杂度：O(h)。
func searchBST(root *TreeNode, val int) *TreeNode {
	// Miss (nil) or hit: both are valid stopping points.
	// 没找到（空）或命中：两种都是合法终止。
	if root == nil || root.Val == val {
		return root
	}

	// Current value is too large; every right-side value is even larger.
	// 当前值偏大，右子树只会更大，只搜左边。
	if root.Val > val {
		return searchBST(root.Left, val)
	}

	// Current value is too small; search only the right subtree.
	// 当前值偏小，只搜右子树。
	return searchBST(root.Right, val)
}

// 2. Iteration: walk left or right in place; the ordered BST path needs no backtracking stack.
// 2. 迭代：按大小原地走向左或右，有序性已确定路径，不需要回溯栈。
// Time: O(h), Space: O(1).
// 时间复杂度：O(h)，空间复杂度：O(1)。
func searchBSTIterative(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val > val {
			root = root.Left
		} else if root.Val < val {
			root = root.Right
		} else {
			// Equal: return the subtree rooted here.
			// 相等：返回以当前节点为根的子树。
			return root
		}
	}
	return nil
}
