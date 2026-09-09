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

func TestCountNodes(t *testing.T) {
	implementations := []struct {
		name string
		fn   func(*TreeNode) int
	}{
		{name: "recursive", fn: countNodes},
		{name: "iterative", fn: countNodesIterative},
		{name: "optimized", fn: countNodesOptimized},
	}
	// Counts near full-level boundaries exercise both the formula and recursive splitting.
	// 满层边界附近的数量，同时检验直接套公式和递归拆分两种情况。
	tests := []struct {
		name string
		size int
	}{
		{name: "empty tree", size: 0},
		{name: "single node", size: 1},
		{name: "left child only", size: 2},
		{name: "full two levels", size: 3},
		{name: "last level one node", size: 4},
		{name: "last level half full", size: 5},
		{name: "six node example", size: 6},
		{name: "full three levels", size: 7},
		{name: "fourth level begins", size: 8},
		{name: "one missing from four levels", size: 14},
		{name: "full four levels", size: 15},
		{name: "fifth level begins", size: 16},
		{name: "full five levels", size: 31},
		{name: "sixth level begins", size: 32},
		{name: "large incomplete last level", size: 50000},
	}
	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					// Level-order array positions 2*i+1 and 2*i+2 are the children of i.
					// 连续层序数组中，i 的左右孩子下标是 2*i+1、2*i+2，天然构成完全二叉树。
					// Allocate fresh nodes for each run; repeated values must not affect the count.
					// 每次创建新节点；使用重复值，确保统计节点数量而不是不同值数量。
					nodes := make([]TreeNode, tt.size)
					for i := range nodes {
						nodes[i].Val = i % 3
						if left := 2*i + 1; left < len(nodes) {
							nodes[i].Left = &nodes[left]
						}
						if right := 2*i + 2; right < len(nodes) {
							nodes[i].Right = &nodes[right]
						}
					}
					var root *TreeNode
					if len(nodes) > 0 {
						root = &nodes[0]
					}
					// The construction size is the expected count, independent of any solution.
					// 构造时的节点数就是期望结果，不依赖任何待测算法计算答案。
					if got := implementation.fn(root); got != tt.size {
						t.Fatalf("countNodes() = %d, want %d", got, tt.size)
					}
				})
			}
		})
	}
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

func TestIsSymmetric(t *testing.T) {
	implementations := []struct {
		name string
		fn   func(*TreeNode) bool
	}{
		{name: "recursive", fn: isSymmetric},
		{name: "iterative", fn: isSymmetricIterative},
	}
	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			// Fresh trees keep implementations independent even if one mistakenly mutates input.
			// 每种实现使用新树，即使某种实现误改输入，也不会干扰另一种实现。
			tests := []struct {
				name string
				root *TreeNode
				want bool
			}{
				{name: "empty tree", want: true},
				{name: "single node", root: &TreeNode{Val: 1}, want: true},
				{name: "left child only", root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, want: false},
				{name: "right child only", root: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, want: false},
				{name: "different child values", root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, want: false},
				{
					name: "full mirror",
					root: &TreeNode{Val: 1,
						Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
						Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}},
					want: true,
				},
				{
					// Equal level values do not imply mirrored nil positions.
					// 每层值相同，也可能因为空节点位置不对称而失败。
					name: "same direction is not a mirror",
					root: &TreeNode{Val: 1,
						Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
						Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}},
					want: false,
				},
				{
					name: "sparse mirror",
					root: &TreeNode{Val: 1,
						Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
						Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
					want: true,
				},
				{
					// A nil/nil outer pair must not hide a later inner mismatch.
					// 外侧同时为空时只能继续，不能漏掉随后内侧的值不匹配。
					name: "nil outer pair before inner mismatch",
					root: &TreeNode{Val: 1,
						Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
						Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}},
					want: false,
				},
				{
					name: "outer value mismatch",
					root: &TreeNode{Val: 1,
						Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}},
						Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 4}}},
					want: false,
				},
			}
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := implementation.fn(tt.root); got != tt.want {
						t.Fatalf("symmetry = %t, want %t", got, tt.want)
					}
				})
			}
		})
	}
}

func TestMaxDepth(t *testing.T) {
	implementations := []struct {
		name string
		fn   func(*TreeNode) int
	}{
		{name: "recursive", fn: maxDepth},
		{name: "iterative", fn: maxDepthIterative},
	}
	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			tests := []struct {
				name string
				root *TreeNode
				want int
			}{
				{name: "empty tree", want: 0},
				{name: "single node", root: &TreeNode{Val: 0}, want: 1},
				{name: "two levels", root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, want: 2},
				{
					name: "example deeper right",
					root: &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}},
					want: 3,
				},
				{
					name: "deeper left",
					root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}},
					want: 3,
				},
				{name: "left chain", root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}, want: 3},
				{name: "right chain", root: &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}, want: 3},
				{
					// Depth counts nodes on one path, regardless of values or direction changes.
					// 深度数的是一条路径上的节点，与节点值、左右转向无关。
					name: "zigzag equal values",
					root: &TreeNode{Val: 0, Left: &TreeNode{Val: 0, Right: &TreeNode{Val: 0, Left: &TreeNode{Val: 0}}}},
					want: 4,
				},
				{
					// Seven nodes occupy only three levels: count layers, not nodes.
					// 七个节点只有三层：检查是否错误地按节点数累加深度。
					name: "full three levels",
					root: &TreeNode{Val: 1,
						Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
						Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}},
					want: 3,
				},
			}
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := implementation.fn(tt.root); got != tt.want {
						t.Fatalf("max depth = %d, want %d", got, tt.want)
					}
				})
			}
		})
	}
}

func TestMinDepth(t *testing.T) {
	implementations := []struct {
		name string
		fn   func(*TreeNode) int
	}{
		{name: "recursive", fn: minDepth},
		{name: "iterative", fn: minDepthIterative},
	}
	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			// Build fresh inputs for each implementation.
			// 每种实现使用独立的新树，避免用例之间相互影响。
			tests := []struct {
				name string
				root *TreeNode
				want int
			}{
				{name: "empty tree", want: 0},
				{name: "single node", root: &TreeNode{Val: 1}, want: 1},
				// A missing child is not a leaf and cannot shorten the path.
				// 缺失的孩子不是叶子，不能通过空的一侧缩短路径。
				{name: "left child only", root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, want: 2},
				{name: "right child only", root: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, want: 2},
				{name: "left chain", root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}, want: 3},
				{name: "right chain", root: &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}, want: 3},
				{
					name: "example shallow left leaf",
					root: &TreeNode{Val: 3, Left: &TreeNode{Val: 9},
						Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}},
					want: 2,
				},
				{
					// A left-first DFS leaf need not be the nearest leaf.
					// 先沿左侧找到的叶子不一定最近，还要比较右侧浅叶子。
					name: "shallow right leaf",
					root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 4}},
					want: 2,
				},
				{
					name: "full three levels",
					root: &TreeNode{Val: 1,
						Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
						Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}},
					want: 3,
				},
				{
					// Repeated values and changing directions do not change path length.
					// 重复值和左右转向不影响路径长度，每个真实节点都要计数。
					name: "zigzag equal values",
					root: &TreeNode{Val: 0, Left: &TreeNode{Val: 0, Right: &TreeNode{Val: 0, Left: &TreeNode{Val: 0}}}},
					want: 4,
				},
			}
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := implementation.fn(tt.root); got != tt.want {
						t.Fatalf("min depth = %d, want %d", got, tt.want)
					}
				})
			}
		})
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
