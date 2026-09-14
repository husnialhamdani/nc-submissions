func canFinish(numCourses int, prerequisites [][]int) bool {
    graph := make([][]int, numCourses)

	for _, prereq := range prerequisites {
		course, pre := prereq[0], prereq[1]
		graph[course] = append(graph[course], pre)
	}

	state := make([]int, numCourses)

	var dfs func(course int) bool
	dfs = func(course int) bool {
		if state[course] == 1{
			return false // cycle
		}

		if state[course] == 2 {
			return true
		}

		state[course] = 1
		for _, pre := range graph[course]{
			if !dfs(pre){
				return false
			}
		}
		state[course] = 2
		return true
	}

	for course:=0; course<numCourses; course++{
		if !dfs(course){
			return false
		}
	}
	return true
}
