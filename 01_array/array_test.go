package _1_array

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"sort"
	"testing"
)

// equalInts compares length and elements, so a nil slice equals an empty slice.
// equalInts 只比较长度和元素，因此 nil 切片与空切片视为相等。
func equalInts(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// sortedCopy returns a sorted copy, used when a solution may reorder its output.
// sortedCopy 返回排序后的副本，用于比较可能改变元素顺序的解法。
func sortedCopy(nums []int) []int {
	out := append([]int(nil), nums...)
	sort.Ints(out)
	return out
}

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"target exists in middle", []int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{"target does not exist", []int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{"target is first element", []int{-1, 0, 3, 5, 9, 12}, -1, 0},
		{"target is last element", []int{-1, 0, 3, 5, 9, 12}, 12, 5},
		{"target smaller than all", []int{-1, 0, 3, 5, 9, 12}, -5, -1},
		{"target larger than all", []int{-1, 0, 3, 5, 9, 12}, 100, -1},
		{"empty slice", []int{}, 1, -1},
		{"single element exists", []int{5}, 5, 0},
		{"single element not exists", []int{5}, 3, -1},
		{"two elements first", []int{4, 7}, 4, 0},
		{"two elements second", []int{4, 7}, 7, 1},
	}

	searchFuncs := []struct {
		name string
		fn   func([]int, int) int
	}{
		{"search1 left closed right open", search1},
		{"search2 left closed right closed", search2},
		{"searchRecursive", searchRecursive},
	}

	for _, sf := range searchFuncs {
		t.Run(sf.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := sf.fn(tt.nums, tt.target)
					if got != tt.want {
						t.Errorf("%s(%v, %d) = %d, want %d", sf.name, tt.nums, tt.target, got, tt.want)
					}
				})
			}
		})
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		val  int
		want []int
	}{
		{"example", []int{3, 2, 2, 3}, 3, []int{2, 2}},
		{"longer example", []int{0, 1, 2, 2, 3, 0, 4, 2}, 2, []int{0, 1, 3, 0, 4}},
		{"empty input", []int{}, 3, []int{}},
		{"single element removed", []int{3}, 3, []int{}},
		{"single element kept", []int{4}, 3, []int{4}},
		{"all removed", []int{2, 2, 2}, 2, []int{}},
		{"none removed", []int{1, 2, 3}, 4, []int{1, 2, 3}},
		{"val at head only", []int{3, 1, 1}, 3, []int{1, 1}},
		{"val at tail only", []int{1, 1, 3}, 3, []int{1, 1}},
	}

	// keepsOrder is false for the opposing-pointer solution, which may reorder survivors.
	// 相向双指针解法可能改变保留元素的顺序，因此 keepsOrder 为 false。
	removeFuncs := []struct {
		name       string
		fn         func([]int, int) int
		keepsOrder bool
	}{
		{"brute force", removeElementBruteForce, true},
		{"fast slow", removeElementFastSlow, true},
		{"two pointers", removeElementTwoPointers, false},
	}

	for _, rf := range removeFuncs {
		t.Run(rf.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					nums := append([]int(nil), tt.nums...)
					gotLen := rf.fn(nums, tt.val)
					if gotLen != len(tt.want) {
						t.Fatalf("got length = %d, want %d", gotLen, len(tt.want))
					}

					got := nums[:gotLen]
					if rf.keepsOrder {
						if !equalInts(got, tt.want) {
							t.Fatalf("got nums[:%d] = %v, want %v", gotLen, got, tt.want)
						}
						return
					}

					if !equalInts(sortedCopy(got), sortedCopy(tt.want)) {
						t.Fatalf("got nums[:%d] = %v, want the same values as %v", gotLen, got, tt.want)
					}
				})
			}
		})
	}
}

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"mixed", []int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{"all negative", []int{-7, -3, -1}, []int{1, 9, 49}},
		{"all positive", []int{1, 2, 3}, []int{1, 4, 9}},
		{"empty", []int{}, []int{}},
		{"single negative", []int{-2}, []int{4}},
		{"single zero", []int{0}, []int{0}},
		{"duplicates around zero", []int{-3, -3, 0, 3, 3}, []int{0, 9, 9, 9, 9}},
	}

	squareFuncs := []struct {
		name string
		fn   func([]int) []int
	}{
		{"sort after squaring", sortedSquares},
		{"two pointers", sortedSquares_TwoPointers},
		{"merge around sign boundary", sortedSquaresMerge},
	}

	for _, sf := range squareFuncs {
		t.Run(sf.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := sf.fn(append([]int(nil), tt.nums...))
					if !equalInts(got, tt.want) {
						t.Errorf("got %v, want %v", got, tt.want)
					}
				})
			}
		})
	}
}

