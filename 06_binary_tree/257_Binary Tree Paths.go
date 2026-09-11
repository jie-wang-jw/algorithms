package _6_binary_tree

/*
257. 二叉树的所有路径 / Binary Tree Paths

题目描述 / Problem Description
给定二叉树根节点 root，以字符串形式返回所有从根到叶子的路径，节点值之间用 "->" 连接。
叶子是左右孩子都为空的节点。答案顺序不限；空树返回空结果。
Given the root of a binary tree, return all root-to-leaf paths as strings with values separated by "->".
A leaf has no children. Return paths in any order; an empty tree has no paths.

示例 / Example
       1
      / \
     2   3
      \
       5
结果 / Result: ["1->2->5", "1->3"]
"1->2" 不能收集，因为 2 还有孩子，不是叶子。
"1->2" is not a complete answer because node 2 has a child and is not a leaf.

解法一：递归传递路径字符串 / Method 1: Recursion with a Path String
定义 traverse(node,path)：进入时 path 保存根到当前节点父亲的路径；先加入当前节点值，得到根到当前节点的路径。
path 非空才添加箭头，所以开头没有多余箭头。到叶子时收集路径，否则分别进入非空孩子。
On entry to traverse(node,path), path describes the route from the root to node's parent.
Append the current value, adding an arrow only if path is nonempty. Collect at a leaf; otherwise recurse into existing children.

关键逻辑：为什么分支不会串在一起？ / Why Do Branches Stay Separate?
Go 字符串不可修改，每层拿到自己的字符串参数。拼接后只是当前调用的 path 指向新字符串，不会修改父调用的 path。
例子中左分支生成 "1->2->5" 后，返回根调用时根的 path 仍为 "1"，所以右分支从 "1" 继续，得到 "1->3"。
这里没有共享可修改的路径缓冲区，因此不必手动删除末尾节点。不要误认为所有回溯写法都无需撤销。
Go strings are immutable and each call has its own parameter. Concatenation changes that call's path, not its parent's path.
After the left branch produces "1->2->5", the root call still has "1", so the right branch produces "1->3".
No shared mutable path buffer is used here, so no explicit removal is needed. This does not mean all backtracking needs no undo step.

解法二：共享路径切片回溯 / Method 2: Backtracking with a Shared Path Slice
用 []string 保存当前路径的节点值。进入节点时 append，处理完该节点的整棵子树后截去最后一个元素。
叶子处用 strings.Join(path,"->") 生成独立结果字符串。先收集，再统一撤销，不要因提前 return 漏掉撤销。
撤销恢复到父节点的路径，所以下一个兄弟分支从正确的公共前缀出发，不会带上上个分支的节点。
Use a shared []string for the current path. Append on entry, then remove the final entry after processing the entire subtree.
At a leaf, strings.Join(path,"->") creates a separate result string. Collect before the common undo step; do not skip it with an early return.
Undo restores the parent's path, letting a sibling start from the shared prefix without the previous branch's nodes.

解法三：栈迭代 / Method 3: Iterative DFS
使用节点栈和路径字符串栈，同步入栈、出栈；相同位置保存一个节点和根到该节点的完整路径。
每轮取出节点与路径，叶子直接收集，否则将孩子及扩展后的路径入栈。
先压右孩子，再压左孩子，利用后进先出让左分支先处理。题目不限顺序，但固定顺序方便理解和测试。
Use synchronized node and path stacks; matching positions hold a node and its full root-to-node path.
Pop both together, collect at a leaf, or push children with their extended paths.
Push right before left to process the left branch first under LIFO. Any result order is valid; consistent order helps inspection.

易错点 / Pitfalls
- 叶子判断是 Left==nil AND Right==nil，有一个孩子就还不是完整路径。
  A leaf requires both children nil; a node with one child is not a path endpoint.
- 将数字转为十进制文字用 strconv.Itoa，不能用 string(node.Val) 代替。
  Use strconv.Itoa for decimal text, not string(node.Val).
- 路径属于每一个分支，不是访问顺序列表；遍历所有节点得到的一串值不能当作所有路径。
  Each branch has its own path; one traversal sequence is not the collection of paths.
- 不修改树，只构造路径结果。所有写法对 nil 根先返回空结果。
  Do not modify the tree. All versions return an empty result for a nil root.

时间与空间复杂度 / Time and Space Complexity
n 为节点数，h 为树高，S 为返回的所有路径字符串的总字符数。以下按题目节点值位数有界分析。
字符串递归：每次拼接可能复制当前前缀，时间上界 O(nh)，最坏 O(n²)。递归中各层的路径字符串可同时保留，
除结果外空间保守上界 O(h²)，不能只算 O(h) 的调用栈；输出另占 O(S)。
共享切片回溯：每个节点入路径、出路径一次，叶子输出共复制 S 个字符，时间 O(n+S)。
除输出外辅助空间 O(h)，包括路径切片与递归栈；输出 O(S)。
栈迭代字符串：拼接时间上界 O(nh)；节点栈最多 O(h)，但每项还保存路径，除输出外空间上界 O(h²)，输出 O(S)。
Let n be the node count, h the height, and S the total output character count, assuming bounded value lengths as in this problem.
String recursion: O(nh) time from prefix copying, up to O(n²). Retained strings across active calls have a conservative O(h²)
auxiliary-space bound, not merely the O(h) call stack; output is O(S).
Shared-slice backtracking: O(n+S) time, O(h) auxiliary space for the path and calls, plus O(S) output.
String-stack iteration: O(nh) time; at most O(h) pending nodes with stored paths give O(h²) auxiliary space, plus O(S) output.

练习 / Practice
先掌握字符串递归，再比较共享切片为何需要撤销。三种实现按上面的编号顺序写在下方。
Start with string recursion, then compare why the shared slice needs an undo step.
The three implementations appear below in the order of the numbered methods above.
*/
import (
	"strconv"
	"strings"
)

