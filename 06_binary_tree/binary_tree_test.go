package _6_binary_tree

import (
	"reflect"
	"testing"
)

type traversalTestCase struct {
	name      string
	root      *TreeNode
	level     [][]int
	preorder  []int
	inorder   []int
	postorder []int
}

func traversalTestCases() []traversalTestCase {
	// Ordinary tree:
	// 普通二叉树：
	//
	//          1
	//        /   \
	//       2     3
	//      / \     \
	//     4   5     6
	ordinary := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 6},
		},
	}

	// Right-skewed tree: 1 -> 2 -> 3.
	// 右偏树：1 -> 2 -> 3。
	rightSkewed := &TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
	}

	// Left-skewed tree: 1 -> 2 -> 3.
	// 左偏树：1 -> 2 -> 3。
	leftSkewed := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}},
	}

	return []traversalTestCase{
		{name: "empty tree", root: nil, level: [][]int{}, preorder: []int{}, inorder: []int{}, postorder: []int{}},
		{name: "single node", root: &TreeNode{Val: 7}, level: [][]int{{7}}, preorder: []int{7}, inorder: []int{7}, postorder: []int{7}},
		{name: "ordinary tree", root: ordinary, level: [][]int{{1}, {2, 3}, {4, 5, 6}}, preorder: []int{1, 2, 4, 5, 3, 6}, inorder: []int{4, 2, 5, 1, 3, 6}, postorder: []int{4, 5, 2, 6, 3, 1}},
		{name: "right skewed tree", root: rightSkewed, level: [][]int{{1}, {2}, {3}}, preorder: []int{1, 2, 3}, inorder: []int{1, 2, 3}, postorder: []int{3, 2, 1}},
		{name: "left skewed tree", root: leftSkewed, level: [][]int{{1}, {2}, {3}}, preorder: []int{1, 2, 3}, inorder: []int{3, 2, 1}, postorder: []int{3, 2, 1}},
	}
}

func TestLevelOrder(t *testing.T) {
	for _, tt := range traversalTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.root); !reflect.DeepEqual(got, tt.level) {
				t.Fatalf("levelOrder() = %v, want %v", got, tt.level)
			}
		})
	}
}

func TestPreorderTraversal(t *testing.T) {
	// Both iterative and recursive implementations must return root-left-right order.
	// 迭代和递归实现都必须返回“根、左、右”的顺序。
	implementations := []struct {
		name string
		fn   func(*TreeNode) []int
	}{
		{name: "iterative", fn: preorderTraversal},
		{name: "recursive", fn: preorderTraversalRecursive},
	}

	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			for _, tt := range traversalTestCases() {
				t.Run(tt.name, func(t *testing.T) {
					if got := implementation.fn(tt.root); !reflect.DeepEqual(got, tt.preorder) {
						t.Fatalf("preorder traversal = %v, want %v", got, tt.preorder)
					}
				})
			}
		})
	}
}

func TestInorderTraversal(t *testing.T) {
	// Both iterative and recursive implementations must return left-root-right order.
	// 迭代和递归实现都必须返回“左、根、右”的顺序。
	implementations := []struct {
		name string
		fn   func(*TreeNode) []int
	}{
		{name: "iterative", fn: inorderTraversal},
		{name: "recursive", fn: inorderTraversalRecursive},
	}

	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			for _, tt := range traversalTestCases() {
				t.Run(tt.name, func(t *testing.T) {
					if got := implementation.fn(tt.root); !reflect.DeepEqual(got, tt.inorder) {
						t.Fatalf("inorder traversal = %v, want %v", got, tt.inorder)
					}
				})
			}
		})
	}
}

func TestInvertTree(t *testing.T) {
	implementations := []struct {
		name string
		fn   func(*TreeNode) *TreeNode
	}{
		{name: "preorder recursive", fn: invertTree},
		{name: "postorder recursive", fn: invertTreePostorderRecursive},
		{name: "preorder iterative", fn: invertTreePreorderIterative},
		{name: "postorder iterative", fn: invertTreePostorderIterative},
		{name: "level order", fn: invertTreeLevelOrder},
	}
	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			// Allocate fresh trees per implementation because inversion mutates the input.
			// 每种实现重新创建测试树，避免原地翻转影响下一个实现的输入。
			// Compare complete structures, including nil children, rather than traversal values alone.
			// 比较包含空孩子位置的完整结构，避免仅比较遍历值而漏掉结构错误。
			tests := []struct {
				name string
				root *TreeNode
				want *TreeNode
			}{
				{name: "empty tree", root: nil, want: nil},
				{name: "single node", root: &TreeNode{Val: 1}, want: &TreeNode{Val: 1}},
				{
					name: "full tree",
					root: &TreeNode{Val: 4,
						Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
						Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}}},
					want: &TreeNode{Val: 4,
						Left:  &TreeNode{Val: 7, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 6}},
						Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}}},
				},
				{
					name: "left chain",
					root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
					want: &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}},
				},
				{
					name: "right chain",
					root: &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}},
					want: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
				},
				{
					name: "equal values asymmetric shape",
					root: &TreeNode{Val: 1, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 1}}},
					want: &TreeNode{Val: 1, Right: &TreeNode{Val: 1, Left: &TreeNode{Val: 1}}},
				},
			}

			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := implementation.fn(tt.root)
					if got != tt.root {
						t.Fatal("invertTree must preserve the root pointer")
					}
					if !reflect.DeepEqual(got, tt.want) {
						t.Fatalf("invertTree() structure mismatch: got %#v, want %#v", got, tt.want)
					}
				})
			}
		})
	}
}

func TestPostorderTraversal(t *testing.T) {
	// Both iterative and recursive implementations must return left-right-root order.
	// 迭代和递归实现都必须返回“左、右、根”的顺序。
	implementations := []struct {
		name string
		fn   func(*TreeNode) []int
	}{
		{name: "iterative", fn: postorderTraversal},
		{name: "recursive", fn: postorderTraversalRecursive},
	}

	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			for _, tt := range traversalTestCases() {
				t.Run(tt.name, func(t *testing.T) {
					if got := implementation.fn(tt.root); !reflect.DeepEqual(got, tt.postorder) {
						t.Fatalf("postorder traversal = %v, want %v", got, tt.postorder)
					}
				})
			}
		})
	}
}
