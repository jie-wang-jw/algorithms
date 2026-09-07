package _6_binary_tree

import "container/list"

/*
226. 翻转二叉树 / Invert Binary Tree

题目描述 / Problem Description
给定二叉树的根节点 root，翻转这棵二叉树，并返回它的根节点。
翻转指得到左右镜像，不是上下颠倒，也不是仅交换节点值。
Given the root of a binary tree, invert the tree and return its root.
Inversion produces a left-right mirror, not an upside-down tree or merely swapped values.

解题思路 / Solution Approach
本文件提供五种解法 / Five implementations:
- invertTree：前序递归，先交换再递归。 / Preorder recursion: swap before recursion.
- invertTreePostorderRecursive：后序递归，先递归再交换。 / Postorder recursion: recurse before swapping.
- invertTreePreorderIterative：入栈前交换，出栈后处理另一边。 / Iterative preorder: swap before pushing, visit the other side after popping.
- invertTreePostorderIterative：用 prev 判断右子树是否完成。 / Iterative postorder: prev tracks completion of the right subtree.
- invertTreeLevelOrder：队列逐节点交换，不必区分层。 / BFS: swap each dequeued node without grouping levels.

下面先解释前序递归 / Preorder recursion explained first:
使用递归。invertTree(root) 的含义是：翻转以 root 为根的整棵子树，并返回这个根。
1. root 为 nil：空树无需处理，返回 nil。
2. 交换 root.Left 和 root.Right，把左右两棵子树整体换位置。
3. 分别递归翻转交换后的左右子树，处理它们内部的左右关系。
4. 返回 root；原地修改连接，根节点本身不变。
Use recursion. invertTree(root) mirrors the entire subtree rooted at root and returns that root.
1. Return nil for an empty subtree.
2. Swap root.Left and root.Right to exchange the two whole subtrees.
3. Recursively invert both children to mirror the internal structure of each subtree.
4. Return root; links change in place, but the root node itself stays the same.

关键逻辑：为什么这样做 / Why This Works
整棵树的镜像 = 根 + 原右子树的镜像（放左边）+ 原左子树的镜像（放右边）。
只交换根的左右孩子，只完成了两棵子树的位置互换，它们内部还没有翻转。
递归恰好补上内部翻转，所以三部分组合后就是整棵树的镜像。
The mirror consists of the root, the mirrored original right subtree on the left,
and the mirrored original left subtree on the right. Swapping children alone relocates
the subtrees but does not mirror their interiors. Recursion completes those interiors.

示例 / Example
       4                 4
      / \               / \
     2   7     ->      7   2
    / \ / \           / \ / \
   1  3 6  9         9  6 3  1

先在 4 处交换：左边变成 7，右边变成 2，但 7 的孩子仍是 6、9，2 的孩子仍是 1、3。
再递归到 7，把 6、9 换成 9、6；递归到 2，把 1、3 换成 3、1，才得到完整镜像。
At 4, swapping puts 7 on the left and 2 on the right, but their children remain 6,9 and 1,3.
Recursion swaps them to 9,6 and 3,1 respectively, completing the mirror.

易错点 / Pitfalls
交换后 root.Left 是原右子树，root.Right 是原左子树；分别调用一次，所以没有重复或遗漏。
递归调用直接修改节点连接，因此这里无需再次接收返回值；不是忽略了翻转结果。
叶子的两个孩子都是 nil，交换不改变它，随后两次空树调用返回，递归就会结束。
After swapping, root.Left is the original right subtree and vice versa; each is processed exactly once.
Recursive calls mutate links directly, so their return values need not be reassigned here.
A leaf has two nil children; swapping changes nothing and both recursive calls terminate immediately.

时间与空间复杂度 / Time and Space Complexity
n 为节点数，h 为树高。时间 O(n)：每个节点只交换一次。
辅助空间 O(h)：递归栈最多保存一条根到叶的调用路径；平衡树 O(log n)，链状树最坏 O(n)。
没有新建树，但递归栈仍占空间，因此不能把辅助空间写成 O(1)。
For n nodes and height h, time is O(n), swapping each node's children once.
Auxiliary space is O(h) for the recursion path: O(log n) when balanced, O(n) in the worst skewed case.
Reusing the tree does not eliminate call-stack space, so auxiliary space is not O(1).
其余三种深度优先版本也都是时间 O(n)、辅助空间 O(h)。后序迭代节点至多入栈、出栈两次，仍为线性。
层序版本时间 O(n)、辅助空间 O(w)，w 为最大层宽；队列可能混合相邻两层，数量仍为 O(w)。
The other three DFS versions also take O(n) time and O(h) auxiliary space. Iterative postorder pushes/pops each node at most twice.
BFS takes O(n) time and O(w) auxiliary space for maximum width w, even while the queue mixes adjacent levels.
*/

