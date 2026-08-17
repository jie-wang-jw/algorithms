package _1_array

import (
	"sort"
)

// 排序法
func sortedSquares(nums []int) []int {
	for i, val := range nums {
		nums[i] *= val
	}
	sort.Ints(nums)
	return nums
}

// 双指针法
func sortedSquares_TwoPointers(nums []int) []int {
	n := len(nums)
	i,j,k := 0,n-1,n-1
	ans := make([]int, n)

	for i<=j {
		lm,rm:=nums[i]*nums[i],nums[j]*nums[j]
		if lm>rm{
			ans[k] = lm
			i++
		}else{
			ans[k] = rm
			j--
		}
		k--
	}
	return ans
}
