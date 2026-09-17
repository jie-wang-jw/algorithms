package _6_binary_tree

/*
404. 左叶子之和 / Sum of Left Leaves

题目描述 / Problem Description
给定二叉树根节点 root，返回所有左叶子节点值的总和。
左叶子必须同时满足：它是父节点的左孩子，且它自己没有左右孩子。
空树和只有根节点的树返回 0；根没有父节点，所以根即使是叶子也不是左叶子。
Given the root of a binary tree, return the sum of all left leaf values.
A left leaf is both its parent's left child and a node with no children of its own.
An empty tree or a tree containing only the root returns 0; the root has no parent and is not a left leaf.

示例 / Example
       3
      / \
     9  20
       /  \
      15   7
左叶子是 9 和 15，结果为 24。7 是右叶子，不计入。
15 虽然在整棵树的右子树里，但它是 20 的左孩子，所以仍是左叶子。
The left leaves are 9 and 15, giving 24. Node 7 is a right leaf and is excluded.
Although 15 lies in the whole tree's right subtree, it is the left child of 20, so it counts.
复杂度记号：n 为节点数，h 为树高，w 为最大层宽；辅助空间不含返回结果。
Notation: n nodes, height h, maximum width w; auxiliary space excludes returned results.
*/

// 1. Recursion: identify a left leaf from its parent, then always search the right subtree as well.
// 1. 递归：从父节点判断左孩子是否为叶子，并且右子树也要继续找左叶子。
// Time: O(n), Space: O(h) for the recursion stack.
// 时间复杂度：O(n)，空间复杂度：O(h)，由递归栈产生。
//
// 左叶子必须同时满足“是父节点的 Left”且“自身无孩子”；因此从父节点检查，不能把所有左孩子都累加。
// A left leaf must be its parent's Left and have no children, so test from the parent rather than summing every left child.
// 右子树内部也可能有左叶子，仍须遍历；单独的根不是左叶子。
// The right subtree may contain left leaves too and must be visited; a lone root is not a left leaf.
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftSum := 0

	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			leftSum = root.Left.Val
		} else {
			leftSum = sumOfLeftLeaves(root.Left)
		}
	}

	rightSum := sumOfLeftLeaves(root.Right)

	return leftSum + rightSum
}

// 2. Iterative DFS: pop a parent, add its left child's value only when that child is a leaf, then push existing children.
// 2. 栈迭代：弹出父节点后，仅当左孩子是叶子才累加，再把非空孩子入栈。
// Time: O(n), Space: O(h) for the stack.
// 时间复杂度：O(n)，空间复杂度：O(h)，由栈产生。
//
// 左叶子必须同时满足“是父节点的 Left”且“自身无孩子”；因此从父节点检查，不能把所有左孩子都累加。
// A left leaf must be its parent's Left and have no children, so test from the parent rather than summing every left child.
// 右子树内部也可能有左叶子，仍须遍历；单独的根不是左叶子。
// The right subtree may contain left leaves too and must be visited; a lone root is not a left leaf.
func sumOfLeftLeavesIterative(root *TreeNode) int {
	if root == nil {
		return 0
	}

	stack := []*TreeNode{root}
	sum := 0

	for len(stack) > 0 {
		last := len(stack) - 1
		node := stack[last]
		stack = stack[:last]

		if node.Left != nil &&
			node.Left.Left == nil &&
			node.Left.Right == nil {
			sum += node.Left.Val
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return sum
}

// 3. Breadth-first search: the same parent-side left-leaf check, taking nodes from the queue front.
// 3. 队列迭代：判断规则与栈版本相同，只是从队首取节点。
// Time: O(n), Space: O(w) for the queue.
// 时间复杂度：O(n)，空间复杂度：O(w)，由队列产生。
//
// 左叶子必须同时满足“是父节点的 Left”且“自身无孩子”；因此从父节点检查，不能把所有左孩子都累加。
// A left leaf must be its parent's Left and have no children, so test from the parent rather than summing every left child.
// 右子树内部也可能有左叶子，仍须遍历；单独的根不是左叶子。
// The right subtree may contain left leaves too and must be visited; a lone root is not a left leaf.
func sumOfLeftLeavesBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	sum := 0

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Left != nil &&
			node.Left.Left == nil &&
			node.Left.Right == nil {
			sum += node.Left.Val
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return sum
}
