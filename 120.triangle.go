func minimumTotal(triangle [][]int) int {
    switch len(triangle) {
    case 0: return 0
    case 1: return triangle[0][0]
    }
    for i := 1; i < len(triangle); i++ {
        for j := range triangle[i] {
            var mn int
            if j == 0 {
                mn = triangle[i-1][0]
            } else if j == len(triangle[i])-1 {
                mn = triangle[i-1][len(triangle[i-1])-1]
            } else {
                mn = min(triangle[i-1][j-1], triangle[i-1][j])
            }
            triangle[i][j] += mn
        }
    }
    return min(triangle[len(triangle)-1]...)
}

func min(nums ...int) int {
    mn := nums[0]
    for i := range nums {
        if nums[i] < mn {
            mn = nums[i]
        }
    }
    return mn
}