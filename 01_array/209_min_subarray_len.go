package _1_array

func minSubArrayLen(target int, nums []int) int {
	i := 0
	l := len(nums)  // æ•°ç»„é•¿åº¦
	sum := 0        // å­æ•°ç»„ä¹‹å’Œ
	result := l + 1 // åˆå§‹åŒ–è¿”å›žé•¿åº¦ä¸ºl+1ï¼Œç›®çš„æ˜¯ä¸ºäº†åˆ¤æ–­â€œä¸å­˜åœ¨ç¬¦åˆæ¡ä»¶çš„å­æ•°ç»„ï¼Œè¿”å›ž0â€çš„æƒ…å†µ

	for j := 0; j < l; j++ {
		sum += nums[j]
		for sum >= target {
			subLength := j - i + 1
			if subLength < result {
				result = subLength
			}

			sum -= nums[i]
			i++
		}
	}
	if result == l+1 {
		return 0
	}
	return result
}

