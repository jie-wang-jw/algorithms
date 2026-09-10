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

关键逻辑：为什么从父节点判断？ / Why Check from the Parent?
只看一个节点自己的 Left、Right，能知道它是不是叶子，但不知道它是父节点的左孩子还是右孩子。
站在父节点处检查 node.Left，就同时知道它在左边；再确认这个孩子的 Left 和 Right 都为空，才算左叶子。
因此判断条件是：左孩子存在 AND 左孩子的左孩子为空 AND 左孩子的右孩子为空。
Inspecting a node's own children identifies a leaf but not whether it is a left or right child.
Checking node.Left from its parent establishes its left-child position; checking both of that child's children establishes leaf status.
The condition is: left child exists AND that child's left child is nil AND its right child is nil.

解法一：递归 / Method 1: Recursion
函数负责统计当前树内部所有父子关系中的左叶子，不把当前根本身当左叶子。
当前节点为空时返回 0。若它的左孩子是叶子，直接取这个孩子的值；否则递归左子树找更深的左叶子。
无论当前左孩子是否是叶子，右子树都要递归，因为右子树内部也可能有左叶子；最后把两边结果相加。
The function sums left leaves identified by parent-child relationships within the current tree; it does not count the current root as a left leaf.
Return 0 for nil. If the left child is a leaf, take its value directly; otherwise recurse left to find deeper left leaves.
Always recurse right as well, since that subtree may contain left leaves. Add both contributions.

为什么不直接递归到叶子后返回它的值？那样没有携带左右身份，会把右叶子也算进去。
为什么不会重复？左叶子只有一个父节点，直接计数后不再递归进入它；其他节点只负责往下找，不把自身值加入。
Returning every encountered leaf's value without tracking its side would also count right leaves.
No duplicate counting occurs: each left leaf has one parent and is counted directly without recursion into it; other nodes only search below.

解法二：栈迭代 / Method 2: Iterative DFS
栈保存待检查的父节点。每次出栈，检查它是否拥有左叶子，有则累加；再将非空孩子入栈。
左叶子即使随后出栈也不会重复累加，因为出栈时检查的是它的左孩子，而它没有孩子。
先压右再压左可让左边先处理，但求和与访问顺序无关。
A stack stores nodes whose children need checking. Pop a node, add its left child's value if that child is a leaf, and push existing children.
Even if that left leaf is later popped, it adds nothing: the check examines its left child, and a leaf has none.
Push right before left to visit left first, although traversal order does not affect the sum.

另一种迭代选择：队列 / Another Iterative Option: Queue
也可以改用层序队列，每次从队首取节点，判断左叶子的逻辑完全相同。
不需要分层计数，因此无需 levelSize 或内层循环。它只是遍历顺序不同，不是新的判断规则。
A BFS queue can replace the stack; remove from the front and use the same left-leaf check.
No level grouping is needed, so no levelSize or inner loop is necessary. Only traversal order changes.

易错点 / Pitfalls
- 左孩子不一定是叶子，不能直接累加所有左孩子的值。
  A left child need not be a leaf; do not sum every left child's value.
- 整棵树左侧的所有叶子不等于左叶子；“左”只看它与直接父节点的关系。
  Leaves in the whole tree's left subtree are not necessarily left leaves; only the immediate parent relationship matters.
- 判断叶子用 AND，不是 OR；只有一个孩子的节点仍不是叶子。
  Leaf status requires both children nil (AND), not just one (OR).
- 节点值可能是 0 或负数，按真实值累加，不按正负筛选。
  Values can be zero or negative; sum their actual values without filtering by sign.

时间与空间复杂度 / Time and Space Complexity
n 为节点数，h 为树高，w 为最大层宽。
递归：时间 O(n)，辅助空间 O(h) 来自调用栈，最坏 O(n)。
栈迭代：时间 O(n)，辅助空间 O(h)，最坏 O(n)。队列版本时间 O(n)，辅助空间 O(w)，最坏 O(n)。
都不修改输入树，返回整数只占 O(1) 空间。
For n nodes, height h, and maximum width w:
Recursion takes O(n) time and O(h) call-stack space. Stack DFS takes O(n) time and O(h) auxiliary space.
Queue BFS takes O(n) time and O(w) auxiliary space. All space bounds are at most O(n).
None modifies the tree; the integer result uses O(1) space.

练习 / Practice
本文件只保留中英文题解，不提供实现或函数骨架。先理解父节点判断，再自行实现递归和迭代。
This file contains explanations only, without implementations or skeletons. Understand the parent-side check, then implement recursion and iteration.
*/

// 1. 递归法
func sumOfLeftLeaves(root *TreeNode) int {
	// An empty tree contributes nothing.
	// 空树没有左叶子，返回 0。
	if root == nil {
		return 0
	}

	leftSum := 0

	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			// The left child is a leaf; count its value directly.
			// 左孩子本身是叶子，直接计入它的值。
			leftSum = root.Left.Val
		} else {
			// The left child is not a leaf; search deeper.
			// 左孩子不是叶子，继续寻找它下面的左叶子。
			leftSum = sumOfLeftLeaves(root.Left)
		}
	}

	// The right subtree may also contain left leaves.
	// 右子树内部也可能有左叶子，不能跳过。
	rightSum := sumOfLeftLeaves(root.Right)

	return leftSum + rightSum
}

// 2. 栈迭代
func sumOfLeftLeavesIterative(root *TreeNode) int {
	if root == nil {
		return 0
	}

	stack := []*TreeNode{root}
	sum := 0

	for len(stack) > 0 {
		// Pop a node and inspect its children.
		// 取出一个节点，检查它的孩子。
		last := len(stack) - 1
		node := stack[last]
		stack = stack[:last]

		// Count only a left child that has no children of its own.
		// 只有左孩子存在且没有自己的孩子，才计入答案。
		if node.Left != nil &&
			node.Left.Left == nil &&
			node.Left.Right == nil {
			sum += node.Left.Val
		}

		// Both subtrees may contain left leaves.
		// 左右子树内部都可能存在左叶子，需要继续检查。
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return sum
}

// 3. 队列迭代
func sumOfLeftLeavesBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	sum := 0

	for len(queue) > 0 {
		// Remove the next node from the queue.
		// 从队首取出下一个节点。
		node := queue[0]
		queue = queue[1:]

		// Determine left-leaf status from its parent.
		// 从父节点判断左孩子是否为叶子。
		if node.Left != nil &&
			node.Left.Left == nil &&
			node.Left.Right == nil {
			sum += node.Left.Val
		}

		// Continue searching both subtrees.
		// 继续检查两棵子树。
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return sum
}
