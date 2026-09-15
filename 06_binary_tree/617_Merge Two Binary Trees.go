package _6_binary_tree

/*
617. 合并二叉树 / Merge Two Binary Trees

题目描述 / Problem Description
将两棵二叉树覆盖合并：重叠节点的值相加，只在一棵树出现的节点直接保留。
必须从两棵树的根开始合并。本文件的原地写法修改 root1 的结构并返回它。
Merge two binary trees by adding overlapping node values and keeping nodes that exist in only one tree.
Merging starts at the roots. The in-place versions mutate root1 and return it.

示例 / Example
   1         2              3
  / \       / \            / \
 3   2     1   3    ->    4   5
/           \   \        / \   \
5            4   7      5   4   7

解法一：前序递归，复用 root1（推荐） / Method 1: Preorder Recursion Reusing root1 (Recommended)
同时传入两个节点。任一为空则合并结果就是另一棵子树。
两节点都存在时，把 root2.Val 加到 root1.Val，再把左右孩子分别合并后接回 root1。
前序最好理解：先处理当前重叠节点，再处理左右。中序、后序也能做，只是加值时机不同，这里不重复实现。
Pass both nodes together. If either is nil, the merge result is the other subtree.
When both exist, add root2.Val into root1.Val and attach the merged children onto root1.
Preorder is the article's recommended order. Inorder and postorder only move the addition; they are omitted here.

解法二：前序递归，新建节点 / Method 2: Preorder Recursion Building a New Tree
不修改输入。两节点都存在时新建节点保存和，并递归构造新的左右孩子。
一侧为空时直接返回另一侧引用；调用方不得再修改那棵输入树。
Do not mutate the inputs. When both nodes exist, allocate a new node for the sum and recurse for children.
If one side is nil, return the other subtree by reference; the caller must not mutate that input afterward.

解法三：队列成对迭代 / Method 3: Iterative Pair Queue
把同时存在的一对节点入队。每轮取出一对，把 node2 的值加到 node1。
两边左孩子都在，继续成对入队；node1 缺左孩子而 node2 有，则把 node2 的左子树直接接过去。右孩子同理。
Enqueue pairs of coexisting nodes. Each round adds node2.Val into node1.
If both left children exist, enqueue that pair; if only node2 has a left child, attach that subtree onto node1. Right children follow the same rule.

时间与空间复杂度 / Time and Space Complexity
n 为两棵树节点数的较大者，h 为较小高度，w 为同时处理的最大宽度。
递归两版时间 O(n)，辅助空间 O(h)。队列迭代时间 O(n)，辅助空间 O(w)。
原地写法不复制整棵树；新建树版本额外分配重叠路径上的新节点。
Let n be the larger node count, h a recursion height, and w the maximum paired width.
Both recursive versions take O(n) time and O(h) auxiliary space. The queue version takes O(n) time and O(w) space.
In-place versions do not copy the whole tree; the new-tree version allocates nodes along overlapping paths.
*/

// 1. Preorder recursion reusing root1 (recommended): add into root1, then merge both children onto it.
// 1. 前序递归复用 root1：推荐；把值加到 root1，再把左右孩子的合并结果接回去。
// Time: O(n), Space: O(h).
// 时间复杂度：O(n)，空间复杂度：O(h)。
//
// 步骤与要点 / Steps and notes:
//  1. A missing tree contributes nothing; keep the other subtree.
//     缺了一棵树，合并结果就是另一棵。
//  2. Both nodes exist, so this position's value is their sum.
//     两个节点重叠，当前值应写成二者之和。
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

// 2. Preorder recursion building a new tree: allocate a sum node when both sides exist.
// 2. 前序递归新建树：两侧都存在时新建求和节点，不修改输入。
// Time: O(n), Space: O(h) plus new overlapping nodes.
// 时间复杂度：O(n)，空间复杂度：O(h)，并额外分配重叠路径上的新节点。
func mergeTreesNew(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	root := &TreeNode{Val: root1.Val + root2.Val}
	root.Left = mergeTreesNew(root1.Left, root2.Left)
	root.Right = mergeTreesNew(root1.Right, root2.Right)
	return root
}

// 3. Iterative pair queue: dequeue two coexisting nodes, add values, and attach or enqueue children.
// 3. 队列成对迭代：取出一对同时存在的节点，加值后再按孩子是否存在入队或直接拼接。
// Time: O(n), Space: O(w).
// 时间复杂度：O(n)，空间复杂度：O(w)。
//
// 步骤与要点 / Steps and notes:
//  1. Both left children exist, so they still need to be merged.
//     左右两棵树的左孩子都在，这对节点还要继续合并。
//  2. Only the second tree has this child; graft that subtree onto the first tree.
//     只有第二棵树有这个孩子，直接接到第一棵树上。
func mergeTreesIterative(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	queue := []*TreeNode{root1, root2}
	for len(queue) > 0 {
		node1 := queue[0]
		node2 := queue[1]
		queue = queue[2:]
		node1.Val += node2.Val

		if node1.Left != nil && node2.Left != nil {
			queue = append(queue, node1.Left, node2.Left)
		}
		if node1.Right != nil && node2.Right != nil {
			queue = append(queue, node1.Right, node2.Right)
		}

		if node1.Left == nil && node2.Left != nil {
			node1.Left = node2.Left
		}
		if node1.Right == nil && node2.Right != nil {
			node1.Right = node2.Right
		}
	}

	return root1
}