// sortedSquares squares in place, while the other two must leave the input untouched.
// sortedSquares 会原地平方，另外两种解法必须保持输入不变。
func TestSortedSquaresDoesNotModifyInput(t *testing.T) {
	squareFuncs := []struct {
		name string
		fn   func([]int) []int
	}{
		{"two pointers", sortedSquares_TwoPointers},
		{"merge around sign boundary", sortedSquaresMerge},
	}

	for _, sf := range squareFuncs {
		t.Run(sf.name, func(t *testing.T) {
			nums := []int{-4, -1, 0, 3, 10}
			want := append([]int(nil), nums...)
			sf.fn(nums)
			if !equalInts(nums, want) {
				t.Errorf("input changed to %v, want %v", nums, want)
			}
		})
	}
}

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		name   string
		target int
		nums   []int
		want   int
	}{
		{"example", 7, []int{2, 3, 1, 2, 4, 3}, 2},
		{"single exact", 4, []int{1, 4, 4}, 1},
		{"not found", 11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},
		{"whole array", 15, []int{1, 2, 3, 4, 5}, 5},
		{"empty input", 1, []int{}, 0},
		{"single element enough", 3, []int{5}, 1},
		{"single element too small", 6, []int{5}, 0},
		{"first element already enough", 4, []int{9, 1, 1}, 1},
		{"last element already enough", 4, []int{1, 1, 9}, 1},
	}

	minFuncs := []struct {
		name string
		fn   func(int, []int) int
	}{
		{"sliding window", minSubArrayLen},
		{"brute force", minSubArrayLenBruteForce},
		{"prefix sum with binary search", minSubArrayLenBinarySearch},
	}

	for _, mf := range minFuncs {
		t.Run(mf.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := mf.fn(tt.target, tt.nums)
					if got != tt.want {
						t.Errorf("%s(%d, %v) = %d, want %d", mf.name, tt.target, tt.nums, got, tt.want)
					}
				})
			}
		})
	}
}

func TestGenerateMatrix(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want [][]int
	}{
		{
			name: "n=1",
			n:    1,
			want: [][]int{{1}},
		},
		{
			name: "n=2",
			n:    2,
			want: [][]int{
				{1, 2},
				{4, 3},
			},
		},
		{
			name: "n=3",
			n:    3,
			want: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
		{
			name: "n=4",
			n:    4,
			want: [][]int{
				{1, 2, 3, 4},
				{12, 13, 14, 5},
				{11, 16, 15, 6},
				{10, 9, 8, 7},
			},
		},
		{
			name: "n=5",
			n:    5,
			want: [][]int{
				{1, 2, 3, 4, 5},
				{16, 17, 18, 19, 6},
				{15, 24, 25, 20, 7},
				{14, 23, 22, 21, 8},
				{13, 12, 11, 10, 9},
			},
		},
	}

	matrixFuncs := []struct {
		name string
		fn   func(int) [][]int
	}{
		{"shrinking boundaries", generateMatrix},
		{"direction simulation", generateMatrixSimulation},
		{"ring by ring", generateMatrixByLayers},
	}

	for _, mf := range matrixFuncs {
		t.Run(mf.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := mf.fn(tt.n)
					if !reflect.DeepEqual(got, tt.want) {
						t.Errorf("%s(%d) = %v, want %v", mf.name, tt.n, got, tt.want)
					}
				})
			}
		})
	}
}

func TestRangeSumBruteForce(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	tests := []struct {
		name  string
		left  int
		right int
		want  int
	}{
		{"middle range", 1, 3, 9},
		{"whole array", 0, 4, 15},
		{"single element", 2, 2, 3},
		{"first element", 0, 0, 1},
		{"last element", 4, 4, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rangeSumBruteForce(nums, tt.left, tt.right)
			if got != tt.want {
				t.Errorf("rangeSumBruteForce(%v, %d, %d) = %d, want %d", nums, tt.left, tt.right, got, tt.want)
			}
		})
	}
}

func TestPrefixSum(t *testing.T) {
	input := `5
1 2 3 4 5
1 3
0 4
2 2
`

	want := `9
15
3
`

	oldStdin := os.Stdin
	oldStdout := os.Stdout

	defer func() {
		os.Stdin = oldStdin
		os.Stdout = oldStdout
	}()

	rIn, wIn, _ := os.Pipe()
	_, _ = wIn.Write([]byte(input))
	_ = wIn.Close()
	os.Stdin = rIn

	rOut, wOut, _ := os.Pipe()
	os.Stdout = wOut

	prefixSum()

	_ = wOut.Close()

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, rOut)

	got := buf.String()
	if got != want {
		t.Errorf("got:\n%s\nwant:\n%s", got, want)
	}
}
