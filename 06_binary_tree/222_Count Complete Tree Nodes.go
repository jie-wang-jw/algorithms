package _6_binary_tree

/*
222. 完全二叉树的节点个数 / Count Complete Tree Nodes

题目描述 / Problem Description
给定一棵完全二叉树的根节点 root，返回节点总数。
完全二叉树：除最后一层外，每层都填满；最后一层节点从左到右连续排列，中间不能有空位。
Given the root of a complete binary tree, return its number of nodes.
Every level except possibly the last is full, and the last level is filled continuously from left to right.

示例 / Example
       1
      / \
     2   3
    / \ /
   4  5 6
答案为 6；最后一层可以缺少右侧节点，但不能跳过左侧位置。
The answer is 6. Missing positions may occur on the right of the last level, not before existing nodes.

解法一：普通递归 / Method 1: Ordinary Recursion
当前树的节点数 = 左子树节点数 + 右子树节点数 + 1。
左右子树互不重叠，加上当前根节点，正好覆盖整棵树；空树返回 0。
区别于 104 和 111：计数要把两边都算入，所以相加，不是取 max 或 min。
Count = left subtree count + right subtree count + 1.
The two subtrees are disjoint, and adding the current root covers the whole tree. An empty tree returns 0.
Unlike depth problems, count both sides by addition rather than selecting max or min.
时间 O(n)，辅助空间 O(h)。本题树完全，h=O(log n)；普通链状树则可达 O(n)。
Time O(n), auxiliary space O(h). Here completeness gives h=O(log n); a general skewed tree may require O(n).

解法二：层序遍历 / Method 2: Breadth-First Traversal
根入队，每取出一个节点就把总数加 1，并将非空孩子入队。
每个节点恰好出队一次，因此累计值就是节点总数。本题不需要分层，无需 levelSize 和内层循环。
Enqueue the root, increment the count for each dequeued node, and enqueue its non-nil children.
Each node is removed exactly once. No level grouping, saved levelSize, or inner loop is needed.
时间 O(n)，辅助空间 O(w)，w 为最大层宽，本题最坏为 O(n)。
Time O(n), auxiliary space O(w) for maximum width w, up to O(n).

解法三：利用完全二叉树判断满子树 / Method 3: Detect Perfect Subtrees
这里“满子树”指每一层都填满的 perfect binary tree。
从当前根一直沿 Left 走，得到最左路径节点数 leftHeight；一直沿 Right 走，得到最右路径节点数 rightHeight。
高度按节点数计算：空树 0，单节点 1。不要把右高度写成“右子树的最大深度”。
Here a perfect subtree has every level completely filled.
Count nodes along the all-left path as leftHeight and along the all-right path as rightHeight.
Heights count nodes: nil has height 0 and a leaf height 1. The all-right path is not the right subtree's maximum depth.

关键逻辑：为什么两边一样高，就能直接计数？ / Why Do Equal Heights Allow Direct Counting?
完全二叉树只可能在最后一层缺少右侧节点。最左路径能到最深层；若最右路径也到同一层，
说明最后一层最右端的位置都存在。因为节点必须从左连续排列，在它前面的所有位置也必然存在。
所以最后一层填满，加上上面本来就填满的层，整棵子树就是满的。
A complete tree can have missing positions only on the right of its last level.
The all-left path reaches its deepest level. If the all-right path reaches that same level, the rightmost last-level position exists.
Left-to-right completeness forces all earlier positions to exist too, so every level of this subtree is full.

高为 h 的满子树，节点数为 1+2+4+...+2^(h-1)=2^h-1，可以直接返回，不再进入内部逐节点统计。
若高度不同，当前子树不满，递归统计左右子树再加根节点；子树仍然完全，所以下一层可继续使用同一判断。
A perfect subtree of height h contains 1+2+4+...+2^(h-1)=2^h-1 nodes, so return that count without visiting its interior.
If heights differ, recurse into both children and add one. Child subtrees remain complete, so the same rule still applies.

示例推演 / Walkthrough
示例根 1 的最左路径是 1->2->4，高 3；最右路径是 1->3，高 2，不能直接套公式。
子树 2 两侧高均为 2，直接得到 2^2-1=3 个节点。
子树 3 两侧高度不同，继续得到左侧 1 个、右侧 0 个，再加它自己，共 2 个。
最终为 3+2+1=6。
At root 1, paths 1->2->4 and 1->3 have heights 3 and 2, so the formula cannot apply yet.
Subtree 2 has equal heights 2 and yields 3 nodes immediately. Subtree 3 yields 1+0+1=2 nodes.
The total is 3+2+1=6.

为什么优化后是 O(log² n)？ / Why O(log² n)?
每次计算两条边界路径最多花 O(h)。完全但不满的子树，其左右孩子中至少有一棵是满子树，能直接计数；
至多只有另一棵继续递归展开。所以每层不是两边都一直展开，总工作量上界为 O(h+(h-1)+...+1)=O(h²)。
完全二叉树 h=O(log n)，所以最坏时间 O(log² n)，递归辅助空间 O(log n)。
Each pair of boundary scans costs O(h). At least one child of a non-perfect complete tree is perfect and is counted directly;
at most the other child continues expanding recursively. The work is bounded by O(h+(h-1)+...+1)=O(h²).
Since h=O(log n), worst-case time is O(log² n), with O(log n) auxiliary call-stack space.

易错点 / Pitfalls
- 等高推出满树依赖“完全二叉树”前提。普通二叉树可能中间缺节点、两端却等高，不能套这个判断。
  Equal boundary heights imply perfection only with completeness. An arbitrary tree can have equal outer heights and interior gaps.
- Go 中 ^ 是按位异或，不是乘方。2 的 h 次方可用 1 << h 表示；计数应写成 (1 << h) - 1。
  Go's ^ is XOR, not exponentiation. Use 1 << h for 2 to the power h, giving (1 << h) - 1 nodes.
- 整数类型必须能容纳节点数及计算中的移位结果；本题给定的规模可使用 int。
  The integer type must accommodate the node count and shift calculation; int suffices for this problem's bounds.
- 前两种方法适用于任意二叉树，但时间为 O(n)；第三种利用题目前提，实现优于 O(n) 的计数。
  The first two methods work on any binary tree in O(n); the third uses completeness for sublinear counting.

练习 / Practice
三种实现按上面的编号顺序写在下方。先自己写出前两种，再推导第三种的等高判断。
The three implementations appear below in the order of the numbered methods above.
Write the first two yourself, then derive the equal-height test of the third.
*/

