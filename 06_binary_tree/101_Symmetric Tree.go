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
复杂度记号：n 为节点数，h 为树高，w 为最大层宽；辅助空间不含返回结果。
Notation: n nodes, height h, maximum width w; auxiliary space excludes returned results.
*/

// 1. Recursively compare mirrored positions (recommended): start from the two children and cross outer with inner.
// 1. 递归比较镜像位置：推荐；从根的两个孩子出发，交叉比较外侧与内侧。
// Time: O(n), Space: O(h) for the recursion stack.
// 时间复杂度：O(n)，空间复杂度：O(h)，由递归栈产生。
//
// 比较的是镜像位置：左的左对右的右，左的右对右的左；不能只比较同层数值而忽略空节点位置。
// Compare mirrored positions: left-left with right-right and left-right with right-left; values alone do not capture structure.
// 两个都空为真，只有一个空或值不同为假；队列/栈版始终成对存取镜像节点，空根视为对称。
// Two nil nodes match; one nil or unequal values fail. Iterative versions process pairs; an empty tree is symmetric.
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isMirror(root.Left, root.Right)
}

// Compare mirrored positions by pairing outer children and inner children, not the same directions.
// 交叉比较镜像位置：外侧 left.Left 对 right.Right，内侧 left.Right 对 right.Left，而不是同方向比较。
// Time: O(n), Space: O(h).
// 时间复杂度：O(n)，空间复杂度：O(h)。
//
// 比较的是镜像位置：左的左对右的右，左的右对右的左；不能只比较同层数值而忽略空节点位置。
// Compare mirrored positions: left-left with right-right and left-right with right-left; values alone do not capture structure.
// 两个都空为真，只有一个空或值不同为假；队列/栈版始终成对存取镜像节点，空根视为对称。
// Two nil nodes match; one nil or unequal values fail. Iterative versions process pairs; an empty tree is symmetric.
func isMirror(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return isMirror(left.Left, right.Right) &&
		isMirror(left.Right, right.Left)
}

// 2. Iterative queue: store mirrored positions as adjacent pairs and dequeue two nodes each round.
// 2. 队列迭代：成对入队、成对出队，每轮取出相邻两个节点作为一对镜像位置。
// Time: O(n), Space: O(w) for the queue.
// 时间复杂度：O(n)，空间复杂度：O(w)，由队列产生。
//
// 比较的是镜像位置：左的左对右的右，左的右对右的左；不能只比较同层数值而忽略空节点位置。
// Compare mirrored positions: left-left with right-right and left-right with right-left; values alone do not capture structure.
// 两个都空为真，只有一个空或值不同为假；队列/栈版始终成对存取镜像节点，空根视为对称。
// Two nil nodes match; one nil or unequal values fail. Iterative versions process pairs; an empty tree is symmetric.
func isSymmetricIterative(root *TreeNode) bool {
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root.Left, root.Right)
	}

	for len(queue) > 0 {
		left, right := queue[0], queue[1]
		queue = queue[2:]

		if left == nil && right == nil {
			continue
		}

		if left == nil || right == nil || left.Val != right.Val {
			return false
		}

		queue = append(queue,
			left.Left, right.Right,
			left.Right, right.Left,
		)
	}

	return true
}

// 3. Iterative stack: keep the same pair rule, but the top is the later-pushed right-side candidate.
// 3. 栈迭代：配对规则不变，但栈顶是后压入的右侧候选，下面一个才是左侧候选。
// Time: O(n), Space: O(h) for the stack.
// 时间复杂度：O(n)，空间复杂度：O(h)，由栈产生。
//
// 比较的是镜像位置：左的左对右的右，左的右对右的左；不能只比较同层数值而忽略空节点位置。
// Compare mirrored positions: left-left with right-right and left-right with right-left; values alone do not capture structure.
// 两个都空为真，只有一个空或值不同为假；队列/栈版始终成对存取镜像节点，空根视为对称。
// Two nil nodes match; one nil or unequal values fail. Iterative versions process pairs; an empty tree is symmetric.
func isSymmetricStack(root *TreeNode) bool {
	var stack []*TreeNode
	if root != nil {
		stack = append(stack, root.Left, root.Right)
	}

	for len(stack) > 0 {
		last := len(stack) - 1
		left, right := stack[last-1], stack[last]
		stack = stack[:last-1]

		if left == nil && right == nil {
			continue
		}

		if left == nil || right == nil || left.Val != right.Val {
			return false
		}

		stack = append(stack,
			left.Left, right.Right,
			left.Right, right.Left,
		)
	}

	return true
}
