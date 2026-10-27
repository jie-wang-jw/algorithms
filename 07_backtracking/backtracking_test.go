package _7_backtracking

import (
	"slices"
	"testing"
)

func normalizeGroups(groups [][]int) [][]int {
	cloned := make([][]int, len(groups))
	for i, group := range groups {
		cloned[i] = slices.Clone(group)
	}
	slices.SortFunc(cloned, slices.Compare[[]int])
	return cloned
}

func TestCombine(t *testing.T) {
	implementations := []struct {
		name string
		fn   func(int, int) [][]int
	}{
		{name: "backtracking", fn: combine},
		{name: "pruned", fn: combinePruned},
	}
	tests := []struct {
		name string
		n, k int
		want [][]int
	}{
		{name: "article example", n: 4, k: 2, want: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
		{name: "k equals n", n: 3, k: 3, want: [][]int{{1, 2, 3}}},
		{name: "k is one", n: 3, k: 1, want: [][]int{{1}, {2}, {3}}},
		{name: "single", n: 1, k: 1, want: [][]int{{1}}},
	}
	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := normalizeGroups(implementation.fn(tt.n, tt.k))
					want := normalizeGroups(tt.want)
					if !slices.EqualFunc(got, want, slices.Equal[[]int]) {
						t.Fatalf("combine(%d,%d) = %v, want %v", tt.n, tt.k, got, want)
					}
				})
			}
		})
	}
}

func TestCombinationSum3(t *testing.T) {
	tests := []struct {
		name string
		k, n int
		want [][]int
	}{
		{name: "seven", k: 3, n: 7, want: [][]int{{1, 2, 4}}},
		{name: "nine", k: 3, n: 9, want: [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
		{name: "impossible", k: 4, n: 1, want: [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalizeGroups(combinationSum3(tt.k, tt.n))
			want := normalizeGroups(tt.want)
			if !slices.EqualFunc(got, want, slices.Equal[[]int]) {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		digits string
		want   []string
	}{
		{digits: "", want: []string{}},
		{digits: "2", want: []string{"a", "b", "c"}},
		{digits: "23", want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	}
	for _, tt := range tests {
		t.Run(tt.digits, func(t *testing.T) {
			got := slices.Clone(letterCombinations(tt.digits))
			want := slices.Clone(tt.want)
			slices.Sort(got)
			slices.Sort(want)
			if !slices.Equal(got, want) {
				t.Fatalf("got %q, want %q", got, want)
			}
		})
	}
}

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{name: "example", candidates: []int{2, 3, 6, 7}, target: 7, want: [][]int{{2, 2, 3}, {7}}},
		{name: "repeats", candidates: []int{2, 3, 5}, target: 8, want: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{name: "none", candidates: []int{2}, target: 1, want: [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalizeGroups(combinationSum(tt.candidates, tt.target))
			want := normalizeGroups(tt.want)
			if !slices.EqualFunc(got, want, slices.Equal[[]int]) {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}

func TestCombinationSum2(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{name: "duplicates", candidates: []int{10, 1, 2, 7, 6, 1, 5}, target: 8, want: [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
		{name: "all twos", candidates: []int{2, 5, 2, 1, 2}, target: 5, want: [][]int{{1, 2, 2}, {5}}},
		{name: "none", candidates: []int{1, 1}, target: 3, want: [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := slices.Clone(tt.candidates)
			got := normalizeGroups(combinationSum2(input, tt.target))
			want := normalizeGroups(tt.want)
			if !slices.EqualFunc(got, want, slices.Equal[[]int]) {
				t.Fatalf("got %v, want %v", got, want)
			}
			if !slices.Equal(input, tt.candidates) {
				t.Fatal("input was modified")
			}
		})
	}
}

func TestSubsets(t *testing.T) {
	got := normalizeGroups(subsets([]int{1, 2, 3}))
	want := normalizeGroups([][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}})
	if !slices.EqualFunc(got, want, slices.Equal[[]int]) {
		t.Fatalf("subsets = %v, want %v", got, want)
	}
	if got := normalizeGroups(subsets([]int{})); !slices.EqualFunc(got, [][]int{{}}, slices.Equal[[]int]) {
		t.Fatalf("empty subsets = %v", got)
	}
}

func TestSubsetsWithDup(t *testing.T) {
	input := []int{1, 2, 2}
	got := normalizeGroups(subsetsWithDup(input))
	want := normalizeGroups([][]int{{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2}})
	if !slices.EqualFunc(got, want, slices.Equal[[]int]) {
		t.Fatalf("subsetsWithDup = %v, want %v", got, want)
	}
	if !slices.Equal(input, []int{1, 2, 2}) {
		t.Fatal("input was modified")
	}
}

func TestPermute(t *testing.T) {
	got := normalizeGroups(permute([]int{1, 2, 3}))
	want := normalizeGroups([][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}})
	if !slices.EqualFunc(got, want, slices.Equal[[]int]) {
		t.Fatalf("permute = %v, want %v", got, want)
	}
}

func TestPermuteUnique(t *testing.T) {
	input := []int{1, 1, 2}
	got := normalizeGroups(permuteUnique(input))
	want := normalizeGroups([][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}})
	if !slices.EqualFunc(got, want, slices.Equal[[]int]) {
		t.Fatalf("permuteUnique = %v, want %v", got, want)
	}
}

func TestPartition(t *testing.T) {
	got := partition("aab")
	want := [][]string{{"a", "a", "b"}, {"aa", "b"}}
	if !equalStringGroups(got, want) {
		t.Fatalf("partition = %v, want %v", got, want)
	}
	if got := partition("a"); !equalStringGroups(got, [][]string{{"a"}}) {
		t.Fatalf("single = %v", got)
	}
}

func TestRestoreIpAddresses(t *testing.T) {
	tests := []struct {
		s    string
		want []string
	}{
		{s: "25525511135", want: []string{"255.255.11.135", "255.255.111.35"}},
		{s: "0000", want: []string{"0.0.0.0"}},
		{s: "101023", want: []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"}},
		{s: "1111", want: []string{"1.1.1.1"}},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := slices.Clone(restoreIpAddresses(tt.s))
			want := slices.Clone(tt.want)
			slices.Sort(got)
			slices.Sort(want)
			if !slices.Equal(got, want) {
				t.Fatalf("got %q, want %q", got, want)
			}
		})
	}
}

func TestSolveNQueens(t *testing.T) {
	got := solveNQueens(4)
	want := [][]string{
		{".Q..", "...Q", "Q...", "..Q."},
		{"..Q.", "Q...", "...Q", ".Q.."},
	}
	if !equalStringGroups(got, want) {
		t.Fatalf("n=4 got %q, want %q", got, want)
	}
	if got := solveNQueens(1); !equalStringGroups(got, [][]string{{"Q"}}) {
		t.Fatalf("n=1 got %q", got)
	}
}

func TestFindSubsequences(t *testing.T) {
	implementations := []struct {
		name string
		fn   func([]int) [][]int
	}{
		{name: "hash set", fn: findSubsequences},
		{name: "array hash", fn: findSubsequencesArray},
	}
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{name: "article example", nums: []int{4, 6, 7, 7}, want: [][]int{{4, 6}, {4, 7}, {4, 6, 7}, {4, 6, 7, 7}, {6, 7}, {6, 7, 7}, {7, 7}, {4, 7, 7}}},
		{name: "unsorted", nums: []int{4, 7, 6, 7}, want: [][]int{{4, 7}, {4, 7, 7}, {4, 6}, {4, 6, 7}, {7, 7}, {6, 7}}},
		{name: "two equals", nums: []int{4, 4, 3, 2, 1}, want: [][]int{{4, 4}}},
		{name: "strictly decreasing", nums: []int{3, 2, 1}, want: [][]int{}},
		{name: "single", nums: []int{1}, want: [][]int{}},
	}
	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := normalizeGroups(implementation.fn(tt.nums))
					want := normalizeGroups(tt.want)
					if !slices.EqualFunc(got, want, slices.Equal[[]int]) {
						t.Fatalf("got %v, want %v", got, want)
					}
				})
			}
		})
	}
}

