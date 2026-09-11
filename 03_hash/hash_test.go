package _3_hash

import (
	"reflect"
	"sort"
	"testing"
)

func normalizeIntGroups(groups [][]int) [][]int {
	for _, group := range groups {
		sort.Ints(group)
	}
	sort.Slice(groups, func(i, j int) bool {
		for k := 0; k < len(groups[i]) && k < len(groups[j]); k++ {
			if groups[i][k] != groups[j][k] {
				return groups[i][k] < groups[j][k]
			}
		}
		return len(groups[i]) < len(groups[j])
	})
	return groups
}

// The problem does not fix the order of the two indices, so compare them sorted.
// 题目不限定两个下标的先后顺序，因此排序后再比较。
func normalizeIndexPair(indices []int) []int {
	if indices == nil {
		return nil
	}
	pair := append([]int(nil), indices...)
	sort.Ints(pair)
	return pair
}

func copyInts(nums []int) []int {
	return append([]int(nil), nums...)
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{name: "example", nums: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		{name: "middle pair", nums: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
		{name: "same value different indices", nums: []int{3, 3}, target: 6, want: []int{0, 1}},
		{name: "last two values", nums: []int{5, 1, 4, 6}, target: 10, want: []int{2, 3}},
		{name: "negative values", nums: []int{-3, 4, 3, 90}, target: 0, want: []int{0, 2}},
		{name: "no solution", nums: []int{1, 2, 3}, target: 100, want: nil},
		{name: "single value cannot pair", nums: []int{4}, target: 8, want: nil},
		{name: "empty input", nums: []int{}, target: 0, want: nil},
	}

	twoSumFuncs := []struct {
		name string
		fn   func([]int, int) []int
	}{
		{name: "hash map solution", fn: twoSum},
		{name: "brute-force solution", fn: twoSumBruteForce},
		{name: "sorting and two pointers solution", fn: twoSumSorted},
	}

	for _, sf := range twoSumFuncs {
		t.Run(sf.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := normalizeIndexPair(sf.fn(copyInts(tt.nums), tt.target))
					if !reflect.DeepEqual(got, tt.want) {
						t.Fatalf("%s(%v, %d) = %v, want %v", sf.name, tt.nums, tt.target, got, tt.want)
					}
				})
			}
		})
	}
}

func TestIsHappy(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{name: "one is happy", n: 1, want: true},
		{name: "example happy number", n: 19, want: true},
		{name: "cycle number", n: 2, want: false},
		{name: "another happy number", n: 7, want: true},
		{name: "cycle entry point", n: 4, want: false},
		{name: "larger happy number", n: 100, want: true},
	}

	isHappyFuncs := []struct {
		name string
		fn   func(int) bool
	}{
		{name: "hash set solution", fn: isHappy},
		{name: "Floyd cycle detection solution", fn: isHappyFloyd},
	}

	for _, hf := range isHappyFuncs {
		t.Run(hf.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := hf.fn(tt.n); got != tt.want {
						t.Fatalf("%s(%d) = %v, want %v", hf.name, tt.n, got, tt.want)
					}
				})
			}
		})
	}
}

func TestGetNext(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{n: 19, want: 82},
		{n: 82, want: 68},
		{n: 100, want: 1},
		{n: 1, want: 1},
	}

	for _, tt := range tests {
		if got := getNext(tt.n); got != tt.want {
			t.Fatalf("getNext(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestIsAnagram(t *testing.T) {
	// All four solutions must agree on lowercase input, which is what the problem guarantees.
	// 四种解法在题目保证的小写字母输入上结果必须一致。
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{name: "same letters different order", s: "anagram", t: "nagaram", want: true},
		{name: "different counts", s: "rat", t: "car", want: false},
		{name: "different lengths", s: "ab", t: "a", want: false},
		{name: "same repeated letters", s: "aabb", t: "baba", want: true},
		{name: "same letter different counts", s: "aab", t: "abb", want: false},
		{name: "single character equal", s: "a", t: "a", want: true},
		{name: "single character different", s: "a", t: "b", want: false},
		{name: "both empty", s: "", t: "", want: true},
	}

	isAnagramFuncs := []struct {
		name string
		fn   func(string, string) bool
	}{
		{name: "two-pass counting solution", fn: isAnagram1},
		{name: "single-pass counting solution", fn: isAnagram2},
		{name: "sorting solution", fn: isAnagramSorted},
		{name: "rune map solution", fn: isAnagramUnicode},
	}

	for _, af := range isAnagramFuncs {
		t.Run(af.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := af.fn(tt.s, tt.t); got != tt.want {
						t.Fatalf("%s(%q, %q) = %v, want %v", af.name, tt.s, tt.t, got, tt.want)
					}
				})
			}
		})
	}
}

