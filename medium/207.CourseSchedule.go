package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)

	for _, prequisite := range prerequisites {
		course, prequi := prequisite[0], prequisite[1]

		graph[course] = append(graph[course], prequi)
	}

	recStack := make(map[int]bool)
	visited := make(map[int]bool)

	for course := range graph {
		if hasCycle(graph, course, visited, recStack) {
			return false
		}
	}

	return true
}

func hasCycle(graph map[int][]int, course int, visited, recStack map[int]bool) bool {
	if recStack[course] {
		return true
	}

	if visited[course] {
		return false
	}

	recStack[course] = true
	visited[course] = true

	for _, neighbor := range graph[course] {
		if hasCycle(graph, neighbor, visited, recStack) {
			return true
		}
	}

	recStack[course] = false

	return false
}
