package _6_binary_tree

/*
501. 二叉搜索树中的众数 / Find Mode in Binary Search Tree

题目描述 / Problem Description
返回 BST 中出现次数最多的值。众数可能有多个，返回顺序不限。空树返回空结果。
Return the values that appear most frequently in a BST. There may be several modes; any order is acceptable.
An empty tree has no modes.

示例 / Examples
      1
       \
        2
       /
      2
中序 1,2,2；2 出现两次，众数是 [2]。
Inorder 1,2,2; 2 appears twice, so the mode list is [2].

若每个值都只出现一次，则所有值都是众数。
If every value appears once, every value is a mode.

关键逻辑 / Key Logic
BST 中序把相同值排在一起，所以可以在遍历时统计连续相等值的长度，而不必先建哈希表。
count 是当前值的连续次数，maxCount 是目前见到的最大次数。
次数等于 maxCount 时把当前值加入答案；次数更大则清空答案并更新 maxCount。
BST inorder groups equal values together, so a linear scan of adjacent equals replaces a frequency map.
count is the run length of the current value; maxCount is the best run seen so far.
Equal the best: append. Beat the best: replace the answer and raise maxCount.

解法一：中序递归（推荐） / Method 1: Inorder Recursion (Recommended)
用前驱节点判断当前值是否延续上一段。空前驱表示第一个节点，count 从 1 开始。
Use the predecessor to decide whether the current value continues a run. A nil predecessor means the first node, so count starts at 1.

解法二：中序迭代 / Method 2: Iterative Inorder
同样的计数规则，用栈模拟中序。
The same counting rule with an explicit inorder stack.

时间与空间复杂度 / Time and Space Complexity
n 为节点数，h 为树高，m 为众数个数。两版时间 O(n)。
除输出外辅助空间 O(h)；输出 O(m)。
For n nodes, height h, and m modes, both take O(n) time.
Auxiliary space is O(h) besides O(m) output.
*/

// 1. Inorder recursion (recommended): equal values form a consecutive run, so update the mode list from that run length.
// 1. 中序递归：推荐；相同值在中序中连成一段，用这一段的长度更新众数列表。
// Time: O(n), Space: O(h) auxiliary plus O(m) output.
// 时间复杂度：O(n)，空间复杂度：辅助 O(h)，输出另占 O(m)。
func findMode(root *TreeNode) []int {
	result := []int{}
	maxCount := 0
	count := 0
	var prev *TreeNode

	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		// Visit root between left and right so equal values stay consecutive.
		// 在左右之间访问根，相同值才会在中序中连成一段。
		updateMode(node, &prev, &count, &maxCount, &result)
		traverse(node.Right)
	}

	traverse(root)
	return result
}

// 2. Iterative inorder: the same run-length counting with an explicit stack.
// 2. 中序迭代：同样按连续相同值计数，改成显式栈。
// Time: O(n), Space: O(h) auxiliary plus O(m) output.
// 时间复杂度：O(n)，空间复杂度：辅助 O(h)，输出另占 O(m)。
func findModeIterative(root *TreeNode) []int {
	result := []int{}
	maxCount := 0
	count := 0
	var prev *TreeNode
	stack := []*TreeNode{}
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		updateMode(cur, &prev, &count, &maxCount, &result)
		cur = cur.Right
	}
	return result
}

// Run-length update: grow the current value's count, then append or replace the mode list.
// 连续计数更新：累加当前值的次数，再追加或重建众数列表。
// Time: O(1) amortized, Space: O(1) besides the result slice.
// 时间复杂度：均摊 O(1)，空间复杂度：除结果切片外 O(1)。
func updateMode(node *TreeNode, prev **TreeNode, count, maxCount *int, result *[]int) {
	// Same value as predecessor: extend the current run; otherwise start a new run of length 1.
	// 与前驱相同则延续当前段；否则从 1 开始新的一段。
	if *prev != nil && node.Val == (*prev).Val {
		*count++
	} else {
		*count = 1
	}

	if *count == *maxCount {
		// Tie for the best frequency: another mode.
		// 与当前最大次数持平：又一个众数。
		*result = append(*result, node.Val)
	} else if *count > *maxCount {
		// Strictly better frequency: reset the mode list.
		// 次数创新高：清空旧答案，只保留当前值。
		*maxCount = *count
		*result = []int{node.Val}
	}
	*prev = node
}
