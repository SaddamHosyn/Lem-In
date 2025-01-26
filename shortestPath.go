package main

import "fmt"

// ShortestPath finds all the possible paths from start to end room using BFS and sorts them in ascending order
func ShortestPath(graph *Graph, start string, end string, path []string) []string {
	path = append(path, start)
	if start == end {
		return path
	}

	shortest := make([]string, 0)
	for _, node := range graph.getRoom(start).Connections {
		if !contains(path, node) && !graph.isVisited(node) {
			newPath := ShortestPath(graph, node, end, path)
			if len(newPath) > 0 && contains(newPath, graph.StartRoomName) && contains(newPath, end) {
				pathArray = append(pathArray, fmt.Sprint(newPath))
			}
		}
	}
	return shortest
}