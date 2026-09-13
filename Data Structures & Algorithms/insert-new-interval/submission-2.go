func insert(intervals [][]int, newInterval []int) [][]int {
    /*
		1. non-overlap
		2. overlap
		3. non-overlap remaining
	*/

	n := len(intervals)
	i := 0
	var result [][]int

	for i < n && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	for i < n && intervals[i][0] <= newInterval[1] {
		if intervals[i][0] < newInterval[0] {
			newInterval[0] = intervals[i][0]
		}
		if intervals[i][1] > newInterval[1] {
			newInterval[1] = intervals[i][1]
		}
		i++
	}

	result = append(result, newInterval)

	for i < n {
		result = append(result, intervals[i])
		i++
	}

	return result
}
