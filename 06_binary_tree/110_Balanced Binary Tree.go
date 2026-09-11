package _6_binary_tree

/*
110. 平衡二叉树 / Balanced Binary Tree

题目描述 / Problem Description
给定二叉树根节点 root，判断它是否为高度平衡二叉树。
高度平衡要求每个节点的左右子树高度差的绝对值都不超过 1，而不只是根节点满足条件。
空树是平衡的；高度按节点数计算，空树高 0、叶子高 1。
Given the root of a binary tree, determine whether it is height-balanced.
At every node, the left and right subtree heights must differ by at most one, not just at the root.
An empty tree is balanced. Heights count nodes: an empty tree has height 0 and a leaf height 1.

示例 / Examples
平衡 / Balanced:          不平衡 / Unbalanced:
       3                        1
      / \                      /
     9  20                    2
       /  \                  /
      15   7                3
左图根两边高度为 1、2，其他节点也满足高度差要求。右图根两边高度为 2、0，差为 2。
The left tree has root-subtree heights 1 and 2, and every other node also satisfies the rule.
The right tree has root-subtree heights 2 and 0, differing by 2.

关键逻辑 / Key Logic
平衡不要求左右节点数量相等，也不要求结构对称。
当前树平衡 = 左子树平衡 AND 右子树平衡 AND 当前左右高度差不超过 1。
只检查根会漏掉子树内部失衡。例如根的两边各挂一条三节点链，根高度差为 0，但链的顶端高度差为 2。
Balance requires neither equal node counts nor mirror symmetry.
A tree is balanced exactly when both child subtrees are balanced and their heights differ by at most one.
Checking only the root misses internal imbalance: attach a three-node chain on each side; the root's height difference is zero, but each chain is unbalanced.

解法一：自顶向下递归 / Method 1: Top-Down Recursion
用求高度函数算出当前左右高度，检查差值，再分别判断左右子树是否平衡。
逻辑直观，但父节点求高度时已经走过子树，随后检查孩子时又求同一批节点的高度，存在重复遍历。
Compute both child heights at each node, check the difference, then recursively check both subtrees.
This is direct but repeats height calculations: descendants scanned for an ancestor's height are scanned again for their own checks.
时间可用 O(n²) 作保守上界（每个节点最多求一次 O(n) 高度）；短路返回会减少实际工作，不表示每棵链都跑满平方次。
辅助空间 O(h)，求高度与判平衡的嵌套调用沿同一树路径展开。
A conservative time bound is O(n²), allowing up to O(n) height work per node; short-circuiting reduces actual work, so a chain need not realize that bound.
Auxiliary space is O(h), with nested calls following tree paths.

解法二：自底向上后序递归（推荐） / Method 2: Bottom-Up Postorder Recursion (Recommended)
把“求高度”和“检查平衡”合并。辅助函数的返回值约定：
- 非负整数：这棵子树平衡，返回它的真实高度。
- -1：这棵子树不平衡，不再把返回值当高度使用。
Combine height calculation and balance checking. The helper returns:
- A nonnegative integer: the subtree is balanced, and this is its height.
- -1: the subtree is unbalanced; this value must not be used as a height.

先递归左子树，再递归右子树；某一边返回 -1 就立即向上返回 -1。
因为只要一个后代已经失衡，整棵树就违反“每个节点都平衡”，父节点高度差再小也无法补救。
两边都正常时比较高度差，超过 1 返回 -1；否则返回 max(leftHeight,rightHeight)+1。
这样正常返回既提供父节点需要的高度，也证明这棵子树内部全部平衡，不必再次遍历验证。
Recurse into the left and right subtrees, propagating -1 immediately if either fails.
Any unbalanced descendant violates the all-nodes requirement; a small height difference at an ancestor cannot repair it.
If both are valid, return -1 for a difference greater than one; otherwise return max(leftHeight,rightHeight)+1.
A normal return supplies the height and certifies balance throughout the subtree, eliminating repeated checks.

为什么用 -1？真实高度不可能为负，因此不会与空树高度 0 或其他合法高度混淆。
必须先判断 -1，再参与高度运算；否则可能把错误标记当成普通高度。
Why -1? Real heights cannot be negative, so it cannot be confused with empty-tree height 0 or another valid height.
Check the sentinel before doing height arithmetic.
时间 O(n)，每个节点最多处理一次；辅助空间 O(h)，最坏 O(n)。
Time O(n), processing each node at most once; auxiliary space O(h), up to O(n).

解法三：后序迭代 / Method 3: Iterative Postorder
沿用之前的后序栈模板，prev 记录最近完成处理的子树根。
先沿左边压栈；查看栈顶时，如果右子树存在且还未完成，就先进入右子树。
否则左右高度已经算好，用 map[*TreeNode]int 读取高度、检查差值、保存当前高度，然后弹出当前节点。
用节点指针作 key，而不是节点值，因为不同节点可以有相同值。
空孩子在 map 中不存在，Go 返回 int 零值 0，正好对应空树高度；非空孩子必须已经处理，不能误把缺失记录当高度 0。
Use a postorder stack and prev, the root of the most recently completed subtree.
Descend left. At the stack top, visit the right subtree first if it exists and has not finished.
Otherwise both heights are available: read them from map[*TreeNode]int, check balance, store the current height, and pop the node.
Keys are node pointers, not values, because different nodes may have equal values.
A missing nil-child entry reads as zero, matching empty-tree height; a non-nil child must already have been processed.
时间平均 O(n)，按哈希操作平均 O(1) 计；辅助空间 O(n)，因为高度表保存已处理节点，不能只算 O(h) 的栈空间。
Expected time O(n), assuming average O(1) map operations; auxiliary space O(n) for the height map plus the O(h) stack.

练习 / Practice
优先掌握后序递归的返回值含义，再练习迭代。三种实现都在下方，建议先自己写一遍再对照。
Master the postorder helper's return contract first, then practice iteration.
All three implementations appear below; write your own version before comparing.
*/

