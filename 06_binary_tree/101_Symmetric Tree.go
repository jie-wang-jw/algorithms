package _6_binary_tree

/*
101. 对称二叉树 / Symmetric Tree
题目链接 / Problem link: https://leetcode.cn/problems/symmetric-tree/

题目描述 / Problem Description
给定二叉树的根节点 root，判断它是否关于中心轴左右对称。
对称要求镜像位置的节点值相同，且结构也互为镜像。
Given the root of a binary tree, determine whether it is symmetric about its center.
Both node values and structure must match at mirrored positions.

示例 / Examples
对称 / Symmetric:
        1
       / \
      2   2
     / \ / \
    3  4 4  3

不对称 / Not symmetric:
        1
       / \
      2   2
       \   \
        3   3
第二棵树虽然每层的值看似对称，但两个 3 都是右孩子，位置并不互为镜像。
In the second tree, both 3 nodes are right children. Their values match, but their positions do not mirror each other.

解法一：递归比较镜像位置（推荐） / Method 1: Recursively Compare Mirrored Positions (Recommended)
不要分别判断左子树和右子树自己是否对称；需要判断它们彼此是否互为镜像。
定义 mirror(left, right)：以这两个节点为根的子树是否互为镜像。
Do not check whether each subtree is independently symmetric. Check whether they mirror each other.
Define mirror(left, right) to mean that the two subtrees are mirror images.

1. 两个节点都为空：对应位置都没有节点，匹配成功。
   Both nil: both mirrored positions are empty, so they match.
2. 只有一个为空：一侧有节点、一侧没有，结构不匹配。
   Only one nil: one side has a node and the other does not, so their structures differ.
3. 两个都存在但值不同：镜像位置的值不匹配。
   Both exist but have different values: mirrored values differ.
4. 值相同：继续比较外侧一对和内侧一对，并要求两对都成立。
   Equal values: compare the outer pair and the inner pair; both must match.

关键逻辑：为什么交叉比较？ / Why Compare Across Sides?
镜像会把左方向变成右方向，因此左子树的左孩子应对应右子树的右孩子（外侧），
左子树的右孩子应对应右子树的左孩子（内侧）。同方向比较判断的是相同结构，不是镜像结构。
A mirror reverses left and right: the left subtree's left child matches the right subtree's right child (outer pair),
and the left subtree's right child matches the right subtree's left child (inner pair).
Comparing the same directions tests identical structure, not mirrored structure.

外侧 / Outer pair: left.Left  <-> right.Right
内侧 / Inner pair: left.Right <-> right.Left

当前两个值相同，加上外侧子树互为镜像、内侧子树互为镜像，就覆盖了这两棵树的全部结构。
所以两组递归结果用 AND 合并，任何一组失败都不能称为对称。
Equal current values plus mirrored outer and inner subtrees cover the whole pair of trees.
Combine both recursive results with AND; either failure breaks symmetry.

从哪里开始？ / Where to Start?
整棵树的根位于中心轴上，不需要找另一个根与它比较；从 root.Left 和 root.Right 开始比较即可。
若 root 为空，返回 true；单节点树的两个孩子都为空，也会返回 true。
The root lies on the axis, so compare root.Left with root.Right.
An empty tree is symmetric; a single-node tree also passes because both children are nil.

易错点 / Pitfalls
- 先处理空指针，再读取 Val、Left、Right，避免空指针访问。
  Handle nil pointers before reading Val, Left, or Right.
- 只比较左右根节点值不够，还要检查下面的结构和节点值。
  Equal child-root values alone do not establish symmetry of their descendants.
- 不要像 226 题一样交换节点。本题只判断，不需要修改输入树。
  Unlike problem 226, this is a read-only check; do not swap nodes.
- 不含空位置的层序值回文不能证明对称，第二个示例就是反例。
  Palindromic level values without nil positions do not prove symmetry, as the second example shows.

解法二：队列迭代 / Method 2: Iterative with a Queue
递归其实只在处理“成对的位置”，所以容器里保存的单位也应该是一对节点，而不是单个节点。
把 root.Left 和 root.Right 相邻入队，此后每轮取出相邻两个节点作为一对镜像位置。
两个都为空只说明这一对匹配，必须 continue 继续检查其余对，不能直接返回 true。
值相同后，按“外侧、内侧”的顺序成对入队：left.Left 与 right.Right，然后 left.Right 与 right.Left。
每轮取走两个，然后放入四个或零个，所以队列长度始终为偶数，取相邻两个不会把不同对拆开。
The recursion only ever compares position pairs, so the container should hold pairs rather than single nodes.
Enqueue root.Left and root.Right next to each other, then remove two adjacent nodes as one mirrored pair each round.
Two nils mean only that this pair matches, so continue checking the rest instead of returning true.
Once the values match, enqueue the outer pair left.Left with right.Right, then the inner pair left.Right with right.Left.
Each round removes two nodes and adds either four or none, so the length stays even and adjacent pairs are never split.

解法三：栈迭代 / Method 3: Iterative with a Stack
把队列换成栈，其余规则完全不变：仍然成对存取，仍然按外侧、内侧配对。
出栈时栈顶是后压入的那个，所以先取的是“右侧候选”，下一个才是“左侧候选”，配对方向不能弄反。
这说明本题与遍历顺序无关：需要的只是“把每一对镜像位置都检查一遍”，深度优先或广度优先都可以。
Replace the queue with a stack and keep every rule: still store pairs, still match outer with outer and inner with inner.
The stack top is the more recently pushed node, so the first value read is the right-side candidate
and the second is the left-side candidate; do not swap that direction.
This shows the problem does not depend on traversal order: it only requires visiting every mirrored pair once.

时间与空间复杂度 / Time and Space Complexity
n 为节点数，h 为树高，w 为最大层宽。三种解法最坏时间均为 O(n)，每个节点最多参与一次对应比较；
发现不匹配都可以提前返回，所以对称失败的树通常远小于这个上界。
解法一辅助空间 O(h)，来自递归调用栈；平衡树为 O(log n)，用一般树高上界可记最坏 O(n)。
解法二辅助空间 O(w)，队列同时保存的是相邻两层的成对候选；最坏 O(n)。
解法三辅助空间 O(h)，栈沿一条路径展开成对候选；最坏 O(n)。
三种解法都只返回布尔值，不新建树，也不修改节点。
For n nodes, height h, and maximum width w, all three methods take O(n) worst-case time,
comparing each node at most once, and all may return early on a mismatch, so failing trees usually do far less work.
Method 1 uses O(h) auxiliary call-stack space: O(log n) when balanced and O(n) in the worst case.
Method 2 uses O(w) auxiliary space, since the queue holds paired candidates from two adjacent levels, up to O(n).
Method 3 uses O(h) auxiliary space, expanding paired candidates along one path, up to O(n).
All three return only a boolean; none constructs or modifies a tree.

练习 / Practice
先掌握 mirror(left,right) 的成对定义，再把它改写成迭代。三种实现按上面的编号顺序写在下方。
Master the paired definition of mirror(left,right) first, then rewrite it iteratively.
The three implementations appear below in the order of the numbered methods above.
*/

