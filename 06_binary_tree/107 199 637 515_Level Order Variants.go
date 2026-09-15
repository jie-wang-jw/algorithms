package _6_binary_tree

/*
102 层序遍历模板的相关题 / Level-Order Template Variants
文章《二叉树层序遍历登场！》用同一套队列模板连做多题。104 与 111 已有独立文件。
The level-order article applies one queue template to several problems. 104 and 111 already have their own files.

107. 二叉树的层序遍历 II / Binary Tree Level Order Traversal II
自底向上返回各层；仍从上到下收集，最后反转层的顺序，层内仍从左到右。
Return levels bottom-up: collect top-down, then reverse the outer slice. Order within a level stays left to right.

199. 二叉树的右视图 / Binary Tree Right Side View
每一层最后一个出队的节点就是该层最右侧节点。
The last node dequeued from each level is that level's rightmost value.

637. 二叉树的层平均值 / Average of Levels in Binary Tree
每一层求和后除以固定的 levelSize。用 float64，避免整数除法截断。
Sum each level and divide by the captured levelSize. Use float64 so integer division does not truncate.

515. 在每个树行中找最大值 / Find Largest Value in Each Tree Row
每一层维护当前最大值，比较该层每个节点。
Track the maximum while scanning one captured level.

429 / 116 / 117 需要不同的节点类型（N 叉孩子表或 Next 指针），本文件只实现仍使用 TreeNode 的四题。
429, 116, and 117 need different node types (N-ary children or Next pointers); this file covers the four TreeNode variants.
*/

// 1. Bottom-up level order: reuse top-down grouping, then reverse the order of levels.
// 1. 自底向上层序：先按从上到下分层，再反转层与层的顺序。
// Time: O(n), Space: O(w).
// 时间复杂度：O(n)，空间复杂度：O(w)。
//
// 步骤与要点 / Steps and notes:
//  1. Reuse the standard top-down level-order result.
//     复用标准的自上而下分层结果。
//  2. Reverse only the outer slice; each level stays left-to-right.
//     只反转外层切片；每一层内部仍保持从左到右。
func levelOrderBottom(root *TreeNode) [][]int {
	levels := levelOrder(root)
	for left, right := 0, len(levels)-1; left < right; left, right = left+1, right-1 {
		levels[left], levels[right] = levels[right], levels[left]
	}
	return levels
}

// 2. Right side view: the last node of each captured level is the rightmost value on that level.
// 2. 右视图：每一层固定数量中的最后一个节点，就是该层最右侧的值。
// Time: O(n), Space: O(w).
// 时间复杂度：O(n)，空间复杂度：O(w)。
//
// 步骤与要点 / Steps and notes:
//  1. An empty tree has no visible values from the right.
//     空树从右侧看不到任何节点。
//  2. Capture the current level size before enqueuing children.
//     先固定本层节点数，再入队孩子，避免把下层混进本层。
//  3. Overwrite last on every dequeue; the final write is the rightmost node.
//     每次出队都覆盖 last；最后一次写入就是该层最右侧节点。
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		var last int
		for range levelSize {
			node := queue[0]
			queue = queue[1:]
			last = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, last)
	}
	return result
}

// 3. Level averages: divide each level's sum by the captured level size using float64.
// 3. 层平均值：用 float64 把该层总和除以固定的层节点数。
// Time: O(n), Space: O(w).
// 时间复杂度：O(n)，空间复杂度：O(w)。
//
// 步骤与要点 / Steps and notes:
//  1. sum accumulates only the nodes already counted in levelSize.
//     sum 只累加已经计入 levelSize 的本层节点。
//  2. float64 avoids truncating the average with integer division.
//     使用 float64，避免整数除法截断平均值。
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	result := []float64{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		sum := 0
		for range levelSize {
			node := queue[0]
			queue = queue[1:]
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, float64(sum)/float64(levelSize))
	}
	return result
}

// 4. Largest value in each row: compare every node in the captured level against a running maximum.
// 4. 每行最大值：在固定的一层里用当前最大值逐个比较。
// Time: O(n), Space: O(w).
// 时间复杂度：O(n)，空间复杂度：O(w)。
//
// 步骤与要点 / Steps and notes:
//  1. Seed best with the first node on this level before scanning the rest.
//     先用本层第一个节点初始化 best，再扫描其余节点。
func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		best := queue[0].Val
		for range levelSize {
			node := queue[0]
			queue = queue[1:]
			if node.Val > best {
				best = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, best)
	}
	return result
}
