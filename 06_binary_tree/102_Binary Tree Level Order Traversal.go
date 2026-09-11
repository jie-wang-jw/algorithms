package _6_binary_tree

/*
102. 二叉树的层序遍历 / Binary Tree Level Order Traversal

题目描述 / Problem Description
给定二叉树的根节点 root，返回节点值的层序遍历结果，即从上到下、从左到右逐层访问所有节点。
Given the root of a binary tree, return its level-order traversal:
visit all nodes level by level, from left to right.

解法一：队列按层处理（推荐） / Method 1: Level-by-Level Queue (Recommended)
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

解法二：递归按深度分组 / Method 2: Recursion Grouped by Depth
层序结果并不要求真的按层推进，只要求“同一深度的值按从左到右放进同一个子数组”。
递归携带当前深度 depth（根为 0），depth 正好是这个节点应写入的 result 下标。
depth==len(result) 说明第一次到达这一层，先追加一个空子数组，再写入当前值。
Level-order output does not require advancing level by level; it only requires that values at the same depth
land in the same subarray, ordered left to right. Carry a depth through the recursion, with 0 at the root;
that depth is exactly the result index this node belongs to.
When depth==len(result), this is the first node of a new level, so append an empty subarray before writing.

为什么下标不会跳号？前序递归到达深度 d 之前，一定已经经过它在 0..d-1 各层的祖先，
所以那些层的子数组已经创建，len(result) 只会恰好等于 d，不会小于 d。
为什么左右顺序正确？同一深度上，先递归左子树意味着左侧节点先被写入，与队列版本的从左到右一致。
Why can no index be skipped? A preorder call reaching depth d has already passed its ancestors at depths 0..d-1,
so those subarrays exist and len(result) equals d exactly rather than falling short.
Why is the order correct? Recursing left first writes left-side nodes first at each depth,
matching the left-to-right order of the queue version.

时间与空间复杂度 / Time and Space Complexity
n 为节点数，w 为最大层宽，h 为树高。
解法一：时间 O(n)，每个节点入队、出队各一次；辅助空间 O(w)，处理过程中队列可以同时包含
当前层剩余节点和下一层节点，但数量仍为 O(w)，最坏 O(n)。
解法二：时间 O(n)，每个节点访问一次并追加一次；辅助空间 O(h) 来自递归栈，最坏 O(n)。
两种解法的返回结果都占 O(n)，包含结果的总空间 O(n)；都不修改输入树。
For n nodes, maximum width w, and height h:
Method 1 takes O(n) time, enqueuing and dequeuing each node once, and O(w) auxiliary space,
because the queue can mix remaining current-level nodes with next-level nodes, up to O(n).
Method 2 takes O(n) time, visiting and appending each node once, and O(h) call-stack space, up to O(n).
Both return O(n) output for O(n) total space, and neither modifies the input tree.

练习 / Practice
先掌握解法一固定 levelSize 的写法，再理解解法二为什么不需要按层推进。两种实现按上面的编号顺序写在下方。
Master the fixed levelSize of method 1 first, then see why method 2 needs no level-by-level advance.
The two implementations appear below in the order of the numbered methods above.
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

// 1. 队列层序：固定 levelSize 分隔相邻两层，推荐
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

	return result
}

// 2. 递归按深度分组：深度就是结果下标
func levelOrderRecursive(root *TreeNode) [][]int {
	// An empty tree produces no levels, so the empty result is already correct.
	// 空树没有任何层，空结果就是正确答案。
	result := [][]int{}

	var traverse func(*TreeNode, int)
	traverse = func(node *TreeNode, depth int) {
		// A nil child occupies no position in any level.
		// 空孩子不占据任何一层中的位置。
		if node == nil {
			return
		}

		// Reaching a new deepest level: create its subarray first.
		// 第一次到达这一层，先为它创建一个空子数组。
		if depth == len(result) {
			result = append(result, []int{})
		}

		// depth is exactly the result index this node belongs to.
		// depth 正好是当前节点应该写入的 result 下标。
		result[depth] = append(result[depth], node.Val)

		// Recurse left first so each level fills from left to right.
		// 先递归左子树，使每一层按从左到右的顺序填充。
		traverse(node.Left, depth+1)
		traverse(node.Right, depth+1)
	}

	// The root belongs to index 0 because depth is used as a slice index.
	// 根对应下标 0，因为这里的 depth 直接当作切片下标使用。
	traverse(root, 0)

	return result
}
