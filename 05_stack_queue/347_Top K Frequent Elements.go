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

解法一：大小为 k 的小顶堆（推荐） / Method 1: Min-Heap of Size k (Recommended)

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

解法二：全排序取前 k 个 / Method 2: Sort All Distinct Values
topKFrequentSorted 先统计频次，再把 m 个不同值按频次降序排序，直接取前 k 个。
sort.Slice 不稳定，所以并列频次之间的先后不确定；题目不要求固定顺序，因此这不影响正确性。
它是最直观的基线，但用了完整的比较排序，不满足题目“严格优于 O(n log n)”的一般性要求。
Count frequencies, sort the m distinct values by descending frequency, and take the first k.
sort.Slice is not stable, so tied frequencies order arbitrarily; the problem fixes no order, so this is fine.
It is the most direct baseline, but a full comparison sort does not meet the problem's better-than-O(n log n) goal.

解法三：桶排序按频次分桶 / Method 3: Bucket Sort by Frequency
topKFrequentBucket 利用一个关键约束：任何值的出现次数只可能落在 1..n。
于是 buckets[count] 保存所有恰好出现 count 次的值，下标本身就编码了频次，不需要任何比较排序。
从 count = n 向下扫描，遇到值就收集，凑满 k 个立即返回。
桶下标单调递减，所以先收集到的频次一定不低于后收集到的，前 k 个自然就是答案。
buckets 长度取 len(nums)+1 而不是 len(nums)，因为要能索引 count == n 这一桶。
Frequencies can only lie in 1..n. So buckets[count] holds every value occurring exactly count times,
and the index itself encodes the frequency, removing the need for any comparison sort.
Scan count downward from n, collecting values until k are gathered, then return immediately.
Bucket indices decrease monotonically, so anything collected earlier has frequency at least as high
as anything collected later, making the first k collected values a correct answer.
buckets has length len(nums)+1, not len(nums), so that the count == n bucket is indexable.

解法四：三路划分快速选择 / Method 4: Quickselect with Three-Way Partition
topKFrequentQuickselect 只把数组划分到“第 k-1 名就位”的程度，不做完整排序。
每轮随机取一个频次作为基准 pivot，把当前区间三路划分：
[left,greater) 频次大于 pivot，[greater,less] 频次等于 pivot，(less,right] 频次小于 pivot。
然后只朝目标下标 k-1 所在的那一段继续：
- k-1 < greater：答案还在大于段里，把 right 收缩到 greater-1。
- k-1 > less：答案在小于段里，把 left 推进到 less+1。
- 否则 k-1 落在等频段，前 k 个元素的频次已经全部合格，段内无需再排序，直接结束。
基准取自区间内的元素，所以等频段一定非空，于是每轮区间都严格变小，循环必然终止。
三路划分是必要的：若只做两路，大量等频值会让划分退化，反复在同一段里空转。
Partition only until rank k-1 is in place, never sorting fully.
Each round picks a random pivot frequency and splits the current range three ways:
[left,greater) above pivot, [greater,less] equal to pivot, and (less,right] below pivot.
Then continue only toward the band containing target index k-1:
- k-1 < greater: the answer lies in the greater band, so shrink right to greater-1.
- k-1 > less: the answer lies in the smaller band, so advance left to less+1.
- otherwise k-1 falls in the equal band, the first k frequencies already qualify, and the loop stops
  without sorting inside that band.
The pivot is taken from inside the range, so the equal band is never empty and the range strictly shrinks,
which guarantees termination.
The three-way split matters: a two-way partition degenerates on many tied frequencies and spins in place.

时间与空间复杂度 / Time and Space Complexity
n 为输入长度，m 为不同数字数量，k<=m<=n；四种解法都先统计频次、都不修改输入、返回结果都占 O(k)。
哈希统计的时间均按平均 O(1) 操作分析。
解法一 topKFrequent：平均时间 O(n + m log(k+1))：统计 O(n)，维护堆 O(m log(k+1))，
最后弹出 k 次 O(k log(k+1))。用 log(k+1) 可正确涵盖 k=1。
辅助空间 O(m+k)：频次表 O(m)，堆 O(k)；堆在弹出前可短暂达到 k+1 个元素。
解法二 topKFrequentSorted：平均时间 O(n + m log(m+1))，辅助空间 O(m)。
解法三 topKFrequentBucket：平均时间 O(n+m)，两次线性遍历加一次桶扫描；
辅助空间 O(n+m)，桶数组固定占 n+1 个槽位，即使 m 很小也无法减少。
解法四 topKFrequentQuickselect：随机化后期望时间 O(n+m)，最坏 O(n+m²)；辅助空间 O(m)。
题目保证答案唯一；结果顺序无需固定，所以并列频次的排列差异不算错误。
For n inputs, m distinct values, and k <= m <= n: all four count frequencies first, leave the input unchanged,
and return an O(k) result. Hash operations are analyzed at average O(1).
Method 1 topKFrequent: average time O(n + m log(k+1)) from counting O(n), heap maintenance O(m log(k+1)),
and k extractions O(k log(k+1)); log(k+1) includes k = 1 correctly.
Auxiliary space O(m+k) for the frequency map and the heap, which temporarily reaches k+1 items.
Method 2 topKFrequentSorted: average time O(n + m log(m+1)), auxiliary space O(m).
Method 3 topKFrequentBucket: average time O(n+m) across two linear passes and one bucket scan;
auxiliary space O(n+m), since the bucket array always occupies n+1 slots even when m is small.
Method 4 topKFrequentQuickselect: expected time O(n+m) with randomization, worst case O(n+m²);
auxiliary space O(m).
The problem guarantees a unique answer set, not a fixed order, so differing arrangements of ties are not errors.
*/

