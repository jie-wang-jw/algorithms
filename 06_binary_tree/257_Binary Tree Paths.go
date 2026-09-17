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
复杂度记号：n 为节点数，h 为树高，w 为最大层宽；辅助空间不含返回结果。
Notation: n nodes, height h, maximum width w; auxiliary space excludes returned results.
*/
import (
	"strconv"
	"strings"
)

// 1. Recursion with a path string: append the current value onto an immutable prefix; each call owns its own string.
// 1. 递归传递路径字符串：把当前值拼到不可变前缀上，每次调用拥有自己的字符串。
// Time: O(nh), Space: O(h²) auxiliary for retained prefixes plus O(S) output, where S is the total output character count.
// 时间复杂度：O(nh)，空间复杂度：辅助 O(h²)，由各层同时保留的路径前缀产生，输出另占 O(S)，S 为所有路径字符串的总字符数。
//
// 每层将当前值追加到字符串路径，只有叶子才保存；字符串不可变，两个递归分支的追加不会互相修改。
// Append the current value and save only at leaves; immutable string concatenation isolates recursive branches.
// 路径越深，重复复制的前缀越长，复杂度须计入字符串复制和输出长度。
// Deeper paths repeatedly copy longer prefixes, so complexity includes string-copy and output lengths.
func binaryTreePaths(root *TreeNode) []string {
	result := []string{}

	var traverse func(*TreeNode, string)
	traverse = func(node *TreeNode, path string) {
		if node == nil {
			return
		}

		if path != "" {
			path += "->"
		}
		path += strconv.Itoa(node.Val)

		if node.Left == nil && node.Right == nil {
			result = append(result, path)
			return
		}

		traverse(node.Left, path)
		traverse(node.Right, path)
	}

	traverse(root, "")
	return result
}

// 2. Backtracking with a shared path slice: append on entry, join at a leaf, then undo the last value before returning.
// 2. 共享切片回溯：进入时加入，叶子处拼接结果，返回前撤销最后一个值。
// Time: O(n+S), Space: O(h) auxiliary plus O(S) output, where S is the total output character count.
// 时间复杂度：O(n+S)，空间复杂度：辅助 O(h)，输出另占 O(S)，S 为所有路径字符串的总字符数。
//
// path 是可复用的根到当前节点序列；进入追加，离开弹出。叶子处 Join 生成独立结果字符串。
// Reuse path by appending on entry and popping on exit; Join at a leaf creates an independent output string.
func binaryTreePathsBacktracking(root *TreeNode) []string {
	result := []string{}
	path := []string{}

	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		path = append(path, strconv.Itoa(node.Val))

		if node.Left == nil && node.Right == nil {
			result = append(result, strings.Join(path, "->"))
		} else {
			traverse(node.Left)
			traverse(node.Right)
		}

		path = path[:len(path)-1]
	}

	traverse(root)
	return result
}

// 3. Iterative DFS: keep synchronized node and path-string stacks so each pending node carries its full root-to-node path.
// 3. 栈迭代：节点栈与路径字符串栈同步，每个待处理节点带着从根到它自己的完整路径。
// Time: O(nh), Space: O(h²) auxiliary plus O(S) output, where S is the total output character count.
// 时间复杂度：O(nh)，空间复杂度：辅助 O(h²)，输出另占 O(S)，S 为所有路径字符串的总字符数。
//
// nodes 与 paths 一一对应；弹出节点同时取其完整路径，只有叶子输出，孩子继承父路径并追加自身。
// Keep nodes aligned with complete path strings; emit at leaves and extend the parent path for each child.
func binaryTreePathsIterative(root *TreeNode) []string {
	result := []string{}
	if root == nil {
		return result
	}

	nodes := []*TreeNode{root}
	paths := []string{strconv.Itoa(root.Val)}

	for len(nodes) > 0 {
		last := len(nodes) - 1
		node, path := nodes[last], paths[last]
		nodes = nodes[:last]
		paths = paths[:last]

		if node.Left == nil && node.Right == nil {
			result = append(result, path)
			continue
		}

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
