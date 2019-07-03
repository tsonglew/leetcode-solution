func reconstructQueue(people [][]int) [][]int {
    if len(people) < 2 {
        return people
    }
    sort.Slice(people, func(i, j int) bool {
        if people[i][0] == people[j][0] {
            return people[i][1] < people[j][1]
        }
        return people[i][0] > people[j][0]
	})
    result := [][]int{people[0]}
    for i := 1; i < len(people); i++ {
        aboveCnt := 0
        for j := 0; j <= len(result); j++ {
            if aboveCnt == people[i][1] {
                result = append(result[0:j], append([][]int{people[i]}, result[j:len(result)]...)...)
                break
            }
            aboveCnt ++
        }
    }
    return result
}