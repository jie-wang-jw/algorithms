package _6_binary_tree

/*
116. 填充每个节点的下一个右侧节点指针 / Populating Next Right Pointers in Each Node
117. 填充每个节点的下一个右侧节点指针 II / Populating Next Right Pointers in Each Node II

题目描述 / Problem Description
116 给定完美二叉树：所有叶子在同一层，每个父节点都有两个孩子。
117 给定任意二叉树。两者都要把每个 next 指向同层右侧节点，没有则置空。
116 is a perfect binary tree: every parent has two children and all leaves share a depth.
117 is any binary tree. In both problems, set next to the next right node on the same level, or nil.

解法一：队列层序（116 与 117 相同） / Method 1: Level-Order Queue (Same for 116 and 117)
102 文章对 117 的说明是：代码、逻辑与 116 的层序写法相同。
每层固定 levelSize，不是最后一个节点就把 next 指向队首尚未出队的同层节点。
The 102 article treats 117 as the same queue logic as 116.
Capture levelSize; if the node is not last on the level, point next at the current queue front.

解法二：前序递归搭线（仅 116） / Method 2: Preorder Wiring (116 Only)
完美二叉树中：左孩子 next 指向右孩子；右孩子 next 指向 cur.next 的左孩子。
前序保证访问 cur 时，上一层的 next 已经连好，跨父节点的线才能接到。
In a perfect tree, left.next = right, and right.next = cur.next.left.
Preorder visits cur after the previous level's next links exist, so the cross-parent link is available.

解法三：沿已连好的 next 走（仅 116，O(1) 额外空间） / Method 3: Walk Established Next Links (116 Only)
从每层最左节点出发，用已经连好的 next 横穿该层，连接下一层的左右孩子。
不使用队列。题目把递归栈不算额外空间；本写法迭代，辅助指针 O(1)。
Start at the leftmost node of each level and walk next to wire the next level's children.
No queue. The problem ignores recursion-stack space; this iterative version uses O(1) extra pointers.

时间与空间复杂度 / Time and Space Complexity
n 为节点数，w 为最大层宽，h 为树高。
解法一时间 O(n)，队列 O(w)。解法二时间 O(n)，递归栈 O(h)。解法三时间 O(n)，辅助 O(1)。
For n nodes, width w, and height h:
method 1 is O(n) time and O(w) queue space; method 2 is O(n) time and O(h) stack space;
method 3 is O(n) time and O(1) extra pointers.
*/

// 1. Level-order queue: the next node in the captured level is the next-right pointer. Used by 116 and 117.
// 1. 队列层序：当前层固定区间里的下一个节点就是 next；116 与 117 共用。
// Time: O(n), Space: O(w) for the queue.
// 时间复杂度：O(n)，空间复杂度：O(w)，由队列产生。
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
func connectII(root *Node) *Node {
	return connect(root)
}

// 2. Preorder wiring for 116: left to right, right to the next parent's left. Perfect trees only.
// 2. 116 的前序搭线：左连右，右连下一个父节点的左孩子；只适用于完美二叉树。
// Time: O(n), Space: O(h) for the recursion stack.
// 时间复杂度：O(n)，空间复杂度：O(h)，由递归栈产生。
func connectRecursive(root *Node) *Node {
	connectTraversal(root)
	return root
}

// Use the next pointer already built on the current level, then recurse left and right.
// 使用当前层已经连好的 next，再递归左右子树。
// Time: O(n), Space: O(h).
// 时间复杂度：O(n)，空间复杂度：O(h)。
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
