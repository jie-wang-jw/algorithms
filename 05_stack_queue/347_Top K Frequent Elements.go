package _5_stack_queue

import (
	"container/heap"
	"math/rand"
	"sort"
)

/*
347. 前 K 个高频元素 / Top K Frequent Elements

题目描述 / Problem Description

给定整数数组 nums 和整数 k，返回出现频率最高的 k 个元素。
答案可以按照任意顺序返回。

Given an integer array nums and an integer k, return the k most
frequent elements. The answer may be returned in any order.
*/

// frequencyHeap is a min-heap of [number, frequency] pairs ordered by frequency, not a fully sorted slice.
// frequencyHeap 是按频率维护父子大小关系的小顶堆，保存 [数字, 出现频率]，不是完全有序的切片。
type frequencyHeap [][2]int

// Report the heap size so container/heap can locate the last slot.
// 返回堆中的元素数量，供 container/heap 定位末尾位置。
// Time: O(1), Space: O(1).
// 时间复杂度：O(1)，空间复杂度：O(1)。
func (h frequencyHeap) Len() int {
	return len(h)
}

// Compare frequencies only so the lower one sits closer to the root, forming a min-heap.
// 只比较出现频率，让较低频率更接近堆顶，从而形成小顶堆。
// Time: O(1), Space: O(1).
// 时间复杂度：O(1)，空间复杂度：O(1)。
func (h frequencyHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

// Exchange two heap slots while container/heap sifts.
// 在上浮和下沉过程中交换堆中的两个元素。
// Time: O(1), Space: O(1).
// 时间复杂度：O(1)，空间复杂度：O(1)。
func (h frequencyHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Append the new pair to the backing slice; heap.Push then sifts it up with Less and Swap.
// 把新元素追加到底层切片末尾；heap.Push 随后用 Less 和 Swap 上浮。
// Time: amortized O(1), worst O(n) for append; backing storage O(n). heap.Push adds O(log n) sift work.
// 追加均摊 O(1)，扩容时 O(n)；底层存储 O(n)。外层 heap.Push 另有 O(log n) 上浮操作。
//
// 这里只向底层切片追加；外层 heap.Push 随后上浮修复堆序。追加均摊 O(1)，扩容时单次 O(n)。
// This hook only appends; heap.Push then sifts upward to restore heap order. Append is amortized O(1), or O(n) when reallocating.
func (h *frequencyHeap) Push(value any) {
	item := value.([2]int)
	*h = append(*h, item)
}

// Remove the last backing-slice element after container/heap has already moved the root there.
// 删除底层切片末尾元素；container/heap 调用前已把原堆顶移到末尾。
// Time: O(1), Space: O(1).
// 时间复杂度：O(1)，空间复杂度：O(1)。
func (h *frequencyHeap) Pop() any {
	oldHeap := *h
	lastIndex := len(oldHeap) - 1

	item := oldHeap[lastIndex]

	*h = oldHeap[:lastIndex]

	return item
}

// 1. Min-heap of size k (recommended): count frequencies, then evict the least frequent root whenever the heap exceeds k.
// 1. 大小为 k 的小顶堆：推荐；先统计频次，堆超过 k 就删除频率最低的堆顶。
// Time: O(n + m log(k+1)) average, Space: O(m+k) for the map and the heap.
// 时间复杂度：平均 O(n + m log(k+1))，空间复杂度：O(m+k)，由频次表和堆产生。
//
// 先统计频次；小顶堆最多保留 k 个候选，超出就删最低频项，留下的始终是已处理值中频次最高的 k 个。
// Count frequencies; keep at most k entries in a min-heap and evict the least frequent, retaining the processed top k.
// 弹出顺序频次递增，所以从结果末尾写入；要求 1<=k<=不同值个数，同频项顺序不保证。
// Heap pops increase in frequency, so fill output backward; require valid k and allow arbitrary tie order.
func topKFrequent(nums []int, k int) []int {
	frequency := make(map[int]int)

	for _, number := range nums {
		frequency[number]++
	}

	minHeap := &frequencyHeap{}
	heap.Init(minHeap)

	for number, count := range frequency {
		heap.Push(minHeap, [2]int{number, count})

		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}

	result := make([]int, k)

	for i := k - 1; i >= 0; i-- {
		item := heap.Pop(minHeap).([2]int)
		result[i] = item[0]
	}

	return result
}

// 2. Sort all distinct values: count frequencies, sort the m values by descending frequency, then take the first k.
// 2. 全排序取前 k 个：统计频次后按频次降序排序，直接取前 k 个。
// Time: O(n + m log(m+1)) average, Space: O(m).
// 时间复杂度：平均 O(n + m log(m+1))，空间复杂度：O(m)。
//
// 按频次对不同值降序排序，取前 k 个；只复制结果值，不修改 nums，要求 1<=k<=不同值个数。
// Sort distinct values by decreasing count and copy the first k; nums is unchanged and k must be valid.
func topKFrequentSorted(nums []int, k int) []int {
	counts := make(map[int]int)
	for _, number := range nums {
		counts[number]++
	}

	values := make([]int, 0, len(counts))
	for number := range counts {
		values = append(values, number)
	}

	sort.Slice(values, func(i, j int) bool {
		return counts[values[i]] > counts[values[j]]
	})

	return append([]int(nil), values[:k]...)
}

// 3. Bucket sort by frequency: put each value in buckets[count], then scan from n downward until k values are collected.
// 3. 桶排序按频次分桶：buckets[count] 保存出现 count 次的值，从 n 向下扫描直到凑满 k 个。
// Time: O(n+m) average, Space: O(n+m) because the bucket array always occupies n+1 slots.
// 时间复杂度：平均 O(n+m)，空间复杂度：O(n+m)，桶数组固定占 n+1 个槽位。
//
// 频次最多为 n，可用频次作桶下标；从高频桶向低频桶收集，恰收满 k 个即停止，同频内部顺序任意。
// Frequency is at most n; use it as a bucket index and collect downward until k values are found, with arbitrary tie order.
// 要求 1<=k<=不同值个数。
// Require 1<=k<=the number of distinct values.
func topKFrequentBucket(nums []int, k int) []int {
	counts := make(map[int]int)
	for _, number := range nums {
		counts[number]++
	}

	buckets := make([][]int, len(nums)+1)

	for number, count := range counts {
		buckets[count] = append(buckets[count], number)
	}

	result := make([]int, 0, k)

	for count := len(nums); count > 0; count-- {
		for _, number := range buckets[count] {
			result = append(result, number)

			if len(result) == k {
				return result
			}
		}
	}

	return result
}

// 4. Quickselect with three-way partition: partition until rank k-1 sits in the equal-frequency band.
// 4. 三路划分快速选择：只划分到第 k-1 名落入等频段为止，不做完整排序。
// Time: O(n+m) expected, O(n+m²) worst case. Space: O(m).
// 时间复杂度：随机化后期望 O(n+m)，最坏 O(n+m²)。空间复杂度：O(m)。
//
// 三路划分维护：左侧频次>pivot，中段==pivot，右侧<pivot；[i,less] 仍未检查。
// Three-way partition keeps >pivot on the left, ==pivot in the middle, <pivot on the right; [i,less] is unexamined.
// 从右侧换回的元素尚未检查，所以该分支不递增 i；第 k 大的位置是 k-1，只继续处理包含它的一段。
// Do not increment i after swapping from the right because the replacement is unexamined; recurse iteratively only toward rank k-1.
// 若 k-1 落入等频段，前 k 项已是答案，无需整体排序；随机枢轴期望线性选择，最坏仍二次。
// If k-1 lies in the equal band, the first k entries suffice without full sorting; randomized selection is expected linear but worst-case quadratic.
// 要求 1<=k<=不同值个数；结果不保证按频次排序。
// Require valid k; the output need not be frequency-sorted.
func topKFrequentQuickselect(nums []int, k int) []int {
	counts := make(map[int]int)
	for _, number := range nums {
		counts[number]++
	}

	items := make([][2]int, 0, len(counts))
	for number, count := range counts {
		items = append(items, [2]int{number, count})
	}

	left, right := 0, len(items)-1

	for left <= right {
		pivot := items[left+rand.Intn(right-left+1)][1]

		greater, i, less := left, left, right

		for i <= less {
			switch {
			case items[i][1] > pivot:
				items[i], items[greater] = items[greater], items[i]
				greater++
				i++

			case items[i][1] < pivot:
				items[i], items[less] = items[less], items[i]
				less--

			default:
				i++
			}
		}

		if k-1 < greater {
			right = greater - 1
			continue
		}

		if k-1 > less {
			left = less + 1
			continue
		}

		break
	}

	result := make([]int, k)
	for i := range k {
		result[i] = items[i][0]
	}

	return result
}
