package _6_binary_tree

/*
题目描述 / Problem Description
给定二叉树的根节点 root，返回节点值的层序遍历结果，即从上到下、从左到右逐层访问所有节点。
Given the root of a binary tree, return its level-order traversal: visit all nodes level by level, from left to right.

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
*/

/*
题目描述 / Problem Description

给定二叉树的根节点 root，返回节点值的层序遍历结果，+
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

	return result
}
