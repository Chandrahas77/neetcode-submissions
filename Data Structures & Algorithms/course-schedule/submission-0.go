func canFinish(numCourses int, prerequisites [][]int) bool {
    //Map reps adjList[courses] = [pre-requisites]
	adjList := make(map[int][]int,numCourses)
	indegree := make([]int,numCourses)
	for _,pre := range prerequisites{
		course,prereq := pre[0], pre[1]
		indegree[course]++
		adjList[prereq] = append(adjList[prereq],course)
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
		for _,neighbor := range adjList[curr]{
			indegree[neighbor]--
			if indegree[neighbor] == 0{
				q = append(q,neighbor)
			}
		}
	}
	return len(res) == numCourses
}
