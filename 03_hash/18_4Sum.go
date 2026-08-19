package _3_hash

import "sort"

/*
a + b + c + d = target
3Sum 是：
固定 1 个数 + left/right 找 2 个数
4Sum 是：
固定 2 个数 + left/right 找 2 个数

I sort the array first.
Then I fix the first two numbers with two loops.
For the remaining part of the array, I use two pointers to find the other two numbers.
If the sum is smaller than the target, I move the left pointer to increase the sum.
If the sum is larger than the target, I move the right pointer to decrease the sum.
I skip duplicate values for each position to avoid duplicate quadruplets.
*/

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	res := [][]int{}

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left, right := j+1, len(nums)-1

			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]

				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})

					b, c := nums[left], nums[right]

					for left < right && nums[left] == b {
						left++
					}

					for left < right && nums[right] == c {
						right--
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}

	return res
}
