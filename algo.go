package main

import (
	"sort"
	"strconv"
	"strings"
)

// BFS preforms a Breadth First Search of a graph from rooms start to end and puts all paths found in the []string paths
func BFS(start, end string, g *Graph, paths *[]string, f func(graph *Graph, start string, end string, path []string) []string) {
	begin := g.getRoom(start)

	for i := 0; i < len(begin.Connections); i++ {
		var shortPath []string
		ShortestPath(g, g.StartRoomName, g.EndRoomName, shortPath)
		var shortStorer string
		if len(pathArray) != 0 {
			shortStorer = pathArray[0]
		}

		for _, v := range pathArray {
			if len(v) < len(shortStorer) {
				shortStorer = v
			}
		}

		if len(pathArray) != 0 {
			shortStorer = shortStorer[1 : len(shortStorer)-1]
		}

		shortStorerSlc := strings.Split(shortStorer, " ")
		shortStorerSlc = shortStorerSlc[1:]

		for z := 0; z < len(shortStorerSlc)-1; z++ {
			g.getRoom(shortStorerSlc[z]).Visited = true
		}

		var pathStr string
		if len(shortStorerSlc) != 0 {
			for i := 0; i < len(shortStorerSlc); i++ {
				if i == len(shortStorerSlc)-1 {
					pathStr += shortStorerSlc[i]
				} else {
					pathStr = pathStr + shortStorerSlc[i] + "-"
				}
			}
		}

		if len(pathStr) != 0 {
			if len(pathStr) != 0 {
				containing := false
				for _, v := range *paths {
					if v == pathStr {
						containing = true
					}
				}
				if !containing {
					*paths = append(*paths, pathStr)
				}
			}
			pathArray = []string{}
		}
	}
}

// DFS preforms a depth first search of a graph and returns the possible paths
func DFS(current, end string, g *Graph, path string, pathList *[]string) {
	curr := g.getRoom(current)
	if current != end {
		curr.Visited = true
	}
	if curr.Roomname == g.EndRoomName {
		path += current
	} else if !(curr.Roomname == g.StartRoomName) {
		path += current + "-"
	}

	if current == end {
		*pathList = append(*pathList, path)
		path = ""
		for i := 0; i < len(g.getRoom(g.StartRoomName).Connections); i++ {
			if g.getRoom(g.StartRoomName).Connections[i] == g.EndRoomName {
				g.getRoom(g.StartRoomName).Connections[i] = ""
			}
		}
		DFS(g.StartRoomName, end, g, path, pathList)
	}
	for i := 0; i < len(curr.Connections); i++ {
		if curr.Connections[i] == g.EndRoomName {
			curr.Connections[0], curr.Connections[i] = curr.Connections[i], curr.Connections[0]
		}
	}
	for _, roomName := range curr.Connections {
		if roomName == "" {
			continue
		}
		currRoom := g.getRoom(roomName)
		if !currRoom.Visited {
			DFS(currRoom.Roomname, end, g, path, pathList)
		}
	}
}

func (graph *Graph) isVisited(str string) bool {
	return graph.getRoom(str).Visited
}

func lenSorter(paths *[]string) {
	sort.Slice(*paths, func(i, j int) bool {
		return len((*paths)[i]) < len((*paths)[j])
	})
}

func AntSender(n int, pathList []string) []string {
	pathLists := make([][]string, len(pathList))
	for i, path := range pathList {
		pathLists[i] = strings.Split(path, "-")
	}

	queue := make([][]string, len(pathList))

	for i := 1; i <= n; i++ {
		minStepsIndex := 0
		minSteps := len(pathLists[0]) + len(queue[0])
		for j, path := range pathLists {
			steps := len(path) + len(queue[j])
			if steps < minSteps {
				minSteps = steps
				minStepsIndex = j
			}
		}
		queue[minStepsIndex] = append(queue[minStepsIndex], strconv.Itoa(i))
	}

	container := make([][][]string, len(queue))
	for i, path := range queue {
		for _, ant := range path {
			adder := make([]string, len(pathLists[i]))
			for j, room := range pathLists[i] {
				adder[j] = "L" + ant + "-" + room
			}
			container[i] = append(container[i], adder)
		}
	}

	finalMoves := []string{}
	for _, paths := range container {
		for j, moves := range paths {
			for k, room := range moves {
				if j+k >= len(finalMoves) {
					finalMoves = append(finalMoves, room+" ")
				} else {
					finalMoves[j+k] += room + " "
				}
			}
		}
	}
	return SortStrings(finalMoves)
}

func Sort(s string) string {
	words := strings.Fields(s)
	sort.Slice(words, func(i, j int) bool {
		num1, _ := strconv.Atoi(strings.Split(words[i], "-")[0][1:])
		num2, _ := strconv.Atoi(strings.Split(words[j], "-")[0][1:])
		return num1 < num2
	})
	return strings.Join(words, " ")
}

func SortStrings(strs []string) []string {
	sortedStrs := make([]string, len(strs))
	for i, s := range strs {
		sortedStr := Sort(s)
		sortedStrs[i] = sortedStr
	}
	return sortedStrs
}

func contains(s []string, name string) bool {
	for _, str := range s {
		if str == name {
			return true
		}
	}
	return false
}

func shorterSearch(DFSSearch, BFSSearch []string) []string {
	if len(DFSSearch) > len(BFSSearch) {
		return BFSSearch
	}
	return DFSSearch
}

