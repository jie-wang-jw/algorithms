package _6_binary_tree

/*
112. 路径总和 / Path Sum

题目描述 / Problem Description
给定二叉树根节点 root 和整数 targetSum，判断是否存在一条从根到叶子的路径，
使路径上所有节点值之和等于 targetSum。叶子节点的左右孩子都为空。
空树不存在根到叶子的路径，即使 targetSum 为 0 也应返回 false。
Given the root of a binary tree and an integer targetSum, determine whether a root-to-leaf path
has a sum equal to targetSum. A leaf has no children.
An empty tree has no root-to-leaf path, even when targetSum is zero.

示例 / Example
       5
      / \
     4   8
    /   / \
   11  13  4
  / \      \
 7   2      1
targetSum=22 时返回 true，路径为 5->4->11->2，总和 5+4+11+2=22。
targetSum=9 不能因为 5+4=9 就成功：4 不是叶子，必须继续走到实际叶子再判断。
For targetSum=22, return true via 5->4->11->2.
For targetSum=9, prefix 5->4 is not enough: node 4 is not a leaf.
复杂度记号：n 为节点数，h 为树高，w 为最大层宽；辅助空间不含返回结果。
Notation: n nodes, height h, maximum width w; auxiliary space excludes returned results.
*/

// 1. Recursion with a remaining target: subtract the current value, then succeed only at a leaf whose remainder is 0.
// 1. 递归传递剩余目标：先扣除当前值，只有到达叶子且剩余为 0 才成功。
// Time: O(n), Space: O(h).
// 时间复杂度：O(n)，空间复杂度：O(h)。
//
// remaining 是扣除当前节点后的待补和；只有到叶子且 remaining==0 才成功，内部节点达标不能提前返回。
// remaining is the target after subtracting the current value; success requires a leaf and zero remaining sum.
// 存在负值，不能因剩余和为负而剪枝；空树没有根到叶路径。
// Negative values forbid pruning solely by a negative remainder; an empty tree has no root-to-leaf path.
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	remaining := targetSum - root.Val

	if root.Left == nil && root.Right == nil {
		return remaining == 0
	}

	return hasPathSum(root.Left, remaining) ||
		hasPathSum(root.Right, remaining)
}

// 2. Iterative DFS: keep synchronized node and sum stacks so each pending node carries its own root-to-node total.
// 2. 栈迭代：节点栈与路径和栈同步，每个待处理节点带着从根到它自己的累计和。
// Time: O(n), Space: O(h).
// 时间复杂度：O(n)，空间复杂度：O(h)。
//
// nodes 与 sums 同步存取，sum 表示根到该节点的累计和；孩子入容器时带上 sum+child.Val。
// Keep nodes and sums aligned; sum is the root-to-node total and each child receives sum+child.Val.
// 只有叶子才能用累计和判答案；负数允许中途超过目标，不能据此剪枝。
// Compare the sum only at leaves; negative values mean exceeding the target early is not grounds for pruning.
func hasPathSumIterative(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	nodes := []*TreeNode{root}
	sums := []int{root.Val}

	for len(nodes) > 0 {
		last := len(nodes) - 1
		node, sum := nodes[last], sums[last]
		nodes = nodes[:last]
		sums = sums[:last]

		if node.Left == nil && node.Right == nil &&
			sum == targetSum {
			return true
		}

		if node.Right != nil {
			nodes = append(nodes, node.Right)
			sums = append(sums, sum+node.Right.Val)
		}
		if node.Left != nil {
			nodes = append(nodes, node.Left)
			sums = append(sums, sum+node.Left.Val)
		}
	}

	return false
}

// 3. Breadth-first search: dequeue a node together with its path sum; no levelSize is needed.
// 3. 队列 BFS：节点与累计和一起出队，本题不问深度，无需按层循环。
// Time: O(n), Space: O(w).
// 时间复杂度：O(n)，空间复杂度：O(w)。
//
// nodes 与 sums 同步存取，sum 表示根到该节点的累计和；孩子入容器时带上 sum+child.Val。
// Keep nodes and sums aligned; sum is the root-to-node total and each child receives sum+child.Val.
// 只有叶子才能用累计和判答案；负数允许中途超过目标，不能据此剪枝。
// Compare the sum only at leaves; negative values mean exceeding the target early is not grounds for pruning.
func hasPathSumBFS(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	nodes := []*TreeNode{root}
	sums := []int{root.Val}

	for len(nodes) > 0 {
		node, sum := nodes[0], sums[0]
		nodes = nodes[1:]
		sums = sums[1:]

		if node.Left == nil && node.Right == nil &&
			sum == targetSum {
			return true
		}

		if node.Left != nil {
			nodes = append(nodes, node.Left)
			sums = append(sums, sum+node.Left.Val)
		}
		if node.Right != nil {
			nodes = append(nodes, node.Right)
			sums = append(sums, sum+node.Right.Val)
		}
	}

	return false
}
