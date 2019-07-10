func findUnsortedSubarray(nums []int) int {
    nums1 := append([]int{}, nums...)
    sort.Ints(nums)
    start, end := -1, -2
    var foundStart bool
    for i := range nums {
        if nums[i] != nums1[i] {
            if foundStart {
                end = i    
            } else {
                start = i
                foundStart = true   
            }
        } 
    }
    return end-start+1
}