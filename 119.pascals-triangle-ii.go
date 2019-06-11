func getRow(rowIndex int) []int {
    result := make([]int, rowIndex+1)
    result[0] = 1
    for cnt := 0; cnt <= rowIndex; cnt++ {
        result[cnt] = 1
        prev := 1
        for i := 1; i < cnt; i++ {
            a := result[i]
            result[i] += prev
            prev = a
        }
    }
    return result
}