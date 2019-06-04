func isInterleave(s1 string, s2 string, s3 string) bool {
    if len(s3) == 0 {
        return true
    }
    if len(s1) + len(s2) != len(s3) {
        return false
    }
    dp := make([][]bool, len(s1)+1)
    for i := range dp {
        dp[i] = make([]bool, len(s2)+1)
    }
    dp[0][0] = true
    for i := range dp {
        for j := range dp[i] {
            if i == 0 && j == 0 {
                continue
            }
            if i > 0 && dp[i-1][j] == true && s1[i-1] == s3[i+j-1] {
                dp[i][j] = true
            } else if j > 0 && dp[i][j-1] == true && s2[j-1] == s3[i+j-1] {
                dp[i][j] = true
            } else {
                dp[i][j] = false
            }
        }
    }
    return dp[len(dp)-1][len(dp[0])-1]
}