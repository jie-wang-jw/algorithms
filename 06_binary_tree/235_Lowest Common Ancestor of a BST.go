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

关键逻辑 / Key Logic
BST 有序性给出唯一搜索路径：从根往下走时，只需比较当前值与 p、q。
- 当前值同时大于 p、q → 两个目标都在左子树，LCA 必在左边。
- 当前值同时小于 p、q → 两个目标都在右子树，LCA 必在右边。
- 否则当前节点就是分叉点：要么 p、q 分居两侧，要么当前节点等于 p 或 q。
第一种“分居两侧”和第二种“自身即祖先”在代码里是同一种返回：直接返回当前节点。
BST order gives a unique walk: compare the current value with p and q.
- Greater than both → both targets lie left; the LCA is in the left subtree.
- Smaller than both → both targets lie right; the LCA is in the right subtree.
- Otherwise this node is the split: either p and q sit on different sides, or the node equals p or q.
Both “different sides” and “self-ancestor” share the same return: the current node.

与 236 的对比 / Contrast with 236 (ordinary binary tree)
236 没有有序性，必须后序搜完整棵左右子树，看两边是否都找到目标，时间 O(n)。
235 利用 BST，每次只走进一侧，路径唯一，时间 O(h)，不必探另一侧。
236 has no order, so it postorder-searches both full subtrees and checks whether both sides found a target: O(n).
235 uses BST order, walks only one child each step along a unique path: O(h), never exploring the other side.

解法一：递归 / Method 1: Recursion
按三种情况进入左、进入右，或返回当前节点。递归深度等于路径长度。
Follow the three cases: recurse left, recurse right, or return the current node.
Recursion depth equals the path length.

解法二：迭代（推荐） / Method 2: Iteration (Recommended)
路径唯一，不需要栈。用循环原地走向左或右，直到当前节点夹在 p、q 之间（含等于）。
The path is unique, so no stack is required. Loop in place until the current node lies between p and q (inclusive).

时间与空间复杂度 / Time and Space Complexity
h 为树高。两版时间 O(h)：平衡约 O(log n)，链状最坏 O(n)。
递归辅助空间 O(h)；迭代 O(1)。
For height h, both take O(h) time: about O(log n) when balanced, O(n) when skewed.
Recursion uses O(h) space; iteration uses O(1).
*/

// 1. Recursion: both targets left → search left; both right → search right; otherwise this node is the split.
// 1. 递归：两个目标都在左就搜左，都在右就搜右，否则当前节点就是分叉点。
// Time: O(h), Space: O(h).
// 时间复杂度：O(h)，空间复杂度：O(h)。
//
// 步骤与要点 / Steps and notes:
//  1. Empty tree has no ancestor.
//     空树没有祖先。
//  2. Both targets are smaller than root, so the LCA must lie in the left subtree.
//     两个目标都比当前值小，LCA 一定在左子树。
//  3. Both targets are larger than root, so the LCA must lie in the right subtree.
//     两个目标都比当前值大，LCA 一定在右子树。
//  4. Split found: p and q diverge here, or root equals p or q.
//     找到分叉：p、q 分居两侧，或当前节点就是 p / q。
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
// 步骤与要点 / Steps and notes:
//  1. Same three-way decision as the recursive version, without a call stack.
//     与递归版相同的三分支判断，但不使用调用栈。
//  2. First node that is not strictly outside both targets is the LCA.
//     第一个不再严格落在两个目标同一侧之外的节点，就是 LCA。
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
