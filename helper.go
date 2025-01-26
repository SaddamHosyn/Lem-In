package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Helper functions for validation and checks
func IsNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
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

func Contains(slice []string, elem string) bool {
	return strings.Contains(strings.Join(slice, "😎"), elem)
}

func NoGo(msg string) {
	fmt.Println("ERROR: invalid data format")
	if msg != "" {
		fmt.Println("\033[101m" + msg + "\033[0m")
	}
	os.Exit(1)
}

// Functions for handling comments and formatting
func RemoveComments(originalFileLines []string) []string {
	var ExtractedLines []string
	for _, line := range originalFileLines {
		if strings.HasPrefix(line, "#") && line != "##end" && line != "##start" {
			continue
		}
		ExtractedLines = append(ExtractedLines, line)
	}
	return ExtractedLines
}

func NoHashInLastLine(s []string) {
	if strings.HasPrefix(s[len(s)-1], "#") {
		NoGo("")
	}
}

// Functions for room extraction and validation
func IsRoom(s string) bool {
	return !((len(strings.Split(s, " ")) != 3) || !IsNumber(strings.Split(s, " ")[1]) || !IsNumber(strings.Split(s, " ")[2]))
}

func ConvertToRoom(roomStr string) *FRoom {
	roomStrSlice := strings.Split(roomStr, " ")
	rName := roomStrSlice[0]
	x, _ := strconv.Atoi(roomStrSlice[1])
	y, _ := strconv.Atoi(roomStrSlice[2])
	return &FRoom{
		Name: rName,
		X:    x,
		Y:    y,
	}
}

func ExtractStartRoom(s []string) {
	for i, line := range s {
		if line == "##start" {
			if i+1 < len(s) && IsRoom(s[i+1]) {
				ah.StartRoom = ConvertToRoom(s[i+1])
			} else {
				NoGo("")
			}
		}
	}
}

func ExtractEndRoom(s []string) {
	for i, line := range s {
		if line == "##end" {
			if i+1 < len(s) && IsRoom(s[i+1]) {
				ah.EndRoom = ConvertToRoom(s[i+1])
			} else {
				NoGo("")
			}
		}
	}
}

func ExtractRooms(s []string) {
	var rooms []*FRoom
	for _, line := range s {
		if IsRoom(line) {
			rooms = append(rooms, ConvertToRoom(line))
		}
	}
	rooms = append(rooms, ah.StartRoom)
	rooms = append(rooms, ah.EndRoom)
	NoDuplicateCoordsOrNames(rooms)
	ah.FRooms = rooms
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

func CheckRoomsInConnectionsPresent(OnlyConnections []string, AllRooms []string) {
	for _, connectionStr := range OnlyConnections {
		roomNames := strings.Split(connectionStr, "-")
		if !Contains(AllRooms, roomNames[0]) || !Contains(AllRooms, roomNames[1]) {
			NoGo("ERROR: room in connection not present in rooms")
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

// Functions for duplicate checks
func NoDuplicateLines(s []string) {
	for i, line := range s {
		for j, line2 := range s {
			if i != j && line == line2 {
				NoGo("Duplicate lines are not allowed")
			}
		}
	}
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

func NoDuplicateCoordsOrNames(s []*FRoom) {
	for i, room := range s {
		for j, room2 := range s {
			if i != j && room.X == room2.X && room.Y == room2.Y {
				NoGo("Duplicate coordinates are not allowed")
			}
			if i != j && room.Name == room2.Name {
				NoGo("Duplicate room names are not allowed")
			}
		}
	}
}

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

// Functions for room deletion
func DeleteStartRoom(s []string) []string {
	var ExtractedLines []string
	startRoomIndex := -1
	for i, line := range s {
		if i == startRoomIndex {
			continue
		}
		if line == "##start" {
			startRoomIndex = i + 1
			continue
		}
		ExtractedLines = append(ExtractedLines, line)
	}
	return ExtractedLines
}

func DeleteEndRoom(s []string) []string {
	var ExtractedLines []string
	endRoomIndex := -1
	for i, line := range s {
		if i == endRoomIndex {
			continue
		}
		if line == "##end" {
			endRoomIndex = i + 1
			continue
		}
		ExtractedLines = append(ExtractedLines, line)
	}
	return ExtractedLines
}

func DeleteAllRooms(s []string) []string {
	var ExtractedLines []string
	for _, line := range s {
		if !IsRoom(line) {
			ExtractedLines = append(ExtractedLines, line)
		}
	}
	return ExtractedLines
}

// Functions for checking room connections
func checkUnconnectedRooms(ah *AntHill) {
	for _, room := range ah.FRooms {
		if len(room.Connections) == 0 {
			NoGo(fmt.Sprintf("The room \"%v\" is not connected to the anthill", room.Name))
		}
	}
}

// Functions for checking dashes and spaces
func No2Dashes(s []string) {
	for _, line := range s {
		if len(strings.Split(line, "-")) > 2 {
			NoGo("2 or more dashes in a line are not allowed")
		}
	}
}

func No3Spaces(s []string) {
	for _, line := range s {
		if len(strings.Split(line, " ")) > 3 {
			NoGo("3 or more spaces in a line are not allowed")
		}
	}
}