func invertTree(root *TreeNode) *TreeNode {
	// An empty subtree is already its own mirror.
	// 空子树的镜像仍是空子树，也是递归的停止条件。
	if root == nil {
		return nil
	}

	// Exchange whole subtrees, not just the child values.
	// 交换的是两棵子树的入口，不只是两个孩子的值。
	root.Left, root.Right = root.Right, root.Left

	// Mirror the interiors of both relocated subtrees, exactly once each.
	// 两棵子树虽然换了位置，内部仍需翻转；左右各处理一次。
	invertTree(root.Left)
	invertTree(root.Right)

	// The root stays in place; its subtree has now been fully mirrored.
	// 根的位置不变，以它为根的整棵子树已经完成镜像翻转。
	return root
}

/*
后序递归 / Postorder recursion
先把原左右子树内部都翻转，再交换它们的位置，同样得到“原右子树的镜像 + 根 + 原左子树的镜像”的布局。
递归返回意味着整棵子树已经完成，不是只访问了它的根节点。
Mirror both original subtrees internally before exchanging their positions.
A recursive return means the entire subtree is finished, not just its root.
*/
func invertTreePostorderRecursive(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// Finish both subtree interiors before exchanging their positions.
	// 先完成两棵子树内部的翻转，再交换它们的位置。
	invertTreePostorderRecursive(root.Left)
	invertTreePostorderRecursive(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

/*
前序迭代 / Iterative preorder
虽然使用沿左侧深入的模板，但交换发生在入栈之前，是先处理根的前序操作。
交换后 Left 是原右子树，Right 是原左子树。先深入前者，出栈后再进入后者，原两棵子树恰好各处理一次。
栈记录的是“回来后还要处理另一边”的节点；出栈不再交换，否则会把刚才的交换撤销。
Despite the left-descent template, swapping before pushing makes this preorder processing.
After swapping, Left is the original right subtree and Right the original left.
Descend into one, then visit the other after popping.
The stack remembers pending right-subtree work. Do not swap again on pop,
which would undo the inversion.
*/
func invertTreePreorderIterative(root *TreeNode) *TreeNode {
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			// Swap on first arrival, then save the node for the other subtree.
			// 首次到达时交换，保存当前节点以便稍后处理另一棵子树。
			node.Left, node.Right = node.Right, node.Left
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// The swapped left subtree is done; continue with the swapped right.
		// 交换后的左子树已完成，接着进入交换后的右子树。
		node = node.Right
	}
	return root
}

/*
后序迭代 / Iterative postorder
prev 是最近完成整棵子树翻转的根，不是最近碰到的节点。
从栈中取出 node 时左子树已完成；Right=nil 表示没有右子树，Right=prev 表示右子树刚完成。
两边都完成才交换当前节点。右子树内部翻转不会改变它的根指针，所以 Right=prev 的判断仍然有效。
否则把当前节点重新压栈，记住“右边完成后还要回来处理我”，再转入右子树。
完成后 node=nil 让下一轮直接返回栈中的祖先，不要再次深入已经翻转的子树。
prev is the root of the most recently completed subtree, not the last node encountered.
On pop, the left subtree is complete. Right=nil means no right subtree; Right=prev means it has just finished.
Only then swap this node. Inverting a child subtree preserves its root pointer, so the pointer comparison remains valid.
Otherwise push the node back to resume after its right subtree. Set node=nil after completion to resume an ancestor instead of descending again.
*/
func invertTreePostorderIterative(root *TreeNode) *TreeNode {
	stack := []*TreeNode{}
	node := root
	// Track subtree completion, not mere arrival.
	// 记录子树完成状态，而不是访问到达状态。
	var prev *TreeNode
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Right == nil || node.Right == prev {
			// Both subtrees are complete; finish their parent.
			// 两棵子树都完成了，现在完成它们的父节点。
			node.Left, node.Right = node.Right, node.Left
			prev = node
			node = nil
		} else {
			// Resume this parent after completing its right subtree.
			// 保存父节点，右子树完成后还要回来处理。
			stack = append(stack, node)
			node = node.Right
		}
	}
	return root
}

/*
层序遍历 / Breadth-first traversal
每个节点出队时交换一次，再将两个孩子入队。交换只改变孩子的位置，不影响两个孩子都被处理。
无需按层记录结果，所以不用保存 length 或增加内层循环；逐节点处理就足够。
queue.Front() 取得队首元素，Remove 删除它并返回其中的值，.(*TreeNode) 把该值断言为存入的节点指针。
Swap each dequeued node once, then enqueue both children. Their order changes, but both are still processed.
No per-level result is needed, so a saved level length and inner loop are unnecessary.
Front obtains the queue element, Remove returns its stored value, and .(*TreeNode) asserts that value to the stored node-pointer type.
*/
func invertTreeLevelOrder(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		node.Left, node.Right = node.Right, node.Left
		// Relocation alone does not mirror the interiors; enqueue both for later work.
		// 子树换位置后内部仍需翻转，因此把两个孩子都入队继续处理。
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return root
}
