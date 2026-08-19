package _3_hash

/*
use a hash map to store numbers I have already seen and their indices.
For each number, I compute its complement, which is target minus the current number.
If the complement is already in the map, I return the two indices.
Otherwise, I store the current number and continue.
*/
func twoSum(nums []int, target int) []int {
	seen := map[int]int{}

	for i, num := range nums {
		need := target - num
		if j, ok := seen[need]; ok {
			return []int{j, i}
		}
		seen[num] = i
	}

	return nil
}
