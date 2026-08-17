package _1_array

func removeElementBruteForce(nums []int, val int) int {
	size := len(nums)
	for i := 0; i < size; i++ {
		if nums[i] == val {
			for j := i + 1; j < size; j++ {
				nums[j-1] = nums[j]
			}
			i--
			size--
		}
	}

	return size
}

func removeElementFastSlow(nums []int, val int) int {
	slow := 0

	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}

func removeElementTwoPointers(nums []int, val int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		for left <= right && nums[left] != val {
			left++
		}

		for left <= right && nums[right] == val {
			right--
		}
		// 找到后用右侧的非 val 覆盖左侧的 val
		if left < right {
			nums[left] = nums[right]
			left++
			right--
		}
	}

	return left
}
