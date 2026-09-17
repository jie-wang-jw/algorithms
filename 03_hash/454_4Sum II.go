package _3_hash

import "sort"

/*
题目描述 / Problem Description
给定四个整数数组 nums1、nums2、nums3 和 nums4，
统计满足 nums1[i] + nums2[j] + nums3[k] + nums4[l] = 0 的下标四元组数量。
Given four integer arrays nums1, nums2, nums3, and nums4,
count index tuples satisfying nums1[i] + nums2[j] + nums3[k] + nums4[l] = 0.
*/

// 1. Split-group hashing: store every A+B sum in a map, then look up -(C+D). Recommended.
// 1. 分组哈希：前两数和存 map，再查找后两数和的相反数。推荐。
// Time: O(n²) expected, Space: O(n²) for the pair-sum map.
// 时间复杂度：期望 O(n²)，空间复杂度：O(n²)，由两数和哈希表产生。
//
// sumCount[x] 保存 A、B 中和为 x 的下标对数量；每个 C、D 下标对贡献 sumCount[-(c+d)] 个四元组。
// sumCount[x] counts A/B index pairs summing to x; each C/D pair contributes sumCount[-(c+d)] quadruples.
// 这里统计的是下标组合而非不同数值组合，重复值的出现次数必须保留。
// Count index combinations, not distinct value tuples; multiplicities must be retained.
func fourSumCount(A []int, B []int, C []int, D []int) int {
	sumCount := map[int]int{}
	count := 0

	for _, a := range A {
		for _, b := range B {
			sumCount[a+b]++
		}
	}

	for _, c := range C {
		for _, d := range D {
			need := -(c + d)
			count += sumCount[need]
		}
	}

	return count
}

// 2. Sort both pair-sum arrays + two pointers: keep multiplicities and add x*y on a match.
// 2. 两组和排序 + 双指针：保留重复次数，命中时累加 x*y。
// Time: O(n² log(n+1)), Space: O(n²).
// 时间复杂度：O(n² log(n+1))，空间复杂度：O(n²)。
//
// 生成并排序两侧所有配对和；一侧从小到大、另一侧从大到小寻找总和为零的配对。
// Sort all pair sums and scan the left side upward and the right downward for zero totals.
// 若相反数分别出现 a、b 次，贡献 a*b 个下标组合；必须整段计数，不能只加一。
// Opposite sums occurring a and b times contribute a*b index combinations, not one.
func fourSumCountSorted(a, b, c, d []int) int {
	leftSums, rightSums := []int{}, []int{}

	for _, x := range a {
		for _, y := range b {
			leftSums = append(leftSums, x+y)
		}
	}
	for _, x := range c {
		for _, y := range d {
			rightSums = append(rightSums, x+y)
		}
	}

	sort.Ints(leftSums)
	sort.Ints(rightSums)

	result := 0

	for i, j := 0, len(rightSums)-1; i < len(leftSums) && j >= 0; {
		sum := leftSums[i] + rightSums[j]

		if sum < 0 {
			i++
		} else if sum > 0 {
			j--
		} else {
			x, y := leftSums[i], rightSums[j]
			startI, startJ := i, j

			for i < len(leftSums) && leftSums[i] == x {
				i++
			}
			for j >= 0 && rightSums[j] == y {
				j--
			}

			result += (i - startI) * (startJ - j)
		}
	}

	return result
}
