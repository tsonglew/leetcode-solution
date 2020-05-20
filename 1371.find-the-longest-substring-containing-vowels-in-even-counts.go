func findTheLongestSubstring(s string) int {
    status := 0
    firstAppear := map[int]int{0:-1}
    maxLen := 0
    for i := range s {
        switch s[i] {
            case 'a': status ^= 1 << 4
            case 'e': status ^= 1 << 3
            case 'i': status ^= 1 << 2
            case 'o': status ^= 1 << 1
            case 'u': status ^= 1
        }
        if v, ok := firstAppear[status]; ok {
            maxLen = max(maxLen, i-v)
        } else {
            firstAppear[status] = i
        }
    }
    return maxLen
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
