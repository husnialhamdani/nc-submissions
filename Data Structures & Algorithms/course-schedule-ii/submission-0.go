func findOrder(numCourses int, prerequisites [][]int) []int {
    graph := make([][]int, numCourses)

	for _, prereq := range prerequisites {
		course, pre := prereq[0], prereq[1]
		graph[course] = append(graph[course], pre)
	}

	state := make([]int, numCourses)
	result := make([]int, 0, numCourses)

	var dfs func(course int) bool
	dfs = func(course int) bool{
		if state[course] == 1{
			return false //cycle
		}

		if state[course] == 2{
			return true
		}

		state[course] = 1
		for _, pre := range graph[course] {
			if !dfs(pre){
				return false
			}
		}
		state[course]=2
		result = append(result, course)
		return true
	}

	for course:=0; course<numCourses; course++{
		if !dfs(course){
			return []int{}
		}
	}
	return result
}
