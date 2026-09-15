package _6_binary_tree

/*
429. N 叉树的层序遍历 / N-ary Tree Level Order Traversal

题目描述 / Problem Description
给定 N 叉树根节点，返回层序遍历：从上到下、从左到右逐层访问。
Given the root of an N-ary tree, return its level-order traversal:
visit nodes level by level from left to right.

示例 / Example
根 1，孩子 3、2、4，节点 3 的孩子 5、6 → [[1],[3,2,4],[5,6]]

解法一：队列按层处理 / Method 1: Level-by-Level Queue
与 102 同一套模板，只是把左右孩子改成遍历 Children。
Same template as 102; enqueue every child instead of only left and right.

解法二：递归按深度分组 / Method 2: Recursion Grouped by Depth
深度当作结果下标，第一次到达该层时追加空子数组。先序遍历孩子，层内仍从左到右。
Use depth as the result index and append a new subarray on first arrival.
Visit children in order so each level fills left to right.

时间与空间复杂度 / Time and Space Complexity
n 为节点数，w 为最大层宽，h 为树高。
解法一时间 O(n)，辅助空间 O(w)。解法二时间 O(n)，辅助空间 O(h)。
For n nodes, width w, and height h:
method 1 takes O(n) time and O(w) auxiliary space; method 2 takes O(n) time and O(h) call-stack space.
*/

// 1. Level-by-level queue: capture levelSize, then enqueue every child.
// 1. 队列层序：先固定 levelSize，再把每个孩子入队。
// Time: O(n), Space: O(w) for the queue.
// 时间复杂度：O(n)，空间复杂度：O(w)，由队列产生。
//
// 步骤与要点 / Steps and notes:
//  1. An empty N-ary tree contributes no levels.
//     空的 N 叉树没有任何层。
//  2. Capture how many nodes belong to the current level.
//     固定属于当前层的节点个数。
//  3. Enqueue every non-nil child; order preserves left-to-right within the next level.
//     把每个非空孩子入队；顺序保证下一层仍从左到右。
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
// 步骤与要点 / Steps and notes:
//  1. First visit to this depth allocates the level slice.
//     第一次到达该深度时，先为这一层分配切片。
//  2. Children are visited in order, so each level fills left to right.
//     按顺序递归孩子，因此每一层仍从左到右填充。
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
