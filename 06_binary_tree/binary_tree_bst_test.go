package _6_binary_tree

import (
	"reflect"
	"slices"
	"testing"
)

func cloneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return &TreeNode{Val: root.Val, Left: cloneTree(root.Left), Right: cloneTree(root.Right)}
}

func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if left := findNode(root.Left, val); left != nil {
		return left
	}
	return findNode(root.Right, val)
}

func TestConstructMaximumBinaryTree(t *testing.T) {
	implementations := []struct {
		name string
		fn   func([]int) *TreeNode
	}{
		{name: "slicing", fn: constructMaximumBinaryTree},
		{name: "index ranges", fn: constructMaximumBinaryTreeIndex},
	}
	tests := []struct {
		name string
		nums []int
		want *TreeNode
	}{
		{name: "empty", nums: []int{}},
		{name: "single", nums: []int{7}, want: &TreeNode{Val: 7}},
		{
			name: "article example",
			nums: []int{3, 2, 1, 6, 0, 5},
			want: &TreeNode{Val: 6,
				Left:  &TreeNode{Val: 3, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 1}}},
				Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 0}}},
		},
		{name: "decreasing left chain", nums: []int{3, 2, 1}, want: &TreeNode{Val: 3, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 1}}}},
		{name: "increasing right spine is left chain", nums: []int{1, 2, 3}, want: &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}}},
	}
	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					input := slices.Clone(tt.nums)
					got := implementation.fn(input)
					if !reflect.DeepEqual(got, tt.want) {
						t.Fatalf("got %#v, want %#v", got, tt.want)
					}
					if !slices.Equal(input, tt.nums) {
						t.Fatal("construction modified the input array")
					}
				})
			}
		})
	}
}

func TestMergeTrees(t *testing.T) {
	type mergeFn func(*TreeNode, *TreeNode) *TreeNode
	implementations := []struct {
		name     string
		fn       mergeFn
		mutates1 bool
	}{
		{name: "preorder reuse root1", fn: mergeTrees, mutates1: true},
		{name: "new tree", fn: mergeTreesNew},
		{name: "queue", fn: mergeTreesIterative, mutates1: true},
	}
	tests := []struct {
		name  string
		root1 *TreeNode
		root2 *TreeNode
		want  *TreeNode
	}{
		{name: "both empty"},
		{name: "only first", root1: &TreeNode{Val: 1}, want: &TreeNode{Val: 1}},
		{name: "only second", root2: &TreeNode{Val: 2}, want: &TreeNode{Val: 2}},
		{
			name:  "article example",
			root1: &TreeNode{Val: 1, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 2}},
			root2: &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 7}}},
			want:  &TreeNode{Val: 3, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 7}}},
		},
		{
			name:  "left chain plus right chain",
			root1: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
			root2: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}},
			want:  &TreeNode{Val: 4, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}},
		},
	}
	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := implementation.fn(cloneTree(tt.root1), cloneTree(tt.root2))
					if !reflect.DeepEqual(got, tt.want) {
						t.Fatalf("got %#v, want %#v", got, tt.want)
					}
				})
			}
		})
	}
}

func TestSearchBST(t *testing.T) {
	implementations := []struct {
		name string
		fn   func(*TreeNode, int) *TreeNode
	}{
		{name: "recursive", fn: searchBST},
		{name: "iterative", fn: searchBSTIterative},
	}
	root := &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 7}}
	tests := []struct {
		name string
		root *TreeNode
		val  int
		want *TreeNode
	}{
		{name: "empty", val: 1},
		{name: "found left subtree", root: root, val: 2, want: root.Left},
		{name: "found root", root: root, val: 4, want: root},
		{name: "missing", root: root, val: 5},
		{name: "single match", root: &TreeNode{Val: 1}, val: 1, want: &TreeNode{Val: 1}},
	}
	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := implementation.fn(tt.root, tt.val)
					if tt.want == nil {
						if got != nil {
							t.Fatalf("got %#v, want nil", got)
						}
						return
					}
					if got != tt.want && !reflect.DeepEqual(got, tt.want) {
						t.Fatalf("got %#v, want %#v", got, tt.want)
					}
				})
			}
		})
	}
}

