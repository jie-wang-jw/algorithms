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
复杂度记号：n 为节点数，h 为树高，w 为最大层宽；辅助空间不含返回结果。
Notation: n nodes, height h, maximum width w; auxiliary space excludes returned results.
*/

// 1. Reverse-inorder recursion (recommended): add values from large to small so each node receives every greater value.
// 1. 反中序递归：推荐；从大到小累加，当前节点就能加上所有比它大的值。
// Time: O(n), Space: O(h).
// 时间复杂度：O(n)，空间复杂度：O(h)。
//
// 按右→根→左访问，遇到当前节点时所有更大值已加入 sum；加上当前原值后赋回，就是所需的“大于等于之和”。
// Visit right→root→left; sum already contains larger original values, so adding the current value yields its greater-or-equal sum.
// 要求原树是值互异的 BST；原地修改 Val，累计和须在 int 范围内。
// Require a unique-valued BST; values are mutated in place and cumulative sums must fit in int.
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

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
//
// 按右→根→左访问，遇到当前节点时所有更大值已加入 sum；加上当前原值后赋回，就是所需的“大于等于之和”。
// Visit right→root→left; sum already contains larger original values, so adding the current value yields its greater-or-equal sum.
// 要求原树是值互异的 BST；原地修改 Val，累计和须在 int 范围内。
// Require a unique-valued BST; values are mutated in place and cumulative sums must fit in int.
func convertBSTIterative(root *TreeNode) *TreeNode {
	sum := 0
	stack := []*TreeNode{}
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Right
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		sum += cur.Val
		cur.Val = sum
		cur = cur.Left
	}
	return root
}