func TestIsAnagramBeyondLowercase(t *testing.T) {
	// isAnagram1 and isAnagram2 index a 26-slot array, so only these two solutions accept
	// characters outside a-z.
	// isAnagram1 和 isAnagram2 使用 26 格数组下标，因此只有这两种解法能接受 a-z 之外的字符。
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{name: "multibyte code points reordered", s: "汉字词", t: "词汉字", want: true},
		{name: "different multibyte characters", s: "汉字", t: "汉汉", want: false},
		{name: "accented letters reordered", s: "café", t: "éfac", want: true},
		{name: "uppercase is a different character", s: "Ab", t: "ab", want: false},
		{name: "digits and symbols", s: "a1!", t: "!1a", want: true},
	}

	unicodeFuncs := []struct {
		name string
		fn   func(string, string) bool
	}{
		{name: "sorting solution", fn: isAnagramSorted},
		{name: "rune map solution", fn: isAnagramUnicode},
	}

	for _, uf := range unicodeFuncs {
		t.Run(uf.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := uf.fn(tt.s, tt.t); got != tt.want {
						t.Fatalf("%s(%q, %q) = %v, want %v", uf.name, tt.s, tt.t, got, tt.want)
					}
				})
			}
		})
	}
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{name: "deduplicate result", nums1: []int{1, 2, 2, 1}, nums2: []int{2, 2}, want: []int{2}},
		{name: "multiple common values", nums1: []int{4, 9, 5}, nums2: []int{9, 4, 9, 8, 4}, want: []int{4, 9}},
		{name: "no intersection", nums1: []int{1, 3, 5}, nums2: []int{2, 4, 6}, want: []int{}},
		{name: "negative numbers", nums1: []int{-1, 0, 1}, nums2: []int{1, -1, -1}, want: []int{-1, 1}},
		{name: "duplicates on both sides", nums1: []int{3, 3, 3}, nums2: []int{3, 3}, want: []int{3}},
		{name: "empty first array", nums1: []int{}, nums2: []int{1, 2}, want: []int{}},
		{name: "both empty", nums1: []int{}, nums2: []int{}, want: []int{}},
	}

	intersectionFuncs := []struct {
		name string
		fn   func([]int, []int) []int
	}{
		{name: "hash set solution", fn: intersection},
		{name: "sorting and two pointers solution", fn: intersectionSorted},
	}

	for _, inf := range intersectionFuncs {
		t.Run(inf.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					// intersectionSorted sorts its inputs, so hand both solutions copies.
					// intersectionSorted 会排序输入，因此统一传入副本。
					got := inf.fn(copyInts(tt.nums1), copyInts(tt.nums2))
					sort.Ints(got)
					sort.Ints(tt.want)

					if !reflect.DeepEqual(got, tt.want) {
						t.Fatalf("%s(%v, %v) = %v, want %v", inf.name, tt.nums1, tt.nums2, got, tt.want)
					}
				})
			}
		})
	}
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{name: "example with duplicates", nums: []int{-1, 0, 1, 2, -1, -4}, want: [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{name: "all zeros deduplicated", nums: []int{0, 0, 0, 0}, want: [][]int{{0, 0, 0}}},
		{name: "duplicate heavy input", nums: []int{-2, 0, 0, 2, 2}, want: [][]int{{-2, 0, 2}}},
		{name: "several triplets", nums: []int{-4, -2, -2, 0, 2, 2, 4}, want: [][]int{{-4, 0, 4}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2}}},
		{name: "no triplet", nums: []int{1, 2, -2, -1}, want: [][]int{}},
		{name: "all positive", nums: []int{1, 2, 3, 4}, want: [][]int{}},
		{name: "less than three numbers", nums: []int{0, 1}, want: [][]int{}},
		{name: "empty input", nums: []int{}, want: [][]int{}},
	}

	threeSumFuncs := []struct {
		name string
		fn   func([]int) [][]int
	}{
		{name: "two pointers solution", fn: threeSum},
		{name: "hash solution", fn: threeSumHash},
	}

	for _, tf := range threeSumFuncs {
		t.Run(tf.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					// Both solutions sort their input, so pass a copy each time.
					// 两种解法都会排序输入，因此每次都传入副本。
					got := normalizeIntGroups(tf.fn(copyInts(tt.nums)))
					want := normalizeIntGroups(deepCopyIntGroups(tt.want))
					if !reflect.DeepEqual(got, want) {
						t.Fatalf("%s(%v) = %v, want %v", tf.name, tt.nums, got, want)
					}
				})
			}
		})
	}
}

func TestFourSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   [][]int
	}{
		{name: "example", nums: []int{1, 0, -1, 0, -2, 2}, target: 0, want: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{name: "all same values deduplicated", nums: []int{2, 2, 2, 2, 2}, target: 8, want: [][]int{{2, 2, 2, 2}}},
		{name: "negative target", nums: []int{-3, -1, 0, 2, 4, 5}, target: 2, want: [][]int{{-3, -1, 2, 4}}},
		// Pruning must not stop at nums[i] > target while nums[i] is still negative.
		// 剪枝不能在 nums[i] 仍为负数时就停止，否则会漏掉这个答案。
		{name: "negative target with negative first value", nums: []int{-9, -3, -2, 4}, target: -10, want: [][]int{{-9, -3, -2, 4}}},
		{name: "duplicate heavy input", nums: []int{0, 0, 0, 0, 0, 0}, target: 0, want: [][]int{{0, 0, 0, 0}}},
		{name: "no quadruplet", nums: []int{1, 2, 3}, target: 6, want: [][]int{}},
		{name: "all positive above target", nums: []int{5, 6, 7, 8}, target: 4, want: [][]int{}},
		// The first value stays below target, so only the pair-level pruning can stop the inner loop.
		// 第一个数没有超过 target，因此只有第二层的剪枝能结束内层循环。
		{name: "pair pruning stops inner loop", nums: []int{-1, 5, 6, 7, 8}, target: 3, want: [][]int{}},
		{name: "empty input", nums: []int{}, target: 0, want: [][]int{}},
	}

	fourSumFuncs := []struct {
		name string
		fn   func([]int, int) [][]int
	}{
		{name: "two pointers solution", fn: fourSum},
		{name: "hash solution", fn: fourSumHash},
	}

	for _, ff := range fourSumFuncs {
		t.Run(ff.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := normalizeIntGroups(ff.fn(copyInts(tt.nums), tt.target))
					want := normalizeIntGroups(deepCopyIntGroups(tt.want))
					if !reflect.DeepEqual(got, want) {
						t.Fatalf("%s(%v, %d) = %v, want %v", ff.name, tt.nums, tt.target, got, want)
					}
				})
			}
		})
	}
}

// normalizeIntGroups sorts in place, so each solution needs its own copy of the expectation.
// normalizeIntGroups 会原地排序，因此每种解法都需要一份独立的期望值副本。
func deepCopyIntGroups(groups [][]int) [][]int {
	copied := make([][]int, len(groups))
	for i, group := range groups {
		copied[i] = copyInts(group)
	}
	return copied
}

func TestFourSumCount(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		c    []int
		d    []int
		want int
	}{
		{name: "example", a: []int{1, 2}, b: []int{-2, -1}, c: []int{-1, 2}, d: []int{0, 2}, want: 2},
		{name: "all zeros", a: []int{0, 0}, b: []int{0}, c: []int{0}, d: []int{0, 0}, want: 4},
		{name: "no tuples", a: []int{1}, b: []int{1}, c: []int{1}, d: []int{1}, want: 0},
		{name: "duplicate values counted separately", a: []int{1, 1}, b: []int{-1, -1}, c: []int{0, 0}, d: []int{0, 0}, want: 16},
		{name: "single tuple", a: []int{-1}, b: []int{-1}, c: []int{1}, d: []int{1}, want: 1},
		{name: "empty array gives no tuples", a: []int{}, b: []int{1}, c: []int{2}, d: []int{3}, want: 0},
	}

	fourSumCountFuncs := []struct {
		name string
		fn   func([]int, []int, []int, []int) int
	}{
		{name: "grouped hash solution", fn: fourSumCount},
		{name: "sorted pair sums solution", fn: fourSumCountSorted},
	}

	for _, cf := range fourSumCountFuncs {
		t.Run(cf.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					if got := cf.fn(tt.a, tt.b, tt.c, tt.d); got != tt.want {
						t.Fatalf("%s(%v, %v, %v, %v) = %d, want %d", cf.name, tt.a, tt.b, tt.c, tt.d, got, tt.want)
					}
				})
			}
		})
	}
}
