package _6_binary_tree

/*
235. 二叉搜索树的最近公共祖先 / Lowest Common Ancestor of a Binary Search Tree

题目描述 / Problem Description
给定二叉搜索树根节点和两个存在于树中的节点 p、q，返回它们的最近公共祖先（LCA）。
最近公共祖先是同时拥有 p 和 q 作为后代的最低节点；节点可以是自己的后代。
Given a BST root and two nodes p and q that exist in the tree, return their lowest common ancestor (LCA).
The LCA is the lowest node that has both p and q as descendants; a node may be a descendant of itself.

示例 / Examples
树 / Tree:
        6
       / \
      2   8
     / \ / \
    0  4 7  9
      / \
     3   5

情形一：分叉在中间节点 / Case 1: split at an internal node
p=2, q=8 → LCA=6（一个在左子树，一个在右子树）
p=2, q=8 → LCA=6 (one target left of 6, one right of 6)

情形二：其中一个就是祖先 / Case 2: one node is the ancestor of the other
p=2, q=4 → LCA=2（4 在 2 的右子树里，2 可以是自己的后代）
p=2, q=4 → LCA=2 (4 sits under 2; a node may be its own descendant)
复杂度记号：n 为节点数，h 为树高，w 为最大层宽；辅助空间不含返回结果。
Notation: n nodes, height h, maximum width w; auxiliary space excludes returned results.
*/

// 1. Recursion: both targets left → search left; both right → search right; otherwise this node is the split.
// 1. 递归：两个目标都在左就搜左，都在右就搜右，否则当前节点就是分叉点。
// Time: O(h), Space: O(h).
// 时间复杂度：O(h)，空间复杂度：O(h)。
//
// 要求 p、q 均存在于值互异的 BST；两值同小则祖先在左树，同大则在右树，首次分叉或命中目标处就是最近公共祖先。
// Require p and q in a unique-valued BST; descend while both lie on one side, stopping at their first split or either target.
func lowestCommonAncestorBST(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestorBST(root.Left, p, q)
	}

	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestorBST(root.Right, p, q)
	}

	return root
}

// 2. Iteration (recommended): walk left or right in place until the current node sits between p and q.
// 2. 迭代：推荐；按大小原地走向左或右，直到当前节点夹在 p 与 q 之间。
// Time: O(h), Space: O(1).
// 时间复杂度：O(h)，空间复杂度：O(1)。
//
// 要求 p、q 均存在于值互异的 BST；两值同小则祖先在左树，同大则在右树，首次分叉或命中目标处就是最近公共祖先。
// Require p and q in a unique-valued BST; descend while both lie on one side, stopping at their first split or either target.
func lowestCommonAncestorBSTIterative(root, p, q *TreeNode) *TreeNode {
	cur := root
	for cur != nil {
		if cur.Val > p.Val && cur.Val > q.Val {
			cur = cur.Left
		} else if cur.Val < p.Val && cur.Val < q.Val {
			cur = cur.Right
		} else {
			return cur
		}
	}
	return nil
}
