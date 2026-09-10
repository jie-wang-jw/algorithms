package _6_binary_tree

/*
112. 路径总和 / Path Sum

题目描述 / Problem Description
给定二叉树根节点 root 和整数 targetSum，判断是否存在一条从根到叶子的路径，
使路径上所有节点值之和等于 targetSum。叶子节点的左右孩子都为空。
空树不存在根到叶子的路径，即使 targetSum 为 0 也应返回 false。
Given the root of a binary tree and an integer targetSum, determine whether a root-to-leaf path
has a sum equal to targetSum. A leaf has no children.
An empty tree has no root-to-leaf path, even when targetSum is zero.

示例 / Example
       5
      / \
     4   8
    /   / \
   11  13  4
  / \      \
 7   2      1
targetSum=22 时返回 true，路径为 5->4->11->2，总和 5+4+11+2=22。
targetSum=9 不能因为 5+4=9 就成功：4 不是叶子，必须继续走到实际叶子再判断。
For targetSum=22, return true via 5->4->11->2.
For targetSum=9, prefix 5->4 is not enough: node 4 is not a leaf.

解法一：递归传递剩余目标 / Method 1: Recursion with a Remaining Target
定义 hasPathSum(node,remaining)：是否存在从 node 到某个叶子的路径，其和等于 remaining；进入时 remaining 还包含当前节点要贡献的值。
空节点返回 false；非空节点先从 remaining 中减去当前值。到叶子时剩余为 0 才成功。
非叶子继续把扣除后的剩余目标传给左右孩子；任意一边成功即可，所以用 OR 合并，而不是 AND。
Define hasPathSum(node,remaining) as whether a node-to-leaf path sums to remaining, including node's own value.
Return false for nil. Subtract the current value; at a leaf, success means the remainder is zero.
Otherwise pass the remainder to both children. Either successful branch suffices, so combine with OR rather than AND.

为什么减去当前值？原目标要求“当前节点值 + 后续路径和 = remaining”，移项后，后续只需要凑 remaining-当前值。
沿示例目标依次变化：22 --经过5--> 17 --经过4--> 13 --经过11--> 2 --经过叶子2--> 0。
两个条件同时成立：到达叶子，并且剩余为 0。
Why subtract? The equation is current value + remaining path sum = remaining, so the child path must sum to remaining-current value.
In the example: 22 --visit 5--> 17 --visit 4--> 13 --visit 11--> 2 --visit leaf 2--> 0.
Both requirements must hold: the node is a leaf and the remainder is zero.

为什么分支不会互相干扰？Go 的 int 参数按值传递，左递归修改自己的目标副本，不会改变父调用给右孩子的目标。
因此无需手动加回当前值；若使用共享可变状态，才需要对应的回溯恢复。
Why are branches independent? Go passes int arguments by value. The left call changes its own copy, not the parent's value used for the right call.
No explicit restoration is needed here; shared mutable state would require undoing changes.

解法二：栈迭代 DFS / Method 2: Iterative DFS
节点栈与路径和栈同步操作，相同下标保存一个节点和从根到该节点（包含该节点）的累计和。
出栈后，如果是叶子且和等于 targetSum，立即返回 true；否则将孩子与“当前和+孩子值”入栈。
节点及其累计和必须一起入栈、一起出栈，不能用一个全局 sum 累加所有访问到的节点，那会把不同分支混在一起。
Use synchronized node and sum stacks. Matching positions store a node and the root-to-node sum including that node.
On pop, return true only for a leaf whose sum equals targetSum. Otherwise push children with current sum + child value.
Push and pop each node with its own sum. A single running total across all visited nodes would mix separate branches.

解法三：队列 BFS / Method 3: Breadth-First Search
把两个栈换成同步队列，从队首取出节点与累计和，判断规则完全相同。
本题不问深度，无需 levelSize 或按层循环。遇到任意一条合法路径即可返回 true，全部检查后仍未找到则返回 false。
Use synchronized queues instead, removing a node and its sum from the front with the same checks.
No depth or level grouping is needed. Return true for any matching root-to-leaf path; return false if all candidates fail.

易错点 / Pitfalls
- 路径必须从根开始，并终止于叶子；不能在中间节点和刚好相等时提前成功。
  The path must start at the root and end at a leaf; a matching internal prefix is insufficient.
- 值可能为负数，不能因为当前和大于目标或剩余小于 0 就剪枝。
  Values may be negative; do not prune simply because the current sum exceeds the target or the remainder is negative.
  例如 5->-3，目标 2，经过 5 时虽然超出，继续走却能成功。
  For 5->-3 and target 2, the prefix exceeds the target but the complete path succeeds.
- 叶子判断必须同时检查左右为空；只缺一侧不表示路径结束。
  Both children must be nil for a leaf; a single missing child is not an endpoint.
- 不需要像 257 一样构造完整路径字符串，本题只判断是否存在，保存累计和或剩余目标即可。
  Unlike problem 257, no path strings are needed; a cumulative sum or remaining target is sufficient.

时间与空间复杂度 / Time and Space Complexity
n 为节点数，h 为树高，w 为最大层宽。三种写法最坏时间 O(n)，找到合法路径时可提前结束。
递归辅助空间 O(h)；栈迭代辅助空间 O(h)，每个待处理节点额外保存一个整数；队列辅助空间 O(w)。
这些空间最坏都可达 O(n)。均不修改树，布尔返回值占 O(1)。
For n nodes, height h, and maximum width w, all methods take O(n) worst-case time and may exit early on success.
Recursion and stack DFS use O(h) auxiliary space; queue BFS uses O(w). Each pending sum is one integer.
All space bounds are at most O(n). The tree is unchanged, and the boolean result uses O(1) space.

练习 / Practice
本文件只提供中英文题解，不包含实现或函数骨架。优先掌握“剩余目标”的递归定义，再练习迭代。
This file contains explanations only, without implementations or skeletons. Start with the remaining-target recursion, then practice iteration.
*/

