func numFriendRequests(ages []int) int {
        ageBuckets := make([]int, 121)
        for i := range ages {
                ageBuckets[ages[i]]++
        }
        result := 0
        for i := range ages {
            if ages[i] <= 14 {
                continue
            }
            result += sum(ageBuckets[max(1, int((ages[i]+2)/2+7)):min(121, ages[i]+1)]) - 1
        }
        return result
}

func sum(l []int) int {
        r := 0
        for i := range l {
                r += l[i]
        }
        return r
}

func max(a, b int) int {
        if a > b {
                return a
        }
        return b
}

func min(a, b int) int {
        if a > b {
                return b
        }
        return a
}