// 1. 递归：交叉比较外侧与内侧，推荐
func isSymmetric(root *TreeNode) bool {
	// An empty tree is symmetric.
	// 空树是对称的。
	if root == nil {
		return true
	}

	// Check whether the two subtrees mirror each other.
	// 判断左右子树彼此是否互为镜像。
	return isMirror(root.Left, root.Right)
}

func isMirror(left, right *TreeNode) bool {
	// Both mirrored positions are empty, so they match.
	// 两个对应位置都没有节点，匹配成功。
	if left == nil && right == nil {
		return true
	}

	// After excluding both nil, either nil means a structural mismatch.
	// 已排除同时为空；此时若有一个为空，就说明结构不匹配。
	if left == nil || right == nil {
		return false
	}

	// Both nodes exist; their values must match.
	// 两个节点都存在，它们的值必须相同。
	if left.Val != right.Val {
		return false
	}

	// Mirror reflection reverses directions: left matches right.
	// 镜像会反转方向，因此左孩子要与另一侧的右孩子比较。
	return isMirror(left.Left, right.Right) &&
		isMirror(left.Right, right.Left)
}

// 2. 队列迭代：成对入队、成对出队
func isSymmetricIterative(root *TreeNode) bool {
	// Consecutive nodes form pairs of mirrored positions.
	// 队列中相邻的两个节点组成一对镜像位置。
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root.Left, root.Right)
	}

	for len(queue) > 0 {
		// The queue length is always even, so take two nodes together.
		// 队列长度始终为偶数，每次取出两个节点。
		left, right := queue[0], queue[1]
		queue = queue[2:]

		// Only this pair is confirmed; other pairs still need checking.
		// 这里只确认当前一对匹配，其他对仍需继续检查。
		if left == nil && right == nil {
			continue
		}

		// Short-circuit evaluation prevents nil pointer access.
		// 短路求值保证先排除空指针，再比较节点值。
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}

		// Enqueue the outer pair, followed by the inner pair.
		// 将外侧一对、内侧一对依次加入队列。
		queue = append(queue,
			left.Left, right.Right,
			left.Right, right.Left,
		)
	}

	// All pairs matched; an empty tree also reaches this return.
	// 所有对应位置都匹配；空树也会直接走到这里。
	return true
}

// 3. 栈迭代：把队列换成栈，配对规则不变
func isSymmetricStack(root *TreeNode) bool {
	// The root lies on the axis, so start from its two children as one pair.
	// 根位于中心轴上，因此直接把它的两个孩子作为第一对入栈。
	var stack []*TreeNode
	if root != nil {
		stack = append(stack, root.Left, root.Right)
	}

	for len(stack) > 0 {
		// The top is the later-pushed node, so it is the right-side candidate.
		// 栈顶是后压入的节点，对应右侧候选；它下面一个才是左侧候选。
		last := len(stack) - 1
		left, right := stack[last-1], stack[last]
		stack = stack[:last-1]

		// This pair matches, but the remaining pairs still need checking.
		// 这一对匹配了，但其余的对仍然必须继续检查。
		if left == nil && right == nil {
			continue
		}

		// Short-circuit evaluation rules out nil before reading Val.
		// 短路求值先排除空指针，再读取节点值。
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}

		// Push the outer pair, then the inner pair, keeping partners adjacent.
		// 先压外侧一对，再压内侧一对，让同一对始终相邻。
		stack = append(stack,
			left.Left, right.Right,
			left.Right, right.Left,
		)
	}

	// Every mirrored pair matched; an empty tree also reaches this return.
	// 所有镜像位置都匹配；空树也会直接走到这里。
	return true
}
