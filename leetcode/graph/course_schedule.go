package graph

func buildGraph(prerequisites [][]int) map[int][]int {
	graph := make(map[int][]int)
	for i := range prerequisites {
		arr, ok := graph[prerequisites[i][1]]
		if !ok {
			arr = make([]int, 0)
		}
		arr = append(arr, prerequisites[i][0])
		graph[prerequisites[i][1]] = arr
	}

	return graph
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := buildGraph(prerequisites)

	visitedMap := make(map[int]bool)
	for i := range graph {
		if isCyclic(visitedMap, graph, i) {
			return false
		}
	}

	return true
}

func isCyclic(visitedMap map[int]bool, graph map[int][]int, start int) bool {
	if _, ok := visitedMap[start]; ok && !visitedMap[start] {
		return false
	}
	if visitedMap[start] {
		return true
	}
	visitedMap[start] = true

	for i := range graph[start] {
		if isCyclic(visitedMap, graph, graph[start][i]) {
			return true
		}
	}

	visitedMap[start] = false

	return false
}
