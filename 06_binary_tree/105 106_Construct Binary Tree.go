package _6_binary_tree

/*
106. 从中序与后序遍历序列构造二叉树
Construct Binary Tree from Inorder and Postorder Traversal

题目描述 / Problem Description
给定同一棵二叉树的中序遍历 inorder 和后序遍历 postorder，构造并返回这棵树。
题目保证节点值互不相同，两组遍历合法且对应同一棵树。
Given the inorder and postorder traversals of the same binary tree, construct and return the tree.
Node values are unique, and the traversals are valid and describe the same tree.

示例 / Example
inorder   = [9,3,15,20,7]
postorder = [9,15,7,20,3]
       3
      / \
     9  20
       /  \
      15   7
复杂度记号：n 为节点数，h 为树高，w 为最大层宽；辅助空间不含返回结果。
Notation: n nodes, height h, maximum width w; auxiliary space excludes returned results.
*/

// 1. Linear root lookup and slice splitting: take the last postorder value as root, then use its inorder index k to split both arrays.
// 1. 递归切片：先理解这一版；后序末尾是根，用中序位置 k 同时切开两组遍历。
// Time: O(n²) worst case, Space: O(h) plus O(n) for the output tree.
// 时间复杂度：最坏 O(n²)，空间复杂度：O(h)，输出树另占 O(n)。
//
// 要求两遍历合法且节点值互异。后序末项是根，根在中序的位置 k 将左右子树分开；左子树大小决定后序分界。
// Require valid traversals with unique values; the last postorder value is the root and its inorder position determines subtree sizes.
// 所有切片/下标范围均左闭右开，右子树后序必须排除末尾根；哈希版只加速定位根，不改变划分规则。
// Ranges are half-open and the right postorder range excludes the final root; the map version only speeds up locating that split.
func buildTreePostOrder(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	rootValue := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootValue}

	k := 0
	for inorder[k] != rootValue {
		k++
	}

	root.Left = buildTreePostOrder(inorder[:k], postorder[:k])

	root.Right = buildTreePostOrder(
		inorder[k+1:],
		postorder[k:len(postorder)-1],
	)

	return root
}

// 2. Linear root lookup and index ranges: split half-open inorder and postorder intervals without creating child slices.
// 2. 下标区间递归：用左闭右开下标切分中序和后序，切割点仍靠线性查找。
// Time: O(n²) worst case, Space: O(h) plus O(n) for the output tree.
// 时间复杂度：最坏 O(n²)，空间复杂度：O(h)，输出树另占 O(n)。
//
// 两组 [left,right) 描述同一棵子树；根是 postorder[postRight-1]，leftSize=delimiterIndex-inLeft。
// Both half-open ranges describe one subtree; the root is postorder[postRight-1] and leftSize=delimiterIndex-inLeft.
// 后序左段取前 leftSize 项，右段取剩余项但排除根；要求输入遍历合法、等长且值互异。
// Take the first leftSize postorder entries for the left subtree and the remainder except root for the right; require valid unique traversals.
func buildTreePostOrderIndex(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	return buildPostorderIndex(inorder, 0, len(inorder), postorder, 0, len(postorder))
}

// Half-open inorder [inLeft, inRight) and postorder [postLeft, postRight) describe the same subtree.
// 左闭右开的中序 [inLeft, inRight) 与后序 [postLeft, postRight) 表示同一棵子树。
// Time: O(n²) worst case, Space: O(h).
// 时间复杂度：最坏 O(n²)，空间复杂度：O(h)。
//
// 两组 [left,right) 描述同一棵子树；根是 postorder[postRight-1]，leftSize=delimiterIndex-inLeft。
// Both half-open ranges describe one subtree; the root is postorder[postRight-1] and leftSize=delimiterIndex-inLeft.
// 后序左段取前 leftSize 项，右段取剩余项但排除根；要求输入遍历合法、等长且值互异。
// Take the first leftSize postorder entries for the left subtree and the remainder except root for the right; require valid unique traversals.
func buildPostorderIndex(inorder []int, inLeft, inRight int, postorder []int, postLeft, postRight int) *TreeNode {
	if postLeft == postRight {
		return nil
	}

	rootValue := postorder[postRight-1]
	root := &TreeNode{Val: rootValue}

	if postRight-postLeft == 1 {
		return root
	}

	delimiterIndex := inLeft
	for delimiterIndex < inRight && inorder[delimiterIndex] != rootValue {
		delimiterIndex++
	}

	leftSize := delimiterIndex - inLeft

	root.Left = buildPostorderIndex(
		inorder, inLeft, delimiterIndex,
		postorder, postLeft, postLeft+leftSize,
	)

	root.Right = buildPostorderIndex(
		inorder, delimiterIndex+1, inRight,
		postorder, postLeft+leftSize, postRight-1,
	)

	return root
}

