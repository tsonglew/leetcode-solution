func findTargetSumWays(nums []int, S int) int {
    output := 0
    sumTarget(nums, 0, S, &output)
    return output
}

func sumTarget(nums []int, i, target int, output *int) {
    if i == len(nums)-1 {
        if nums[i] == target {
            (*output)++
        }
        if nums[i] == -target {
            (*output)++
        }
        return
    }
    sumTarget(nums, i+1, target-nums[i], output)
    sumTarget(nums, i+1, target+nums[i], output)
}