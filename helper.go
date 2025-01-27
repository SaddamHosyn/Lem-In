package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Functions for handling comments and formatting
func RemoveComments(originalFileLines []string) []string {
	var NewLines []string
	for _, line := range originalFileLines  {
		if strings.HasPrefix(line, "#") && line != "##start" && line != "##end" {
     continue
		}
		NewLines = append(NewLines, line)
	}
  return NewLines
}


func Error(msg string) {
	fmt.Println("Error: InValid data format")
	if msg != "" {
		fmt.Println("\033[101m" + msg + "\033[0m")
	}
	os.Exit(1)
}

// Helper functions for validation and checks
func IsNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func DashesInLine(s []string) {

	for _, line := range s {
	 if len(strings.Split(line, "-")) > 2 {
		 Error("2  or more dashes are not allowed")
	 }
	}
} 

// DoubleLines checks if there are duplicate lines in the slice
func DoubleLines(s []string) {
	seen := make(map[string]bool)

	for _, line := range s {
	 if seen[line] {
		 Error("Duplicate lines are not allowed")
	 }
	 seen[line] = true
	}
}

func NoHashInLastLine(s []string) {
	if strings.HasPrefix(s[len(s)-1], "#") {
		Error("")
	}
}

func ExtractStartRoom(s []string) {

	found := false 
	for i, line := range s {
		if line == "##start" {
			if found {
				Error("Multiple ##start declarations are not allowed")
			}
			found = true 

			if i+1 >= len(s) || !IsItARoom(s[i+1]) {
				Error("Invalid start room declaration")
			}

			ah.StartRoom = ConvertToRoom(s[i+1])
		}
	}

	if !found {
		Error("No ##start declaration found")
	}
}

func ExtractEndRoom(s []string) {
	for i, line := range s {
		if line == "##end" {
			if i+1 < len(s) && IsItARoom(s[i+1]) {
				ah.EndRoom = ConvertToRoom(s[i+1])
			} else {
				Error("")
			}
		}
	}
}


// DeleteStartRoom deletes the start room from the slice
func DeleteStartRoom(s []string) []string {

	var ExtractedLines []string
	skipNext := false

	for _, line := range s {
		if skipNext {
			skipNext = false
			continue
		}
		if line == "##start" {
			skipNext = true
			continue
		}
		ExtractedLines = append(ExtractedLines, line)
	}

	return ExtractedLines
}

func ExtractRooms(s []string) {
	var rooms []*FRoom
	for _, line := range s {
		if IsItARoom(line) {
			rooms = append(rooms, ConvertToRoom(line))

		}
	}
	rooms = append(rooms, ah.StartRoom, ah.EndRoom)
	NoDuplicateCoordsOrNames(rooms)
	ah.FRooms = rooms
}

// NoDuplicateCoordsOrNames checks if there are duplicate coordinates in the slice
func NoDuplicateCoordsOrNames(s []*FRoom) {
	for i, room := range s {
		for j, room2 := range s {
			if i != j && room.X == room2.X && room.Y == room2.Y {
				Error("Duplicate coordinates are not allowed")
			}
			if i != j && room.Name == room2.Name {
				Error("Duplicate room names are not allowed")
			}
		}
	}
}



// DeleteEndRoom deletes the end room from the slice
func DeleteEndRoom(s []string) []string {
	var ExtractedLines []string
	skipNext := false

	for _, line := range s {
		if skipNext {
			skipNext = false
			continue
		}
		if line == "##end" {
			skipNext = true
			continue
		}
		ExtractedLines = append(ExtractedLines, line)
	}

	return ExtractedLines
}


func ConvertToRoom(RoomStr string) *FRoom {

	RoomStrSlice := strings.Split(RoomStr, " ")

	if len(RoomStrSlice) != 3 {
	 Error("Invalid room format: " + RoomStr)
	 return nil
	}

	roomname := RoomStrSlice[0]

	x, err := strconv.Atoi(RoomStrSlice[1])
	if err != nil {
	 Error("Invalid x coordinate:" + RoomStrSlice[1] )
	}


	y, err := strconv.Atoi(RoomStrSlice[2])
	if err != nil {
	 Error("Invalid x coordinate:" + RoomStrSlice[2] )
	}

	return &FRoom{
	 Name: roomname,
	 X:    x,
	 Y:    y,
}
}

