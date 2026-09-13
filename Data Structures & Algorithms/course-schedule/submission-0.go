func canFinish(numCourses int, prerequisites [][]int) bool {
    
    preMap := make(map[int][]int)
    for i:=0; i<numCourses; i++ {
        preMap[i] = []int{}
    }

    for _, prereq := range prerequisites {
        crs, pre := prereq[0], prereq[1]
        preMap[crs] = append(preMap[crs], pre)
    }

    // store all course along the current DFS path
    visiting := make(map[int]bool)
    var dfs func(int) bool
    dfs = func(crs int) bool {
        if visiting[crs] { // cycle detected
            return false 
        }

        if len(preMap[crs]) == 0 {
            return true
        }

        visiting[crs] = true
        for _, pre := range preMap[crs] {
            if !dfs(pre) {
                return false
            }
        }

        visiting[crs] = false
        preMap[crs] = []int{}
        return true
    }

    for c:=0; c<numCourses; c++ {
        if !dfs(c) {
            return false
        }
    }
    return true
}
