package _6_binary_tree

/*
669. 修剪二叉搜索树 / Trim a Binary Search Tree

题目描述 / Problem Description
给定 BST 根节点和闭区间 [low, high]，删除所有值不在区间内的节点，保持 BST 性质，返回新根。
Trim a BST so that every remaining node value lies in the closed interval [low, high], and return the new root.

示例 / Examples
修剪前 / Before, [low,high]=[1,3]:     修剪后 / After:
      3                                    3
     / \                                  /
    0   4                                2
     \                                  /
      2                                1
     /
    1

0 < low，整棵以 0 为根的左支里小于 1 的部分被丢掉，只保留落在 [1,3] 内的 1、2、3；4 > high 整侧丢弃。
0 is below low, so values under it that stay below 1 are dropped; only 1, 2, 3 remain in [1,3]; 4 exceeds high and its side is discarded.

另一例 / Another example, [low,high]=[1,2]:
      3                                    2
     / \                                  /
    0   4                                1
     \
      2
     /
    1
根 3 > high，答案只可能在左子树；继续修剪后得到以 2 为根的树。
Root 3 exceeds high, so the answer can only come from the left; further trimming yields the tree rooted at 2.

关键逻辑 / Key Logic（三种递归情形）
利用 BST：左子树所有值 < 根 < 右子树所有值。
1. root.Val < low
   根和整棵左子树都小于 low，全部不合法；合法节点只可能在右子树。
   → 直接返回 trim(root.Right)，根本身也被丢弃。
2. root.Val > high
   根和整棵右子树都大于 high，全部不合法；合法节点只可能在左子树。
   → 直接返回 trim(root.Left)。
3. low <= root.Val <= high
   根保留；左右仍可能含越界节点，分别修剪后接回。
   → root.Left = trim(root.Left); root.Right = trim(root.Right); return root

ASCII 示意 root < low 时整侧丢弃 / When root < low, drop the whole left side:
      0 (< low)                 合法区只在右边 / keep searching right only
     / \                   →
   ...  2
       /
      1

Use BST order: every left value < root < every right value.
1. root.Val < low — root and its entire left subtree are too small; only the right can survive → return trim(right).
2. root.Val > high — root and its entire right subtree are too large; only the left can survive → return trim(left).
3. In range — keep the root and trim both children, then reattach.

解法一：递归（推荐） / Method 1: Recursion (Recommended)
按上面三种情况返回右子树修剪结果、左子树修剪结果，或接回修剪后的左右孩子。
一棵越界子树在比较根值时整侧被跳过，不必逐个访问其中每个节点的“删除”，但最坏仍可能访问全部节点。
Return the trimmed right subtree, the trimmed left subtree, or the current node with trimmed children.
An out-of-range side is skipped wholesale when the root comparison fires; the worst case still visits every node.

解法二：迭代先定位新根，再修剪两侧 / Method 2: Iteration
先把根移到区间内（或空）：根太小就改走右孩子，太大就改走左孩子。
新根确定后，沿左链删掉仍小于 low 的节点（用其右孩子顶替），沿右链删掉仍大于 high 的节点（用其左孩子顶替）。
Move the root into range (or to nil): go right when too small, left when too large.
Then walk the left spine dropping values below low (replace with that node's right child),
and the right spine dropping values above high (replace with that node's left child).

时间与空间复杂度 / Time and Space Complexity
n 为节点数，h 为树高。两版时间 O(n)。递归辅助空间 O(h)；迭代 O(1)。
For n nodes and height h, both take O(n) time. Recursion uses O(h) space; iteration uses O(1).
*/

// 1. Recursion (recommended): drop the whole left or right subtree when the root is outside [low, high].
// 1. 递归：推荐；根落在区间外时整侧子树都可以丢掉，只修剪可能合法的一侧。
// Time: O(n), Space: O(h).
// 时间复杂度：O(n)，空间复杂度：O(h)。
//
// 步骤与要点 / Steps and notes:
//  1. Empty subtree contributes nothing.
//     空子树无需修剪。
//  2. Root and its entire left subtree are below low; only the right side can remain.
//     根和整棵左子树都小于 low，合法节点只可能在右子树。
//  3. Root and its entire right subtree are above high; only the left side can remain.
//     根和整棵右子树都大于 high，合法节点只可能在左子树。
//  4. Root is in range; trim both children and reattach the results.
//     根在区间内：分别修剪左右孩子并接回。
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < low {
		return trimBST(root.Right, low, high)
	}

	if root.Val > high {
		return trimBST(root.Left, low, high)
	}

	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}

// 2. Iteration: first move the root into range, then prune illegal nodes on each remaining spine.
// 2. 迭代：先把根移进区间，再分别删掉左右链上越界的节点。
// Time: O(n), Space: O(1).
// 时间复杂度：O(n)，空间复杂度：O(1)。
//
// 步骤与要点 / Steps and notes:
//  1. Slide the root into [low, high], or to nil if the whole tree is out of range.
//     先把根移进 [low, high]；整棵树都越界则变为 nil。
//  2. Left spine: any child below low is replaced by that child's right subtree.
//     左链：小于 low 的左孩子用它的右子树顶替（左子树只会更小）。
//  3. Right spine: any child above high is replaced by that child's left subtree.
//     右链：大于 high 的右孩子用它的左子树顶替（右子树只会更大）。
func trimBSTIterative(root *TreeNode, low int, high int) *TreeNode {
	for root != nil && (root.Val < low || root.Val > high) {
		if root.Val < low {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	if root == nil {
		return nil
	}

	cur := root
	for cur.Left != nil {
		if cur.Left.Val < low {
			cur.Left = cur.Left.Right
		} else {
			cur = cur.Left
		}
	}

	cur = root
	for cur.Right != nil {
		if cur.Right.Val > high {
			cur.Right = cur.Right.Left
		} else {
			cur = cur.Right
		}
	}
	return root
}
