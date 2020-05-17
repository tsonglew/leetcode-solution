// method 1. DFS
type Course struct {
	Prevs []int
	NextCnt int
	travelled bool
}
func findOrder(numCourses int, prerequisites [][]int) []int {
	courses := make([]*Course, numCourses)
	result := []int{}
	for i := 0; i < numCourses; i++ {
		courses[i] = &Course{[]int{}, 0, false}
	}
	for i := 0; i < len(prerequisites); i++ {
		relation := prerequisites[i]
		courses[relation[0]].Prevs = append(
			courses[relation[0]].Prevs,
			relation[1],
		)
		courses[relation[1]].NextCnt++
	}
	lastCourses := []int{}
	for i := range courses {
		if courses[i].NextCnt == 0 {
			lastCourses = append(lastCourses, i)
		}
	}
	if len(lastCourses) == 0 {
		return result
	}
	for i := range lastCourses {
		if !dfs(lastCourses[i], &result, &courses) {
			return []int{}
		}
	}
	for i := range courses {
		if courses[i] != nil {
			return []int{}
		}
	}
	return result
}

func dfs(courseIdx int, result *[]int, courses *[]*Course) bool {
	if (*courses)[courseIdx] == nil {
		return true
	}
	if (*courses)[courseIdx].travelled {
		return false
	}
	(*courses)[courseIdx].travelled = true
	for i := range (*courses)[courseIdx].Prevs {
		if !dfs((*courses)[courseIdx].Prevs[i], result, courses) {
			return false
		}
	}
	(*courses)[courseIdx] = nil
	*result = append(*result, courseIdx)
	return true
}

// method 2. kahn algorithm
func findOrder(numCourses int, prerequisites [][]int) []int {
    afterCourses := make([][]int, numCourses)
    for i := range afterCourses {
        afterCourses[i] = []int{}
    }
    preCount := make([]int, numCourses)
    for i := range prerequisites {
        afterCourses[prerequisites[i][1]] = append(afterCourses[prerequisites[i][1]], prerequisites[i][0])
        preCount[prerequisites[i][0]] ++
    }
    q := []int{}
    for i, v := range preCount {
        if v == 0 {
            q = append(q, i)
        }
    }
    result := []int{}
    for len(q) > 0 {
        nq := []int{}
        for _, n := range q {
            result = append(result, n)
            for _, j := range afterCourses[n] {
                preCount[j] --
                if preCount[j] == 0 {
                    nq = append(nq, j)
                }
            }
        }
        q = nq
    }
    if len(result) == numCourses {
        return result
    }
    return []int{}
}
