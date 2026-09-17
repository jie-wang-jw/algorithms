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
复杂度记号：n 为节点数，h 为树高，w 为最大层宽；辅助空间不含返回结果。
Notation: n nodes, height h, maximum width w; auxiliary space excludes returned results.
*/

// 1. Inorder recursion (recommended): equal values form a consecutive run, so update the mode list from that run length.
// 1. 中序递归：推荐；相同值在中序中连成一段，用这一段的长度更新众数列表。
// Time: O(n), Space: O(h) auxiliary plus O(m) output.
// 时间复杂度：O(n)，空间复杂度：辅助 O(h)，输出另占 O(m)。
//
// BST 中序非递减，相同值必连续；count 是当前连续段长度，prev 是中序前驱，maxCount 是已见最大频次。
// BST inorder is nondecreasing, making equal values consecutive; count tracks the current run, prev the prior node, maxCount the best run.
// count 超过历史最大值就清空旧答案并替换，相等则追加；只在访问节点时更新，不能在入栈时更新。
// Replace answers when count exceeds the maximum and append on equality; update at inorder visitation, not when pushing the node.
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
//
// BST 中序非递减，相同值必连续；count 是当前连续段长度，prev 是中序前驱，maxCount 是已见最大频次。
// BST inorder is nondecreasing, making equal values consecutive; count tracks the current run, prev the prior node, maxCount the best run.
// count 超过历史最大值就清空旧答案并替换，相等则追加；只在访问节点时更新，不能在入栈时更新。
// Replace answers when count exceeds the maximum and append on equality; update at inorder visitation, not when pushing the node.
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
//
// BST 中序非递减，相同值必连续；count 是当前连续段长度，prev 是中序前驱，maxCount 是已见最大频次。
// BST inorder is nondecreasing, making equal values consecutive; count tracks the current run, prev the prior node, maxCount the best run.
// count 超过历史最大值就清空旧答案并替换，相等则追加；只在访问节点时更新，不能在入栈时更新。
// Replace answers when count exceeds the maximum and append on equality; update at inorder visitation, not when pushing the node.
func updateMode(node *TreeNode, prev **TreeNode, count, maxCount *int, result *[]int) {
	if *prev != nil && node.Val == (*prev).Val {
		*count++
	} else {
		*count = 1
	}

	if *count == *maxCount {
		*result = append(*result, node.Val)
	} else if *count > *maxCount {
		*maxCount = *count
		*result = []int{node.Val}
	}
	*prev = node
}
