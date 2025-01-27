package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Graph-related methods
func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <filename>")
		return
	}
	orgLines, err := ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Remove the comments from the original file lines
	NewLines := RemoveComments(orgLines)

	// check length of slice to be minimum 6: 1st line is number of ants, 2nd  and 3rd line is start room, 4th and 5th line is end room, 6th line is a link
	if len(NewLines) < 6 {
		Error("")
	}

	// check if first line is a number
	if !IsNumber(NewLines[0]) {
		Error("")
	}

	// convert first line to int and store in AntNum
	anthill.Ants, _ = strconv.Atoi(NewLines[0])
	NewLines = NewLines[1:]

	// check if number of ants is valid
	if anthill.Ants <= 0 {
		Error("Number of ants is invalid")
	}

	DashesInLine(NewLines)
	DoubleLines(NewLines)
	NoHashInLastLine(NewLines)

	// extract start room
	ExtractStartRoom(NewLines)
	NewLines = DeleteStartRoom(NewLines)

	// extract end room
	ExtractEndRoom(NewLines)
	NewLines = DeleteEndRoom(NewLines)

	// extract rooms
	ExtractRooms(NewLines)
	OnlyLinks := DeleteAllRooms(NewLines)

	// check if any room is there in the connections that is not in the rooms
	CheckRoomsConnections(OnlyLinks, GetAllRoomNames(&anthill))

	// Add Connections to the rooms where a connection is in the format "room1-room2" and room1 and room2 are in the rooms
	AddLinks(OnlyLinks)

	CheckUnconnectedRooms(&anthill)

	lines := validateFileGiveMeStrings()
	//create a graph for the DFS search
	gfordfs := &Graph{Rooms: []*Room{}}
	if err := PopulateGraph(lines, gfordfs); err != nil {
		fmt.Print(err)
		return
	}
	gforbfs := CopyFullGraph(gfordfs)
	// Print the contents of the slice with a new line after each element
	fmt.Println(strings.Join(orgLines, "\n") + "\n")

	allPathsDFS, allPathsBFS := []string{}, []string{}
	var path string
	DFS(gfordfs.StartRoomName, gfordfs.EndRoomName, gfordfs, path, &allPathsDFS)
	BFS(gforbfs.StartRoomName, gforbfs.EndRoomName, gforbfs, &allPathsBFS, ShortestPath)
	lenSorter(&allPathsBFS)
	lenSorter(&allPathsDFS)
	antNum := gfordfs.Ants
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
