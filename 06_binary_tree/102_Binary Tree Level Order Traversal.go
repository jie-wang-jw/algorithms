package _6_binary_tree

/*
题目描述 / Problem Description
给定二叉树的根节点 root，返回节点值的层序遍历结果，即从上到下、从左到右逐层访问所有节点。
Given the root of a binary tree, return its level-order traversal:
visit all nodes level by level, from left to right.

解题思路 / Solution Approach
层序遍历本质上是广度优先搜索（BFS），需要使用队列。
Level-order traversal is a breadth-first search (BFS), so we use a queue.
执行过程：
Process:
1. 将根节点加入队列。
   Add the root to the queue.
2. 每轮先记录当前队列长度 levelSize。
   At the start of each round, record the current queue length as levelSize.
3. 从队列中取出当前层的 levelSize 个节点。
   Remove exactly levelSize nodes belonging to the current level.
4. 保存这些节点的值，并将它们的左右子节点加入队列。
   Save their values and enqueue their left and right children.
5. 当前层处理完成后，将结果加入 result。
   After finishing the current level, append it to result.

关键逻辑：为什么这样做 / Why This Works
每轮开始时队列里恰好只有当前层，因此先保存的 levelSize 才是当前层数量。处理中孩子不断加入队尾，使队列混合两层；只弹出原来的 levelSize 个，
剩下的就恰好是下一层。若每次用变化的 len(queue) 判断本层结束，会把层的边界弄错。父节点从左到右出队，每个父节点又先左后右入队，所以孩子也按从左到右排列。
At each outer-loop start, the queue contains exactly one level, so its saved length is the level size. Enqueuing children mixes two levels;
removing exactly the original levelSize nodes leaves precisely the next level. Using the changing queue length would lose that boundary.
Parents leave left to right and enqueue left child before right, preserving the next level's left-to-right order.

时间与空间复杂度 / Time and Space Complexity
n 为节点数，w 为最大层宽。时间 O(n)，每个节点入队、出队各一次。辅助空间 O(w)，
处理过程中队列可以同时包含当前层剩余节点和下一层节点，但数量仍为 O(w)。
返回结果 O(n)，包含结果的总空间 O(n)。
For n nodes and maximum width w, time is O(n), enqueuing and dequeuing each node once.
Auxiliary space O(w): the queue can mix remaining current-level nodes with next-level
nodes but stays O(w). Output and total space are O(n).
*/

/*
题目描述 / Problem Description

给定二叉树的根节点 root，返回节点值的层序遍历结果，
即从上到下、从左到右逐层访问节点。

Given the root of a binary tree, return its level-order traversal,
visiting nodes level by level from left to right.

解题思路 / Solution Approach

使用队列进行广度优先搜索。
Use a queue to perform breadth-first search.

每轮先保存当前队列长度，它就是当前层的节点数量。
At the beginning of each round, save the current queue length,
which is the number of nodes in the current level.

处理当前层节点时，将它们的子节点加入队列，
这些子节点会在下一轮中处理。

While processing the current level, enqueue its children so they
can be processed in the next round.

// Ordinary tree:
	// 普通二叉树：
	//
	//          1
	//        /   \
	//       2     3
	//      / \     \
	//     4   5     6
*/

func levelOrder(root *TreeNode) [][]int {
	// An empty tree has no levels.
	// 空树没有任何层。
	if root == nil {
		return [][]int{}
	}

	// result stores the values of all completed levels.
	// result 保存所有已经处理完成的层。
	result := make([][]int, 0)

	// Start BFS by adding the root to the queue.
	// 将根节点加入队列，开始广度优先搜索。
	queue := []*TreeNode{root}

	// Continue until every node has been processed.
	// 持续处理，直到队列中没有节点。
	for len(queue) > 0 {
		// Save the current queue length before adding any children.
		// 添加子节点之前，先保存当前层的节点数量。
		levelSize := len(queue)

		// level stores the values in the current level.
		// level 保存当前层的所有节点值。
		level := make([]int, 0, levelSize)

		// Process exactly the nodes belonging to the current level.
		// 只处理属于当前层的 levelSize 个节点。
		for range levelSize {
			// Remove the node at the front of the queue.
			// 取出并删除队首节点。
			node := queue[0]
			queue = queue[1:]

			// Save the current node's value.
			// 保存当前节点的值。
			level = append(level, node.Val)

			// Add the left child for the next level.
			// 将左子节点加入队列，等待下一层处理。
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			// Add the right child for the next level.
			// 将右子节点加入队列，等待下一层处理。
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// The current level is complete.
		// 当前层已经处理完成。
		result = append(result, level)
	}

	//fmt.Println(result)
	return result
}