// 3. Index map and ranges (recommended): locate the root in average O(1), then split half-open intervals by leftSize = k-inLeft.
// 3. 哈希表 + 区间递归：优化查找；平均 O(1) 定位根，再用 leftSize = k-inLeft 切分左闭右开区间。
// Time: O(n) expected, Space: O(n+h) = O(n) for the map and the recursion stack.
// 时间复杂度：平均 O(n)，空间复杂度：O(n+h)=O(n)，由索引表和递归栈产生。
//
// 要求两遍历合法且节点值互异。后序末项是根，根在中序的位置 k 将左右子树分开；左子树大小决定后序分界。
// Require valid traversals with unique values; the last postorder value is the root and its inorder position determines subtree sizes.
// 所有切片/下标范围均左闭右开，右子树后序必须排除末尾根；哈希版只加速定位根，不改变划分规则。
// Ranges are half-open and the right postorder range excludes the final root; the map version only speeds up locating that split.
func buildTreePostOrderOptimized(inorder []int, postorder []int) *TreeNode {
	positions := make(map[int]int, len(inorder))
	for i, value := range inorder {
		positions[value] = i
	}

	var build func(int, int, int, int) *TreeNode
	build = func(inLeft, inRight, postLeft, postRight int) *TreeNode {
		if inLeft == inRight {
			return nil
		}

		rootValue := postorder[postRight-1]
		k := positions[rootValue]
		leftSize := k - inLeft

		root := &TreeNode{Val: rootValue}

		root.Left = build(
			inLeft, k,
			postLeft, postLeft+leftSize,
		)

		root.Right = build(
			k+1, inRight,
			postLeft+leftSize, postRight-1,
		)

		return root
	}

	return build(0, len(inorder), 0, len(postorder))
}

/*
105. 从前序与中序遍历序列构造二叉树
Construct Binary Tree from Preorder and Inorder Traversal

题目描述 / Problem Description
给定同一棵二叉树的前序遍历 preorder 和中序遍历 inorder，构造并返回这棵树。
题目保证节点值互不相同，输入合法。105 与 106 的题解放在本文件中对照学习。
Given valid preorder and inorder traversals of the same binary tree with unique values, reconstruct the tree.
This file includes both 105 and 106 for comparison.

核心区别 / Core Difference
105 前序：根、左子树、右子树，所以第一个元素是根。
106 后序：左子树、右子树、根，所以最后一个元素是根。
两题都依靠中序“左、根、右”确定左右分界和左子树大小，不能直接从中间平分。
105 preorder is root-left-right, so the first element is the root.
106 postorder is left-right-root, so the last element is the root.
Both use inorder's left-root-right layout to determine child boundaries and subtree sizes, not midpoint splitting.

示例拆分 / Example Split
preorder = [3,9,20,15,7]
inorder  = [9,3,15,20,7]
前序第一个值 3 是根。在中序中，3 左边只有 9，所以左子树有 1 个节点。
前序跳过根后取 1 个元素 [9] 作为左子树，其余 [20,15,7] 是右子树。
左子树：前序 [9]，中序 [9]；右子树：前序 [20,15,7]，中序 [15,20,7]。
子树继续按同样规则构造，得到根 3、左孩子 9、右孩子 20，20 的左右孩子为 15、7。
The first preorder value 3 is the root. Only 9 precedes it in inorder, so the left subtree has one node.
Skip the preorder root, take [9] for the left subtree, and use [20,15,7] for the right.
Left: preorder [9], inorder [9]. Right: preorder [20,15,7], inorder [15,20,7].
Recursing gives root 3 with children 9 and 20; node 20 has children 15 and 7.
*/

// 1. Linear lookup and slice splitting: take preorder[0] as root, then use its inorder index k as the left-subtree size.
// 1. 递归切片：前序第一个值是根，用中序位置 k 作为左子树节点数去切前序。
// Time: O(n²) worst case, Space: O(h) plus O(n) for the output tree.
// 时间复杂度：最坏 O(n²)，空间复杂度：O(h)，输出树另占 O(n)。
//
// 要求合法、等长且值互异的遍历；前序首项是根，中序根位置给出 leftSize，前序根后的 leftSize 项属于左树。
// Require valid equal-length unique traversals; preorder starts with root and the inorder split determines the left subtree's size.
// 右树使用余下前序和根右侧中序；半开区间均排除根，哈希版只把找根位置从扫描改为查表。
// Use the remaining preorder and right inorder ranges for the right subtree; exclude the root, with a map only accelerating lookup.
func buildTreePreorder(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	rootValue := preorder[0]
	root := &TreeNode{Val: rootValue}

	k := 0
	for inorder[k] != rootValue {
		k++
	}

	root.Left = buildTreePreorder(
		preorder[1:k+1],
		inorder[:k],
	)

	root.Right = buildTreePreorder(
		preorder[k+1:],
		inorder[k+1:],
	)

	return root
}

