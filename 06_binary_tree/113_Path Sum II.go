package _6_binary_tree

/*
113. 路径总和 II / Path Sum II

题目描述 / Problem Description
给定二叉树根节点 root 和整数 targetSum，返回所有从根到叶子、节点值之和等于 targetSum 的路径。
每条路径保存节点值，整体返回 [][]int，顺序不限。空树返回空结果。
Given a binary tree root and an integer targetSum, return all root-to-leaf paths whose values sum to targetSum.
Return node values as [][]int in any order. An empty tree returns no paths.

示例 / Example
       5
      / \
     4   8
    /   / \
   11  13  4
  / \     / \
 7   2   5   1
targetSum=22，结果为 [[5,4,11,2],[5,8,4,5]]。
For targetSum=22, return [[5,4,11,2],[5,8,4,5]].

与 112、257 的联系 / Relation to Problems 112 and 257
112 只判断存在性，找到一条即可结束；113 要收集全部匹配路径，不能提前结束整个搜索。
257 收集所有根到叶子的路径；113 增加路径和筛选，保存整数切片而非字符串。
112 checks existence and may stop at one match; 113 must collect all matches.
257 collects every root-to-leaf path; 113 filters by sum and stores integer slices rather than strings.

解法一：递归回溯（推荐） / Method 1: Recursive Backtracking (Recommended)
path 保存当前根到节点的路径；递归参数 remaining 表示从当前节点开始还需凑出的和。
进入节点时，把当前值加入 path，并从 remaining 中扣除它。
只有到达叶子且 remaining=0，才能复制 path 加入结果；非叶子分别探索左右子树。
离开节点前，统一删除 path 最后一个元素，恢复父节点的路径，无论这一支是否找到答案。
path tracks the root-to-node route; remaining is the sum still required starting at the current node.
On entry, append the current value and subtract it from remaining.
Only at a leaf with remaining=0, copy path into the result; otherwise explore both children.
Before leaving, remove the final path item to restore the parent route, regardless of whether a match was found.

关键逻辑：为什么必须复制 path？ / Why Copy the Path?
Go 切片包含底层数组的引用，直接把 path 加入结果只复制切片描述，不复制元素。
回溯后再追加兄弟节点可能覆盖同一个数组，导致已保存的答案被修改。
所以收集时新建等长切片并 copy，使每个答案拥有独立元素；缩短 path 长度不能代替复制。
A Go slice references a backing array. Appending path to the result copies its descriptor, not its elements.
Later sibling appends may overwrite that same array and corrupt earlier answers.
Allocate an equal-length slice and copy the elements when collecting. Shortening path alone does not isolate saved results.

为什么恢复路径，却不用恢复 remaining？ / Why Undo the Path but Not remaining?
path 是闭包共享的可变切片，下一分支必须回到共同父节点的路径，所以要删除当前节点。
remaining 是按值传递的整数，每次调用有自己的副本，不会改变父调用给另一分支的目标。
path is shared mutable state, so remove the current node to let a sibling inherit only the common parent route.
remaining is an integer passed by value; each call owns its copy and cannot change the parent's target for another branch.

解法二：栈迭代，保存独立路径 / Method 2: Iterative DFS with Independent Paths
同步维护节点栈、累计和栈和路径栈，相同下标描述同一个待处理状态。
出栈时，叶子且和匹配就收集；否则为每个非空孩子新建路径，复制父路径后加入孩子值，并保存新的累计和。
每个状态拥有独立路径，不用手动回溯；代价是即使路径最终不匹配，也会复制中间前缀。
Maintain synchronized node, sum, and path stacks. Matching positions describe one pending state.
Collect matching leaves; otherwise allocate a separate path per child, copy the parent prefix, append the child value, and store the new sum.
Independent paths need no explicit undo, but copying occurs even along paths that ultimately do not match.

其他迭代选择 / Other Iterative Options
相同状态可改用队列 BFS，从队首取出；判断规则不变，但宽树会同时保留较多路径。
也可保存父指针及累计和，到匹配叶子才重建路径，减少中间复制；初学优先掌握直接的回溯写法。
The same states can use a BFS queue, retaining more simultaneous paths on wide trees.
Alternatively, parent pointers and sums allow reconstruction only at matching leaves; prioritize direct backtracking when learning.

易错点 / Pitfalls
- 和相等但不是叶子不能收集；叶子必须左右孩子都为空。
  A matching sum at an internal node is insufficient; a leaf has both children nil.
- 值可以是负数，不能因为 remaining<0 就剪枝。
  Negative values mean remaining<0 is not a valid pruning rule.
- 找到一条后仍要恢复路径、继续其他分支，不能使用 112 的短路 OR 来收集全部答案。
  Restore the path and explore other branches after a match; 112's short-circuit OR cannot collect all answers.
- 不同节点路径可能有相同值序列，不能用集合擅自去重。
  Distinct node paths can have identical value sequences; do not deduplicate them with a set.

时间与空间复杂度 / Time and Space Complexity
n 为节点数，h 为树高，w 为最大层宽，S 为所有返回路径中整数的总数。
递归回溯：时间 O(n+S)，辅助空间 O(h)，输出另占 O(S)。节点各处理一次，答案复制成本为 S。
独立路径栈迭代：时间上界 O(nh)，辅助空间上界 O(h²)，输出另占 O(S)。
独立路径 BFS：时间上界 O(nh)，辅助空间上界 O(wh)，输出另占 O(S)。
For n nodes, height h, maximum width w, and S total output elements:
Recursive backtracking takes O(n+S) time, O(h) auxiliary space, plus O(S) output.
Independent-path stack DFS has O(nh) time and O(h²) auxiliary-space bounds, plus O(S) output.
Independent-path BFS has O(nh) time and O(wh) auxiliary-space bounds, plus O(S) output.

练习 / Practice
本文件仅保留中英文题解，不提供实现或函数骨架。代码在聊天中展示，供自行手写练习。
This file contains explanations only, without implementations or skeletons. Code is shown in chat for manual practice.
*/

