package _6_binary_tree

/*
LeetCode provides this definition.
LeetCode 已经提供这个结构，不需要重复提交。

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
LeetCode 116 / 117 use a next-right pointer. This is not TreeNode.
力扣 116 / 117 需要右侧 next 指针，不能复用 TreeNode。

type Node struct {
    Val   int
    Left  *Node
    Right *Node
    Next  *Node
}
*/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/*
LeetCode 429 uses an N-ary children list. This is not TreeNode.
力扣 429 使用 N 叉孩子列表，不能复用 TreeNode。

type NaryNode struct {
    Val      int
    Children []*NaryNode
}
*/

type NaryNode struct {
	Val      int
	Children []*NaryNode
}