// 1. 普通递归：左右数量相加
func countNodes(root *TreeNode) int {
	// An empty tree has no nodes.
	// 空树没有节点。
	if root == nil {
		return 0
	}

	// Count both disjoint subtrees, then include the root.
	// 左右子树都要计数，再加上根节点。
	leftCount := countNodes(root.Left)
	rightCount := countNodes(root.Right)

	return leftCount + rightCount + 1
}

// 2. 层序遍历：每出队一个节点就计数
func countNodesIterative(root *TreeNode) int {
	// An empty tree has no nodes.
	// 空树没有节点。
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	count := 0

	for len(queue) > 0 {
		// Every node is removed and counted exactly once.
		// 每个节点恰好出队一次，也就恰好计数一次。
		node := queue[0]
		queue = queue[1:]
		count++

		// Enqueue existing children for later counting.
		// 非空孩子入队，等待后续计数。
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return count
}

// 3. 满子树公式：利用完全二叉树性质，推荐
func countNodesOptimized(root *TreeNode) int {
	// An empty subtree contributes zero nodes.
	// 空子树贡献 0 个节点。
	if root == nil {
		return 0
	}

	leftHeight, rightHeight := 0, 0

	// Count nodes along the all-left path.
	// 沿最左路径统计高度，按节点数计。
	for node := root; node != nil; node = node.Left {
		leftHeight++
	}

	// Count nodes along the all-right path.
	// 沿最右路径统计高度，不是求右子树的最大深度。
	for node := root; node != nil; node = node.Right {
		rightHeight++
	}

	// Equal boundary heights imply a perfect subtree because it is complete.
	// 因为树完全，两端等高就说明每层都填满，可以直接套公式。
	if leftHeight == rightHeight {
		return (1 << leftHeight) - 1
	}

	// Otherwise, count child subtrees using the same optimization.
	// 当前子树不满，就分别统计左右子树，继续尝试优化。
	return countNodesOptimized(root.Left) +
		countNodesOptimized(root.Right) + 1
}