// 1. 递归回溯：推荐 O(n+S)O(h)
func pathSum(root *TreeNode, targetSum int) [][]int {
	result := [][]int{}
	path := []int{}

	var traverse func(*TreeNode, int)
	traverse = func(node *TreeNode, remaining int) {
		if node == nil {
			return
		}

		// Include this node in the current path.
		// 把当前节点加入路径，并扣除它贡献的值。
		path = append(path, node.Val)
		remaining -= node.Val

		if node.Left == nil && node.Right == nil {
			if remaining == 0 {
				// Copy the elements so later backtracking cannot change this answer.
				// 复制元素，避免后续回溯修改已经保存的答案。
				saved := make([]int, len(path))
				copy(saved, path)
				result = append(result, saved)
			}
		} else {
			// Search both branches; one match does not finish the whole search.
			// 两边都要搜索，找到一条不能结束整题。
			traverse(node.Left, remaining)
			traverse(node.Right, remaining)
		}

		// Restore the parent's path before visiting another branch.
		// 返回前恢复父节点的路径，避免干扰其他分支。
		path = path[:len(path)-1]
	}

	traverse(root, targetSum)
	return result
}

// 2. 栈迭代：每个分支保存独立路径 O(nh) 上界O(h²) 上界
func pathSumIterative(root *TreeNode, targetSum int) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	nodes := []*TreeNode{root}
	sums := []int{root.Val}
	paths := [][]int{{root.Val}}

	for len(nodes) > 0 {
		// Pop the node together with its sum and independent path.
		// 节点、累计和、独立路径一起出栈。
		last := len(nodes) - 1
		node, sum, path := nodes[last], sums[last], paths[last]
		nodes = nodes[:last]
		sums = sums[:last]
		paths = paths[:last]

		if node.Left == nil && node.Right == nil {
			if sum == targetSum {
				// This leaf path is independent and will not be modified again.
				// 这条叶子路径独立保存，之后不会再修改，可以直接收集。
				result = append(result, path)
			}
			continue
		}

		// Push right before left so the left branch is processed first.
		// 先右后左入栈，使左分支先处理。
		for _, child := range []*TreeNode{node.Right, node.Left} {
			if child == nil {
				continue
			}

			// Give each child its own backing array.
			// 每个孩子分配独立的底层数组，避免兄弟路径相互覆盖。
			nextPath := make([]int, len(path)+1)
			copy(nextPath, path)
			nextPath[len(path)] = child.Val

			nodes = append(nodes, child)
			sums = append(sums, sum+child.Val)
			paths = append(paths, nextPath)
		}
	}

	return result
}
