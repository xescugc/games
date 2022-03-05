package main

import "github.com/gorilla/websocket"

type Room struct {
	Name        string
	Connections map[string]*websocket.Conn
	Game        *Game
}

func NewRoom(name, sid string, conn *websocket.Conn) *Room {
	return &Room{
		Name: name,
		Connections: map[string]*websocket.Conn{
			sid: conn,
		},
		Game: NewGame(sid),
	}
}

func (r *Room) AddPlayer(sid string, conn *websocket.Conn) {
	r.Connections[sid] = conn
	r.Game.AddPlayer(sid)
}

func (r *Room) RemovePlayer(sid string) {
	delete(r.Connections, sid)
	r.Game.RemovePlayer(sid)
}
