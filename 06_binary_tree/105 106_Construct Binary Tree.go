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

关键逻辑 / Key Logic
中序是“左子树、根、右子树”，后序是“左子树、右子树、根”。
后序最后一个值一定是当前子树的根，但后序本身不能直接告诉我们左右各有几个节点。
在中序中找到这个根，左边全部属于左子树，右边全部属于右子树，于是得到左子树节点数 leftSize。
因为后序也先放左子树，取前 leftSize 个就是左子树后序；剩下除根之外的部分是右子树后序。
对子树继续使用同样规则，直到区间为空返回 nil。
Inorder is left-root-right, while postorder is left-right-root.
The final postorder value identifies the subtree root but does not directly give child-subtree sizes.
Find the root in inorder: everything before it belongs to the left subtree and everything after it to the right.
The resulting leftSize tells us that the first leftSize postorder values describe the left subtree.
The rest, excluding the root, describes the right subtree. Repeat recursively until an interval is empty.

示例拆分 / Example Split
根为 3。中序根位置左边只有 9，因此左子树大小为 1。
左子树：中序 [9]，后序 [9]。
右子树：中序 [15,20,7]，后序 [15,7,20]；其根为 20，再拆成 15 和 7。
The root is 3. Only 9 precedes it in inorder, so the left subtree has one node.
Left subtree: inorder [9], postorder [9].
Right subtree: inorder [15,20,7], postorder [15,7,20]; its root is 20, with children 15 and 7.

解法一：线性查找根并切分 / Method 1: Linear Root Lookup and Slice Splitting
每次取 postorder 最后一个元素作为根，在 inorder 中扫描根的位置 k。
左右中序分别为 inorder[:k]、inorder[k+1:]；左右后序分别为 postorder[:k]、postorder[k:len(postorder)-1]。
这里 k 也是当前左子树节点数，因为传入的切片已经从当前子树开头开始。
递归返回的是子树根指针，分别接到新节点的 Left 和 Right 上，最后返回新根。
Take the final postorder value, then scan inorder for its position k.
Split inorder into [:k] and [k+1:], and postorder into [:k] and [k:len(postorder)-1].
Here k is also the left-subtree size because each input slice begins at the current subtree.
Attach the returned child-root pointers to the new root's Left and Right, then return that root.

解法二：线性查找根 + 下标区间 / Method 2: Linear Root Lookup and Index Ranges
思路与解法一相同：后序末尾是根，在中序里线性扫描切割点。
区别是不再切出子切片，而是用左闭右开区间 [inLeft,inRight) 与 [postLeft,postRight) 描述同一棵子树。
空区间 postLeft==postRight 返回 nil；区间只剩一个节点时直接返回根。
切割点 delimiterIndex 是全局中序下标，左子树大小是 delimiterIndex-inLeft，不能直接拿 delimiterIndex 去切后序。
文章强调区间不变量必须全程一致：四个子区间都坚持左闭右开，才不会把根算进孩子，或漏掉节点。
Same idea as method 1: the last postorder value is the root, found by a linear inorder scan.
Instead of creating child slices, describe the same subtree with half-open ranges [inLeft,inRight) and [postLeft,postRight).
Return nil when postLeft==postRight; return the root immediately when the interval has one node.
delimiterIndex is a global inorder index, so the left-subtree size is delimiterIndex-inLeft, not delimiterIndex.
Keep the half-open invariant on every child range so the root is excluded exactly once.

解法三：哈希表定位根 + 下标区间（推荐） / Method 3: Index Map and Ranges (Recommended)
先保存“节点值 -> 中序下标”，后续平均 O(1) 找到根的位置，避免每层重复扫描。
区间切分与解法二相同；只是把线性查找换成哈希表。
空区间 inLeft==inRight 返回 nil；当前根为 postorder[postRight-1]。
设根的全局中序位置为 k，左子树大小必须算 k-inLeft，而不是直接用 k。
Precompute value -> inorder index for average O(1) root lookup.
Range splitting matches method 2; only the delimiter search becomes a map lookup.
Return nil when inLeft==inRight; the root value is postorder[postRight-1].
For global inorder root index k, the left-subtree size is k-inLeft, not k.