// 2. Linear lookup and index ranges: split half-open preorder and inorder intervals without creating child slices.
// 2. 下标区间递归：用左闭右开下标切分前序和中序，切割点仍靠线性查找。
// Time: O(n²) worst case, Space: O(h) plus O(n) for the output tree.
// 时间复杂度：最坏 O(n²)，空间复杂度：O(h)，输出树另占 O(n)。
//
// 两组半开区间表示同一子树；根为 preorder[preLeft]，从中序根位置算 leftSize，前序左段从 preLeft+1 开始。
// Two half-open ranges represent one subtree; root is preorder[preLeft], with leftSize from inorder and left preorder starting at preLeft+1.
// 每个子范围都排除根，空范围返回 nil；输入须是合法、等长且值互异的遍历。
// Exclude the root from both child ranges; empty ranges return nil. Require valid equal-length unique traversals.
func buildTreePreorderIndex(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	return buildPreorderIndex(preorder, 0, len(preorder), inorder, 0, len(inorder))
}

// Half-open preorder [preLeft, preRight) and inorder [inLeft, inRight) describe the same subtree.
// 左闭右开的前序 [preLeft, preRight) 与中序 [inLeft, inRight) 表示同一棵子树。
// Time: O(n²) worst case, Space: O(h).
// 时间复杂度：最坏 O(n²)，空间复杂度：O(h)。
//
// 两组半开区间表示同一子树；根为 preorder[preLeft]，从中序根位置算 leftSize，前序左段从 preLeft+1 开始。
// Two half-open ranges represent one subtree; root is preorder[preLeft], with leftSize from inorder and left preorder starting at preLeft+1.
// 每个子范围都排除根，空范围返回 nil；输入须是合法、等长且值互异的遍历。
// Exclude the root from both child ranges; empty ranges return nil. Require valid equal-length unique traversals.
func buildPreorderIndex(preorder []int, preLeft, preRight int, inorder []int, inLeft, inRight int) *TreeNode {
	if preLeft == preRight {
		return nil
	}

	rootValue := preorder[preLeft]
	root := &TreeNode{Val: rootValue}

	if preRight-preLeft == 1 {
		return root
	}

	delimiterIndex := inLeft
	for delimiterIndex < inRight && inorder[delimiterIndex] != rootValue {
		delimiterIndex++
	}

	leftSize := delimiterIndex - inLeft

	root.Left = buildPreorderIndex(
		preorder, preLeft+1, preLeft+1+leftSize,
		inorder, inLeft, delimiterIndex,
	)

	root.Right = buildPreorderIndex(
		preorder, preLeft+1+leftSize, preRight,
		inorder, delimiterIndex+1, inRight,
	)

	return root
}

// 3. Index map and ranges: the root is preorder[preLeft]; skip it and take leftSize = k-inLeft preorder elements.
// 3. 哈希表 + 区间递归：根是 preorder[preLeft]，跳过它再取 leftSize = k-inLeft 个前序元素。
// Time: O(n) expected, Space: O(n+h) = O(n) for the map and the recursion stack.
// 时间复杂度：平均 O(n)，空间复杂度：O(n+h)=O(n)，由索引表和递归栈产生。
//
// 要求合法、等长且值互异的遍历；前序首项是根，中序根位置给出 leftSize，前序根后的 leftSize 项属于左树。
// Require valid equal-length unique traversals; preorder starts with root and the inorder split determines the left subtree's size.
// 右树使用余下前序和根右侧中序；半开区间均排除根，哈希版只把找根位置从扫描改为查表。
// Use the remaining preorder and right inorder ranges for the right subtree; exclude the root, with a map only accelerating lookup.
func buildTreePreorderOptimized(preorder []int, inorder []int) *TreeNode {
	positions := make(map[int]int, len(inorder))
	for i, value := range inorder {
		positions[value] = i
	}

	var build func(int, int, int, int) *TreeNode
	build = func(preLeft, preRight, inLeft, inRight int) *TreeNode {
		if preLeft == preRight {
			return nil
		}

		rootValue := preorder[preLeft]
		k := positions[rootValue]

		leftSize := k - inLeft
		root := &TreeNode{Val: rootValue}

		root.Left = build(
			preLeft+1, preLeft+1+leftSize,
			inLeft, k,
		)

		root.Right = build(
			preLeft+1+leftSize, preRight,
			k+1, inRight,
		)

		return root
	}

	return build(0, len(preorder), 0, len(inorder))
}
