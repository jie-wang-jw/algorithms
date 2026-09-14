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

关键逻辑 / Key Logic
插入 = 按 BST 搜索，走到本应放置 val 的空链接，新建节点挂上。
不会替换已有节点；题目保证 val 不与现有值冲突。
插入后树可能更不平衡，本题不要求旋转维护平衡。
Insertion is a BST search that stops at the empty link where val belongs, then allocates a node there.
Existing nodes are never replaced; val is guaranteed distinct.
The tree may become more skewed; this problem does not require rebalancing rotations.

解法一：递归 / Method 1: Recursion
空位置就是插入点，新建节点返回给父节点接上。
当前值大于 val 则插入左子树，否则插入右子树，并把返回的子树根赋回孩子指针。
An empty child is the insertion site: allocate a node and return it to the parent.
Insert left when the current value is greater than val, otherwise right, and assign the returned subtree.

解法二：迭代（推荐） / Method 2: Iteration (Recommended)
先记住父节点，再按大小向下走。走到空时把新节点接到父节点对应的一侧。
空树直接返回新节点。
Keep walking left or right. When the next child is nil, attach the new node there.
An empty tree returns the new node itself.

时间与空间复杂度 / Time and Space Complexity
h 为树高。两版时间 O(h)。递归辅助空间 O(h)；迭代 O(1)。
For height h, both take O(h) time. Recursion uses O(h) space; iteration uses O(1).
*/

// 1. Recursion: replace a nil child with the new node, otherwise recurse into the ordered side and assign the result.
// 1. 递归：空孩子处新建节点，否则进入有序的一侧并把返回的子树接回去。
// Time: O(h), Space: O(h).
// 时间复杂度：O(h)，空间复杂度：O(h)。
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	// Empty link: this is the insertion site.
	// 空链接：这里就是插入点。
	if root == nil {
		return &TreeNode{Val: val}
	}

	// Attach on the side that preserves BST order, then return the unchanged root.
	// 挂到能保持 BST 有序的一侧，再返回（子树根仍是）当前根。
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
func insertIntoBSTIterative(root *TreeNode, val int) *TreeNode {
	node := &TreeNode{Val: val}

	// Empty tree: the new node is the root.
	// 空树：新节点就是根。
	if root == nil {
		return node
	}

	cur := root
	for cur != nil {
		if cur.Val > val {
			// Next left is empty: attach here; otherwise keep walking left.
			// 左孩子为空则挂上；否则继续向左。
			if cur.Left == nil {
				cur.Left = node
				return root
			}
			cur = cur.Left
		} else {
			// Symmetric case on the right spine.
			// 右链上的对称情况。
			if cur.Right == nil {
				cur.Right = node
				return root
			}
			cur = cur.Right
		}
	}
	return root
}