func TestIsValidBST(t *testing.T) {
	implementations := []struct {
		name string
		fn   func(*TreeNode) bool
	}{
		{name: "inorder array", fn: isValidBST},
		{name: "inorder prev", fn: isValidBSTInorder},
		{name: "iterative", fn: isValidBSTIterative},
	}
	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{name: "empty", want: true},
		{name: "single", root: &TreeNode{Val: 1}, want: true},
		{name: "valid", root: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, want: true},
		{name: "equal child invalid", root: &TreeNode{Val: 1, Left: &TreeNode{Val: 1}}, want: false},
		{
			name: "right grandchild too small",
			root: &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}}},
			want: false,
		},
		{name: "min int root", root: &TreeNode{Val: int(^uint(0)>>1) * 0}, want: true},
		{name: "left chain valid", root: &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}}, want: true},
		{name: "right chain valid", root: &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}, want: true},
	}
	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := implementation.fn(tt.root); got != tt.want {
						t.Fatalf("isValidBST() = %t, want %t", got, tt.want)
					}
				})
			}
		})
	}
}

func TestGetMinimumDifference(t *testing.T) {
	implementations := []struct {
		name string
		fn   func(*TreeNode) int
	}{
		{name: "recursive", fn: getMinimumDifference},
		{name: "iterative", fn: getMinimumDifferenceIterative},
	}
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{name: "empty"},
		{name: "single", root: &TreeNode{Val: 5}},
		{name: "example", root: &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 6}}, want: 1},
		{name: "right chain", root: &TreeNode{Val: 1, Right: &TreeNode{Val: 8, Right: &TreeNode{Val: 10}}}, want: 2},
		{name: "sparse large gap then one", root: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, want: 1},
	}
	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := implementation.fn(tt.root); got != tt.want {
						t.Fatalf("min diff = %d, want %d", got, tt.want)
					}
				})
			}
		})
	}
}

func TestFindMode(t *testing.T) {
	implementations := []struct {
		name string
		fn   func(*TreeNode) []int
	}{
		{name: "recursive", fn: findMode},
		{name: "iterative", fn: findModeIterative},
	}
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{name: "empty", want: []int{}},
		{name: "single", root: &TreeNode{Val: 1}, want: []int{1}},
		{name: "one mode", root: &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 2}}}, want: []int{2}},
		{name: "all unique", root: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, want: []int{1, 2, 3}},
		{name: "two modes", root: &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 2}}}, want: []int{1, 2}},
	}
	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := slices.Clone(implementation.fn(tt.root))
					want := slices.Clone(tt.want)
					slices.Sort(got)
					slices.Sort(want)
					if !slices.Equal(got, want) {
						t.Fatalf("modes = %v, want %v", got, want)
					}
				})
			}
		})
	}
}

func TestLowestCommonAncestor(t *testing.T) {
	implementations := []struct {
		name string
		fn   func(root, p, q *TreeNode) *TreeNode
	}{
		{name: "postorder", fn: lowestCommonAncestor},
		{name: "parents", fn: lowestCommonAncestorParents},
	}
	root := &TreeNode{Val: 3,
		Left:  &TreeNode{Val: 5, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}}},
		Right: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 8}},
	}
	tests := []struct {
		name string
		p, q int
		want int
	}{
		{name: "split at root", p: 5, q: 1, want: 3},
		{name: "ancestor is p", p: 5, q: 4, want: 5},
		{name: "same subtree deep", p: 7, q: 4, want: 2},
		{name: "root and leaf", p: 3, q: 8, want: 3},
	}
	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					p := findNode(root, tt.p)
					q := findNode(root, tt.q)
					got := implementation.fn(root, p, q)
					if got == nil || got.Val != tt.want || got != findNode(root, tt.want) {
						t.Fatalf("LCA = %#v, want node %d", got, tt.want)
					}
				})
			}
		})
	}
}

