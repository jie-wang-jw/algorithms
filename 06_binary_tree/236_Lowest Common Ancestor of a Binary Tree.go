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

解法一：后序递归（推荐） / Method 1: Postorder Recursion (Recommended)
在一个节点处先看左右子树有没有找到 p 或 q。
左右都非空：p、q 分居两侧，当前节点就是 LCA。
只有一侧非空：两个目标都在这一侧，把这一侧的返回值继续向上传。
当前节点本身就是 p 或 q：即使另一侧还没看完，这个节点也已经是候选祖先。
Search both children first.
Both nonempty: p and q sit on different sides, so the current node is the LCA.
Only one nonempty: both targets are in that subtree; propagate that result upward.
If the current node is p or q, it is already an ancestor candidate.

解法二：父指针迭代 / Method 2: Parent Pointers
先遍历整棵树记录每个节点的父指针，再把 p 到根的路径放入集合，沿 q 的父链向上，第一个已出现的节点就是 LCA。
Record parent pointers, store the path from p to the root, then walk from q until a recorded ancestor is found.

时间与空间复杂度 / Time and Space Complexity
n 为节点数，h 为树高。递归时间 O(n)，辅助空间 O(h)。
父指针法时间 O(n)，辅助空间 O(n)。
For n nodes and height h, recursion takes O(n) time and O(h) space.
The parent-pointer version takes O(n) time and O(n) space.
*/

// 1. Postorder recursion (recommended): a node is the LCA iff p and q are discovered in different subtrees, or it is p or q itself.
// 1. 后序递归：推荐；左右都找到目标，或当前节点就是 p/q，则当前节点是 LCA。
// Time: O(n), Space: O(h).
// 时间复杂度：O(n)，空间复杂度：O(h)。
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// Empty branch, or this node is p/q: stop and report upward.
	// 空分支，或当前就是 p/q：停止并向上回报。
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// Targets found on both sides: this node is the lowest common ancestor.
	// 左右都找到目标：当前节点就是最近公共祖先。
	if left != nil && right != nil {
		return root
	}

	// Only one side hit: both targets lie in that subtree (or only one was under this node).
	// 只有一侧非空：两个目标都在那一侧（或本子树只覆盖其中一个）。
	if left != nil {
		return left
	}
	return right
}

// 2. Parent pointers: walk from p to the root, then from q until that path is met.
// 2. 父指针：先记录 p 到根的路径，再从 q 向上走到这条路径上。
// Time: O(n), Space: O(n).
// 时间复杂度：O(n)，空间复杂度：O(n)。
func lowestCommonAncestorParents(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// BFS/level walk to record every node's parent until both p and q are known.
	// 层序记录每个节点的父指针，直到 p、q 都已入表。
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

	// Mark the path from p to the root, then climb from q until a marked ancestor appears.
	// 先标记 p 到根的路径，再从 q 向上爬，遇到已标记的祖先即为 LCA。
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
