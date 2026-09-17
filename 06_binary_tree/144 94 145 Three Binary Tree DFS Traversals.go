package _6_binary_tree

/*
144. 二叉树的前序遍历 / Binary Tree Preorder Traversal
题目描述 / Problem Description
给定二叉树根节点 root，返回节点值的前序遍历结果。
Given the root of a binary tree, return its preorder traversal.
前序遍历顺序：
Preorder traversal order:
根 → 左 → 右
Root → Left → Right
复杂度记号：n 为节点数，h 为树高，w 为最大层宽；辅助空间不含返回结果。
Notation: n nodes, height h, maximum width w; auxiliary space excludes returned results.
*/

/*
题目描述 / Problem Description

返回二叉树的前序遍历结果。
Return the preorder traversal of a binary tree.
*/

// 1. Dedicated preorder iteration: visit the node, then push right before left so LIFO processes the left child first.
// 1. 前序专用迭代：先访问根，再先压右孩子、后压左孩子，利用后进先出让左孩子先出栈。
// Time: O(n), Space: O(h) for the stack.
// 时间复杂度：O(n)，空间复杂度：O(h)，由显式栈产生。
//
// 前序是根→左→右；栈后进先出，所以访问根后先压右再压左，才能先弹出左子树。
// Preorder is root→left→right; push right before left because the stack is LIFO.
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := make([]int, 0)
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		last := len(stack) - 1
		node := stack[last]
		stack = stack[:last]

		result = append(result, node.Val)

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return result
}

// 2. Preorder recursion: record the root before the two child calls.
// 2. 前序递归：在两次递归之前记录根。
// Time: O(n), Space: O(h) for the recursion stack.
// 时间复杂度：O(n)，空间复杂度：O(h)，由递归栈产生。
//
// 三种 DFS 的区别只在写入根值的时机：前序在两个递归前，中序在其间，后序在其后；空节点直接返回。
// The DFS orders differ only in when the root is emitted: before, between, or after child calls; nil ends recursion.
func preorderTraversalRecursive(root *TreeNode) []int {
	result := make([]int, 0)

	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		result = append(result, node.Val)
		traverse(node.Left)
		traverse(node.Right)
	}

	traverse(root)
	return result
}

/*
94. 二叉树的中序遍历 / Binary Tree Inorder Traversal

题目描述 / Problem Description
给定二叉树根节点 root，返回按照“左、根、右”顺序访问的节点值。
Given the root of a binary tree, return its node values in left-root-right order.
*/

// 1. Dedicated inorder iteration: push along the left spine, then visit on pop before entering the right subtree.
// 1. 中序专用迭代：沿左侧压栈，弹出时记录根，再进入右子树。
// Time: O(n), Space: O(h) for the stack.
// 时间复杂度：O(n)，空间复杂度：O(h)，由显式栈产生。
//
// 先沿左链入栈，直到无左孩子；弹栈时左子树已完成，才能访问当前节点，再转入其右子树。
// Push the left spine; popping means the left subtree is complete, so visit the node and then its right subtree.
// current 为空但栈未空时还有祖先待处理，所以外层条件用“或”。
// A nil current with a nonempty stack still has ancestors pending, hence the outer OR condition.
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	current := root

	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		last := len(stack) - 1
		current = stack[last]
		stack = stack[:last]

		result = append(result, current.Val)
		current = current.Right
	}

	return result
}

// 2. Inorder recursion: record the root between the two child calls.
// 2. 中序递归：在两次递归之间记录根。
// Time: O(n), Space: O(h) for the recursion stack.
// 时间复杂度：O(n)，空间复杂度：O(h)，由递归栈产生。
//
// 三种 DFS 的区别只在写入根值的时机：前序在两个递归前，中序在其间，后序在其后；空节点直接返回。
// The DFS orders differ only in when the root is emitted: before, between, or after child calls; nil ends recursion.
func inorderTraversalRecursive(root *TreeNode) []int {
	result := make([]int, 0)

	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		traverse(node.Left)
		result = append(result, node.Val)
		traverse(node.Right)
	}

	traverse(root)
	return result
}

/*
145. 二叉树的后序遍历 / Binary Tree Postorder Traversal

题目描述 / Problem Description
给定二叉树根节点 root，返回按照“左、右、根”顺序访问的节点值。
Given the root of a binary tree, return its node values in left-right-root order.
*/

