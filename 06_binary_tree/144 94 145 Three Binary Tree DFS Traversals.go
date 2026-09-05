package _6_binary_tree

/*
144. 二叉树的前序遍历 / Binary Tree Preorder Traversal
题目描述 / Problem Description
给定二叉树根节点 root，返回节点值的前序遍历结果。
Given the root of a binary tree, return its preorder traversal.
前序遍历顺序：
Preorder traversal order:
根 → 左 → 右
Root → Left → Right
解题思路 / Solution Approach
使用栈保存等待访问的节点。
Use a stack to store nodes waiting to be visited.
每次弹出一个节点后：
After popping a node:
1. 访问当前节点；
   Visit the current node;
2. 先将右子节点压栈；
   Push the right child first;
3. 再将左子节点压栈。
   Push the left child second.
因为栈是后进先出，所以左子节点会先被处理。
Because the stack is LIFO, the left child is processed first.

时间与空间复杂度 / Time and Space Complexity
n 为节点数，h 为树高。前序、中序、后序的递归和迭代共六种实现，时间均为 O(n)；后序迭代额外反转 O(n)，总阶不变。辅助空间均为 O(h)，包括显式栈或递归调用栈；平衡树 h=O(log n)，最坏 h=n。显式栈可保存路径上的待访问兄弟节点，并非只保存一条路径本身。每种解法返回结果占 O(n)，包含结果的总空间 O(n)。
For n nodes and height h, all six recursive/iterative preorder, inorder, and postorder implementations take O(n) time; reversing iterative postorder adds O(n) without changing the bound. Auxiliary space O(h) includes explicit or recursive stacks; h=O(log n) for a balanced tree and up to n otherwise. Explicit stacks may hold pending siblings along a path. Each output and total space take O(n).
*/

/*
题目描述 / Problem Description

返回二叉树的前序遍历结果。
Return the preorder traversal of a binary tree.

解题思路 / Solution Approach

使用栈按照“根、左、右”的顺序访问节点。
Use a stack to visit nodes in root-left-right order.

因为栈是后进先出，所以先压入右子节点，再压入左子节点。
Because the stack is LIFO, push the right child before the left child.
*/

func preorderTraversal(root *TreeNode) []int {
	// An empty tree has no traversal result.
	// 空树没有遍历结果。
	if root == nil {
		return []int{}
	}

	result := make([]int, 0)
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		// Pop the top node.
		// 弹出栈顶节点。
		last := len(stack) - 1
		node := stack[last]
		stack = stack[:last]

		// Visit the root before its children.
		// 在左右子节点之前访问根节点。
		result = append(result, node.Val)

		// Push right first so it is processed after the left child.
		// 先压入右子节点，使它在左子节点之后处理。
		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		// Push left last so it is processed next.
		// 后压入左子节点，使它下一步被处理。
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return result
}

func preorderTraversalRecursive(root *TreeNode) []int {
	result := make([]int, 0)

	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		// Root → Left → Right
		// 根 → 左 → 右
		result = append(result, node.Val)
		traverse(node.Left)
		traverse(node.Right)
	}

	traverse(root)
	return result
}

/*
94. 二叉树的中序遍历 / Binary Tree Inorder Traversal

题目描述 / Problem Description
给定二叉树根节点 root，返回按照“左、根、右”顺序访问的节点值。
Given the root of a binary tree, return its node values in left-root-right order.

解题思路 / Solution Approach
迭代法沿左侧路径不断压栈；无法继续向左时，弹出并访问栈顶节点，再进入其右子树。
The iterative solution pushes nodes along the left path. When no further left node exists,
pop and visit the top node, then enter its right subtree.

// Ordinary tree:
	// 普通二叉树：
	//
	//          1
	//        /   \
	//       2     3
	//      / \     \
	//     4   5     6
*/

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	current := root

	// Continue while an unexplored node or a saved ancestor remains.
	// 只要还有未探索节点或栈中还有祖先节点，就继续处理。
	for current != nil || len(stack) > 0 {
		// Reach the leftmost node and save every ancestor on the path.
		// 不断向左移动，并保存路径上的每个祖先节点。
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		last := len(stack) - 1
		current = stack[last]
		stack = stack[:last]

		// Visit the root after its left subtree.
		// 左子树处理完成后访问根节点。
		result = append(result, current.Val)
		current = current.Right
	}

	return result
}

func inorderTraversalRecursive(root *TreeNode) []int {
	result := make([]int, 0)

	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		// Left -> Root -> Right
		// 左 -> 根 -> 右
		traverse(node.Left)
		result = append(result, node.Val)
		traverse(node.Right)
	}

	traverse(root)
	return result
}

/*
145. 二叉树的后序遍历 / Binary Tree Postorder Traversal

题目描述 / Problem Description
给定二叉树根节点 root，返回按照“左、右、根”顺序访问的节点值。
Given the root of a binary tree, return its node values in left-right-root order.

解题思路 / Solution Approach
迭代法先通过栈得到“根、右、左”的顺序，再反转结果得到“左、右、根”。
The iterative solution first produces root-right-left order with a stack, then reverses it to obtain left-right-root order.
*/

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := make([]int, 0)
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		last := len(stack) - 1
		node := stack[last]
		stack = stack[:last]

		// Temporarily record nodes in root-right-left order.
		// 暂时按照“根、右、左”的顺序记录节点。
		result = append(result, node.Val)

		// Push left before right so right is processed first.
		// 先压入左子节点、再压入右子节点，使右子节点先处理。
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	// Reverse root-right-left into left-right-root.
	// 将“根、右、左”反转为“左、右、根”。
	for left, right := 0, len(result)-1; left < right; left, right = left+1, right-1 {
		result[left], result[right] = result[right], result[left]
	}
	/*	// 项目代码更简洁
		// More concise for project code.
		slices.Reverse(result)
	*/

	return result
}

func postorderTraversalRecursive(root *TreeNode) []int {
	result := make([]int, 0)

	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		// Left -> Right -> Root
		// 左 -> 右 -> 根
		traverse(node.Left)
		traverse(node.Right)
		result = append(result, node.Val)
	}

	traverse(root)
	return result
}
