package _7_backtracking

import "sort"

/*
332. 重新安排行程 / Reconstruct Itinerary

题目描述 / Problem Description
机票 [from, to] 必须全部使用一次。行程从 JFK 出发。若有多种行程，返回字典序最小的一条。
机场代码为三个大写字母。题目保证至少存在一种行程。
Use every ticket [from, to] exactly once. The itinerary starts at JFK.
If several itineraries exist, return the lexicographically smallest. Airports are three uppercase letters.
At least one valid itinerary exists.

示例 / Example
[["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]] → ["JFK","MUC","LHR","SFO","SJC"]
*/

// 1. Backtracking with remaining counts: try destinations in lexical order and stop at the first complete path.
// 1. 回溯加剩余次数：按字典序尝试目的地，找到第一条走完所有机票的路径就返回。
// Time: factorial search in the worst case; space O(e^2) conservative, including destination lists retained by recursive frames.
// 时间：最坏为机票排列量级；空间保守上界 O(e²)，含递归各层保留的目的地列表，e 为机票数。
//
// targets[from][to] 是剩余同航线机票数，不能用布尔值，否则丢失重复机票；每次选票减一，失败回退加一。
// Store remaining ticket multiplicities, not booleans; decrement on choice and restore on backtracking.
// 目的地按字典序尝试，首个用完全部票的完整路线就是最小字典序；题目保证存在从 JFK 出发的有效路线。
// Try destinations lexicographically; the first complete route is minimal. Assume a valid itinerary from JFK exists.
func findItinerary(tickets [][]string) []string {
	targets := map[string]map[string]int{}
	for _, ticket := range tickets {
		if targets[ticket[0]] == nil {
			targets[ticket[0]] = map[string]int{}
		}
		targets[ticket[0]][ticket[1]]++
	}

	result := []string{"JFK"}
	var backtrack func() bool
	backtrack = func() bool {
		if len(result) == len(tickets)+1 {
			return true
		}
		from := result[len(result)-1]
		dests := make([]string, 0, len(targets[from]))
		for dest := range targets[from] {
			dests = append(dests, dest)
		}
		sort.Strings(dests)
		for _, dest := range dests {
			if targets[from][dest] == 0 {
				continue
			}
			targets[from][dest]--
			result = append(result, dest)
			if backtrack() {
				return true
			}
			result = result[:len(result)-1]
			targets[from][dest]++
		}
		return false
	}

	backtrack()
	return result
}

// 面试首选 / Interview first choice：332 重新安排行程。
//
// 2. Hierholzer: post-order append airports, then reverse to obtain the Euler path.
// 2. Hierholzer：后序加入机场，再反转得到欧拉路径。
// Time: O(e log e), Space: O(e).
// 时间复杂度：O(e log e)，空间复杂度：O(e)。
//
// 邻接表倒序保存，尾部弹出当前最小目的地；每张票是一条独立边，弹出即消费一次。
// Store destinations in reverse order so popping takes the smallest; each ticket is a separate edge consumed exactly once.
// 走到无剩余出边才记录机场，得到路线的逆序；这样提前遇到的终点会被放到最终路线末尾，不会截断其他剩余边。
// Record airports only after exhausting their outgoing edges; reversing this postorder puts early dead ends at the end of the final route.
// 依赖存在从 JFK 出发、用完全部边的欧拉路径；不是对任意机票集合的合法性检查。
// Assume an Euler trail from JFK using every edge; this routine does not validate arbitrary ticket sets.
func findItineraryHierholzer(tickets [][]string) []string {
	targets := map[string][]string{}
	for _, ticket := range tickets {
		targets[ticket[0]] = append(targets[ticket[0]], ticket[1])
	}
	for from := range targets {
		sort.Strings(targets[from])
		dests := targets[from]
		for i, j := 0, len(dests)-1; i < j; i, j = i+1, j-1 {
			dests[i], dests[j] = dests[j], dests[i]
		}
		targets[from] = dests
	}

	route := []string{}
	var dfs func(string)
	dfs = func(from string) {
		for len(targets[from]) > 0 {
			last := len(targets[from]) - 1
			next := targets[from][last]
			targets[from] = targets[from][:last]
			dfs(next)
		}
		route = append(route, from)
	}
	dfs("JFK")

	for i, j := 0, len(route)-1; i < j; i, j = i+1, j-1 {
		route[i], route[j] = route[j], route[i]
	}
	return route
}
