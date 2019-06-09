func lengthOfLISMethod1(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    liss := make([]int, len(nums))
    result := 1
    liss[0] = 1
    for i := 1; i < len(liss); i++ {
        liss[i] = 1
        for j := i-1;j >= 0; j-- {
            if nums[j] < nums[i] && liss[j] + 1 > liss[i] {
                liss[i] = liss[j] + 1
            }
            if nums[j] == nums[i] - 1 || nums[i] == nums[j] {
                break
            }
        }
        if liss[i] > result {
            result = liss[i]
        }
    } 
    return result
}