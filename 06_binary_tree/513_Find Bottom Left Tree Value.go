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

解法一：正常层序遍历（推荐） / Method 1: Standard BFS (Recommended)
用队列按层处理，孩子先左后右入队，因此每层开始时，队首就是这一层最左节点。
每轮先把队首值保存为答案，再记录 levelSize，恰好处理这一层的节点。
后面的层会覆盖前面的答案；队列耗尽后，最后保存的就是最深层的最左值。
Process levels with a queue, enqueuing left children before right children.
At each level's start, the front is its leftmost node. Save that value, capture levelSize, then process exactly that level.
Each deeper level replaces the answer, so the final saved value is the deepest level's leftmost value.

为什么必须保存 levelSize？处理当前层时会加入下一层孩子，固定数量才能准确分隔两层。
为什么每层队首最左？父节点从左到右出队，每个父节点又先左后右添加孩子，下一层顺序也保持从左到右。
Why save levelSize? Enqueuing children changes the queue; a fixed count separates the current level from the next.
Why is the front leftmost? Parents are processed left to right and enqueue left before right, preserving the next level's order.

解法二：先左后右的 DFS 递归 / Method 2: Left-First Recursive DFS
递归携带当前深度 depth，记录已经遇到的最大深度 deepest 和答案 answer。
先访问当前节点，再递归左子树、右子树。只有 depth>deepest 时才更新答案。
先左后右保证同一深度第一次访问到的节点最靠左；相同深度不能再次更新，否则会被右侧节点覆盖。
更深的节点即使在右子树，也应覆盖旧答案，因为深度优先于左右位置。
Carry depth through recursion and track the deepest depth seen and its answer.
Visit the node, then recurse left before right. Update only when depth>deepest.
Left-first traversal reaches the leftmost node at each depth first; equal-depth updates would overwrite it with nodes farther right.
A deeper node must replace a shallower answer even when it lies in the right subtree.

不必只判断叶子：最深层节点一定没有孩子，否则其孩子会处于更深层。因此遍历全部节点并记录最大深度即可。
No leaf-only check is required: a deepest-level node cannot have a child, which would be even deeper.
Tracking the greatest depth across all nodes suffices.

解法三：反向层序遍历 / Method 3: Right-to-Left BFS
仍用队列，但先右孩子、后左孩子入队，不再按层分组。每次出队都更新答案，最终返回最后出队节点的值。
队列依旧先处理浅层再处理深层；先右后左让每层从右往左处理，所以最后出队的恰好是最深层最左节点。
这依赖“队列 + 先右后左”的组合。普通先左后右 BFS 的最后节点是最深层最右节点，不能混用。
Use a queue but enqueue right before left, without grouping levels. Update the answer on every dequeue.
BFS still visits shallower levels first, while reversed child order visits each level right to left.
The final dequeued node is therefore the deepest level's leftmost node.
This requires a queue and right-before-left insertion; standard BFS ends at the deepest level's rightmost node instead.

时间与空间复杂度 / Time and Space Complexity
n 为节点数，h 为树高，w 为最大层宽。
三种方法时间均为 O(n)，每个节点访问一次。
两种 BFS 辅助空间 O(w)，递归 DFS 辅助空间 O(h)；最坏都可达 O(n)。不修改树，返回整数占 O(1)。
For n nodes, height h, and maximum width w, all methods take O(n) time.
Both BFS methods use O(w) auxiliary space; recursive DFS uses O(h). Each is at most O(n).
The tree is unchanged and the integer result uses O(1) space.

练习 / Practice
优先练习第一种，理解按层覆盖答案的逻辑。三种实现按上面的编号顺序写在下方。
Start with standard BFS and understand per-level answer replacement.
The three implementations appear below in the order of the numbered methods above.
*/

// 1. 层序遍历：推荐
func findBottomLeftValue(root *TreeNode) int {
	// The problem guarantees a nonempty tree.
	// 题目保证树非空，用根节点值初始化答案。
	answer := root.Val
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		// The queue front is the leftmost node of this level.
		// 每层开始时，队首就是这一层最左的节点。
		answer = queue[0].Val
		levelSize := len(queue)

		for range levelSize {
			node := queue[0]
			queue = queue[1:]

			// Enqueue left before right to preserve left-to-right order.
			// 先左后右入队，保持下一层从左到右的顺序。
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

// 2. 递归 DFS：先左后右
func findBottomLeftValueDFS(root *TreeNode) int {
	answer := root.Val
	deepest := 0

	var traverse func(*TreeNode, int)
	traverse = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		// Keep the first node encountered at each new deepest level.
		// 每次发现更深的一层，只记录第一次遇到的节点。
		if depth > deepest {
			deepest = depth
			answer = node.Val
		}

		// Visit left first so equal-depth nodes are encountered left to right.
		// 先左后右，保证同一深度先遇到左侧节点。
		traverse(node.Left, depth+1)
		traverse(node.Right, depth+1)
	}

	traverse(root, 1)
	return answer
}

// 3. 反向层序遍历：更简短
func findBottomLeftValueReverseBFS(root *TreeNode) int {
	queue := []*TreeNode{root}
	answer := root.Val

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		answer = node.Val

		// Visit each level from right to left.
		// 先右后左入队，让每层从右往左访问。
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
	}

	// The last visited node is the bottom-left node.
	// 最后访问的节点就是最深层最左节点。
	return answer
}
