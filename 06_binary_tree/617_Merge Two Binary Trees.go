package _6_binary_tree

/*
617. 合并二叉树 / Merge Two Binary Trees

题目描述 / Problem Description
将两棵二叉树覆盖合并：重叠节点的值相加，只在一棵树出现的节点直接保留。
必须从两棵树的根开始合并。本文件的原地写法修改 root1 的结构并返回它。
Merge two binary trees by adding overlapping node values and keeping nodes that exist in only one tree.
Merging starts at the roots. The in-place versions mutate root1 and return it.

示例 / Example
   1         2              3
  / \       / \            / \
 3   2     1   3    ->    4   5
/           \   \        / \   \
5            4   7      5   4   7
复杂度记号：n 为节点数，h 为树高，w 为最大层宽；辅助空间不含返回结果。
Notation: n nodes, height h, maximum width w; auxiliary space excludes returned results.
*/

// 1. Preorder recursion reusing root1 (recommended): add into root1, then merge both children onto it.
// 1. 前序递归复用 root1：推荐；把值加到 root1，再把左右孩子的合并结果接回去。
// Time: O(n), Space: O(h).
// 时间复杂度：O(n)，空间复杂度：O(h)。
//
// 重叠位置累加到第一棵树；仅一侧存在时直接复用该子树，不复制节点，所以结果可能与第二棵树共享节点。
// Add overlapping values into the first tree; reuse one-sided subtrees without copying, so the result can share nodes with the second input.
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

// 2. Preorder recursion building a new tree: allocate a sum node when both sides exist.
// 2. 前序递归新建树：两侧都存在时新建求和节点，不修改输入。
// Time: O(n), Space: O(h) plus new overlapping nodes.
// 时间复杂度：O(n)，空间复杂度：O(h)，并额外分配重叠路径上的新节点。
//
// 只有两侧都非空时才创建新节点并相加；只有一侧存在时直接返回原子树，所以这是部分复制，不是完整深拷贝。
// Allocate only where both inputs exist; a one-sided subtree is returned directly, so this is a partial copy, not a deep copy.
func mergeTreesNew(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	root := &TreeNode{Val: root1.Val + root2.Val}
	root.Left = mergeTreesNew(root1.Left, root2.Left)
	root.Right = mergeTreesNew(root1.Right, root2.Right)
	return root
}

// 3. Iterative pair queue: dequeue two coexisting nodes, add values, and attach or enqueue children.
// 3. 队列成对迭代：取出一对同时存在的节点，加值后再按孩子是否存在入队或直接拼接。
// Time: O(n), Space: O(w).
// 时间复杂度：O(n)，空间复杂度：O(w)。
//
// 重叠位置累加到第一棵树；仅一侧存在时直接复用该子树，不复制节点，所以结果可能与第二棵树共享节点。
// Add overlapping values into the first tree; reuse one-sided subtrees without copying, so the result can share nodes with the second input.
func mergeTreesIterative(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	queue := []*TreeNode{root1, root2}
	for len(queue) > 0 {
		node1 := queue[0]
		node2 := queue[1]
		queue = queue[2:]
		node1.Val += node2.Val

		if node1.Left != nil && node2.Left != nil {
			queue = append(queue, node1.Left, node2.Left)
		}
		if node1.Right != nil && node2.Right != nil {
			queue = append(queue, node1.Right, node2.Right)
		}

		if node1.Left == nil && node2.Left != nil {
			node1.Left = node2.Left
		}
		if node1.Right == nil && node2.Right != nil {
			node1.Right = node2.Right
		}
	}

	return root1
}
