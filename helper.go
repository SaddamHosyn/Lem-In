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
	for _, line := range originalFileLines {
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

			anthill.StartRoom = ConvertToRoom(s[i+1])
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
				anthill.EndRoom = ConvertToRoom(s[i+1])
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
	rooms = append(rooms, anthill.StartRoom, anthill.EndRoom)
	NoDuplicateCoordsOrNames(rooms)
	anthill.FRooms = rooms
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
		Error("Invalid x coordinate:" + RoomStrSlice[1])
	}

	y, err := strconv.Atoi(RoomStrSlice[2])
	if err != nil {
		Error("Invalid x coordinate:" + RoomStrSlice[2])
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

func CheckRoomsConnections(connections, rooms []string) {
	roomMap := make(map[string]bool)
	for _, room := range rooms {
		roomMap[room] = true
	}

	for _, conn := range connections {
		parts := strings.Split(conn, "-")
		if len(parts) != 2 || !roomMap[parts[0]] || !roomMap[parts[1]] {
			Error("ERROR: room in connection not present in rooms")
		}
	}
}

func GetAllRoomNames(anthill *AntHill) []string {
	var roomNames []string
	for _, room := range anthill.FRooms {
		roomNames = append(roomNames, room.Name)
	}
	return roomNames
}

func GetRoomByName(name string) *FRoom {
	for _, room := range anthill.FRooms {
		if room.Name == name {
			return room
		}
	}
	return nil
}

// Functions for handling connections
func AddLinks(OnlyConnections []string) {
	for _, connection := range OnlyConnections {
		parts := strings.Split(connection, "-")
		if len(parts) != 2 {
			Error("Invalid connection format")
		}
		room1 := GetRoomByName(parts[0])
		room2 := GetRoomByName(parts[1])

		// Add connections both ways
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

func checkDuplicateCoordinates(lines []string) {
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
func CheckUnconnectedRooms(ah *AntHill) {
	for _, room := range ah.FRooms {
		if len(room.Connections) == 0 {
			Error(fmt.Sprintf("Room %q has no connections", room.Name))
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
	if len(os.Args) < 2 {
		log.Fatal("ERROR: No file provided")
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("ERROR: File does not exist")
		}
		log.Fatal("ERROR: File reading error", err)
	}

	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	for len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1] // Remove empty lines
	}

	if len(lines) < 6 {
		log.Fatal("ERROR: Not enough lines in the file")
	}

	numAnts, err := strconv.Atoi(lines[0])
	if err != nil || numAnts <= 0 {
		log.Fatal("ERROR: Invalid number of ants")
	}

	startCount, endCount := 0, 0
	for i, line := range lines {
		if len(line) < 3 && i > 0 { // Skip the first line (number of ants)
			log.Fatal("ERROR: Line is too short at line", i+1)
		}

		if strings.Contains(line, "##start") {
			startCount++
			if i+1 >= len(lines) || !isValidRoomName(lines[i+1]) {
				log.Fatal("ERROR: Invalid room name after ##start")
			}
		}
		if strings.Contains(line, "##end") {
			endCount++
			if i+1 >= len(lines) || !isValidRoomName(lines[i+1]) {
				log.Fatal("ERROR: Invalid room name after ##end")
			}
		}
	}

	if startCount != 1 || endCount != 1 {
		log.Fatal("ERROR: Missing or duplicate ##start or ##end")
	}

	// Remove comments (lines starting with #, except ##start and ##end)
	processline := make([]string, 0, len(lines))
	for _, line := range lines {
		if !strings.HasPrefix(line, "#") || strings.HasPrefix(line, "##start") || strings.HasPrefix(line, "##end") {
			processline = append(processline, line)
		}
	}

	checkConnectionInTheEnd(processline)
	checkDuplicateCoordinates(processline)

	return processline
}

func checkConnectionInTheEnd(lines []string) {
	// Remove lines containing "-" from the end until no more are found
	for len(lines) > 0 && strings.Contains(lines[len(lines)-1], "-") {
		lines = lines[:len(lines)-1]
	}

	// Check if any remaining lines contain "-"
	for _, line := range lines {
		if strings.Contains(line, "-") {
			log.Fatal("ERROR: invalid data format. Invalid connection, all connections must be continuous at the end")
		}
	}
}

// Functions for graph manipulation
func CopyFullGraph(g *Graph) *Graph {
	newGraph := &Graph{
		Rooms:         make([]*Room, 0, len(g.Rooms)),
		StartRoomName: g.StartRoomName,
		EndRoomName:   g.EndRoomName,
		Ants:          g.Ants,
	}

	for _, room := range g.Rooms {
		newRoom := &Room{
			Roomname:    room.Roomname,
			Connections: append([]string{}, room.Connections...), // Copy connections
			Visited:     room.Visited,
		}
		newGraph.Rooms = append(newGraph.Rooms, newRoom)
	}

	return newGraph
}

func PopulateGraph(lines []string, g *Graph) error {
	ants, err := strconv.Atoi(lines[0])
	if err != nil || ants <= 0 {
		return errors.New("ERROR: invalid data format. Number of ants must be greater than 0")
	}
	g.Ants = ants

	var start, end bool
	for _, line := range lines[1:] {
		switch {
		case line == "##start":
			start = true
		case line == "##end":
			end = true
		case strings.Contains(line, " "):
			if !isValidRoomName(line) {
				return errors.New("ERROR: invalid data format. Room name or coordinates invalid")
			}
			g.AddRoom(strings.Split(line, " ")[0])
			if start {
				g.StartRoomName = g.Rooms[len(g.Rooms)-1].Roomname
				start = false
			} else if end {
				g.EndRoomName = g.Rooms[len(g.Rooms)-1].Roomname
				end = false
			}
		case strings.Contains(line, "-"):
			rooms := strings.Split(line, "-")
			if rooms[0] == rooms[1] {
				return errors.New("ERROR: invalid data format. Connection from the same room to itself")
			}
			g.AddTunnels(rooms[0], rooms[1])
		}
	}

	return nil
}
