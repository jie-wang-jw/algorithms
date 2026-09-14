package _6_binary_tree

/*
538. 把二叉搜索树转换为累加树 / Convert BST to Greater Tree

题目描述 / Problem Description
把 BST 的每个节点改成：原值加上树中所有大于它的节点值之和。返回同一棵树的根。
本题与 1038. Binary Search Tree to Greater Sum Tree 相同。
Replace every BST node with its original value plus the sum of all greater node values, and return the same root.
This is the same problem as 1038.

示例 / Examples
转换前 / Before:              转换后 / After:
      4                            30
     / \                          /  \
    1   6                       36    21
   / \ / \                     / \   /  \
  0  2 5  7                   36 35 26  15
      \    \                       \      \
       3    8                       33     8

反中序（右→根→左）从大到小：8,7,6,5,4,3,2,1,0；累加和依次覆盖每个节点。
Reverse inorder (right→root→left) visits 8,7,6,5,4,3,2,1,0 largest-first; the running sum overwrites each node.

关键逻辑 / Key Logic
BST 的反中序（右、根、左）按从大到小访问。维护累加和 sum：先走右子树把更大的值加完，
再把当前值加进 sum，并用 sum 覆盖当前节点，最后走左子树。这样每个节点都加上了所有比它大的值。
Reverse inorder (right-root-left) visits values from largest to smallest. Keep a running sum:
finish the right subtree first so every greater value is included, add the current value into sum,
overwrite the node, then visit the left subtree.

解法一：反中序递归（推荐） / Method 1: Reverse-Inorder Recursion (Recommended)
用闭包保存 sum，按右、根、左递归并原地改值。
A closure holds sum; recurse right-root-left and overwrite values in place.

解法二：反中序迭代 / Method 2: Iterative Reverse Inorder
用栈先一路向右压栈，弹出时累加并改值，再转向左孩子。
Push along the right spine, add and overwrite on pop, then move to the left child.

时间与空间复杂度 / Time and Space Complexity
n 为节点数，h 为树高。两版时间 O(n)，辅助空间 O(h)。原地修改，不新建树。
For n nodes and height h, both take O(n) time and O(h) auxiliary space. The tree is updated in place.
*/

// 1. Reverse-inorder recursion (recommended): add values from large to small so each node receives every greater value.
// 1. 反中序递归：推荐；从大到小累加，当前节点就能加上所有比它大的值。
// Time: O(n), Space: O(h).
// 时间复杂度：O(n)，空间复杂度：O(h)。
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		// Larger values first: finish the right subtree before touching this node.
		// 先处理更大的值：右子树走完再改当前节点。
		traverse(node.Right)
		sum += node.Val
		node.Val = sum
		traverse(node.Left)
	}
	traverse(root)
	return root
}

// 2. Iterative reverse inorder: push the right spine, add on pop, then visit the left child.
// 2. 反中序迭代：先沿右链压栈，弹出时累加改值，再处理左孩子。
// Time: O(n), Space: O(h).
// 时间复杂度：O(n)，空间复杂度：O(h)。
func convertBSTIterative(root *TreeNode) *TreeNode {
	sum := 0
	stack := []*TreeNode{}
	cur := root
	for cur != nil || len(stack) > 0 {
		// Descend right, mirroring the usual left-descending inorder stack.
		// 向右深入，对称于普通中序向左压栈。
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Right
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Pop order is largest-to-smallest; overwrite with the running greater-sum.
		// 弹出顺序从大到小；用累加和覆盖当前值。
		sum += cur.Val
		cur.Val = sum
		cur = cur.Left
	}
	return root
}
