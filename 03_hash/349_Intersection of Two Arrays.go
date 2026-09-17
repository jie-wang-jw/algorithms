package _3_hash

import "sort"

/*
题目描述 / Problem Description
给定两个整数数组 nums1 和 nums2，返回它们的交集。结果中的每个元素必须唯一，返回顺序不限。
Given two integer arrays nums1 and nums2, return their intersection.
Every result value must be unique, and the order does not matter.
*/

// 1. Hash set: store nums1, then record a nums2 hit and delete it so each value appears once. Recommended.
// 1. 哈希集合：先存 nums1，再扫 nums2，命中后删除以保证结果唯一。推荐。
// Time: O(n+m) expected, Space: O(n) for the set.
// 时间复杂度：期望 O(n+m)，空间复杂度：O(n)，由集合产生。
//
// set 保存 nums1 中尚未输出的值；命中后删除，使 nums2 的重复值不会产生重复答案。
// set contains nums1 values not yet emitted; deletion after a match suppresses duplicates from nums2.
func intersection(nums1 []int, nums2 []int) []int {
	set := map[int]bool{}
	result := []int{}

	for _, num := range nums1 {
		set[num] = true
	}

	for _, num := range nums2 {
		if set[num] {
			result = append(result, num)
			delete(set, num)
		}
	}

	return result
}

// 2. Sort + two pointers: advance the smaller side; mutates both inputs.
// 2. 排序 + 双指针：较小一侧前进，相等则记录；会修改输入顺序。
// Time: O(n log n+m log m), Space: O(log(n+1)+log(m+1)) for the sort stacks plus O(r) output.
// 时间复杂度：O(n log n+m log m)，空间复杂度：排序栈 O(log(n+1)+log(m+1))，结果 O(r)。
//
// 排序两数组后，较小端不可能与对面当前及后续值相等，可直接前进；相等时两端同时前进。
// After sorting both inputs, the smaller value cannot match the other side's remaining values; advance it, or both on equality.
// 相等值连续出现，只有与结果末项不同时才追加；会修改两份输入顺序。
// Equal values are adjacent, so append only if different from the last result; both input orders are modified.
func intersectionSorted(nums1, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	result := []int{}

	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			if len(result) == 0 || result[len(result)-1] != nums1[i] {
				result = append(result, nums1[i])
			}
			i++
			j++
		}
	}

	return result
}
