package _6_binary_tree

/*
701. 二叉搜索树中的插入操作 / Insert into a Binary Search Tree

题目描述 / Problem Description
给定 BST 根节点和要插入的值 val（保证与现有值不同），插入后仍保持 BST，返回新树的根。
插入位置不唯一；本文件把新节点接到搜索路径尽头的空孩子上，不刻意再平衡。
Given a BST root and a distinct value val, insert it so the tree remains a BST and return the root.
The insertion site is not unique; these implementations attach the new node at the empty child where the search ends.

示例 / Examples
插入前 / Before, val=5:     插入后 / After:
      4                          4
     / \                        / \
    2   7                      2   7
   / \                        / \ /
  1   3                      1  3 5

搜索 5：4 < 5 → 右到 7；7 > 5 → 左为空，把 5 接到 7.Left。
Search for 5: 4 < 5 → right to 7; 7 > 5 → left is empty, attach 5 as 7.Left.
复杂度记号：n 为节点数，h 为树高，w 为最大层宽；辅助空间不含返回结果。
Notation: n nodes, height h, maximum width w; auxiliary space excludes returned results.
*/

// 1. Recursion: replace a nil child with the new node, otherwise recurse into the ordered side and assign the result.
// 1. 递归：空孩子处新建节点，否则进入有序的一侧并把返回的子树接回去。
// Time: O(h), Space: O(h).
// 时间复杂度：O(h)，空间复杂度：O(h)。
//
// 按 BST 比较沿唯一搜索路径找空位，只新增叶子，不重排原节点；递归版必须接回子调用返回的新根。
// Follow the BST search path to an empty slot and add one leaf; recursive calls must reconnect the returned subtree root.
// 按题意待插入值原先不存在；当前代码未单独处理重复值。
// Assume the inserted value is absent as required by the problem; duplicate values are not separately handled.
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

// 2. Iteration (recommended): walk to the empty insertion site and attach the new node to its parent.
// 2. 迭代：推荐；走到空插入点，把新节点接到父节点对应的一侧。
// Time: O(h), Space: O(1).
// 时间复杂度：O(h)，空间复杂度：O(1)。
//
// 按 BST 比较沿唯一搜索路径找空位，只新增叶子，不重排原节点；递归版必须接回子调用返回的新根。
// Follow the BST search path to an empty slot and add one leaf; recursive calls must reconnect the returned subtree root.
// 按题意待插入值原先不存在；当前代码未单独处理重复值。
// Assume the inserted value is absent as required by the problem; duplicate values are not separately handled.
func insertIntoBSTIterative(root *TreeNode, val int) *TreeNode {
	node := &TreeNode{Val: val}

	if root == nil {
		return node
	}

	cur := root
	for cur != nil {
		if cur.Val > val {
			if cur.Left == nil {
				cur.Left = node
				return root
			}
			cur = cur.Left
		} else {
			if cur.Right == nil {
				cur.Right = node
				return root
			}
			cur = cur.Right
		}
	}
	return root
}
