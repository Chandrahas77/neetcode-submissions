func findOrder(numCourses int, prerequisites [][]int) []int {
    indegree := make([]int,numCourses)
	adjL := make(map[int][]int,numCourses)
	for _,pre := range prerequisites{
		course, prereq := pre[0],pre[1]
		adjL[prereq] = append(adjL[prereq],course)
		indegree[course]++
	}
	q := []int{}
	for i := range indegree{
		if indegree[i] == 0{
			q = append(q,i)
		}
	}
	res := []int{}
	for len(q) > 0{
		curr := q[0]
		q = q[1:]
		res = append(res,curr)
		for _,neighbor := range adjL[curr]{
			indegree[neighbor]--
			if indegree[neighbor] == 0{
				q = append(q,neighbor)
			}
		}
	}
	if len(res) != numCourses{
		return []int{}
	}
 	return res
}
