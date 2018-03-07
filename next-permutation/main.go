func nextPermutation(nums []int) []int {
    length := len(nums)
    if length <= 1 {
        return nums
    }
    for i := length - 1; i > 0; i-- {
        if nums[i] > nums[i-1] {
            for j := length-1; j >= i; j-- {
                if nums[j] > nums[i-1] {
                    nums[j], nums[i-1] = nums[i-1], nums[j]
                    sort(nums, i, length-1)
                    return nums
                }
            }
        }
    }
    sort(nums, 0, length-1)
    return nums
}

func sort(nums []int, s, e int) {
    for i := s; i < e; i++ {
        for j := e; j > s; j-- {
            if nums[j] < nums[j-1] {
                nums[j], nums[j-1] = nums[j-1], nums[j]
            }
        }
    }
}
