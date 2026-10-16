package _8_greedy

import "sort"

/*
56. 合并区间 / Merge Intervals

题目描述 / Problem Description
合并所有重叠区间。端点相接也算重叠，例如 [1,4] 与 [4,5] 合并为 [1,5]。
Merge overlapping intervals. Touching endpoints merge, e.g. [1,4] and [4,5] become [1,5].

示例 / Examples
[[1,3],[2,6],[8,10],[15,18]] → [[1,6],[8,10],[15,18]]
[[1,4],[4,5]] → [[1,5]]
*/

// sortIntervalsByStart
//
// 原地按区间左端点升序排列；只调整外层切片的顺序，不复制各区间。
// Sort intervals in place by their left endpoints, reordering the outer slice without copying interval rows.
func sortIntervalsByStart(intervals [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
}

// 面试首选 / Interview first choice：56 合并区间。
//
// 1. Sort by start and merge into the last interval already stored in the result.
// 1. 按左边界排序，有重叠就更新结果集最后一个区间的右边界。
// Time: O(n log n), Space: O(n) for the answer plus sort.
// 时间复杂度：O(n log n)，空间复杂度：答案 O(n)，另加排序。
//
// 按起点排序后，只需与最后一个合并区间比较；相交或接触就扩展右端，分离时才能确定前一个结果。
// Sort by start and compare only with the latest merged interval; extend overlaps/touching endpoints and finalize on a gap.
// 新结果用独立的端点切片，原端点不改，但排序会改变输入外层顺序；跟踪边界版要在循环后补最后一段。
// Allocate result endpoint slices without changing original endpoints, but sorting reorders input; append the final tracked interval after the loop.
func merge(intervals [][]int) [][]int {
	result := [][]int{}
	if len(intervals) == 0 {
		return result
	}
	sortIntervalsByStart(intervals)
	result = append(result, []int{intervals[0][0], intervals[0][1]})
	for i := 1; i < len(intervals); i++ {
		if result[len(result)-1][1] >= intervals[i][0] {
			result[len(result)-1][1] = max(result[len(result)-1][1], intervals[i][1])
		} else {
			result = append(result, []int{intervals[i][0], intervals[i][1]})
		}
	}
	return result
}

// 2. Sort by start and keep the active merged range in left/right until a gap appears.
// 2. 按左边界排序，用 left/right 维护当前合并段，断开时写入结果。
// Time: O(n log n), Space: O(n) for the answer plus sort.
// 时间复杂度：O(n log n)，空间复杂度：答案 O(n)，另加排序。
//
// 按起点排序后，只需与最后一个合并区间比较；相交或接触就扩展右端，分离时才能确定前一个结果。
// Sort by start and compare only with the latest merged interval; extend overlaps/touching endpoints and finalize on a gap.
// 新结果用独立的端点切片，原端点不改，但排序会改变输入外层顺序；跟踪边界版要在循环后补最后一段。
// Allocate result endpoint slices without changing original endpoints, but sorting reorders input; append the final tracked interval after the loop.
func mergeTrackBounds(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	sortIntervalsByStart(intervals)
	res := make([][]int, 0, len(intervals))
	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if right < intervals[i][0] {
			res = append(res, []int{left, right})
			left, right = intervals[i][0], intervals[i][1]
		} else {
			right = max(right, intervals[i][1])
		}
	}
	res = append(res, []int{left, right})
	return res
}