区间对应 / Corresponding Ranges
左中序 / Left inorder:   [inLeft, k)
右中序 / Right inorder:  [k+1, inRight)
左后序 / Left postorder: [postLeft, postLeft+leftSize)
右后序 / Right postorder:[postLeft+leftSize, postRight-1)
左后序必须恰好取 leftSize 个；右后序从其后开始，到当前根之前结束。
两个中序区间排除根位置 k，两个后序区间排除最后的根，确保不重复构造、不漏节点。
Take exactly leftSize elements for left postorder; right postorder starts afterward and ends before the root.
Excluding k from inorder and the final root from postorder prevents duplicate construction and omissions.

易错点 / Pitfalls
- 不能直接把两组遍历从中间平分，左右子树大小不一定相同。
  Do not split the arrays at their midpoints; child subtrees need not have equal sizes.
- 节点值互不相同，才可以用唯一中序位置确定左右分界；这里不需要处理不合法输入。
  Unique node values allow one inorder position per root; invalid inputs are outside the problem's contract.
- 本文为左右子树传独立区间，先构造左还是右都可以。只有共享一个从后往前的后序游标时，才必须先右后左。
  Independent child ranges allow either construction order. A shared reverse-postorder cursor instead requires right before left.
- Go 的切片操作只创建视图，不会自动复制元素；三种方法均只读输入，新建的是结果树节点。
  Go subslicing creates views rather than copying elements. All three methods read the inputs and allocate output tree nodes.

时间与空间复杂度 / Time and Space Complexity
n 为节点数，h 为结果树高度。
线性查找切片递归：最坏时间 O(n²)，根查找可累计 n+(n-1)+...+1；辅助空间 O(h)，输出树 O(n)。
Go 切片是视图，不会像 C++ vector 那样按层复制元素，因此解法一、解法二的辅助空间同为 O(h)。
下标区间线性查找：最坏时间仍为 O(n²)；用下标表达切分，对应文章避免重复定义数组的写法。
哈希表区间递归：平均时间 O(n)，每个节点一次定位和构造；辅助空间 O(n+h)=O(n)，包括索引表和递归栈，输出树另占 O(n)。
For n nodes and resulting height h:
Slice recursion with linear lookup takes O(n²) worst-case time and O(h) auxiliary space, plus O(n) output tree space.
Go slices are views and do not copy elements the way C++ vectors do, so methods 1 and 2 share O(h) auxiliary space.
Index-range linear lookup remains O(n²) worst-case; it matches the article's version that splits by indices instead of new arrays.
Map-based range recursion takes expected O(n) time and O(n+h)=O(n) auxiliary space for the map and calls, plus O(n) output.

练习 / Practice
先练习画出根与四个子区间，再按切片、下标、哈希表的顺序阅读 buildTreePostOrder、buildTreePostOrderIndex 和 buildTreePostOrderOptimized。
Practice identifying the root and the four child ranges first,
then read buildTreePostOrder, buildTreePostOrderIndex, and buildTreePostOrderOptimized in that order.
*/

// 1. Linear root lookup and slice splitting: take the last postorder value as root, then use its inorder index k to split both arrays.
// 1. 递归切片：先理解这一版；后序末尾是根，用中序位置 k 同时切开两组遍历。
// Time: O(n²) worst case, Space: O(h) plus O(n) for the output tree.
// 时间复杂度：最坏 O(n²)，空间复杂度：O(h)，输出树另占 O(n)。
func buildTreePostOrder(inorder []int, postorder []int) *TreeNode {
	// An empty traversal represents an empty subtree.
	// 遍历为空，说明这棵子树不存在。
	if len(inorder) == 0 {
		return nil
	}

	// Postorder ends with the current subtree's root.
	// 后序最后一个值就是当前子树的根。
	rootValue := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootValue}

	// Find the root in inorder; k also equals the left subtree size.
	// 在中序中找到根；根前面有 k 个节点，所以左子树大小也是 k。
	k := 0
	for inorder[k] != rootValue {
		k++
	}

	// The first k postorder values belong to the left subtree.
	// 后序前 k 个值属于左子树，与左侧中序对应。
	root.Left = buildTreePostOrder(inorder[:k], postorder[:k])

	// The remaining values before the root belong to the right subtree.
	// 后序剩下的部分去掉最后的根，就是右子树的后序。
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
func buildTreePostOrderIndex(inorder []int, postorder []int) *TreeNode {
	// Empty input has no subtree to construct.
	// 输入为空，没有可构造的子树。
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	// Keep the half-open invariant [begin, end) on every recursive call.
	// 每一层都坚持左闭右开区间 [begin, end)。
	return buildPostorderIndex(inorder, 0, len(inorder), postorder, 0, len(postorder))
}

