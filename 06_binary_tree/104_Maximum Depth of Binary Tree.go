package _6_binary_tree

/*
104. 二叉树的最大深度 / Maximum Depth of Binary Tree

题目描述 / Problem Description
给定二叉树的根节点 root，返回它的最大深度。
最大深度是从根节点到最远叶子节点的路径上的节点数量，不是边的数量。
空树深度为 0，只有根节点的树深度为 1。
Given the root of a binary tree, return its maximum depth.
Maximum depth is the number of nodes, not edges, on the longest path from the root to a leaf.
An empty tree has depth 0; a single-node tree has depth 1.

示例 / Example
       3
      / \
     9  20
       /  \
      15   7
最大深度为 3，例如路径 3 -> 20 -> 15 包含三个节点。
The maximum depth is 3: for example, path 3 -> 20 -> 15 contains three nodes.
复杂度记号：n 为节点数，h 为树高，w 为最大层宽；辅助空间不含返回结果。
Notation: n nodes, height h, maximum width w; auxiliary space excludes returned results.
*/

// 1. Postorder recursion on heights (recommended): take the larger child height and add one for the current node.
// 1. 后序递归求高度：推荐；先求左右高度，再取较大值加一。
// Time: O(n), Space: O(h) for the recursion stack.
// 时间复杂度：O(n)，空间复杂度：O(h)，由递归栈产生。
//
// 递归返回以当前节点为根的高度：空节点为 0，否则取两个孩子高度较大者再加当前这一层。
// Return subtree height: nil contributes 0; otherwise take the larger child height plus the current level.
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	return max(leftDepth, rightDepth) + 1
}

// 2. Level-order traversal: process a fixed levelSize of nodes, then increment depth once per finished level.
// 2. 层序遍历：每轮固定 levelSize 处理一层，整层结束后深度加一。
// Time: O(n), Space: O(w) for the queue.
// 时间复杂度：O(n)，空间复杂度：O(w)，由队列产生。
//
// 固定每层节点数后全部出队，再令 depth++；队列处理完多少层，最大深度就是多少。
// Process each frozen layer before incrementing depth; the number of completed layers is the maximum depth.
func maxDepthIterative(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	depth := 0

	for len(queue) > 0 {
		levelSize := len(queue)

		for range levelSize {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		depth++
	}

	return depth
}

// 3. Preorder backtracking on depths: pass a top-down level and keep the maximum seen so far.
// 3. 前序回溯记录深度：自顶向下传递层数，并用当前深度更新答案。
// Time: O(n), Space: O(h) for the recursion stack.
// 时间复杂度：O(n)，空间复杂度：O(h)，由递归栈产生。
//
// depth 是根到当前节点的节点数，根为 1；每到非空节点就更新全局最大值，孩子深度为 depth+1。
// depth counts nodes from root to current node, starting at 1; update the maximum and pass depth+1 to children.
func maxDepthPreorder(root *TreeNode) int {
	answer := 0

	var traverse func(*TreeNode, int)
	traverse = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		answer = max(answer, depth)

		traverse(node.Left, depth+1)
		traverse(node.Right, depth+1)
	}

	traverse(root, 1)

	return answer
}
