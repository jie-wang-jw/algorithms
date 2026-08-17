package _3_hash

/*
I use a hash set to store all numbers from the first array.
Then I iterate through the second array.
If a number exists in the set, I add it to the result and remove it from the set to avoid duplicates.
*/

func intersection(nums1 []int, nums2 []int) []int {
	set := map[int]bool{}
	result := []int{}

	for _, num := range nums1 {
		set[num] = true
	}

	for _, num := range nums2 {
		if set[num] {
			result = append(result, num)
			//set[num] = false
			delete(set, num)
		}
	}

	return result
}
