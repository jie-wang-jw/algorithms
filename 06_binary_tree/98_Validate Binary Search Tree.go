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

解法一：中序写入数组 / Method 1: Inorder Array
中序收集全部节点值，再检查是否严格递增。最直观，额外 O(n) 数组。
Collect every inorder value, then check that the slice is strictly increasing.

解法二：中序递归比较前驱（推荐） / Method 2: Inorder Recursion with a Predecessor (Recommended)
中序过程中保存上一个访问的节点。当前值必须严格大于前驱，否则立刻失败。
用节点指针而不是 int 最小值作初始前驱，避免根值等于 int 最小值时误判。
Keep the previously visited node during inorder. The current value must be strictly greater than that predecessor.
Use a node pointer, not a minimum int, so a root equal to math.MinInt still works.

解法三：中序迭代 / Method 3: Iterative Inorder
用栈模拟中序，同样用前驱指针检查严格递增。
Simulate inorder with a stack and the same predecessor check.

时间与空间复杂度 / Time and Space Complexity
n 为节点数，h 为树高。三版时间都是 O(n)。
数组版辅助空间 O(n)；递归版 O(h)；迭代版 O(h)。
For n nodes and height h, all three take O(n) time.
The array version uses O(n) extra space; recursion and iteration use O(h).
*/

// 1. Inorder array: convert the tree to values, then reject any non-increasing adjacent pair.
// 1. 中序数组：先把树转成序列，再检查是否存在非严格递增的相邻元素。
// Time: O(n), Space: O(n).
// 时间复杂度：O(n)，空间复杂度：O(n)。
func isValidBST(root *TreeNode) bool {
	values := make([]int, 0)
	collectInorder(root, &values)

	// Adjacent equals or inversions both invalidate a BST.
	// 相邻相等或逆序都不是合法 BST。
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
func isValidBSTInorder(root *TreeNode) bool {
	var prev *TreeNode
	var valid func(*TreeNode) bool
	valid = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		// Left subtree must already be a valid increasing sequence.
		// 左子树必须已经构成合法递增序列。
		if !valid(node.Left) {
			return false
		}

		// Current value must be strictly greater than the previous inorder value.
		// 当前值必须严格大于中序前驱。
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
func isValidBSTIterative(root *TreeNode) bool {
	stack := []*TreeNode{}
	var prev *TreeNode
	cur := root
	for cur != nil || len(stack) > 0 {
		// Descend left, saving ancestors to visit after the left subtree.
		// 先向左深入，祖先稍后在左子树完成后访问。
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
