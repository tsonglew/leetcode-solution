func findCircleNum(M [][]int) int {
    if len(M) == 0 {
        return 0
    }
    result := 0
    for i := range M {
        if M[i][i] == 0 {
            continue
        }
        result += 1
        M[i][i] = 0
        dfs(M, i)
    }
    return result
}

func dfs(M [][]int, i int) {
    for j := range M[i] {
        if M[i][j] == 0 {
            continue
        }
        M[i][j] = 0
        M[j][i] = 0
        dfs(M, j)
    }
}
