func containsNearbyDuplicate(nums []int, k int) bool {
    appear := make(map[int]int)
    for i, n := range nums {
        if p, ok := appear[n]; ok && i-p <= k {
                return true
        }
        appear[n] = i
    }
    return false
}