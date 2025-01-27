package main

import (
	"log"
)

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
	if fromRoom.Roomname == g.EndRoomName {
		toRoom.Connections = append(toRoom.Connections, fromRoom.Roomname)
	} else if toRoom.Roomname == g.EndRoomName {
		fromRoom.Connections = append(fromRoom.Connections, toRoom.Roomname)
	} else if toRoom.Roomname == g.StartRoomName {
		toRoom.Connections = append(toRoom.Connections, fromRoom.Roomname)
	} else if fromRoom.Roomname == g.StartRoomName {
		fromRoom.Connections = append(fromRoom.Connections, toRoom.Roomname)
	} else {
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
