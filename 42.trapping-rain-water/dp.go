func trap(height []int) int {
    maxLeft := make([]int, len(height))
    maxRight := make([]int, len(height))
    currentMxLeft, currentMxRight := 0, 0
    for i := range height {
        maxLeft[i] = currentMxLeft
        maxRight[len(height)-1-i] = currentMxRight
        currentMxLeft = max(currentMxLeft, height[i])
        currentMxRight = max(currentMxRight, height[len(height)-1-i])
    }
    ans := 0
    for i := range height {
        ans += max(min(maxLeft[i], maxRight[i])-height[i], 0)
    }
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