// 1. Reversal-based postorder iteration: produce root-right-left with a stack, then reverse the result to left-right-root.
// 1. 后序迭代：先用栈得到“根、右、左”，再整体反转为“左、右、根”。
// Time: O(n), Space: O(h) for the stack.
// 时间复杂度：O(n)，空间复杂度：O(h)，由显式栈产生。
//
// 先压左再压右，得到根→右→左的访问结果；整体反转恰好是左→右→根的后序。
// Push left before right to collect root→right→left; reversing the full result yields left→right→root.
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := make([]int, 0)
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		last := len(stack) - 1
		node := stack[last]
		stack = stack[:last]

		result = append(result, node.Val)

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	for left, right := 0, len(result)-1; left < right; left, right = left+1, right-1 {
		result[left], result[right] = result[right], result[left]
	}

	return result
}

// 2. Postorder recursion: record the root after both child calls return.
// 2. 后序递归：在两次递归之后记录根。
// Time: O(n), Space: O(h) for the recursion stack.
// 时间复杂度：O(n)，空间复杂度：O(h)，由递归栈产生。
//
// 三种 DFS 的区别只在写入根值的时机：前序在两个递归前，中序在其间，后序在其后；空节点直接返回。
// The DFS orders differ only in when the root is emitted: before, between, or after child calls; nil ends recursion.
func postorderTraversalRecursive(root *TreeNode) []int {
	result := make([]int, 0)

	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		traverse(node.Left)
		traverse(node.Right)
		result = append(result, node.Val)
	}

	traverse(root)
	return result
}