// 1. 大小为 k 的小顶堆：推荐

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
func (h *frequencyHeap) Push(value any) {
	item := value.([2]int)
	*h = append(*h, item)
}

// Pop removes and returns the last element of the underlying slice.
// container/heap moves the root to the end before calling this method.
// Pop 删除并返回底层切片末尾的元素。
// container/heap 调用此方法前，会先把原堆顶移动到切片末尾。
func (h *frequencyHeap) Pop() any {
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

// 2. 全排序取前 k 个：最直观的基线
func topKFrequentSorted(nums []int, k int) []int {
	// counts maps each number to its occurrence count.
	// counts 记录每个数字的出现次数。
	counts := make(map[int]int)
	for _, number := range nums {
		counts[number]++
	}

	// values holds the distinct numbers; there are m of them, not n.
	// values 保存所有不同的数字，共 m 个，而不是 n 个。
	values := make([]int, 0, len(counts))
	for number := range counts {
		values = append(values, number)
	}

	// Sort by descending frequency; sort.Slice is unstable, so ties order arbitrarily.
	// 按频次降序排序；sort.Slice 不稳定，所以并列频次的先后不确定。
	sort.Slice(values, func(i, j int) bool {
		return counts[values[i]] > counts[values[j]]
	})

	// Copy the prefix so the result does not alias the larger values slice.
	// 复制前缀，避免返回值与更长的 values 切片共享底层数组。
	return append([]int(nil), values[:k]...)
}

// 3. 桶排序按频次分桶：不需要比较排序
func topKFrequentBucket(nums []int, k int) []int {
	// counts maps each number to its occurrence count.
	// counts 记录每个数字的出现次数。
	counts := make(map[int]int)
	for _, number := range nums {
		counts[number]++
	}

	// A count lies in 1..len(nums), so index len(nums) must exist: hence len(nums)+1 slots.
	// 出现次数的范围是 1..len(nums)，需要能索引 len(nums)，因此长度取 len(nums)+1。
	buckets := make([][]int, len(nums)+1)

	// The bucket index encodes the frequency, so no value needs to store its own count.
	// 桶下标本身就编码了频次，因此桶里只需保存数字。
	for number, count := range counts {
		buckets[count] = append(buckets[count], number)
	}

	result := make([]int, 0, k)

	// Scan from the highest possible frequency downward.
	// 从可能的最高频次开始向下扫描。
	for count := len(nums); count > 0; count-- {
		for _, number := range buckets[count] {
			result = append(result, number)

			// Bucket indices only decrease, so the first k collected values are already correct.
			// 桶下标只会递减，所以先收集到的 k 个数字已经是答案，可以立即返回。
			if len(result) == k {
				return result
			}
		}
	}

	// Reached only when k exceeds the distinct count, which the problem excludes.
	// 只有 k 超过不同数字个数时才会走到这里，而题目排除了这种输入。
	return result
}

// 4. 三路划分快速选择：只让第 k-1 名就位
func topKFrequentQuickselect(nums []int, k int) []int {
	// counts maps each number to its occurrence count.
	// counts 记录每个数字的出现次数。
	counts := make(map[int]int)
	for _, number := range nums {
		counts[number]++
	}

	// Each item stores [number, frequency]; partitioning reorders this local slice only.
	// 每一项保存 [数字, 出现频率]；划分只重排这个局部切片，不触及输入。
	items := make([][2]int, 0, len(counts))
	for number, count := range counts {
		items = append(items, [2]int{number, count})
	}

	// [left,right] is the range that may still contain the element of rank k-1.
	// [left,right] 是仍可能包含第 k-1 名的区间。
	left, right := 0, len(items)-1

	for left <= right {
		// The pivot is taken from inside the range, so the equal band is never empty.
		// 基准取自区间内部，因此等频段一定非空，区间每轮都严格变小。
		pivot := items[left+rand.Intn(right-left+1)][1]

		// [left,greater) > pivot; [greater,i) == pivot; (less,right] < pivot.
		// 三段分别保存大于、等于、小于基准的频次，中间未检查部分继续扫描。
		greater, i, less := left, left, right

		for i <= less {
			switch {
			case items[i][1] > pivot:
				// Grow the greater band; the swapped-in element is already checked.
				// 扩大大于段；换过来的元素已经检查过，因此 i 也前进。
				items[i], items[greater] = items[greater], items[i]
				greater++
				i++

			case items[i][1] < pivot:
				// Grow the smaller band; the swapped-in element is still unchecked, so i stays.
				// 扩大小于段；换过来的元素尚未检查，因此 i 不前进。
				items[i], items[less] = items[less], items[i]
				less--

			default:
				// Equal to the pivot: it already belongs to the middle band.
				// 与基准相等：它已经位于中间段，直接前进。
				i++
			}
		}

		if k-1 < greater {
			// Rank k-1 lies in the greater band, so discard everything from greater onward.
			// 第 k-1 名落在大于段，丢弃 greater 及其之后的部分。
			right = greater - 1
			continue
		}

		if k-1 > less {
			// Rank k-1 lies in the smaller band, so discard everything up to less.
			// 第 k-1 名落在小于段，丢弃 less 及其之前的部分。
			left = less + 1
			continue
		}

		// Rank k-1 falls in the equal band: the first k frequencies all qualify already.
		// 第 k-1 名落在等频段：前 k 个元素的频次都已合格，段内无需再排序。
		break
	}

	// Positions 0..k-1 now hold the k highest frequencies, in unspecified internal order.
	// 此时下标 0..k-1 保存频次最高的 k 个元素，它们内部的顺序不作保证。
	result := make([]int, k)
	for i := range k {
		result[i] = items[i][0]
	}

	return result
}
