package _6_binary_tree

/*
513. 找树左下角的值 / Find Bottom Left Tree Value

题目描述 / Problem Description
给定非空二叉树的根节点 root，返回最深一层中最左边节点的值。
优先比较深度；只有深度相同时，才选择更靠左的节点。题目保证至少有一个节点。
Given the root of a nonempty binary tree, return the value of the leftmost node in its deepest level.
Depth takes priority; choose the leftmost only among nodes at that depth. At least one node is guaranteed.

示例 / Example
       1
      / \
     2   3
    /   / \
   4   5   6
      /
     7
结果为 7，而不是 4：4 更靠左，但 7 所在层更深。
The answer is 7, not 4: node 4 is farther left, but node 7 is deeper.

关键区别 / Key Distinctions
- 不是找最小节点值，节点值大小不参与位置比较。
  This is not the minimum value; values do not determine positions.
- 不是一直沿左孩子走到底，最深层可能在右子树中。
  Repeatedly following left children may miss the deepest level in the right subtree.
- 不是找“左叶子”。右链 1->2->3 的答案是 3，虽然它是右孩子。
  This is not a left-leaf search. A right-only chain 1->2->3 returns 3 even though it is a right child.
复杂度记号：n 为节点数，h 为树高，w 为最大层宽；辅助空间不含返回结果。
Notation: n nodes, height h, maximum width w; auxiliary space excludes returned results.
*/

// 1. Standard BFS (recommended): save the queue front at each level start so later levels overwrite with a deeper leftmost value.
// 1. 正常层序遍历：推荐；每层开始时保存队首，更深层会覆盖答案，最终就是最深层最左值。
// Time: O(n), Space: O(w) for the queue.
// 时间复杂度：O(n)，空间复杂度：O(w)，由队列产生。
//
// 要求根非空。左孩子先入队，保证每层队首是该层最左节点；逐层覆盖 answer，最终留下最深层最左值。
// Require nonnil root; enqueue left first so each level's front is leftmost, then overwrite answer per level to retain the deepest.
func findBottomLeftValue(root *TreeNode) int {
	answer := root.Val
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		answer = queue[0].Val
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
	}

	return answer
}

// 2. Left-first recursive DFS: update the answer only when the current depth strictly exceeds the deepest seen so far.
// 2. 先左后右的 DFS 递归：只有当前深度严格大于已见最大深度时才更新答案。
// Time: O(n), Space: O(h) for the recursion stack.
// 时间复杂度：O(n)，空间复杂度：O(h)，由递归栈产生。
//
// 要求根非空。先左后右遍历，同层先遇到的就是最左；仅 depth>deepest 才更新，不能用 >= 覆盖同层答案。
// Require nonnil root; left-first DFS encounters each depth's leftmost node first, so update only for depth>deepest, never equality.
func findBottomLeftValueDFS(root *TreeNode) int {
	answer := root.Val
	deepest := 0

	var traverse func(*TreeNode, int)
	traverse = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		if depth > deepest {
			deepest = depth
			answer = node.Val
		}

		traverse(node.Left, depth+1)
		traverse(node.Right, depth+1)
	}

	traverse(root, 1)
	return answer
}

// 3. Right-to-left BFS: enqueue right before left and keep the last dequeued value, which is the deepest leftmost node.
// 3. 反向层序遍历：先右后左入队，最后出队的节点就是最深层最左节点。
// Time: O(n), Space: O(w) for the queue.
// 时间复杂度：O(n)，空间复杂度：O(w)，由队列产生。
//
// 要求根非空。每层先右后左出队，最后一个出队节点必在最深层最左端；不断覆盖 answer 即可。
// Require nonnil root; right-first BFS makes the final dequeued node the deepest leftmost one, so retain the last value.
func findBottomLeftValueReverseBFS(root *TreeNode) int {
	queue := []*TreeNode{root}
	answer := root.Val

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		answer = node.Val

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
	}

	return answer
}
