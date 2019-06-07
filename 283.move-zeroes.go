func moveZeroes(nums []int)  {
    zerosCnt := 0
    for i := range nums {
        if nums[i] == 0 {
            zerosCnt ++
        } else {
            nums[i-zerosCnt] = nums[i]
        }
    }
    for i := len(nums)-zerosCnt; i < len(nums); i++ {
        nums[i] = 0
    }
}