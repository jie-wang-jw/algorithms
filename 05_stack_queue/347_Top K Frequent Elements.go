package _5_stack_queue

import "container/heap"

/*
题目描述 / Problem Description

给定整数数组 nums 和整数 k，返回出现频率最高的 k 个元素。
答案可以按照任意顺序返回。

Given an integer array nums and an integer k, return the k most
frequent elements. The answer may be returned in any order.

解题思路 / Solution Approach

首先使用哈希表统计每个数字的出现频率。
First, use a hash map to count the frequency of each value.

然后维护一个大小为 k 的小顶堆。
Then maintain a min-heap of size k.

当堆的大小超过 k 时，删除频率最低的堆顶元素。
Whenever the heap size exceeds k, remove the least frequent root.

最后堆中保留的就是频率最高的 k 个元素。
The heap finally contains the k most frequent elements.

为什么使用小顶堆？ / Why Use a Min-Heap?
堆中只保留当前频率最高的 k 个元素。
The heap stores only the current top k frequent elements.
小顶堆的堆顶是这 k 个元素中频率最低的元素。
The root of the min-heap is the least frequent among those k elements.
每次加入新的数字和频率后：
After inserting a new value-frequency pair:
- 如果堆的大小不超过 k，继续保留。
  If the heap size does not exceed k, keep all elements.
- 如果堆的大小超过 k，删除堆顶。
  If the heap size exceeds k, remove the root.
这样每次删除的都是当前候选中频率最低的元素。
This always removes the least frequent current candidate.
遍历结束后，堆中剩下的就是频率最高的 k 个元素。
After processing all frequencies, the heap contains exactly the top k frequent elements.

关键逻辑：为什么这样做 / Why This Works
处理一个新候选前，堆里已是此前候选的前 k 名（不足 k 个则全保留）。
加入新候选后，从至多 k+1 个中删掉最小频率，就仍是前 k 名；以前淘汰的值不可能因新候选加入而挤进前 k，因此无需找回。
堆不是整个切片排好序，只要求父节点频率不大于子节点，沿父子路径可知根最小。
heap.Push 先调用自定义 Push 追加，再用 Less/Swap 上浮；
heap.Pop 先把根移到末尾并调整剩余堆，再调用自定义 Pop 删除末尾，所以自定义 Pop 不需要自己找最小值。
Before each candidate, the heap contains the best k seen so far, or all if fewer.
Adding one and discarding the minimum of at most k+1 preserves that property.
Previously discarded candidates cannot improve their rank when more candidates arrive.
A heap is not a fully sorted slice: parent frequencies are no larger than children's,
making the root minimal. heap.Push calls the custom Push to append, then uses Less/Swap to sift up.
heap.Pop moves the root to the end, repairs the remaining heap, then calls the custom Pop to remove the last item.

时间与空间复杂度 / Time and Space Complexity
n 为输入长度，m 为不同数字数量。平均时间 O(n + m log(k+1))：统计 O(n)，维护堆 O(m log(k+1))，
最后弹出 k 次 O(k log(k+1))，且 k<=m。用 log(k+1) 可正确涵盖 k=1。辅助空间 O(m+k)，
返回结果 O(k)；堆在弹出前可短暂达到 k+1 个元素。
For n inputs and m distinct values, average time is O(n + m log(k+1)): counting O(n),
heap maintenance O(m log(k+1)), and extraction O(k log(k+1)), with k<=m. log(k+1)
includes k=1 correctly. Auxiliary space O(m+k), output O(k);
the heap temporarily reaches k+1 items.
*/

// frequencyHeap is a min-heap ordered by frequency.
// frequencyHeap 是按频率维护父子大小关系的小顶堆，不是完全有序的切片。
//
// Each item stores [number, frequency].
// 每个元素保存 [数字, 出现频率]。
type frequencyHeap [][2]int

// Len returns the number of elements in the heap.
// Len 返回堆中的元素数量。
func (h frequencyHeap) Len() int {
	return len(h)
}

// Less places the lower-frequency element closer to the root.
// Less 让出现频率较低的元素更接近堆顶。
func (h frequencyHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

// Swap exchanges two heap elements.
// Swap 交换堆中的两个元素。
func (h frequencyHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push adds a new element to the end of the underlying slice.
// Push 将新元素加入底层切片末尾。
func (h *frequencyHeap) Push(value interface{}) {
	item := value.([2]int)
	*h = append(*h, item)
}

// Pop removes and returns the last element of the underlying slice.
// container/heap moves the root to the end before calling this method.
// Pop 删除并返回底层切片末尾的元素。
// container/heap 调用此方法前，会先把原堆顶移动到切片末尾。
func (h *frequencyHeap) Pop() interface{} {
	oldHeap := *h
	lastIndex := len(oldHeap) - 1

	// Save the element that will be returned.
	// 保存即将返回的元素。
	item := oldHeap[lastIndex]

	// Remove the last element from the slice.
	// 从切片中删除最后一个元素。
	*h = oldHeap[:lastIndex]

	return item
}

func topKFrequent(nums []int, k int) []int {
	// frequency maps each number to its occurrence count.
	// frequency 记录每个数字的出现次数。
	frequency := make(map[int]int)

	for _, number := range nums {
		frequency[number]++
	}

	// Each completed iteration keeps at most k pairs; insertion may temporarily produce k+1.
	// 每轮结束最多保留 k 对，插入后可暂时达到 k+1 对。
	minHeap := &frequencyHeap{}
	heap.Init(minHeap)

	for number, count := range frequency {
		// Store the number and its frequency together.
		// 将数字和它的出现频率一起加入堆中。
		heap.Push(minHeap, [2]int{number, count})

		// If the heap contains more than k candidates,
		// remove the candidate with the lowest frequency.
		// 如果堆中候选元素超过 k 个，
		// 就删除出现频率最低的候选元素。
		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}

	// The heap now contains exactly the top k frequent values.
	// 此时堆中正好保存频率最高的 k 个数字。
	result := make([]int, k)

	// Pop from the min-heap and write from right to left.
	// The result is therefore ordered from higher to lower frequency.
	// 依次弹出小顶堆，并从结果数组右侧向左写入。
	// 因此结果按频率从高到低排列；相同频率之间的顺序不作保证。
	for i := k - 1; i >= 0; i-- {
		item := heap.Pop(minHeap).([2]int)
		result[i] = item[0]
	}

	return result
}
