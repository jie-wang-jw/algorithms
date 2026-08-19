package _3_hash

import "sort"

/*排序 + 固定一个数 + 双指针
Sort + fix one number + two pointers
sort the array first.
Then I fix one number and use two pointers on the remaining range.
If the sum is too small, I move the left pointer to increase it.
If the sum is too large, I move the right pointer to decrease it.
I also skip duplicate values to avoid returning the same triplet multiple times.
*/

func threeSum(nums []int) [][]int {
	/*先排序。排序以后，相同数字会挨在一起，也方便双指针移动。*/
	sort.Ints(nums)

	res := [][]int{}

	/*i 是固定第一个数 a 的位置。
	为什么是 len(nums)-2？
	因为后面还要留两个位置给 left 和 right*/
	for i := 0; i < len(nums)-2; i++ {
		a := nums[i]

		/*如果 a 已经大于 0，后面的数也只会更大。*/
		if a > 0 {
			break
		}

		/*这是给 a 去重 如果当前 a 和前一个 a 一样，就跳过。*/
		if i > 0 && a == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1

		for left < right {
			b, c := nums[left], nums[right]
			sum := a + b + c

			if sum == 0 {
				res = append(res, []int{a, b, c})

				/*跳过重复的 b c*/
				for left < right && nums[left] == b {
					left++
				}

				for left < right && nums[right] == c {
					right--
				}
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}

	return res
}
