package _6_binary_tree

/*
98. 验证二叉搜索树 / Validate Binary Search Tree

题目描述 / Problem Description
判断二叉树是否为有效二叉搜索树：左子树所有值都小于根，右子树所有值都大于根，左右子树自身也是 BST。
相等值不合法。空树是 BST。
Determine whether a binary tree is a valid BST: every left-subtree value is less than the root,
every right-subtree value is greater, and both children are BSTs. Equal values are invalid. An empty tree is a BST.

示例 / Examples
合法 / Valid:     非法 / Invalid:
    2                  5
   / \                / \
  1   3              1   4
                        / \
                       3   6

右图：4 是 5 的右孩子，看似“右更大”，但 4 的左孩子 3 仍在 5 的右子树里，3 < 5，破坏 BST。
In the right tree, 4 looks like a larger right child of 5, yet 3 sits in 5's right subtree and 3 < 5, so the tree is invalid.

陷阱 / Pitfalls
不能只比较根和它的左右孩子。必须保证整棵左子树小于根、整棵右子树大于根。
中序遍历 BST 会得到严格递增序列，验证就变成检查这个顺序。
Do not only compare a node with its immediate children.
Inorder traversal of a BST is strictly increasing, so validation checks that order.
复杂度记号：n 为节点数，h 为树高，w 为最大层宽；辅助空间不含返回结果。
Notation: n nodes, height h, maximum width w; auxiliary space excludes returned results.
*/

// 1. Inorder array: convert the tree to values, then reject any non-increasing adjacent pair.
// 1. 中序数组：先把树转成序列，再检查是否存在非严格递增的相邻元素。
// Time: O(n), Space: O(n).
// 时间复杂度：O(n)，空间复杂度：O(n)。
//
// 严格 BST 等价于中序严格递增；必须比较相邻中序值，而不是只看每个节点与直接孩子的大小。
// A strict BST has strictly increasing inorder; compare neighboring inorder values, not only each parent with its immediate children.
// 相等也判非法；prev 节点是否为空标记“还没前驱”，避免用某个整数哨兵误排除极值。
// Equal values are invalid too; predecessor-node presence avoids numeric sentinels that could reject extreme values.
func isValidBST(root *TreeNode) bool {
	values := make([]int, 0)
	collectInorder(root, &values)

	for i := 1; i < len(values); i++ {
		if values[i] <= values[i-1] {
			return false
		}
	}
	return true
}

// Inorder collector: append node values in left-root-right order.
// 中序收集：按左、根、右的顺序把节点值写入切片。
// Time: O(n), Space: O(h).
// 时间复杂度：O(n)，空间复杂度：O(h)。
//
// 按左→根→右写入同一个结果切片；指针参数使追加后的切片头能回传给调用方。
// Append left→root→right into one slice; the pointer parameter propagates its updated slice header to the caller.
func collectInorder(node *TreeNode, values *[]int) {
	if node == nil {
		return
	}
	collectInorder(node.Left, values)
	*values = append(*values, node.Val)
	collectInorder(node.Right, values)
}

// 2. Inorder recursion with a predecessor (recommended): each value must exceed the previously visited node.
// 2. 中序递归比较前驱：推荐；每个值都必须严格大于刚刚访问过的节点。
// Time: O(n), Space: O(h).
// 时间复杂度：O(n)，空间复杂度：O(h)。
//
// 严格 BST 等价于中序严格递增；必须比较相邻中序值，而不是只看每个节点与直接孩子的大小。
// A strict BST has strictly increasing inorder; compare neighboring inorder values, not only each parent with its immediate children.
// 相等也判非法；prev 节点是否为空标记“还没前驱”，避免用某个整数哨兵误排除极值。
// Equal values are invalid too; predecessor-node presence avoids numeric sentinels that could reject extreme values.
func isValidBSTInorder(root *TreeNode) bool {
	var prev *TreeNode
	var valid func(*TreeNode) bool
	valid = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		if !valid(node.Left) {
			return false
		}

		if prev != nil && node.Val <= prev.Val {
			return false
		}
		prev = node
		return valid(node.Right)
	}
	return valid(root)
}

// 3. Iterative inorder: the same predecessor check, using an explicit stack.
// 3. 中序迭代：同样用前驱检查，改成显式栈。
// Time: O(n), Space: O(h).
// 时间复杂度：O(n)，空间复杂度：O(h)。
//
// 严格 BST 等价于中序严格递增；必须比较相邻中序值，而不是只看每个节点与直接孩子的大小。
// A strict BST has strictly increasing inorder; compare neighboring inorder values, not only each parent with its immediate children.
// 相等也判非法；prev 节点是否为空标记“还没前驱”，避免用某个整数哨兵误排除极值。
// Equal values are invalid too; predecessor-node presence avoids numeric sentinels that could reject extreme values.
func isValidBSTIterative(root *TreeNode) bool {
	stack := []*TreeNode{}
	var prev *TreeNode
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if prev != nil && cur.Val <= prev.Val {
			return false
		}
		prev = cur
		cur = cur.Right
	}
	return true
}
