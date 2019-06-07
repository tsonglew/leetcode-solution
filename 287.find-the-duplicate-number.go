import "sort"
func findDuplicate(nums []int) int {
    sort.Ints(nums)
    for i := range nums {
        if nums[i] == nums[i+1] {
            return nums[i]
        }
    }
    return -1
}