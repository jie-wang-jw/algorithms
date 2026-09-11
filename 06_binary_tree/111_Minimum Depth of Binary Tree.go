package _6_binary_tree

/*
111. 二叉树的最小深度 / Minimum Depth of Binary Tree

题目描述 / Problem Description
给定二叉树的根节点 root，返回从根节点到最近叶子节点的最短路径上的节点数量。
叶子节点是左右孩子都为空的节点；空树的最小深度为 0。
Given the root of a binary tree, return the number of nodes on the shortest path from the root to a leaf.
A leaf has neither a left nor a right child. An empty tree has minimum depth 0.

示例 / Examples
       3                 1
      / \                 \
     9  20                 2
       /  \                 \
      15   7                 3
左图最小深度为 2，路径为 3 -> 9；右图最小深度为 3，路径为 1 -> 2 -> 3。
The left tree has minimum depth 2 via 3 -> 9; the right tree has minimum depth 3 via 1 -> 2 -> 3.

关键逻辑：为什么不能直接把 104 的 max 换成 min？
Key Logic: Why Not Simply Replace max from Problem 104 with min?
最短路径必须终止于实际存在的叶子节点，不能终止于空指针。
右图根节点的左孩子为空，但根有右孩子，所以根不是叶子，路径不能停在根这里。
若直接取 min(0,2)+1，会错误返回 1；这个 0 代表没有子树，不代表找到了一条到叶子的路径。
The shortest path must end at an actual leaf, not a nil pointer.
The right example's root has no left child but does have a right child, so it is not a leaf.
Taking min(0,2)+1 would incorrectly return 1. That zero means absence of a subtree, not a valid path to a leaf.

解法一：递归（推荐） / Method 1: Recursion (Recommended)
定义 minDepth(node)：返回以 node 为根，到该子树最近叶子的路径节点数。
Define minDepth(node) as the number of nodes from node to the nearest leaf in its subtree.

1. 当前节点为空：返回 0。
   If the current node is nil, return 0.
2. 左孩子为空：只能沿右子树找叶子，返回右子树最小深度加 1。
   If the left child is nil, use the right subtree's minimum depth plus one.
3. 右孩子为空：只能沿左子树找叶子，返回左子树最小深度加 1。
   If the right child is nil, use the left subtree's minimum depth plus one.
4. 两个孩子都存在：两边都有通向叶子的路径，才能取较小深度再加 1。
   If both children exist, both offer a path to a leaf; choose the smaller depth and add one.

为什么加 1？子树返回的深度从孩子开始算，还没有包含当前节点自己。
为什么不单独判断叶子？叶子的左孩子为空，会进入第 2 步；右孩子也为空，得到 0+1=1，结果自然正确。
Add one because child-subtree depths exclude the current node.
A separate leaf case is unnecessary: step 2 returns the nil right subtree's depth plus one, giving 0+1=1.

对比 104 / Comparison with Problem 104
最大深度选最长路径，空子树的 0 不会压过非空子树的正深度，所以直接取 max 即可。
最小深度选最短合法路径，空子树的 0 却会被 min 优先选中，所以必须先排除缺失的那一边。
For maximum depth, zero from a missing subtree cannot beat a positive depth under max.
For minimum depth, min would incorrectly favor that zero, so exclude the missing side first.

解法二：层序遍历 / Method 2: Breadth-First Search
从第 1 层开始，逐层出队检查。第一次遇到左右孩子都为空的节点，就返回当前层数。
队列按深度从小到大处理，因此第一个叶子的深度就是最小深度，不需要遍历更深的节点。
判断叶子必须是两个孩子都为空，不能用“其中一个为空”。固定 levelSize，处理一整层后再增加深度。
Visit levels starting at depth 1. Return the current depth upon encountering the first node with both children nil.
BFS processes shallower nodes first, so its first leaf has minimum depth; deeper nodes need not be visited.
A leaf requires both children to be nil, not merely one. Save levelSize and increment depth only after finishing that level.

时间与空间复杂度 / Time and Space Complexity
n 为节点数，h 为树高，w 为最大层宽。
递归：最坏时间 O(n)，辅助空间 O(h) 来自调用栈，最坏 O(n)。不修改树，返回整数。
层序：最坏时间 O(n)，可在第一个叶子处提前返回；辅助空间 O(w)，最坏 O(n)。
For n nodes, height h, and maximum width w:
Recursion takes O(n) worst-case time and O(h) call-stack space, up to O(n). It returns an integer without modifying the tree.
BFS takes O(n) worst-case time, with early return at the first leaf, and O(w) auxiliary space, up to O(n).

练习 / Practice
先想清楚“为什么不能直接把 max 换成 min”，再写两种实现。两种实现按上面的编号顺序写在下方。
Settle the question of why max cannot simply become min, then write both versions.
The two implementations appear below in the order of the numbered methods above.
*/

// 1. 递归：先排除缺失的一边
func minDepth(root *TreeNode) int {
	// An empty tree has depth 0.
	// 空树深度为 0。
	if root == nil {
		return 0
	}

	// No left subtree: the path must continue through the right.
	// 左子树不存在，只能沿右子树寻找叶子。
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}

	// No right subtree: the path must continue through the left.
	// 右子树不存在，只能沿左子树寻找叶子。
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}

	// Both sides contain leaves, so choose the shorter valid path.
	// 两边都有通向叶子的路径，选择较短的一边。
	leftDepth := minDepth(root.Left)
	rightDepth := minDepth(root.Right)

	// Include the current node.
	// 子树深度还没包含当前节点，所以加 1。
	return min(leftDepth, rightDepth) + 1
}

// 2. 层序遍历：第一个叶子就是最近叶子
func minDepthIterative(root *TreeNode) int {
	// An empty tree has depth 0.
	// 空树深度为 0。
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	depth := 1

	for len(queue) > 0 {
		// Fix the number of nodes belonging to the current level.
		// 固定当前层的节点数量，避免把孩子算进同一层。
		levelSize := len(queue)

		for range levelSize {
			node := queue[0]
			queue = queue[1:]

			// BFS reaches shallower levels first, so the first leaf is nearest.
			// BFS 先检查浅层，因此第一个叶子就是最近的叶子。
			if node.Left == nil && node.Right == nil {
				return depth
			}

			// Children belong to the next level.
			// 孩子属于下一层，入队等待后续处理。
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// Advance only after finishing the entire current level.
		// 当前整层处理完后，才进入下一层。
		depth++
	}

	// A finite nonempty tree always has a leaf, so this is unreachable.
	// 有限非空二叉树一定有叶子，正常情况下不会执行到这里。
	return 0
}
