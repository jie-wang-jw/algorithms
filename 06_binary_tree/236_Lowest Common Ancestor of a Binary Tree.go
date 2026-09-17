package _6_binary_tree

/*
236. 二叉树的最近公共祖先 / Lowest Common Ancestor of a Binary Tree

题目描述 / Problem Description
给定二叉树根节点和两个存在于树中的节点 p、q，返回它们的最近公共祖先。
最近公共祖先是同时拥有 p 和 q 作为后代的最低节点，节点可以是自己的后代。
Given a binary tree root and two nodes p and q that exist in the tree, return their lowest common ancestor.
A node may be a descendant of itself.

示例 / Examples
        3
       / \
      5   1
     / \ / \
    6  2 0  8
      / \
     7   4

p=5, q=1 → LCA=3（分居 3 的左右子树）
p=5, q=4 → LCA=5（4 在 5 下面，5 可以是自己的后代）
For p=5,q=1 the ancestor is 3; for p=5,q=4 it is 5.

与 235 的对比 / Contrast with 235 (BST)
235 利用 BST 有序性，每次只走进一侧，时间 O(h)。
本题是普通二叉树，没有大小关系可用，必须在两侧都搜索，后序合并左右结果，时间 O(n)。
235 uses BST order and walks only one child per step: O(h).
This problem is an ordinary binary tree with no value order, so both sides are searched and merged in postorder: O(n).
复杂度记号：n 为节点数，h 为树高，w 为最大层宽；辅助空间不含返回结果。
Notation: n nodes, height h, maximum width w; auxiliary space excludes returned results.
*/

// 1. Postorder recursion (recommended): a node is the LCA iff p and q are discovered in different subtrees, or it is p or q itself.
// 1. 后序递归：推荐；左右都找到目标，或当前节点就是 p/q，则当前节点是 LCA。
// Time: O(n), Space: O(h).
// 时间复杂度：O(n)，空间复杂度：O(h)。
//
// 要求 p、q 都在树中；递归返回该子树找到的目标或已确定的祖先，nil 表示未找到。
// Require both targets in the tree; return a found target or established ancestor, with nil meaning neither was found.
// 左右都非空说明目标分居两侧，当前根就是祖先；仅一侧非空就向上传递该结果，命中自身也直接返回。
// Two nonnil child results split the targets at the current root; propagate a single result, and return immediately on a target node.
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}
	return right
}

// 2. Parent pointers: walk from p to the root, then from q until that path is met.
// 2. 父指针：先记录 p 到根的路径，再从 q 向上走到这条路径上。
// Time: O(n), Space: O(n).
// 时间复杂度：O(n)，空间复杂度：O(n)。
//
// 先建立节点地址→父节点映射，标记 p 到根的祖先链；q 从自身向上遇到的第一个已标记节点就是最近公共祖先。
// Build parent links by node identity and mark p's ancestor chain; the first marked node on q's upward path is the LCA.
// 按题意要求 p、q 均存在且非空，值相同不能代替节点身份。
// Assume both targets exist and are nonnil; equal values do not replace node identity.
func lowestCommonAncestorParents(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	parent := map[*TreeNode]*TreeNode{root: nil}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		_, hasP := parent[p]
		_, hasQ := parent[q]
		if hasP && hasQ {
			break
		}
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			parent[node.Left] = node
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			parent[node.Right] = node
			queue = append(queue, node.Right)
		}
	}

	seen := map[*TreeNode]bool{}
	for cur := p; cur != nil; cur = parent[cur] {
		seen[cur] = true
	}
	for cur := q; cur != nil; cur = parent[cur] {
		if seen[cur] {
			return cur
		}
	}
	return nil
}
