package _6_binary_tree

import "container/list"

/*
226. 翻转二叉树 / Invert Binary Tree

题目描述 / Problem Description
给定二叉树的根节点 root，翻转这棵二叉树，并返回它的根节点。
翻转指得到左右镜像，不是上下颠倒，也不是仅交换节点值。
Given the root of a binary tree, invert the tree and return its root.
Inversion produces a left-right mirror, not an upside-down tree or merely swapped values.
复杂度记号：n 为节点数，h 为树高，w 为最大层宽；辅助空间不含返回结果。
Notation: n nodes, height h, maximum width w; auxiliary space excludes returned results.
*/

// 1. Preorder recursion (recommended): swap the two child pointers first, then invert each relocated subtree.
// 1. 前序递归：推荐；先交换左右孩子指针，再分别翻转换位后的两棵子树。
// Time: O(n), Space: O(h) for the recursion stack; auxiliary space is not O(1).
// 时间复杂度：O(n)，空间复杂度：O(h)，由递归栈产生，不能写成 O(1)。
//
// 每个节点交换左右孩子恰好一次，整棵树即成镜像；先交换再遍历或先遍历再交换都可，但不能重复处理同一子树。
// Swap each node's children exactly once to mirror the tree; either pre- or postorder works without revisiting a subtree.
// 操作原地修改节点连接，返回原根引用；空树保持为空。
// Rewire the original tree in place and return its root; nil remains nil.
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

// 2. Postorder recursion: invert both child interiors first, then swap their positions.
// 2. 后序递归：先翻转左右子树内部，再交换它们的位置。
// Time: O(n), Space: O(h) for the recursion stack.
// 时间复杂度：O(n)，空间复杂度：O(h)，由递归栈产生。
//
// 每个节点交换左右孩子恰好一次，整棵树即成镜像；先交换再遍历或先遍历再交换都可，但不能重复处理同一子树。
// Swap each node's children exactly once to mirror the tree; either pre- or postorder works without revisiting a subtree.
// 操作原地修改节点连接，返回原根引用；空树保持为空。
// Rewire the original tree in place and return its root; nil remains nil.
func invertTreePostorderRecursive(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTreePostorderRecursive(root.Left)
	invertTreePostorderRecursive(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

// 3. Iterative preorder: swap on first arrival, push the node, then descend into the swapped left child.
// 3. 前序迭代：首次到达时交换，压栈后再深入交换后的左子树，出栈后处理另一边。
// Time: O(n), Space: O(h) for the stack.
// 时间复杂度：O(n)，空间复杂度：O(h)，由显式栈产生。
//
// 每个节点交换左右孩子恰好一次，整棵树即成镜像；先交换再遍历或先遍历再交换都可，但不能重复处理同一子树。
// Swap each node's children exactly once to mirror the tree; either pre- or postorder works without revisiting a subtree.
// 操作原地修改节点连接，返回原根引用；空树保持为空。
// Rewire the original tree in place and return its root; nil remains nil.
func invertTreePreorderIterative(root *TreeNode) *TreeNode {
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			node.Left, node.Right = node.Right, node.Left
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Right
	}
	return root
}

// 4. Iterative postorder: swap a node only after both child subtrees are done, using prev as the last finished root.
// 4. 后序迭代：用 prev 记录最近完成的子树根，左右都完成后才交换当前节点。
// Time: O(n), Space: O(h) for the stack.
// 时间复杂度：O(n)，空间复杂度：O(h)，由显式栈产生。
//
// prev 记录最近完成的子树根；右树不存在或等于 prev 时，两边已完成，才交换当前节点并出栈。
// prev marks the last completed subtree; only swap and finish a node when its right subtree is absent or already completed.
// 否则将当前节点放回栈，先处理右树；完成后 node=nil 防止重新走入已翻转的孩子。
// Otherwise restore the parent to the stack and visit its right subtree; set node=nil after finishing to avoid revisiting flipped children.
func invertTreePostorderIterative(root *TreeNode) *TreeNode {
	stack := []*TreeNode{}
	node := root
	var prev *TreeNode
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Right == nil || node.Right == prev {
			node.Left, node.Right = node.Right, node.Left
			prev = node
			node = nil
		} else {
			stack = append(stack, node)
			node = node.Right
		}
	}
	return root
}

// 5. Breadth-first traversal: swap each dequeued node once, then enqueue both children without grouping levels.
// 5. 层序遍历：每个节点出队时交换一次，再将两个孩子入队，不必按层分组。
// Time: O(n), Space: O(w) for the queue.
// 时间复杂度：O(n)，空间复杂度：O(w)，由队列产生。
//
// 每个节点交换左右孩子恰好一次，整棵树即成镜像；先交换再遍历或先遍历再交换都可，但不能重复处理同一子树。
// Swap each node's children exactly once to mirror the tree; either pre- or postorder works without revisiting a subtree.
// 操作原地修改节点连接，返回原根引用；空树保持为空。
// Rewire the original tree in place and return its root; nil remains nil.
func invertTreeLevelOrder(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return root
}