func TestLowestCommonAncestorBST(t *testing.T) {
	implementations := []struct {
		name string
		fn   func(root, p, q *TreeNode) *TreeNode
	}{
		{name: "recursive", fn: lowestCommonAncestorBST},
		{name: "iterative", fn: lowestCommonAncestorBSTIterative},
	}
	root := &TreeNode{Val: 6,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}},
		Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 9}},
	}
	tests := []struct {
		name string
		p, q int
		want int
	}{
		{name: "split at root", p: 2, q: 8, want: 6},
		{name: "ancestor is p", p: 2, q: 4, want: 2},
		{name: "q smaller than p", p: 8, q: 7, want: 8},
	}
	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := implementation.fn(root, findNode(root, tt.p), findNode(root, tt.q))
					if got == nil || got.Val != tt.want {
						t.Fatalf("LCA = %#v, want %d", got, tt.want)
					}
				})
			}
		})
	}
}

func TestInsertIntoBST(t *testing.T) {
	implementations := []struct {
		name string
		fn   func(*TreeNode, int) *TreeNode
	}{
		{name: "recursive", fn: insertIntoBST},
		{name: "iterative", fn: insertIntoBSTIterative},
	}
	tests := []struct {
		name string
		root *TreeNode
		val  int
		want *TreeNode
	}{
		{name: "empty", val: 1, want: &TreeNode{Val: 1}},
		{
			name: "insert right of seven's left",
			root: &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 7}},
			val:  5,
			want: &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 5}}},
		},
		{name: "insert left of single", root: &TreeNode{Val: 2}, val: 1, want: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}},
	}
	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := implementation.fn(cloneTree(tt.root), tt.val)
					if !reflect.DeepEqual(got, tt.want) {
						t.Fatalf("got %#v, want %#v", got, tt.want)
					}
				})
			}
		})
	}
}

func TestDeleteNode(t *testing.T) {
	implementations := []struct {
		name string
		fn   func(*TreeNode, int) *TreeNode
	}{
		{name: "recursive", fn: deleteNode},
		{name: "iterative", fn: deleteNodeIterative},
	}
	base := func() *TreeNode {
		return &TreeNode{Val: 5,
			Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}},
			Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 7}}}
	}
	tests := []struct {
		name string
		root *TreeNode
		key  int
		want *TreeNode
	}{
		{name: "empty", key: 1},
		{name: "missing", root: &TreeNode{Val: 1}, key: 2, want: &TreeNode{Val: 1}},
		{name: "delete leaf", root: base(), key: 2, want: &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 7}}}},
		{name: "delete one child", root: base(), key: 6, want: &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 7}}},
		{name: "delete two children", root: base(), key: 3, want: &TreeNode{Val: 5, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 7}}}},
		{name: "delete root two children", root: base(), key: 5, want: &TreeNode{Val: 6, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 7}}},
	}
	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := implementation.fn(cloneTree(tt.root), tt.key)
					if !reflect.DeepEqual(got, tt.want) {
						t.Fatalf("got %#v, want %#v", got, tt.want)
					}
				})
			}
		})
	}
}

func TestTrimBST(t *testing.T) {
	implementations := []struct {
		name string
		fn   func(*TreeNode, int, int) *TreeNode
	}{
		{name: "recursive", fn: trimBST},
		{name: "iterative", fn: trimBSTIterative},
	}
	tests := []struct {
		name      string
		root      *TreeNode
		low, high int
		want      *TreeNode
	}{
		{name: "empty", low: 1, high: 2},
		{name: "keep all", root: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, low: 1, high: 3, want: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}},
		{name: "drop left leaf", root: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 2}}, low: 1, high: 2, want: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}},
		{name: "replace root", root: &TreeNode{Val: 3, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 4}}, low: 1, high: 2, want: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}},
	}
	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := implementation.fn(cloneTree(tt.root), tt.low, tt.high)
					if !reflect.DeepEqual(got, tt.want) {
						t.Fatalf("got %#v, want %#v", got, tt.want)
					}
				})
			}
		})
	}
}

