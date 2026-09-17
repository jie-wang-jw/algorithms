package _6_binary_tree

/*
110. 平衡二叉树 / Balanced Binary Tree

题目描述 / Problem Description
给定二叉树根节点 root，判断它是否为高度平衡二叉树。
高度平衡要求每个节点的左右子树高度差的绝对值都不超过 1，而不只是根节点满足条件。
空树是平衡的；高度按节点数计算，空树高 0、叶子高 1。
Given the root of a binary tree, determine whether it is height-balanced.
At every node, the left and right subtree heights must differ by at most one, not just at the root.
An empty tree is balanced. Heights count nodes: an empty tree has height 0 and a leaf height 1.

示例 / Examples
平衡 / Balanced:          不平衡 / Unbalanced:
       3                        1
      / \                      /
     9  20                    2
       /  \                  /
      15   7                3
左图根两边高度为 1、2，其他节点也满足高度差要求。右图根两边高度为 2、0，差为 2。
The left tree has root-subtree heights 1 and 2, and every other node also satisfies the rule.
The right tree has root-subtree heights 2 and 0, differing by 2.
复杂度记号：n 为节点数，h 为树高，w 为最大层宽；辅助空间不含返回结果。
Notation: n nodes, height h, maximum width w; auxiliary space excludes returned results.
*/

// 1. Top-down recursion: check this node's height difference, then recurse into both subtrees; height work is repeated.
// 1. 自顶向下递归：先判断当前高度差，再分别判断左右子树；求高度会重复遍历。
// Time: O(n²) conservative, Space: O(h).
// 时间复杂度：保守上界 O(n²)，空间复杂度：O(h)。
//
// 每个节点都必须满足左右高度差<=1，不能只检查根；分别调用 maxDepth 再检查孩子，会重复计算子树高度。
// Every node must have child heights differing by at most one; repeated maxDepth calls recompute subtree heights.
func isBalancedTopDown(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftHeight := maxDepth(root.Left)
	rightHeight := maxDepth(root.Right)
	diff := leftHeight - rightHeight

	if diff > 1 || diff < -1 {
		return false
	}

	return isBalancedTopDown(root.Left) &&
		isBalancedTopDown(root.Right)
}

// 2. Bottom-up postorder recursion (recommended): a nonnegative helper result means the whole tree is balanced.
// 2. 自底向上后序递归：推荐；辅助函数返回非负高度即整棵树平衡。
// Time: O(n), Space: O(h).
// 时间复杂度：O(n)，空间复杂度：O(h)。
//
// 返回值同时编码高度和失败：空树为 0，失衡为 -1；孩子已失衡时立即把 -1 向上传播。
// Encode height or failure in one result: nil is 0 and imbalance is -1, propagated immediately from either child.
// 仅在两孩子都平衡时比较高度差，再返回较高者+1；后序把每个高度只计算一次。
// Compare heights only after both children succeed; return max+1. Postorder computes each height once.
func isBalanced(root *TreeNode) bool {
	return balancedHeight(root) != -1
}

// Postorder height helper: return the real height, or -1 as a sentinel once any subtree is unbalanced.
// 后序求高辅助函数：子树平衡时返回真实高度，一旦失衡就返回哨兵 -1，该值不能再当高度使用。
// Time: O(n), Space: O(h).
// 时间复杂度：O(n)，空间复杂度：O(h)。
//
// 返回值同时编码高度和失败：空树为 0，失衡为 -1；孩子已失衡时立即把 -1 向上传播。
// Encode height or failure in one result: nil is 0 and imbalance is -1, propagated immediately from either child.
// 仅在两孩子都平衡时比较高度差，再返回较高者+1；后序把每个高度只计算一次。
// Compare heights only after both children succeed; return max+1. Postorder computes each height once.
func balancedHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := balancedHeight(root.Left)

	if leftHeight == -1 {
		return -1
	}

	rightHeight := balancedHeight(root.Right)
	if rightHeight == -1 {
		return -1
	}

	diff := leftHeight - rightHeight
	if diff > 1 || diff < -1 {
		return -1
	}

	return max(leftHeight, rightHeight) + 1
}

// 3. Iterative postorder: finish both children, then read heights from a node-pointer map and check the difference.
// 3. 后序迭代：左右孩子都完成后，用节点指针当 key 的高度表检查差值。
// Time: O(n) expected, Space: O(n) for the height map plus O(h) for the stack.
// 时间复杂度：平均 O(n)，空间复杂度：O(n)，由高度表产生，外加 O(h) 的栈。
//
// 栈模拟后序，prev 记录最近完成的子树根；右孩子未完成时先转右，不能提前计算父节点高度。
// Simulate postorder with a stack; prev marks the last completed subtree, so visit an unfinished right child before evaluating its parent.
// heights 保存已完成子树高度，nil 查表为 0；两边高度已知后检查差值并存入当前高度。
// heights stores completed subtree heights, with missing nil keys yielding 0; compare and store only after both children finish.
func isBalancedIterative(root *TreeNode) bool {
	stack := []*TreeNode{}
	heights := make(map[*TreeNode]int)
	node := root

	var prev *TreeNode

	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]

		if node.Right != nil && node.Right != prev {
			node = node.Right
			continue
		}

		leftHeight := heights[node.Left]
		rightHeight := heights[node.Right]
		diff := leftHeight - rightHeight

		if diff > 1 || diff < -1 {
			return false
		}

		heights[node] = max(leftHeight, rightHeight) + 1
		stack = stack[:len(stack)-1]
		prev = node

		node = nil
	}

	return true
}
