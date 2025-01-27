package main

// The Graph structure keeps track of all rooms the ant can take, the start and end rooms of the path and the number of ants
type Graph struct {
	Rooms         []*Room
	StartRoomName string
	EndRoomName   string
	Ants          int
}

// Room represents a room in an ant hill
type FRoom struct {
	Name        string   // name of the room
	X           int      // x-coordinate of the room
	Y           int      // y-coordinate of the room
	Connections []*FRoom // list of rooms connected to this room
	Visited     bool     // whether this room has been visited or not
	Distance    int      // distance of this room from the start room
}

// AntHill represents an ant hill with rooms, a start room, an end room, and the number of ants
type AntHill struct {
	FRooms    []*FRoom // list of all rooms in the ant hill
	StartRoom *FRoom   // start room of the ant hill
	EndRoom   *FRoom   // end room of the ant hill
	Ants      int      // number of ants in the ant hill
}

// The Room structure keeps track of the roomname, The rooms that the the current room is connected to and if the room has been visited before
type Room struct {
	Roomname    string
	Connections []string
	Visited     bool
}

var anthill AntHill // global variable representing the ant hill

var pathArray []string
