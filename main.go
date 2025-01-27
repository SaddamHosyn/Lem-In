package main

import (
	"fmt"
	"log"
	"os"
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
	anthill.Ants, _ = strconv.Atoi(NewLines[0])
	NewLines  = NewLines[1:]

	// check if number of ants is valid
	if anthill.Ants <= 0 {
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
	OnlyLinks := DeleteAllRooms(NewLines )

	// check if any room is there in the connections that is not in the rooms
	CheckRoomsConnections(OnlyLinks, GetAllRoomNames(&anthill))

	// Add Connections to the rooms where a connection is in the format "room1-room2" and room1 and room2 are in the rooms
	AddLinks(OnlyLinks)

	CheckUnconnectedRooms(&anthill)

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
