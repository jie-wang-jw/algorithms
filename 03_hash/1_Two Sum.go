package _3_hash

/*
Use a hash map to store numbers I have already seen and their indices.
使用哈希表保存已经见过的数字及其下标。

For each number, I compute its complement, which is target minus the current number.
对每个数字计算它需要的另一个数：target - 当前数字。

If the complement is already in the map, I return the two indices.
如果需要的数已经在哈希表中，就返回两个下标。

Otherwise, I store the current number and continue.
否则保存当前数字并继续扫描。
*/
func twoSum(nums []int, target int) []int {
	// Key: a seen number; value: its index.
	// key 是已经见过的数字，value 是该数字的下标。
	seen := map[int]int{}

	for i, num := range nums {
		// num + need = target.
		// 当前数字 num 需要搭配 need 才能得到 target。
		need := target - num
		// ok tells us whether need already exists in the map.
		// ok 表示 need 是否已经存在于哈希表中。
		if j, ok := seen[need]; ok {
			return []int{j, i}
		}
		// Store after checking so the same element cannot be used twice.
		// 先查找、后存入，避免同一个元素被使用两次。
		seen[num] = i
	}

	// No valid pair was found.
	// 没有找到满足条件的两个数。
	return nil
}
