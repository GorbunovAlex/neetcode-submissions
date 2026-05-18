func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)
	for _, n := range nums {
		counts[n]++
	}
	keys := make([]int, 0, len(counts))
for k := range counts {
    keys = append(keys, k)
}
	sort.Slice(keys, func(i, j int) bool {
    return counts[keys[i]] > counts[keys[j]]
})
return keys[:k]
}
