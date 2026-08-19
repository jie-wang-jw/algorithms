package _3_hash

/*排序 + 固定一个数 + 双指针
Sort + fix one number + two pointers
sort the array first.
Then I fix one number and use two pointers on the remaining range.
If the sum is too small, I move the left pointer to increase it.
If the sum is too large, I move the right pointer to decrease it.
I also skip duplicate values to avoid returning the same triplet multiple times.
*/
