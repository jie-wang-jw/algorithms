package _6_binary_tree

/*
429. N 叉树的层序遍历 / N-ary Tree Level Order Traversal

题目描述 / Problem Description
给定 N 叉树根节点，返回层序遍历：从上到下、从左到右逐层访问。
Given the root of an N-ary tree, return its level-order traversal:
visit nodes level by level from left to right.

示例 / Example
根 1，孩子 3、2、4，节点 3 的孩子 5、6 → [[1],[3,2,4],[5,6]]
复杂度记号：n 为节点数，h 为树高，w 为最大层宽；辅助空间不含返回结果。
Notation: n nodes, height h, maximum width w; auxiliary space excludes returned results.
*/

// 1. Level-by-level queue: capture levelSize, then enqueue every child.
// 1. 队列层序：先固定 levelSize，再把每个孩子入队。
// Time: O(n), Space: O(w) for the queue.
// 时间复杂度：O(n)，空间复杂度：O(w)，由队列产生。
//
// 每轮先固定 levelSize=len(queue)，只弹出这些节点；新入队的孩子属于下一层，不能计入本轮。
// Freeze levelSize before each round; newly enqueued children belong to the next level, not the current one.
// 按从左到右的孩子顺序入队，确保同层输出顺序；空树返回空层列表。
// Enqueue children left to right to preserve order within each level; empty input produces no levels.
func naryLevelOrder(root *NaryNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*NaryNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, 0, levelSize)
		for range levelSize {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			for _, child := range node.Children {
				if child != nil {
					queue = append(queue, child)
				}
			}
		}
		result = append(result, level)
	}
	return result
}

// 2. Recursion grouped by depth: depth is the result index; first arrival creates the subarray.
// 2. 递归按深度分组：深度就是结果下标，第一次到达该层时先追加空子数组。
// Time: O(n), Space: O(h) for the recursion stack.
// 时间复杂度：O(n)，空间复杂度：O(h)，由递归栈产生。
//
// depth 是从根起的零基层号；首次到达新深度才新增结果层，再把节点值加入 result[depth]。
// depth is the zero-based level; create a result row on first reaching that depth, then append to result[depth].
// DFS 不按层访问，但按从左到右递归保证每层加入顺序正确。
// DFS does not visit by level, but left-to-right recursion preserves order within each result row.
func naryLevelOrderRecursive(root *NaryNode) [][]int {
	result := [][]int{}

	var traverse func(*NaryNode, int)
	traverse = func(node *NaryNode, depth int) {
		if node == nil {
			return
		}
		if depth == len(result) {
			result = append(result, []int{})
		}
		result[depth] = append(result[depth], node.Val)
		for _, child := range node.Children {
			traverse(child, depth+1)
		}
	}

	traverse(root, 0)
	return result
}
