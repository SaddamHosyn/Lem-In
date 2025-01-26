package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Graph-related methods
func (g *Graph) AddRoom(name string) {
	g.Rooms = append(g.Rooms, &Room{Roomname: name, Connections: []string{}, Visited: false})
}

func (g *Graph) AddTunnels(from, to string) {
	fromRoom := g.getRoom(from)
	toRoom := g.getRoom(to)
	if fromRoom == nil || toRoom == nil {
		log.Fatalf("No Room is present (%v-%v)", from, to)
	}
	if contains(fromRoom.Connections, to) || contains(toRoom.Connections, from) {
		log.Fatalf("Error: Link Duplication (%v --- %v)", from, to)
	}
	switch {
	case fromRoom.Roomname == g.EndRoomName:
		toRoom.Connections = append(toRoom.Connections, fromRoom.Roomname)
	case toRoom.Roomname == g.EndRoomName:
		fromRoom.Connections = append(fromRoom.Connections, toRoom.Roomname)
	case toRoom.Roomname == g.StartRoomName:
		toRoom.Connections = append(toRoom.Connections, fromRoom.Roomname)
	case fromRoom.Roomname == g.StartRoomName:
		fromRoom.Connections = append(fromRoom.Connections, toRoom.Roomname)
	default:
		fromRoom.Connections = append(fromRoom.Connections, toRoom.Roomname)
		toRoom.Connections = append(toRoom.Connections, fromRoom.Roomname)
	}
}

func (g *Graph) getRoom(name string) *Room {
	for _, room := range g.Rooms {
		if room.Roomname == name {
			return room
		}
	}
	return nil
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <filename>")
		return
	}
	originalFileLines, err := ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Remove the comments from the original file lines
	NewLines := RemoveComments(originalFileLines)

	// check length of slice to be minimum 6: 1st line is number of ants, 2nd  and 3rd line is start room, 4th and 5th line is end room, 6th line is a link
	if len(NewLines ) < 6 {
		Error("")
	}

	// check if first line is a number
	if !IsNumber(NewLines[0]) {
		Error("")
	}

	// convert first line to int and store in AntNum
	ah.Ants, _ = strconv.Atoi(NewLines[0])
	NewLines  = NewLines[1:]

	// check if number of ants is valid
	if ah.Ants <= 0 {
		Error("Number of ants is invalid")
	}

   DashesInLine(NewLines)
	 DoubleLines(NewLines)
	NoHashInLastLine(NewLines)

	// extract start room
	ExtractStartRoom(NewLines )
	NewLines  = DeleteStartRoom(NewLines )

	// extract end room
	ExtractEndRoom(NewLines )
	NewLines  = DeleteEndRoom(NewLines )

	// extract rooms
	ExtractRooms(NewLines )
	OnlyConnections := DeleteAllRooms(NewLines )

	// check if any room is there in the connections that is not in the rooms
	CheckRoomsInConnectionsPresent(OnlyConnections, GetAllRoomNames(&ah))

	// Add Connections to the rooms where a connection is in the format "room1-room2" and room1 and room2 are in the rooms
	AddConnections(OnlyConnections)

	checkUnconnectedRooms(&ah)

	/////////////////////////////////////////////////

	lines := validateFileGiveMeStrings()
	_ = lines
	//create a graph for the DFS search
	gdfs := &Graph{Rooms: []*Room{}}
	if err := PopulateGraph(lines, gdfs); err != nil {
		fmt.Print(err)
		return
	}
	gbfs := DeepCopyGraph(gdfs)
	// Print the contents of the slice with a new line after each element
	fmt.Println(strings.Join(originalFileLines, "\n") + "\n")

	allPathsDFS, allPathsBFS := []string{}, []string{}
	var path string
	DFS(gdfs.StartRoomName, gdfs.EndRoomName, gdfs, path, &allPathsDFS)
	BFS(gbfs.StartRoomName, gbfs.EndRoomName, gbfs, &allPathsBFS, ShortestPath)
	lenSorter(&allPathsBFS)
	lenSorter(&allPathsDFS)
	antNum := gdfs.Ants
	DFSSearch := AntSender(antNum, allPathsDFS)
	BFSSearch := AntSender(antNum, allPathsBFS)

	if len(DFSSearch) == 0 || len(BFSSearch) == 0 {
		if len(DFSSearch) == 0 {
			fmt.Println(fmt.Errorf("DFS Search Failed").Error())
		}
		if len(BFSSearch) == 0 {
			fmt.Println(fmt.Errorf("BFS Search Failed").Error())
		}
		return
	}

	for _, step := range shorterSearch(DFSSearch, BFSSearch) {
		fmt.Println(step)
	}
}
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
	return SortLexicalStrings(finalMoves)
}

func LexicalLsort(s string) string {
	words := strings.Fields(s)
	sort.Slice(words, func(i, j int) bool {
		num1, _ := strconv.Atoi(strings.Split(words[i], "-")[0][1:])
		num2, _ := strconv.Atoi(strings.Split(words[j], "-")[0][1:])
		return num1 < num2
	})
	return strings.Join(words, " ")
}

func SortLexicalStrings(strs []string) []string {
	sortedStrs := make([]string, len(strs))
	for i, s := range strs {
		sortedStr := LexicalLsort(s)
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

