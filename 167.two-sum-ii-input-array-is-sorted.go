func twoSum(numbers []int, target int) []int {
    idx1, idx2 := 0, len(numbers) - 1
    for idx1 < idx2 {
        sum := numbers[idx1] + numbers[idx2]
        if sum > target {
            idx2--
        } else if sum < target {
            idx1++
        } else {
            break
        }
    }
    return []int{idx1+1, idx2+1}
}