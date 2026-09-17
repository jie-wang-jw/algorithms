package _6_binary_tree

/*
450. 删除二叉搜索树中的节点 / Delete Node in a BST

题目描述 / Problem Description
在 BST 中删除值为 key 的节点（若存在），保持 BST 性质，返回新根。
删除后可用中序后继或前驱顶替；本文件采用 Carl 文章写法：把左子树接到右子树最左节点上，再返回右孩子。
Delete the node valued key from a BST if it exists, keep the BST property, and return the new root.
Either the inorder successor or predecessor may replace a two-child node. This file follows Carl's article:
attach the left subtree onto the leftmost node of the right subtree, then return the right child.

示例 / Examples
删除前 / Before (delete 3):     删除后 / After:
      5                              5
     / \                            / \
    3   6                          4   6
   / \   \                        /     \
  2   4   7                      2       7

3 有两个孩子。右子树最左节点是中序后继 4；把 3 的左子树（2）接到 4 的左边，再用右子树根 4 顶替 3。
3 has two children. The leftmost node of its right subtree is inorder successor 4;
attach 3's left subtree (2) under 4, then replace 3 with the right-subtree root 4.

三种删除情形的 ASCII / Three deletion cases

情形一：叶子 / Case 1: leaf
  删 2：左右皆空 → 返回 nil，父节点对应孩子变空。
  Delete 2: both children nil → return nil; the parent's link becomes empty.
      5                    5
     / \                  / \
    3   6       →        3   6
   / \   \                \   \
  2   4   7                4   7

情形二：只有一个孩子 / Case 2: one child
  删 6：只有右孩子 7 → 直接返回 7，顶替 6 的位置。
  Delete 6: only right child 7 → return 7 in place of 6.
      5                    5
     / \                  / \
    3   6       →        3   7
   / \   \              / \
  2   4   7            2   4

情形三：两个孩子（Carl 接法） / Case 3: two children (Carl's attach)
  删 3：右子树为 4（叶子），最左节点仍是 4；successor.Left = 原左子树 2，返回右子树根 4。
  Delete 3: right subtree is leaf 4, still the leftmost; set successor.Left to old left (2), return 4.
      5                    5
     / \                  / \
    3   6       →        4   6
   / \   \              /     \
  2   4   7            2       7

为什么中序后继可用？ / Why the inorder successor works
中序后继是右子树中最小的值，严格大于被删节点，且小于右子树其余节点。
它原来没有左孩子（否则还能更左），因此可以把被删节点的整棵左子树安全挂到它的 Left 上：
左子树所有值 < 被删值 < 后继，挂上后仍满足 BST。整棵右子树根再顶替被删节点，右子树内部结构不变。
等价效果：用后继“吸收”左子树后，右子树整体上移。不必单独拷贝后继值再删后继节点。
The inorder successor is the smallest value in the right subtree: strictly greater than the deleted node,
and smaller than every other right-subtree value. It has no left child (or a further left would exist),
so the deleted node's entire left subtree can hang from its Left safely:
every left value < deleted value < successor, so the BST order holds. Returning the right-subtree root
lifts that whole side into the deleted node's place without reshaping the right interior.
Effectively the successor absorbs the left subtree, then the right subtree moves up —
no separate “copy successor value, then delete the successor node” step.
复杂度记号：n 为节点数，h 为树高，w 为最大层宽；辅助空间不含返回结果。
Notation: n nodes, height h, maximum width w; auxiliary space excludes returned results.
*/

// 1. Recursion (recommended): search by BST order, then replace the node with nil, one child, or the rewired right subtree.
// 1. 递归：推荐；按 BST 找到目标后，用空、单孩子或改接后的右子树替换它。
// Time: O(h), Space: O(h).
// 时间复杂度：O(h)，空间复杂度：O(h)。
//
// 按 BST 大小关系只递归可能含 key 的一侧；递归返回删除后的子树根，必须重新接到 root.Left/Right。
// Use BST ordering to recurse into one side; reconnect the returned subtree root because deletion may replace that root.
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	}

	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	}

	return deleteRoot(root)
}

// 2. Iteration: find the target and parent, then apply the same replacement rule.
// 2. 迭代：先找到目标及其父节点，再使用相同的替换规则。
// Time: O(h), Space: O(1).
// 时间复杂度：O(h)，空间复杂度：O(1)。
//
// 搜索时记录 parent；删除后把替代子树接回原父节点对应边。若 parent==nil，删除的是整树根，直接返回替代根。
// Track parent during search and reconnect the replacement; if parent is nil, the deleted node is the root and its replacement is returned.
func deleteNodeIterative(root *TreeNode, key int) *TreeNode {
	var parent *TreeNode
	cur := root

	for cur != nil && cur.Val != key {
		parent = cur
		if key < cur.Val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}

	if cur == nil {
		return root
	}

	replacement := deleteRoot(cur)

	if parent == nil {
		return replacement
	}

	if parent.Left == cur {
		parent.Left = replacement
	} else {
		parent.Right = replacement
	}
	return root
}

// Replacement for a found node: return nil, the only child, or the right subtree after attaching the left subtree to its successor.
// 已找到节点的替换结果：叶子返回 nil，单孩子直接顶上，双孩子把左子树接到右子树最左后再返回右子树。
// Time: O(h), Space: O(1).
// 时间复杂度：O(h)，空间复杂度：O(1)。
//
// 要求 root 非空；少于两个孩子时直接返回另一孩子。两孩子都在时，把左树接到右树最左节点的空 Left，再返回右树根。
// Require nonnil root; return the existing child if one is absent. Otherwise attach the left subtree to the right subtree's leftmost node.
// 左树所有值都小于原根，也小于右树所有值，因此这个接法仍满足 BST，且不丢失任何其他节点。
// All left values are smaller than the old root and every right value, preserving BST order without losing other nodes.
func deleteRoot(root *TreeNode) *TreeNode {
	if root.Left == nil {
		return root.Right
	}

	if root.Right == nil {
		return root.Left
	}

	successor := root.Right
	for successor.Left != nil {
		successor = successor.Left
	}

	successor.Left = root.Left
	return root.Right
}
