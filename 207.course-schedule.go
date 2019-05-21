func canFinish(numCourses int, prerequisites [][]int) bool {
	courses := make([][]int, numCourses)
	for i := range prerequisites {
		pair := prerequisites[i]
		if courses[pair[0]] == nil {
			courses[pair[0]] = []int{}
		}
		courses[pair[0]] = append(courses[pair[0]], pair[1])
	}
	for i := range courses {
		visited := make([]bool, numCourses)
		if !travel(courses, visited, i) {
			return false
		}
	}
	return true
}

func travel(courses [][]int, visited []bool, i int) bool {
	if visited[i] {
		return false
	}
	visited[i] = true
	for j := range courses[i] {
		if !travel(courses, visited, courses[i][j]) {
			visited[i] = false
			return false
		}
	}
	visited[i] = false
	return true
}