func TestSortedArrayToBST(t *testing.T) {
	implementations := []struct {
		name string
		fn   func([]int) *TreeNode
	}{
		{name: "index", fn: sortedArrayToBST},
		{name: "slicing", fn: sortedArrayToBSTSlice},
	}
	tests := []struct {
		name string
		nums []int
		want *TreeNode
	}{
		{name: "empty", nums: []int{}},
		{name: "single", nums: []int{1}, want: &TreeNode{Val: 1}},
		{name: "three", nums: []int{1, 2, 3}, want: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}},
		{
			name: "leetcode example",
			nums: []int{-10, -3, 0, 5, 9},
			want: &TreeNode{Val: 0, Left: &TreeNode{Val: -3, Left: &TreeNode{Val: -10}}, Right: &TreeNode{Val: 9, Left: &TreeNode{Val: 5}}},
		},
	}
	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					input := slices.Clone(tt.nums)
					got := implementation.fn(input)
					if !reflect.DeepEqual(got, tt.want) {
						t.Fatalf("got %#v, want %#v", got, tt.want)
					}
					if !isValidBSTInorder(got) {
						t.Fatal("result is not a BST")
					}
					if !isBalanced(got) {
						t.Fatal("result is not height-balanced")
					}
				})
			}
		})
	}
}

func TestConvertBST(t *testing.T) {
	implementations := []struct {
		name string
		fn   func(*TreeNode) *TreeNode
	}{
		{name: "recursive", fn: convertBST},
		{name: "iterative", fn: convertBSTIterative},
	}
	tests := []struct {
		name string
		root *TreeNode
		want *TreeNode
	}{
		{name: "empty"},
		{name: "single", root: &TreeNode{Val: 1}, want: &TreeNode{Val: 1}},
		{
			name: "seven nodes",
			root: &TreeNode{Val: 4,
				Left:  &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 2}},
				Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 7}}},
			want: &TreeNode{Val: 22,
				Left:  &TreeNode{Val: 25, Left: &TreeNode{Val: 25}, Right: &TreeNode{Val: 24}},
				Right: &TreeNode{Val: 13, Left: &TreeNode{Val: 18}, Right: &TreeNode{Val: 7}}},
		},
	}
	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := implementation.fn(cloneTree(tt.root))
					if !reflect.DeepEqual(got, tt.want) {
						t.Fatalf("got %#v, want %#v", got, tt.want)
					}
				})
			}
		})
	}
}

func TestLevelOrderVariants(t *testing.T) {
	example := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	if got := levelOrderBottom(example); !reflect.DeepEqual(got, [][]int{{15, 7}, {9, 20}, {3}}) {
		t.Fatalf("levelOrderBottom = %v", got)
	}
	if got := rightSideView(example); !slices.Equal(got, []int{3, 20, 7}) {
		t.Fatalf("rightSideView = %v", got)
	}
	if got := averageOfLevels(example); !reflect.DeepEqual(got, []float64{3, 14.5, 11}) {
		t.Fatalf("averageOfLevels = %v", got)
	}
	if got := largestValues(example); !slices.Equal(got, []int{3, 20, 15}) {
		t.Fatalf("largestValues = %v", got)
	}
	if got := levelOrderBottom(nil); !reflect.DeepEqual(got, [][]int{}) {
		t.Fatalf("empty bottom = %v", got)
	}
	if got := rightSideView(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}}); !slices.Equal(got, []int{1, 2}) {
		t.Fatalf("left child still visible from the right when no right sibling: %v", got)
	}
}
