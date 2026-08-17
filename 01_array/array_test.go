package _1_array

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"testing"
)

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
		{"empty slice", []int{}, 1, -1},
		{"single element exists", []int{5}, 5, 0},
		{"single element not exists", []int{5}, 3, -1},
	}

	searchFuncs := []struct {
		name string
		fn   func([]int, int) int
	}{
		{"search1 left closed right open", search1},
		{"search2 left closed right closed", search2},
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
		fn   func([]int, int) int
		nums []int
		val  int
		want []int
	}{
		{"brute force", removeElementBruteForce, []int{3, 2, 2, 3}, 3, []int{2, 2}},
		{"fast slow", removeElementFastSlow, []int{3, 2, 2, 3}, 3, []int{2, 2}},
		{"two pointers", removeElementTwoPointers, []int{3, 2, 2, 3}, 3, []int{2, 2}},
		{"fast slow more cases", removeElementFastSlow, []int{0, 1, 2, 2, 3, 0, 4, 2}, 2, []int{0, 1, 3, 0, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := append([]int(nil), tt.nums...)
			gotLen := tt.fn(nums, tt.val)
			if gotLen != len(tt.want) {
				t.Fatalf("got length = %d, want %d", gotLen, len(tt.want))
			}
			if !reflect.DeepEqual(nums[:gotLen], tt.want) {
				t.Fatalf("got nums[:%d] = %v, want %v", gotLen, nums[:gotLen], tt.want)
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
	}

	squareFuncs := []struct {
		name string
		fn   func([]int) []int
	}{
		{"brute force", sortedSquares},
		{"two pointers", sortedSquares_TwoPointers},
	}

	for _, sf := range squareFuncs {
		t.Run(sf.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := sf.fn(append([]int(nil), tt.nums...))
					if !reflect.DeepEqual(got, tt.want) {
						t.Errorf("got %v, want %v", got, tt.want)
					}
				})
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minSubArrayLen(tt.target, tt.nums)
			if got != tt.want {
				t.Errorf("minSubArrayLen(%d, %v) = %d, want %d", tt.target, tt.nums, got, tt.want)
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateMatrix(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateMatrix(%d) = %v, want %v", tt.n, got, tt.want)
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
