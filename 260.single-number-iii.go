func singleNumber(nums []int) []int {
    xorAll := xor(nums)
    xorFlag := 1
    for xorAll != 0 && xorAll % 2 == 0 {
        xorFlag <<= 1
        xorAll >>= 1
    }
    r1, r2 := 0, 0
    for i := range nums {
        if nums[i] | xorFlag == nums[i] {
            r1 ^= nums[i]
        } else {
            r2 ^= nums[i]
        }
    }
    return []int{r1, r2}
}

func xor(nums []int) int {
    res := 0
    for i := range nums {
        res ^= nums[i]
    }
    return res
}
