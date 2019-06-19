func countBits(num int) []int {
    result := make([]int, num+1)
    result[0] = 0
    if num == 0 {
        return result
    }
    result[1] = 1
    if num == 1 {
        return result
    }
    powNum := 1
    powNext := 2
    for i := 2; i < num+1; i++ {
        if i == powNext {
            result[i] = 1
            powNum = powNext
            powNext *= 2
        } else {
            result[i] = result[i % powNum]+1
        }
    }
    return result
}