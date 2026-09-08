package _6_binary_tree

/*
101. 对称二叉树 / Symmetric Tree
题目链接 / Problem link: https://leetcode.cn/problems/symmetric-tree/

题目描述 / Problem Description
给定二叉树的根节点 root，判断它是否关于中心轴左右对称。
对称要求镜像位置的节点值相同，且结构也互为镜像。
Given the root of a binary tree, determine whether it is symmetric about its center.
Both node values and structure must match at mirrored positions.

示例 / Examples
对称 / Symmetric:
        1
       / \
      2   2
     / \ / \
    3  4 4  3

不对称 / Not symmetric:
        1
       / \
      2   2
       \   \
        3   3
第二棵树虽然每层的值看似对称，但两个 3 都是右孩子，位置并不互为镜像。
In the second tree, both 3 nodes are right children. Their values match, but their positions do not mirror each other.

解题思路：递归比较镜像位置 / Approach: Recursively Compare Mirrored Positions
不要分别判断左子树和右子树自己是否对称；需要判断它们彼此是否互为镜像。
定义 mirror(left, right)：以这两个节点为根的子树是否互为镜像。
Do not check whether each subtree is independently symmetric. Check whether they mirror each other.
Define mirror(left, right) to mean that the two subtrees are mirror images.

1. 两个节点都为空：对应位置都没有节点，匹配成功。
   Both nil: both mirrored positions are empty, so they match.
2. 只有一个为空：一侧有节点、一侧没有，结构不匹配。
   Only one nil: one side has a node and the other does not, so their structures differ.
3. 两个都存在但值不同：镜像位置的值不匹配。
   Both exist but have different values: mirrored values differ.
4. 值相同：继续比较外侧一对和内侧一对，并要求两对都成立。
   Equal values: compare the outer pair and the inner pair; both must match.

关键逻辑：为什么交叉比较？ / Why Compare Across Sides?
镜像会把左方向变成右方向，因此左子树的左孩子应对应右子树的右孩子（外侧），
左子树的右孩子应对应右子树的左孩子（内侧）。同方向比较判断的是相同结构，不是镜像结构。
A mirror reverses left and right: the left subtree's left child matches the right subtree's right child (outer pair),
and the left subtree's right child matches the right subtree's left child (inner pair).
Comparing the same directions tests identical structure, not mirrored structure.

外侧 / Outer pair: left.Left  <-> right.Right
内侧 / Inner pair: left.Right <-> right.Left

当前两个值相同，加上外侧子树互为镜像、内侧子树互为镜像，就覆盖了这两棵树的全部结构。
所以两组递归结果用 AND 合并，任何一组失败都不能称为对称。
Equal current values plus mirrored outer and inner subtrees cover the whole pair of trees.
Combine both recursive results with AND; either failure breaks symmetry.

从哪里开始？ / Where to Start?
整棵树的根位于中心轴上，不需要找另一个根与它比较；从 root.Left 和 root.Right 开始比较即可。
若 root 为空，返回 true；单节点树的两个孩子都为空，也会返回 true。
The root lies on the axis, so compare root.Left with root.Right.
An empty tree is symmetric; a single-node tree also passes because both children are nil.

易错点 / Pitfalls
- 先处理空指针，再读取 Val、Left、Right，避免空指针访问。
  Handle nil pointers before reading Val, Left, or Right.
- 只比较左右根节点值不够，还要检查下面的结构和节点值。
  Equal child-root values alone do not establish symmetry of their descendants.
- 不要像 226 题一样交换节点。本题只判断，不需要修改输入树。
  Unlike problem 226, this is a read-only check; do not swap nodes.
- 不含空位置的层序值回文不能证明对称，第二个示例就是反例。
  Palindromic level values without nil positions do not prove symmetry, as the second example shows.

时间与空间复杂度 / Time and Space Complexity
n 为节点数，h 为树高。最坏时间 O(n)，每个节点最多参与一次对应比较；发现不匹配可提前返回。
辅助空间 O(h)，来自递归调用栈；平衡树为 O(log n)，用一般树高上界可记最坏 O(n)。
返回布尔值，不新建树，也不修改节点。
For n nodes and height h, worst-case time is O(n), with early return on a mismatch.
Auxiliary call-stack space is O(h): O(log n) for a balanced tree, bounded by O(n) in general.
The result is a boolean; no tree is constructed or modified.

练习 / Practice
请在下方自行实现 isSymmetric 和镜像比较函数；本文件暂不提供实现或代码骨架。
Implement isSymmetric and its mirror-comparison helper below. No implementation or skeleton is provided yet.
*/

func isSymmetric(root *TreeNode) bool {
	// An empty tree is symmetric.
	// 空树是对称的。
	if root == nil {
		return true
	}

	// Check whether the two subtrees mirror each other.
	// 判断左右子树彼此是否互为镜像。
	return isMirror(root.Left, root.Right)
}

func isMirror(left, right *TreeNode) bool {
	// Both mirrored positions are empty, so they match.
	// 两个对应位置都没有节点，匹配成功。
	if left == nil && right == nil {
		return true
	}

	// After excluding both nil, either nil means a structural mismatch.
	// 已排除同时为空；此时若有一个为空，就说明结构不匹配。
	if left == nil || right == nil {
		return false
	}

	// Both nodes exist; their values must match.
	// 两个节点都存在，它们的值必须相同。
	if left.Val != right.Val {
		return false
	}

	// Mirror reflection reverses directions: left matches right.
	// 镜像会反转方向，因此左孩子要与另一侧的右孩子比较。
	return isMirror(left.Left, right.Right) &&
		isMirror(left.Right, right.Left)
}

func isSymmetricIterative(root *TreeNode) bool {
	// Consecutive nodes form pairs of mirrored positions.
	// 队列中相邻的两个节点组成一对镜像位置。
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root.Left, root.Right)
	}

	for len(queue) > 0 {
		// The queue length is always even, so take two nodes together.
		// 队列长度始终为偶数，每次取出两个节点。
		left, right := queue[0], queue[1]
		queue = queue[2:]

		// Only this pair is confirmed; other pairs still need checking.
		// 这里只确认当前一对匹配，其他对仍需继续检查。
		if left == nil && right == nil {
			continue
		}

		// Short-circuit evaluation prevents nil pointer access.
		// 短路求值保证先排除空指针，再比较节点值。
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}

		// Enqueue the outer pair, followed by the inner pair.
		// 将外侧一对、内侧一对依次加入队列。
		queue = append(queue,
			left.Left, right.Right,
			left.Right, right.Left,
		)
	}

	// All pairs matched; an empty tree also reaches this return.
	// 所有对应位置都匹配；空树也会直接走到这里。
	return true
}