// 1. 自顶向下递归：直接判断
func isBalancedTopDown(root *TreeNode) bool {
	// An empty tree is balanced.
	// 空树是平衡的。
	if root == nil {
		return true
	}

	leftHeight := maxDepth(root.Left)
	rightHeight := maxDepth(root.Right)
	diff := leftHeight - rightHeight

	// The current node must satisfy the height constraint.
	// 当前节点的左右高度差必须不超过 1。
	if diff > 1 || diff < -1 {
		return false
	}

	// Its descendants must also be balanced.
	// 当前节点满足还不够，左右子树内部也必须平衡。
	return isBalancedTopDown(root.Left) &&
		isBalancedTopDown(root.Right)
}

// 2. 自底向上递归：推荐
func isBalanced(root *TreeNode) bool {
	// Any nonnegative result means the entire tree is balanced.
	// 返回非负高度，说明整棵树平衡。
	return balancedHeight(root) != -1
}

func balancedHeight(root *TreeNode) int {
	// An empty subtree is balanced and has height 0.
	// 空子树平衡，高度为 0。
	if root == nil {
		return 0
	}

	leftHeight := balancedHeight(root.Left)

	// An unbalanced descendant makes the whole tree unbalanced.
	// 左子树内部已经失衡，整棵树就不可能平衡。
	if leftHeight == -1 {
		return -1
	}

	rightHeight := balancedHeight(root.Right)
	if rightHeight == -1 {
		return -1
	}

	// Both subtrees are balanced; now check the current node.
	// 两棵子树内部都平衡，再检查当前节点的左右高度差。
	diff := leftHeight - rightHeight
	if diff > 1 || diff < -1 {
		return -1
	}

	// Return both a valid height and an implicit balance confirmation.
	// 返回真实高度，同时表示这棵子树已通过平衡检查。
	return max(leftHeight, rightHeight) + 1
}

// 3. 后序迭代：用栈代替递归
func isBalancedIterative(root *TreeNode) bool {
	stack := []*TreeNode{}
	heights := make(map[*TreeNode]int)
	node := root

	// The root of the most recently completed subtree.
	// 最近完成处理的子树根节点。
	var prev *TreeNode

	for node != nil || len(stack) > 0 {
		// Descend left and save ancestors for later processing.
		// 先向左深入，保存之后还要处理的祖先节点。
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		// Peek without popping: the right subtree may still be pending.
		// 先查看栈顶，不急着弹出，因为右子树可能还没处理。
		node = stack[len(stack)-1]

		if node.Right != nil && node.Right != prev {
			node = node.Right
			continue
		}

		// Both children are finished; nil children have height 0.
		// 两个孩子都处理完了；空孩子对应的 map 零值恰好为 0。
		leftHeight := heights[node.Left]
		rightHeight := heights[node.Right]
		diff := leftHeight - rightHeight

		if diff > 1 || diff < -1 {
			return false
		}

		// Save the height for the parent before completing this node.
		// 保存当前高度，供父节点使用，然后完成当前节点。
		heights[node] = max(leftHeight, rightHeight) + 1
		stack = stack[:len(stack)-1]
		prev = node

		// Resume a saved ancestor instead of descending again.
		// 下一轮返回栈中的祖先，不再重复进入当前子树。
		node = nil
	}

	return true
}
