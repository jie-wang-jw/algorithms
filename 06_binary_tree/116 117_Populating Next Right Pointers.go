package _6_binary_tree

/*
116. 填充每个节点的下一个右侧节点指针 / Populating Next Right Pointers in Each Node
117. 填充每个节点的下一个右侧节点指针 II / Populating Next Right Pointers in Each Node II

题目描述 / Problem Description
116 给定完美二叉树：所有叶子在同一层，每个父节点都有两个孩子。
117 给定任意二叉树。两者都要把每个 next 指向同层右侧节点，没有则置空。
116 is a perfect binary tree: every parent has two children and all leaves share a depth.
117 is any binary tree. In both problems, set next to the next right node on the same level, or nil.
复杂度记号：n 为节点数，h 为树高，w 为最大层宽；辅助空间不含返回结果。
Notation: n nodes, height h, maximum width w; auxiliary space excludes returned results.
*/

// 1. Level-order queue: the next node in the captured level is the next-right pointer. Used by 116 and 117.
// 1. 队列层序：当前层固定区间里的下一个节点就是 next；116 与 117 共用。
// Time: O(n), Space: O(w) for the queue.
// 时间复杂度：O(n)，空间复杂度：O(w)，由队列产生。
//
// 固定 levelSize 后，同层非末节点的 Next 指向出队后的 queue[0]；本层末节点不能指向下一层。
// Freeze levelSize; each nonfinal node points to the next queue front, while the level's final node must not point into the next level.
// 适用于普通二叉树；此实现不重置层末 Next，沿用题目“初始所有 Next 为 nil”的前提。
// Works for general binary trees; this implementation leaves final Next fields unchanged, assuming they initially are nil.
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if i != levelSize-1 {
				node.Next = queue[0]
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}

// connectII is 117: the 102 article uses the same queue logic as 116.
// connectII 对应 117：102 文章与 116 使用同一套层序逻辑。
// Time: O(n), Space: O(w) for the queue.
// 时间复杂度：O(n)，空间复杂度：O(w)，由队列产生。
//
// 固定 levelSize 后，同层非末节点的 Next 指向出队后的 queue[0]；本层末节点不能指向下一层。
// Freeze levelSize; each nonfinal node points to the next queue front, while the level's final node must not point into the next level.
// 适用于普通二叉树；此实现不重置层末 Next，沿用题目“初始所有 Next 为 nil”的前提。
// Works for general binary trees; this implementation leaves final Next fields unchanged, assuming they initially are nil.
func connectII(root *Node) *Node {
	return connect(root)
}

// 2. Preorder wiring for 116: left to right, right to the next parent's left. Perfect trees only.
// 2. 116 的前序搭线：左连右，右连下一个父节点的左孩子；只适用于完美二叉树。
// Time: O(n), Space: O(h) for the recursion stack.
// 时间复杂度：O(n)，空间复杂度：O(h)，由递归栈产生。
//
// 仅适用于完美二叉树：左孩子连接右兄弟，右孩子连接父节点 Next 的左孩子；普通缺口树不满足此规则。
// Require a perfect tree: link left to its sibling and right to the next parent's left child; arbitrary gaps break this rule.
// 先连好父层，才能借父节点 Next 跨子树连下一层；调用时所有 Next 按题意初始为 nil。
// Parent-level links must exist before using Next to cross subtrees below; assume all Next links start nil.
func connectRecursive(root *Node) *Node {
	connectTraversal(root)
	return root
}

// Use the next pointer already built on the current level, then recurse left and right.
// 使用当前层已经连好的 next，再递归左右子树。
// Time: O(n), Space: O(h).
// 时间复杂度：O(n)，空间复杂度：O(h)。
//
// 仅适用于完美二叉树：左孩子连接右兄弟，右孩子连接父节点 Next 的左孩子；普通缺口树不满足此规则。
// Require a perfect tree: link left to its sibling and right to the next parent's left child; arbitrary gaps break this rule.
// 先连好父层，才能借父节点 Next 跨子树连下一层；调用时所有 Next 按题意初始为 nil。
// Parent-level links must exist before using Next to cross subtrees below; assume all Next links start nil.
func connectTraversal(cur *Node) {
	if cur == nil {
		return
	}
	if cur.Left != nil {
		cur.Left.Next = cur.Right
	}
	if cur.Right != nil {
		if cur.Next != nil {
			cur.Right.Next = cur.Next.Left
		} else {
			cur.Right.Next = nil
		}
	}
	connectTraversal(cur.Left)
	connectTraversal(cur.Right)
}

// 3. Constant extra space for 116: walk each level through next and wire the children below.
// 3. 116 的常数额外空间：沿着 next 走完当前层，同时连接下一层孩子。
// Time: O(n), Space: O(1).
// 时间复杂度：O(n)，空间复杂度：O(1)。
//
// 仅适用于完美二叉树：左孩子连接右兄弟，右孩子连接父节点 Next 的左孩子；普通缺口树不满足此规则。
// Require a perfect tree: link left to its sibling and right to the next parent's left child; arbitrary gaps break this rule.
// 先连好父层，才能借父节点 Next 跨子树连下一层；调用时所有 Next 按题意初始为 nil。
// Parent-level links must exist before using Next to cross subtrees below; assume all Next links start nil.
func connectConstant(root *Node) *Node {
	if root == nil {
		return root
	}
	for cur := root; cur.Left != nil; cur = cur.Left {
		for node := cur; node != nil; node = node.Next {
			node.Left.Next = node.Right
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}
	return root
}