// 3. Unified preorder iteration: a following nil means record-only; push right, left, then root+nil.
// 3. 前序统一迭代：节点后紧跟 nil 表示只记录；入栈次序为右、左、根+nil。
// Time: O(n), Space: O(h) for the marked stack.
// 时间复杂度：O(n)，空间复杂度：O(h)，由带 nil 标记的栈产生。
//
// nil 是“输出其下一栈顶节点”的标记，不代表树中的空孩子；无标记节点只负责安排后续访问。
// nil marks an emit action for the node beneath it, not a missing child; unmarked nodes schedule future work.
// 按目标遍历顺序的逆序压栈，并把 node,nil 当作一次输出动作，区分“经过节点”和“记录节点”。
// Push actions in reverse traversal order; node,nil is one emit action, separating expansion from output.
func preorderTraversalUnified(root *TreeNode) []int {
	result := make([]int, 0)
	stack := []*TreeNode{}

	if root != nil {
		stack = append(stack, root)
	}

	for len(stack) > 0 {
		last := len(stack) - 1
		node := stack[last]
		stack = stack[:last]

		if node == nil {
			last = len(stack) - 1
			node = stack[last]
			stack = stack[:last]
			result = append(result, node.Val)
			continue
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		stack = append(stack, node, nil)
	}

	return result
}

// 3. Unified inorder iteration: the same nil-marker rule, pushing right, root+nil, then left.
// 3. 中序统一迭代：同一条 nil 标记规则，入栈次序为右、根+nil、左。
// Time: O(n), Space: O(h) for the marked stack.
// 时间复杂度：O(n)，空间复杂度：O(h)，由带 nil 标记的栈产生。
//
// nil 是“输出其下一栈顶节点”的标记，不代表树中的空孩子；无标记节点只负责安排后续访问。
// nil marks an emit action for the node beneath it, not a missing child; unmarked nodes schedule future work.
// 按目标遍历顺序的逆序压栈，并把 node,nil 当作一次输出动作，区分“经过节点”和“记录节点”。
// Push actions in reverse traversal order; node,nil is one emit action, separating expansion from output.
func inorderTraversalUnified(root *TreeNode) []int {
	result := make([]int, 0)
	stack := []*TreeNode{}

	if root != nil {
		stack = append(stack, root)
	}

	for len(stack) > 0 {
		last := len(stack) - 1
		node := stack[last]
		stack = stack[:last]

		if node == nil {
			last = len(stack) - 1
			node = stack[last]
			stack = stack[:last]
			result = append(result, node.Val)
			continue
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		stack = append(stack, node, nil)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return result
}

// 3. Unified postorder iteration: the same nil-marker rule, pushing root+nil, right, then left, so no final reversal.
// 3. 后序统一迭代：同一条 nil 标记规则，入栈次序为根+nil、右、左，因此不必再反转结果。
// Time: O(n), Space: O(h) for the marked stack.
// 时间复杂度：O(n)，空间复杂度：O(h)，由带 nil 标记的栈产生。
//
// nil 是“输出其下一栈顶节点”的标记，不代表树中的空孩子；无标记节点只负责安排后续访问。
// nil marks an emit action for the node beneath it, not a missing child; unmarked nodes schedule future work.
// 按目标遍历顺序的逆序压栈，并把 node,nil 当作一次输出动作，区分“经过节点”和“记录节点”。
// Push actions in reverse traversal order; node,nil is one emit action, separating expansion from output.
func postorderTraversalUnified(root *TreeNode) []int {
	result := make([]int, 0)
	stack := []*TreeNode{}

	if root != nil {
		stack = append(stack, root)
	}

	for len(stack) > 0 {
		last := len(stack) - 1
		node := stack[last]
		stack = stack[:last]

		if node == nil {
			last = len(stack) - 1
			node = stack[last]
			stack = stack[:last]
			result = append(result, node.Val)
			continue
		}

		stack = append(stack, node, nil)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return result
}

// 4. Morris preorder: thread the predecessor's right pointer and record on the first arrival, then restore the tree.
// 4. 前序 Morris：借用前驱右指针当线索，第一次到达就记录，走完左子树后再拆掉线索。
// Time: O(n), Space: O(1) auxiliary; right pointers are rewritten temporarily.
// 时间复杂度：O(n)，空间复杂度：辅助 O(1)，会临时改写右指针。
//
// predecessor 是左子树最右节点；第一次遇见时把其空 Right 临时连回 current，代替递归返回路径。
// predecessor is the left subtree's rightmost node; temporarily point its nil Right to current as a return thread.
// 第二次沿线索回到 current 时拆掉临时边，再转右树；前序第一次记录根，中序第二次记录根。
// On returning via the thread, remove it and visit the right subtree; preorder emits on first arrival, inorder on return.
// 无线索可建时直接记录并转右；完整遍历会恢复原树，遍历中会暂时修改指针。
// Without a left subtree, emit and move right; full traversal restores the tree after temporary pointer changes.
func preorderTraversalMorris(root *TreeNode) []int {
	result := make([]int, 0)
	current := root

	for current != nil {
		if current.Left == nil {
			result = append(result, current.Val)
			current = current.Right
			continue
		}

		predecessor := current.Left

		for predecessor.Right != nil && predecessor.Right != current {
			predecessor = predecessor.Right
		}

		if predecessor.Right == nil {
			result = append(result, current.Val)

			predecessor.Right = current
			current = current.Left
			continue
		}

		predecessor.Right = nil
		current = current.Right
	}

	return result
}

// 4. Morris inorder: the same threading walk, but record on the second arrival after the left subtree is finished.
// 4. 中序 Morris：走法与前序相同，但第二次到达、左子树走完后才记录。
// Time: O(n), Space: O(1) auxiliary; right pointers are rewritten temporarily.
// 时间复杂度：O(n)，空间复杂度：辅助 O(1)，会临时改写右指针。
//
// predecessor 是左子树最右节点；第一次遇见时把其空 Right 临时连回 current，代替递归返回路径。
// predecessor is the left subtree's rightmost node; temporarily point its nil Right to current as a return thread.
// 第二次沿线索回到 current 时拆掉临时边，再转右树；前序第一次记录根，中序第二次记录根。
// On returning via the thread, remove it and visit the right subtree; preorder emits on first arrival, inorder on return.
// 无线索可建时直接记录并转右；完整遍历会恢复原树，遍历中会暂时修改指针。
// Without a left subtree, emit and move right; full traversal restores the tree after temporary pointer changes.
func inorderTraversalMorris(root *TreeNode) []int {
	result := make([]int, 0)
	current := root

	for current != nil {
		if current.Left == nil {
			result = append(result, current.Val)
			current = current.Right
			continue
		}

		predecessor := current.Left
		for predecessor.Right != nil && predecessor.Right != current {
			predecessor = predecessor.Right
		}

		if predecessor.Right == nil {
			predecessor.Right = current
			current = current.Left
			continue
		}

		predecessor.Right = nil
		result = append(result, current.Val)
		current = current.Right
	}

	return result
}