// 1. 递归：传递“还差多少” O(n)O(h)
func hasPathSum(root *TreeNode, targetSum int) bool {
	// An empty tree has no root-to-leaf path.
	// 空树没有根到叶子的路径，即使目标为 0 也不能成功。
	if root == nil {
		return false
	}

	// The current value is used; the children must supply the remainder.
	// 当前值已经计入，后面的路径只需要凑剩余目标。
	remaining := targetSum - root.Val

	// Success requires both reaching a leaf and matching the sum.
	// 必须同时满足：已经到叶子，并且剩余目标为 0。
	if root.Left == nil && root.Right == nil {
		return remaining == 0
	}

	// Either branch may supply a valid path.
	// 任意一边找到合法路径即可，所以用 ||。
	return hasPathSum(root.Left, remaining) ||
		hasPathSum(root.Right, remaining)
}

// 2. 栈迭代：节点和累计和一起保存 O(n)O(h)
func hasPathSumIterative(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	// Matching positions store a node and its root-to-node sum.
	// 相同下标保存一个节点和从根到该节点的累计和。
	nodes := []*TreeNode{root}
	sums := []int{root.Val}

	for len(nodes) > 0 {
		last := len(nodes) - 1
		node, sum := nodes[last], sums[last]
		nodes = nodes[:last]
		sums = sums[:last]

		// A matching internal node is not enough; it must be a leaf.
		// 中间节点的和相等还不够，必须是叶子。
		if node.Left == nil && node.Right == nil &&
			sum == targetSum {
			return true
		}

		// Extend this path independently for each child.
		// 每个孩子分别延伸当前路径，保存自己的累计和。
		if node.Right != nil {
			nodes = append(nodes, node.Right)
			sums = append(sums, sum+node.Right.Val)
		}
		if node.Left != nil {
			nodes = append(nodes, node.Left)
			sums = append(sums, sum+node.Left.Val)
		}
	}

	return false
}

// 3. 队列 BFS O(n)O(w)
func hasPathSumBFS(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	nodes := []*TreeNode{root}
	sums := []int{root.Val}

	for len(nodes) > 0 {
		// Remove a node together with its own path sum.
		// 节点与它对应的路径和一起出队。
		node, sum := nodes[0], sums[0]
		nodes = nodes[1:]
		sums = sums[1:]

		// Return as soon as a complete matching path is found.
		// 找到完整且和匹配的根到叶子路径即可返回。
		if node.Left == nil && node.Right == nil &&
			sum == targetSum {
			return true
		}

		if node.Left != nil {
			nodes = append(nodes, node.Left)
			sums = append(sums, sum+node.Left.Val)
		}
		if node.Right != nil {
			nodes = append(nodes, node.Right)
			sums = append(sums, sum+node.Right.Val)
		}
	}

	return false
}
