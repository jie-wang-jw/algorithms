package _3_hash

import "sort"

/*
题目描述 / Problem Description
给定两个整数数组 nums1 和 nums2，返回它们的交集。结果中的每个元素必须唯一，返回顺序不限。
Given two integer arrays nums1 and nums2, return their intersection.
Every result value must be unique, and the order does not matter.

解题思路 / Solution Approach
先把 nums1 的元素放入哈希集合，再遍历 nums2。命中集合时加入结果并从集合删除，
确保相同数字只加入一次。
Put all values from nums1 into a hash set, then scan nums2.
When a value is found, append it and remove it from the set so it can appear only once.

时间与空间复杂度 / Time and Space Complexity
n、m 为两个数组长度，r 为交集大小。平均时间 O(n+m)，分别扫描一次；辅助空间 O(n)，
集合保存第一个数组的不同数字。结果占 O(r)，r <= min(n,m)，总额外空间仍为 O(n)。
For input lengths n and m and intersection size r, average time is O(n+m) for two scans.
Auxiliary space O(n) stores distinct values from the first array. Output takes O(r),
where r <= min(n,m), so total extra space remains O(n).

补充解法：排序 + 双指针 / Alternative: Sorting and Two Pointers
intersectionSorted 排序两个输入。小值不可能匹配对方当前及后续更大值，故只移动较小一侧。
相等时记录并移动两侧；仅当与上次答案不同才追加，保证结果唯一。
时间 O(n log n+m log m)，辅助排序栈 O(log(n+1)+log(m+1))，输出 O(r)；会修改两个输入的顺序。
Sort both inputs. Advance the smaller value: it cannot match the other side's current or later values.
On equality advance both, appending only a value different from the previous output.
Time O(n log n+m log m), sorting stack O(log(n+1)+log(m+1)), output O(r); mutates both inputs.
*/

/*
I use a hash set to store all numbers from the first array.
我使用哈希集合保存第一个数组中的所有数字。
Then I iterate through the second array.
然后遍历第二个数组。
If a number exists in the set, I add it to the result and remove it from the set to avoid duplicates.
如果数字存在于集合中，就把它加入结果，并从集合删除以避免重复结果。
*/

// 1. 哈希集合：推荐
func intersection(nums1 []int, nums2 []int) []int {
	// The map acts as a set; only key existence matters.
	// 这里把 map 当作集合使用，只关心 key 是否存在。
	set := map[int]bool{}
	result := []int{}

	for _, num := range nums1 {
		// Repeated values simply overwrite true, so the set remains unique.
		// 重复数字只会再次写入 true，因此集合中仍然只有一份。
		set[num] = true
	}

	for _, num := range nums2 {
		// A true value means num appeared in nums1 and has not been used yet.
		// true 表示 num 在 nums1 中出现过，并且还没有加入结果。
		if set[num] {
			result = append(result, num)
			// Writing false instead of deleting the key has the same effect: the value is already used.
			// 把 value 改写成 false 和删除 key 的效果相同，都表示这个数字已经用过。
			//set[num] = false
			// Delete after use so duplicates in nums2 cannot be appended again.
			// 使用后删除，这样 nums2 中重复的 num 不会再次加入结果。
			delete(set, num)
		}
	}

	return result
}

// 2. 排序 + 双指针：会修改输入顺序
func intersectionSorted(nums1, nums2 []int) []int {
	// Sorting both sides lets a single forward scan compare values in increasing order.
	// 两侧都排序后，就可以按数值递增的顺序单向扫描比较。
	sort.Ints(nums1)
	sort.Ints(nums2)

	result := []int{}

	// i walks nums1 and j walks nums2; neither pointer ever moves backward.
	// i 扫描 nums1，j 扫描 nums2，两个指针都不会回退。
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] < nums2[j] {
			// The smaller value cannot match nums2[j] or anything larger after it, so drop it.
			// 较小的值配不上 nums2[j]，也配不上它后面更大的值，直接丢弃。
			i++
		} else if nums1[i] > nums2[j] {
			// Symmetric case: nums2[j] is too small to appear in the rest of nums1.
			// 对称情况：nums2[j] 太小，不会出现在 nums1 的剩余部分中。
			j++
		} else {
			// Equal values are part of the intersection, but each value may be recorded only once.
			// 相等的值属于交集，但每个值只能记录一次。
			// Duplicates are adjacent after sorting, so comparing with the last output is enough.
			// 排序后重复值相邻，因此只需要和上一个答案比较即可去重。
			if len(result) == 0 || result[len(result)-1] != nums1[i] {
				result = append(result, nums1[i])
			}
			// This value is settled on both sides, so advance both pointers.
			// 这个值在两侧都处理完了，所以两个指针同时前进。
			i++
			j++
		}
	}

	return result
}