func DeleteAllRooms(s []string) []string {
	var ExtractedLines []string
	for _, line := range s {
		if !IsItARoom(line) {
			ExtractedLines = append(ExtractedLines, line)
		}
	}
	return ExtractedLines
}


func CheckRoomsConnections(OnlyConnections []string, AllRooms []string) {
	for _, connectionStr := range OnlyConnections {
		// Split the connection string into room names
		roomNames := strings.Split(connectionStr, "-")
		if !roomExists(AllRooms, roomNames[0]) || !roomExists(AllRooms, roomNames[1]) {
			Error("ERROR: room in connection not present in rooms")
		}
	}
}

func GetAllRoomNames(ah *AntHill) []string {
	var roomNames []string
	for _, room := range ah.FRooms {
		roomNames = append(roomNames, room.Name)
	}
	return roomNames
}

func GetRoomByName(name string) *FRoom {
	for _, room := range ah.FRooms {
		if room.Name == name {
			return room
		}
	}
	return nil
}

// Functions for handling connections
func AddConnections(OnlyConnections []string) {
	for _, connection := range OnlyConnections {
		room1Name := strings.Split(connection, "-")[0]
		room2Name := strings.Split(connection, "-")[1]
		room1 := GetRoomByName(room1Name)
		room2 := GetRoomByName(room2Name)
		room1.Connections = append(room1.Connections, room2)
		room2.Connections = append(room2.Connections, room1)
	}
}


func isValidRoomName(name string) bool {
	words := strings.Fields(name)
	if len(words) != 3 {
		return false
	}
	_, err := strconv.Atoi(words[1])
	if err != nil {
		return false
	}
	_, err = strconv.Atoi(words[2])
	return err == nil
}


func chkDuplicateCoords(lines []string) {
	coords := make(map[string]bool)
	countCoord := 0
	for _, s := range lines {
		if strings.Contains(s, " ") {
			countCoord++
			words := strings.Fields(s)
			coords[words[1]+" "+words[2]] = true
		}
	}
	if len(coords) != countCoord {
		log.Fatal("ERROR: invalid data format. Duplicate coordinates")
	}
}

// Functions for checking room connections
func checkUnconnectedRooms(ah *AntHill) {
	for _, room := range ah.FRooms {
		if len(room.Connections) == 0 {
			Error(fmt.Sprintf("The room \"%v\" is not connected to the anthill", room.Name))
		}
	}
}

// No # in last line, or it is a start or end room
func HashInLastLine(s []string) {
	if strings.HasPrefix(s[len(s)-1], "#") {
		Error("")
	}
}

// IsRoom checks if a string is a room
func IsItARoom(s string) bool {
	parts := strings.Split(s, " ")
	if len(parts) != 3 {
		return false
	}
	if !IsNumber(parts[1]) || !IsNumber(parts[2]) {
		return false
	}
	return true
}



// Helper function to check if a room exists in the list of all rooms
func roomExists(rooms []string, room string) bool {
	for _, r := range rooms {
		if r == room {
			return true
		}
	}
	return false
}




// validateFileGiveMeStrings validates the input file and returns its contents as a slice of strings.
// It checks if the file exists, reads its content, and performs several validations:
// - Ensures the file contains at least six lines, representing the number of ants, start room, end room, and at least one connection.
// - Checks if the number of ants is a positive integer.
// - Verifies that each line (except the first) has a minimum length of 3 characters.
// - Ensures there is exactly one '##start' and one '##end' identifier, each followed by a valid room name.
// - Removes lines starting with '#' that are not '##start' or '##end'.
// - Calls helper functions to check for continuous connections and duplicate coordinates.
// If any validation fails, the function logs a fatal error and terminates the program.

// Functions for handling file validation

