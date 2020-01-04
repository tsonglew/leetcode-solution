func majorityElement(nums []int) []int {
    if len(nums) <= 1 {
        return nums
    }
    candidates := []int{nums[0], nums[1]}
    candidatesCnt := []int{0, 0}
    for _, n := range nums {
        if n == candidates[0] {
            candidatesCnt[0]++
            continue
        }
        if n == candidates[1] {
            candidatesCnt[1]++
            continue
        }

        if candidatesCnt[0] == 0 {
            candidates[0] = n            
            candidatesCnt[1] = 1
            continue
        }
        if candidatesCnt[1] == 0 {
            candidates[1] = n
            candidatesCnt[1] = 1
            continue
        }
        candidatesCnt[0]--
        candidatesCnt[1]--
    }
    result := make(map[int]bool)
    candidatesCnt[0], candidatesCnt[1] = 0, 0
    for _, n := range nums {
        switch n {
        case candidates[0]:
            candidatesCnt[0]++
            if candidatesCnt[0] > len(nums)/3 {
                result[candidates[0]] = true
            }
        case candidates[1]:
            candidatesCnt[1]++
            if candidatesCnt[1] > len(nums)/3 {
                result[candidates[1]] = true
            }
        }
    }
    element := make([]int, 0)
    for k, _ := range result {
        element = append(element, k)
    }
    return element
}