// Half-open inorder [inLeft, inRight) and postorder [postLeft, postRight) describe the same subtree.
// 左闭右开的中序 [inLeft, inRight) 与后序 [postLeft, postRight) 表示同一棵子树。
// Time: O(n²) worst case, Space: O(h).
// 时间复杂度：最坏 O(n²)，空间复杂度：O(h)。
func buildPostorderIndex(inorder []int, inLeft, inRight int, postorder []int, postLeft, postRight int) *TreeNode {
	// Equal postorder bounds mean this interval has no nodes.
	// 后序区间两端相等，说明当前子树为空。
	if postLeft == postRight {
		return nil
	}

	// The last postorder value is the current subtree root.
	// 后序区间最后一个值就是当前子树的根。
	rootValue := postorder[postRight-1]
	root := &TreeNode{Val: rootValue}

	// A single-node interval is already a complete leaf.
	// 区间只剩一个节点，它就是叶子，不必再切分。
	if postRight-postLeft == 1 {
		return root
	}

	// Scan the current inorder interval for the delimiter.
	// 在当前中序区间里线性查找切割点。
	delimiterIndex := inLeft
	for delimiterIndex < inRight && inorder[delimiterIndex] != rootValue {
		delimiterIndex++
	}

	// leftSize is the inorder distance from inLeft to the delimiter, not the global index itself.
	// 左子树大小是切割点相对 inLeft 的距离，不能直接用全局下标 delimiterIndex。
	leftSize := delimiterIndex - inLeft

	// Left postorder takes exactly leftSize values starting at postLeft.
	// 左后序从 postLeft 起恰好取 leftSize 个值。
	root.Left = buildPostorderIndex(
		inorder, inLeft, delimiterIndex,
		postorder, postLeft, postLeft+leftSize,
	)

	// Right postorder starts after the left subtree and excludes the current root.
	// 右后序紧接左后序，并排除当前根所在的末尾位置。
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
func buildTreePostOrderOptimized(inorder []int, postorder []int) *TreeNode {
	// Unique values map to unique inorder positions.
	// 节点值互不相同，每个值对应唯一的中序位置。
	positions := make(map[int]int, len(inorder))
	for i, value := range inorder {
		positions[value] = i
	}

	var build func(int, int, int, int) *TreeNode
	build = func(inLeft, inRight, postLeft, postRight int) *TreeNode {
		// Half-open intervals are empty when their boundaries meet.
		// 左闭右开区间的两端相等，表示空子树。
		if inLeft == inRight {
			return nil
		}

		// The final value in this postorder interval is its root.
		// 当前后序区间最后一个值就是根。
		rootValue := postorder[postRight-1]
		k := positions[rootValue]
		leftSize := k - inLeft

		root := &TreeNode{Val: rootValue}

		// Split both traversals using the same left-subtree size.
		// 用相同的左子树节点数切分两组遍历。
		root.Left = build(
			inLeft, k,
			postLeft, postLeft+leftSize,
		)

		// Exclude the root from both right-subtree intervals.
		// 右子树区间要排除中序中的根和后序末尾的根。
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

解法一：线性查找根并切片 / Method 1: Linear Lookup and Slice Splitting
取 preorder[0] 为根，在当前 inorder 中找到位置 k，k 就是左子树节点数。
左前序为 preorder[1:k+1]：跳过根后恰好取 k 个元素；右前序为 preorder[k+1:]。
左中序为 inorder[:k]，右中序为 inorder[k+1:]。分别递归构造后连接到新根。
Take preorder[0] as root and find its position k in the current inorder slice; k is the left-subtree size.
Left preorder is preorder[1:k+1], taking exactly k elements after the root; right preorder is preorder[k+1:].
Left inorder is inorder[:k], and right inorder is inorder[k+1:]. Recursively build and attach both children.

解法二：线性查找根 + 下标区间 / Method 2: Linear Lookup and Index Ranges
思路与解法一相同：preorder[preLeft] 是根，中序线性扫描切割点。
用左闭右开区间表示同一棵子树，不再切出子切片。
根仍是 preorder[preLeft]，左子树大小 leftSize=delimiterIndex-inLeft。
左前序从 preLeft+1 起取 leftSize 个；右前序紧接其后直到 preRight。
Same idea as method 1: preorder[preLeft] is the root, found by a linear inorder scan.
Describe the subtree with half-open ranges instead of creating child slices.
The root remains preorder[preLeft], and leftSize=delimiterIndex-inLeft.
Left preorder takes leftSize elements starting at preLeft+1; right preorder starts immediately afterward.

解法三：哈希表 + 下标区间 / Method 3: Index Map and Ranges
用哈希表保存值到全局中序下标。每次用两个左闭右开区间表示同一棵子树。
根为 preorder[preLeft]，中序位置为 k，左子树大小为 leftSize=k-inLeft。
区间切分与解法二相同；只是把线性查找换成哈希表。
Map values to global inorder indices. Use two half-open ranges for the same subtree.
The root is preorder[preLeft], its inorder index is k, and leftSize=k-inLeft.
Range splitting matches method 2; only the delimiter search becomes a map lookup.

左前序 / Left preorder:  [preLeft+1, preLeft+1+leftSize)
右前序 / Right preorder: [preLeft+1+leftSize, preRight)
左中序 / Left inorder:   [inLeft, k)
右中序 / Right inorder:  [k+1, inRight)

为什么前序边界多一个 +1？preLeft 本身是根，必须跳过它，从 preLeft+1 开始数 leftSize 个节点。
右子树紧接左子树结束处开始，两边都不包含根；中序边界仍与 106 相同。
Why the extra +1? preLeft contains the root; skip it and count leftSize elements starting at preLeft+1.
The right subtree starts immediately afterward. Inorder boundaries remain the same as in 106.

易错点 / Pitfalls
- 105 和 106 在 LeetCode 上都叫 buildTree，但同一个 Go 包不能重复定义同名函数。
  本项目给 105 使用 buildTreePreorder、buildTreePreorderIndex 和 buildTreePreorderOptimized；提交 105 时将入口改名为 buildTree。
  Both problems use buildTree on LeetCode, but a Go package cannot define it twice.
  Use buildTreePreorder, buildTreePreorderIndex, and buildTreePreorderOptimized here; rename the entry to buildTree when submitting 105.
- 线性切片法 k 是当前切片内的位置；下标法和哈希表法的切割点是原数组下标，必须减去 inLeft 才是大小。
  Slice-based k is local, while index-based and map-based delimiters are global; subtract inLeft to get the subtree size.
- 空区间返回 nil。每次递归排除当前根，规模才会缩小，也不会重复创建根。
  Return nil for an empty range. Excluding the current root makes recursive ranges shrink and avoids duplicate construction.
- Go 切片不复制元素，三种方法都只读输入，构造新树。
  Go subslicing does not copy elements. All three methods read inputs and construct a new tree.

时间与空间复杂度 / Time and Space Complexity
n 为节点数，h 为树高。线性查找切片与下标两版最坏时间都是 O(n²)，辅助空间 O(h)。
哈希表版本平均时间 O(n)，辅助空间 O(n+h)=O(n)。三种版本的输出树都另占 O(n)。
For n nodes and height h, both linear-lookup versions take O(n²) worst-case time and O(h) auxiliary space.
Map lookup takes expected O(n) time and O(n+h)=O(n) auxiliary space. All three output trees use O(n) space.

练习 / Practice
105 的三种实现按切片、下标、哈希表顺序写在本段之后：buildTreePreorder、buildTreePreorderIndex、buildTreePreorderOptimized；106 的实现保留在文件前半部分。
The three 105 implementations follow this section as buildTreePreorder, buildTreePreorderIndex, and buildTreePreorderOptimized,
while the 106 implementations remain in the first half of the file.
*/

// 1. Linear lookup and slice splitting: take preorder[0] as root, then use its inorder index k as the left-subtree size.
// 1. 递归切片：前序第一个值是根，用中序位置 k 作为左子树节点数去切前序。
// Time: O(n²) worst case, Space: O(h) plus O(n) for the output tree.
// 时间复杂度：最坏 O(n²)，空间复杂度：O(h)，输出树另占 O(n)。
func buildTreePreorder(preorder []int, inorder []int) *TreeNode {
	// Empty traversals represent an empty subtree.
	// 遍历为空，说明当前子树不存在。
	if len(preorder) == 0 {
		return nil
	}

	// Preorder starts with the current subtree's root.
	// 前序第一个值就是当前子树的根。
	rootValue := preorder[0]
	root := &TreeNode{Val: rootValue}

	// k is the root position and the number of nodes in the left subtree.
	// k 是根在当前中序切片中的位置，也是左子树节点数。
	k := 0
	for inorder[k] != rootValue {
		k++
	}

	// Skip the root, then take exactly k nodes for the left subtree.
	// 跳过前序中的根，再取 k 个节点构造左子树。
	root.Left = buildTreePreorder(
		preorder[1:k+1],
		inorder[:k],
	)

	// Everything after those k nodes belongs to the right subtree.
	// 前序剩余部分属于右子树，中序则取根右侧部分。
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
func buildTreePreorderIndex(preorder []int, inorder []int) *TreeNode {
	// Empty input has no subtree to construct.
	// 输入为空，没有可构造的子树。
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// Keep the half-open invariant [begin, end) on every recursive call.
	// 每一层都坚持左闭右开区间 [begin, end)。
	return buildPreorderIndex(preorder, 0, len(preorder), inorder, 0, len(inorder))
}

// Half-open preorder [preLeft, preRight) and inorder [inLeft, inRight) describe the same subtree.
// 左闭右开的前序 [preLeft, preRight) 与中序 [inLeft, inRight) 表示同一棵子树。
// Time: O(n²) worst case, Space: O(h).
// 时间复杂度：最坏 O(n²)，空间复杂度：O(h)。
func buildPreorderIndex(preorder []int, preLeft, preRight int, inorder []int, inLeft, inRight int) *TreeNode {
	// Equal preorder bounds mean this interval has no nodes.
	// 前序区间两端相等，说明当前子树为空。
	if preLeft == preRight {
		return nil
	}

	// Use preLeft, not 0: this interval may start in the middle of preorder.
	// 根取 preLeft 处的值，不能写死成下标 0，因为子区间不一定从数组开头开始。
	rootValue := preorder[preLeft]
	root := &TreeNode{Val: rootValue}

	// A single-node interval is already a complete leaf.
	// 区间只剩一个节点，它就是叶子，不必再切分。
	if preRight-preLeft == 1 {
		return root
	}

	// Scan the current inorder interval for the delimiter.
	// 在当前中序区间里线性查找切割点。
	delimiterIndex := inLeft
	for delimiterIndex < inRight && inorder[delimiterIndex] != rootValue {
		delimiterIndex++
	}

	// leftSize is the inorder distance from inLeft to the delimiter.
	// 左子树大小是切割点相对 inLeft 的距离。
	leftSize := delimiterIndex - inLeft

	// Skip the root at preLeft, then take leftSize preorder values.
	// 跳过 preLeft 处的根，再取 leftSize 个前序值构造左子树。
	root.Left = buildPreorderIndex(
		preorder, preLeft+1, preLeft+1+leftSize,
		inorder, inLeft, delimiterIndex,
	)

	// The right subtree starts immediately after the left preorder interval.
	// 右子树紧接左前序结束处，中序则跳过根的位置。
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
func buildTreePreorderOptimized(preorder []int, inorder []int) *TreeNode {
	// Map each unique value to its inorder index.
	// 保存每个节点值对应的中序下标，避免重复扫描。
	positions := make(map[int]int, len(inorder))
	for i, value := range inorder {
		positions[value] = i
	}

	var build func(int, int, int, int) *TreeNode
	build = func(preLeft, preRight, inLeft, inRight int) *TreeNode {
		// Both ranges are half-open; equal boundaries mean an empty subtree.
		// 两组区间都左闭右开，边界相等表示空子树。
		if preLeft == preRight {
			return nil
		}

		// The first preorder value is the root.
		// 当前前序区间的第一个值就是根。
		rootValue := preorder[preLeft]
		k := positions[rootValue]

		// k is a global index; subtract the current inorder start.
		// k 是全局下标，减去当前中序起点才是左子树节点数。
		leftSize := k - inLeft
		root := &TreeNode{Val: rootValue}

		// Skip the root and take leftSize preorder elements.
		// 跳过根，取 leftSize 个前序元素构造左子树。
		root.Left = build(
			preLeft+1, preLeft+1+leftSize,
			inLeft, k,
		)

		// The right subtree starts immediately after the left subtree.
		// 右子树紧接左子树之后，中序部分则跳过根的位置。
		root.Right = build(
			preLeft+1+leftSize, preRight,
			k+1, inRight,
		)

		return root
	}

	return build(0, len(preorder), 0, len(inorder))
}
