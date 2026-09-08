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

解题思路：递归 / Approach: Recursion
定义 maxDepth(node)：返回以 node 为根的这棵子树的最大深度，计数从 node 自己开始。
它不是 node 在原树中处于第几层，也不需要知道它的父节点是谁。
Define maxDepth(node) as the maximum depth of the subtree rooted at node, counting from node itself.
It is not node's level in the original tree and requires no information about its parent.

1. node 为空时返回 0，表示这边没有节点。
   Return 0 for nil: there are no nodes on this side.
2. 递归求左子树最大深度 leftDepth 和右子树最大深度 rightDepth。
   Recursively obtain leftDepth and rightDepth for the two subtrees.
3. 返回 1 + max(leftDepth, rightDepth)。
   Return 1 + max(leftDepth, rightDepth).

关键逻辑：为什么取较大值再加 1？ / Why Take the Maximum and Add One?
从当前节点向下走的一条路径，只能进入左子树或右子树，不能同时走两边。
因此找最长路径应选两边较大的深度，而不是把两边相加。
子树返回的深度从孩子开始计数，还没包含当前节点，所以最后要加上当前节点这 1 个。
A downward path from the current node can enter either child subtree, not both.
Choose the larger subtree depth rather than adding the two depths.
Each child result counts from that child and excludes the current node, so add one for the current node.

按示例理解递归返回 / Following the Example's Return Values
- 9、15、7 都是叶子，左右子树深度均为 0，返回 1。
  Leaves 9, 15, and 7 each have two empty subtrees and return 1.
- 20 得到左右深度 1、1，返回 1+max(1,1)=2。
  Node 20 receives depths 1 and 1, returning 1+max(1,1)=2.
- 3 得到左右深度 1、2，返回 1+max(1,2)=3。
  Node 3 receives depths 1 and 2, returning 1+max(1,2)=3.
调用时从根向下深入，返回时从叶子往上汇总。需要先知道左右结果才能算当前结果，因此是后序思路。
Calls descend from the root, while results propagate upward from the leaves.
Both child results must be known before computing the current result, making this a postorder approach.

易错点 / Pitfalls
- 返回 0 的是空节点，不是叶子；叶子也算一个节点，所以返回 1。
  Nil returns 0; a leaf still counts as one node and returns 1.
- 只有一个孩子时，空的那边为 0，取 max 会选择实际存在的路径，无需特殊分支。
  With one child, the empty side contributes 0 and max selects the existing path without a special case.
- 不需要修改树，也不需要使用全局变量累计层数。
  No tree mutation or global depth counter is needed.

时间与空间复杂度 / Time and Space Complexity
n 为节点数，h 为树高。时间 O(n)，每个节点只计算一次深度。
辅助空间 O(h)，递归栈同时保存的是一条根到叶路径上的调用，而不是所有已经访问过的节点。
平衡树为 O(log n)，链状树最坏为 O(n)。返回整数占 O(1)。
For n nodes and height h, time is O(n), computing each node's depth once.
Auxiliary space is O(h): active calls follow a root-to-leaf path, not all nodes ever visited.
This is O(log n) for a balanced tree and O(n) for a skewed tree. The integer result uses O(1) space.

另一种思路：层序遍历 / Alternative: Level-Order Traversal
沿用 102 题队列写法，每轮保存当前队列长度，恰好处理这一层的节点，再将深度加 1。
全部节点处理完时，经过的层数就是最大深度。必须固定本层节点数，避免把新入队的孩子算入同一层。
时间 O(n)，辅助空间 O(w)，w 为最大层宽，最坏 O(n)。
Reuse the BFS queue from problem 102: save the level size, process exactly that many nodes, and increment depth once.
The number of completed levels is the maximum depth. Fix the level size to keep newly enqueued children in the next level.
Time is O(n), with O(w) auxiliary space for maximum width w, up to O(n).

练习 / Practice
请在下方自行实现 maxDepth；本文件仅保留题解，不提供实现代码或函数骨架。
Implement maxDepth below. This file contains explanations only, without implementation code or a function skeleton.
*/

func maxDepth(root *TreeNode) int {
	// An empty tree contains no nodes, so its depth is 0.
	// 空树没有节点，深度为 0。
	if root == nil {
		return 0
	}

	// Get each subtree's depth, counting from the child itself.
	// 分别求左右子树深度，都从各自的孩子节点开始计数。
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	// Choose the longer path and count the current node as well.
	// 选择较长的一边，再算上当前节点自己。
	return max(leftDepth, rightDepth) + 1
}

func maxDepthIterative(root *TreeNode) int {
	// An empty tree has no levels.
	// 空树没有层，深度为 0。
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	depth := 0

	for len(queue) > 0 {
		// At this moment, the queue contains exactly the current level.
		// 此时队列中恰好只有当前层的节点，先固定它们的数量。
		levelSize := len(queue)

		for range levelSize {
			// Remove one node from the current level.
			// 取出当前层的一个节点。
			node := queue[0]
			queue = queue[1:]

			// Children belong to the next level.
			// 孩子属于下一层，入队等待下一轮处理。
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// One entire level is finished.
		// 一整层处理完成，深度增加 1。
		depth++
	}

	return depth
}