func TestFindItinerary(t *testing.T) {
	implementations := []struct {
		name string
		fn   func([][]string) []string
	}{
		{name: "backtracking counts", fn: findItinerary},
		{name: "Hierholzer", fn: findItineraryHierholzer},
	}
	tests := []struct {
		name    string
		tickets [][]string
		want    []string
	}{
		{name: "example one", tickets: [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}, want: []string{"JFK", "MUC", "LHR", "SFO", "SJC"}},
		{name: "example two", tickets: [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}, want: []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"}},
		{name: "dead end first in lexical order", tickets: [][]string{{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"}}, want: []string{"JFK", "NRT", "JFK", "KUL"}},
		{name: "single ticket", tickets: [][]string{{"JFK", "ATL"}}, want: []string{"JFK", "ATL"}},
	}
	for _, implementation := range implementations {
		t.Run(implementation.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					input := cloneTickets(tt.tickets)
					got := implementation.fn(input)
					if !slices.Equal(got, tt.want) {
						t.Fatalf("got %q, want %q", got, tt.want)
					}
				})
			}
		})
	}
}

func TestSolveSudoku(t *testing.T) {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	want := [][]byte{
		{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
		{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
		{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
		{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
		{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
		{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
		{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
		{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
		{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
	}

	solveSudoku(board)
	if !equalByteBoard(board, want) {
		t.Fatalf("solveSudoku filled\n%q\nwant\n%q", board, want)
	}

	already := [][]byte{
		{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
		{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
		{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
		{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
		{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
		{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
		{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
		{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
		{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
	}
	clone := cloneBoard(already)
	solveSudoku(clone)
	if !equalByteBoard(clone, already) {
		t.Fatal("a completed board must stay unchanged")
	}
}

func cloneTickets(tickets [][]string) [][]string {
	cloned := make([][]string, len(tickets))
	for i, ticket := range tickets {
		cloned[i] = slices.Clone(ticket)
	}
	return cloned
}

func cloneBoard(board [][]byte) [][]byte {
	cloned := make([][]byte, len(board))
	for i := range board {
		cloned[i] = slices.Clone(board[i])
	}
	return cloned
}

func equalByteBoard(got, want [][]byte) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if !slices.Equal(got[i], want[i]) {
			return false
		}
	}
	return true
}

func equalStringGroups(got, want [][]string) bool {
	if len(got) != len(want) {
		return false
	}
	g := append([][]string(nil), got...)
	w := append([][]string(nil), want...)
	slices.SortFunc(g, slices.Compare[[]string])
	slices.SortFunc(w, slices.Compare[[]string])
	return slices.EqualFunc(g, w, slices.Equal[[]string])
}