// 1. 递归传递字符串：最容易理解  T:O(nh) S:O(h²)
func binaryTreePaths(root *TreeNode) []string {
	result := []string{}

	var traverse func(*TreeNode, string)
	traverse = func(node *TreeNode, path string) {
		if node == nil {
			return
		}

		// On entry, path ends at the parent; now include this node.
		// 进入时 path 到父节点为止，现在加入当前节点。
		if path != "" {
			path += "->"
		}
		path += strconv.Itoa(node.Val)

		// Only a leaf completes a root-to-leaf path.
		// 只有到达叶子，才形成完整的根到叶子路径。
		if node.Left == nil && node.Right == nil {
			result = append(result, path)
			return
		}

		// Each call extends its own path without changing the parent's.
		// 每次调用延伸自己的路径，不会修改父调用的路径。
		traverse(node.Left, path)
		traverse(node.Right, path)
	}

	traverse(root, "")
	return result
}

// 2. 共享切片回溯：显式“加入、撤销” T:O(n+S) S:O(h)
func binaryTreePathsBacktracking(root *TreeNode) []string {
	result := []string{}
	path := []string{}

	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		// Choose: include this node in the shared path.
		// 做选择：把当前节点加入共享路径。
		path = append(path, strconv.Itoa(node.Val))

		if node.Left == nil && node.Right == nil {
			// Create a result string independent of later path changes.
			// 生成独立结果字符串，不受之后路径修改的影响。
			result = append(result, strings.Join(path, "->"))
		} else {
			traverse(node.Left)
			traverse(node.Right)
		}

		// Undo: restore the parent's path before returning.
		// 撤销选择：返回前恢复父节点的路径。
		path = path[:len(path)-1]
	}

	traverse(root)
	return result
}

// 3. 栈迭代 T:O(nh) S:O(h²)
func binaryTreePathsIterative(root *TreeNode) []string {
	result := []string{}
	if root == nil {
		return result
	}

	// Matching stack positions describe the same pending node.
	// 两个栈的相同位置，对应同一个待处理节点及其完整路径。
	nodes := []*TreeNode{root}
	paths := []string{strconv.Itoa(root.Val)}

	for len(nodes) > 0 {
		last := len(nodes) - 1
		node, path := nodes[last], paths[last]
		nodes = nodes[:last]
		paths = paths[:last]

		// The stored path is complete when this node is a leaf.
		// 当前节点是叶子时，保存的路径就是完整答案。
		if node.Left == nil && node.Right == nil {
			result = append(result, path)
			continue
		}

		// Push right first so the left branch is processed first.
		// 先压右边，让左分支先出栈处理。
		if node.Right != nil {
			nodes = append(nodes, node.Right)
			paths = append(paths,
				path+"->"+strconv.Itoa(node.Right.Val))
		}

		if node.Left != nil {
			nodes = append(nodes, node.Left)
			paths = append(paths,
				path+"->"+strconv.Itoa(node.Left.Val))
		}
	}

	return result
}
