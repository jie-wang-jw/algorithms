package _6_binary_tree

/*
669. 修剪二叉搜索树 / Trim a Binary Search Tree

题目描述 / Problem Description
给定 BST 根节点和闭区间 [low, high]，删除所有值不在区间内的节点，保持 BST 性质，返回新根。
Trim a BST so that every remaining node value lies in the closed interval [low, high], and return the new root.

示例 / Examples
修剪前 / Before, [low,high]=[1,3]:     修剪后 / After:
      3                                    3
     / \                                  /
    0   4                                2
     \                                  /
      2                                1
     /
    1

0 < low，整棵以 0 为根的左支里小于 1 的部分被丢掉，只保留落在 [1,3] 内的 1、2、3；4 > high 整侧丢弃。
0 is below low, so values under it that stay below 1 are dropped; only 1, 2, 3 remain in [1,3]; 4 exceeds high and its side is discarded.

另一例 / Another example, [low,high]=[1,2]:
      3                                    2
     / \                                  /
    0   4                                1
     \
      2
     /
    1
根 3 > high，答案只可能在左子树；继续修剪后得到以 2 为根的树。
Root 3 exceeds high, so the answer can only come from the left; further trimming yields the tree rooted at 2.
复杂度记号：n 为节点数，h 为树高，w 为最大层宽；辅助空间不含返回结果。
Notation: n nodes, height h, maximum width w; auxiliary space excludes returned results.
*/

// 1. Recursion (recommended): drop the whole left or right subtree when the root is outside [low, high].
// 1. 递归：推荐；根落在区间外时整侧子树都可以丢掉，只修剪可能合法的一侧。
// Time: O(n), Space: O(h).
// 时间复杂度：O(n)，空间复杂度：O(h)。
//
// root.Val<low 时，整棵左树都更小，可直接丢弃，仅修剪右树；超过 high 时同理仅保留左侧候选。
// If root.Val<low, its entire left subtree is smaller and can be discarded; symmetrically keep only the left candidate above high.
// 根在范围内时保留它并接回两侧修剪结果；要求 BST 且 low<=high，原地修改连接。
// Retain an in-range root and reconnect trimmed children; require a BST and low<=high, with in-place rewiring.
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < low {
		return trimBST(root.Right, low, high)
	}

	if root.Val > high {
		return trimBST(root.Left, low, high)
	}

	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}

// 2. Iteration: first move the root into range, then prune illegal nodes on each remaining spine.
// 2. 迭代：先把根移进区间，再分别删掉左右链上越界的节点。
// Time: O(n), Space: O(1).
// 时间复杂度：O(n)，空间复杂度：O(1)。
//
// 先沿 BST 方向找到合法根；此后左边界只可能偏小，右边界只可能偏大。
// First find an in-range root by BST direction; afterward only the left boundary can be too small and the right too large.
// 左孩子偏小时用其右树替代，原左树必全无效；右边界对称处理。替代后仍检查同一位置直到合法。
// Replace an undersized left child by its right subtree, discarding its invalid left subtree; handle the right boundary symmetrically and recheck replacements.
func trimBSTIterative(root *TreeNode, low int, high int) *TreeNode {
	for root != nil && (root.Val < low || root.Val > high) {
		if root.Val < low {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	if root == nil {
		return nil
	}

	cur := root
	for cur.Left != nil {
		if cur.Left.Val < low {
			cur.Left = cur.Left.Right
		} else {
			cur = cur.Left
		}
	}

	cur = root
	for cur.Right != nil {
		if cur.Right.Val > high {
			cur.Right = cur.Right.Left
		} else {
			cur = cur.Right
		}
	}
	return root
}