func validateFileGiveMeStrings() []string {
	f, err := os.Stat(os.Args[1])
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("ERROR: invalid data format. File does not exist")
		}
	}
	data, err := os.ReadFile(f.Name())
	if err != nil {
		log.Fatal("ERROR: invalid data format. File reading error", err)
	}
	data = []byte(strings.ReplaceAll(string(data), "\r\n", "\n"))
	lines := strings.Split(string(data), "\n")
	for len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	if len(lines) < 6 {
		log.Fatal("ERROR: invalid data format. Not enough lines")
	}

	for i := 1; i < len(lines); i++ {
		if len(lines[i]) < 3 {
			log.Fatal("ERROR: invalid data format. Line is too short", lines[i], "at line number", i)
		}
	}

	numA, err := strconv.Atoi(lines[0])
	if err != nil {
		log.Fatal("ERROR: invalid data format. First line is not a number")
	}
	if numA <= 0 {
		log.Fatal("ERROR: invalid data format. Number of ants is negative or zero")
	}

	startCount := 0
	endCount := 0
	for _, s := range lines {
		if strings.Contains(s, "##start") {
			startCount++
		}
		if strings.Contains(s, "##end") {
			endCount++
		}
	}
	if startCount > 1 || endCount > 1 {
		log.Fatal("ERROR: invalid data format. More than one ##start or ##end")
	}
	if startCount == 0 || endCount == 0 {
		log.Fatal("ERROR: invalid data format. No ##start or ##end")
	}

	for i, s := range lines {
		if strings.HasPrefix(s, "#") && len(s) > 1 && !strings.Contains(s, "##start") && !strings.Contains(s, "##end") {
			lines = append(lines[:i], lines[i+1:]...)
		}
	}

	for i, s := range lines {
		if strings.Contains(s, "##start") && i+1 < len(lines) {
			if !isValidRoomName(lines[i+1]) {
				log.Fatal("ERROR: invalid data format. Invalid room name")
			}
			break
		}
		if strings.Contains(s, "##end") && i+1 < len(lines) {
			if !isValidRoomName(lines[i+1]) {
				log.Fatal("ERROR: invalid data format. Invalid room name")
			}
			break
		}
	}
	chkConsRInTheEnd(lines)
	chkDuplicateCoords(lines)
	return lines
}

func chkConsRInTheEnd(lines []string) {
	for {
		if strings.Contains((lines[len(lines)-1]), "-") {
			lines = lines[:len(lines)-1]
		} else {
			break
		}
	}
	for _, s := range lines {
		if strings.Contains(s, "-") {
			log.Fatal("ERROR: invalid data format. Invalid connection, all connections have to be continuous in the end")
		}
	}
}

// Functions for graph manipulation
func DeepCopyGraph(g *Graph) *Graph {
	newGraph := &Graph{Rooms: []*Room{}}
	for _, room := range g.Rooms {
		newGraph.Rooms = append(newGraph.Rooms, &Room{
			Roomname:    room.Roomname,
			Connections: make([]string, len(room.Connections)),
			Visited:     room.Visited,
		})
		copy(newGraph.Rooms[len(newGraph.Rooms)-1].Connections, room.Connections)
	}
	newGraph.StartRoomName = g.StartRoomName
	newGraph.EndRoomName = g.EndRoomName
	newGraph.Ants = g.Ants
	return newGraph
}

func PopulateGraph(lines []string, g *Graph) error {
	var err error
	g.Ants, err = strconv.Atoi(lines[0])
	if err != nil {
		return err
	}
	if g.Ants == 0 {
		return errors.New("ERROR: invalid data format. Number of ants must be greater than 0")
	}

	start := false
	end := false
	for _, line := range lines[1:] {
		space := strings.Split(line, " ")

		if len(space) > 1 {
			if !isValidRoomName(line) {
				log.Fatalf("ERROR: invalid data format. Room name or room coordinates invalid")
			}
			g.AddRoom(space[0])
		}

		if start {
			g.StartRoomName = g.Rooms[len(g.Rooms)-1].Roomname
			start = false
		} else if end {
			g.EndRoomName = g.Rooms[len(g.Rooms)-1].Roomname
			end = false
		}

		hyphen := strings.Split(line, "-")
		if len(hyphen) > 1 {
			if hyphen[0] == hyphen[1] {
				log.Fatalf("ERROR: invalid data format. You have a connection from the same room to same room.\n")
			}
			g.AddTunnels(hyphen[0], hyphen[1])
		}
		switch line {
		case "##start":
			start = true
		case "##end":
			end = true
		}
	}

	return nil
}





