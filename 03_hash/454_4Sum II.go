package _3_hash

/* split the four arrays into two pairs.
First, I count all possible sums from nums1 and nums2 using a hash map.
Then, for each sum from nums3 and nums4, I look for its negative value in the map.
The frequency in the map tells me how many valid tuples we can form.*/

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
