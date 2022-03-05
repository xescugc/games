package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
	"github.com/xescugc/games/tiles/message"
)

var upgrader = websocket.Upgrader{}
var (
	rooms        = make(map[string]*Room)
	playersRooms = make(map[string]string)
	players      = make(map[string]*websocket.Conn)
	addressID    = make(map[string]string)

	port string
)

func init() {
	flag.StringVar(&port, "port", ":5555", "The port of the application with the ':'")
}

func main() {
	flag.Parse()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go startGameLoop(ctx)
	http.HandleFunc("/ws", wsHandler)
	log.Printf("Staring server at %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, _ := upgrader.Upgrade(w, r, nil)
	defer ws.Close()

	for {
		var msg message.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			sid := addressID[ws.RemoteAddr().String()]
			delete(players, sid)
			r := rooms[playersRooms[sid]]
			r.RemovePlayer(sid)
			ws.Close()
			break
		}

		switch msg.Type {
		case message.Open:
			sid := uuid.Must(uuid.NewV4())
			err = ws.WriteJSON(message.NewConnectMessage(sid.String()))
			if err != nil {
				log.Fatal(err)
			}
			players[sid.String()] = ws
			addressID[ws.RemoteAddr().String()] = sid.String()
		case message.JoinRoom:
			r, ok := rooms[msg.JoinRoom.Room]
			if !ok {
				r = NewRoom(msg.JoinRoom.Room, msg.JoinRoom.SessionID, ws)
				rooms[msg.JoinRoom.Room] = r
			} else {
				r.AddPlayer(msg.JoinRoom.SessionID, ws)
			}
			playersRooms[msg.JoinRoom.SessionID] = r.Name
		case message.Move:
			rooms[playersRooms[msg.Move.SessionID]].Game.MovePlayer(msg.Move.SessionID, msg.Move.Direction)
		}
	}
}

func startGameLoop(ctx context.Context) {
	ticker := time.NewTicker(time.Second / 50)
	for {
		select {
		case <-ticker.C:
			for _, r := range rooms {
				for _, con := range r.Connections {
					err := con.WriteJSON(message.NewUpdateMessage(r.Game.Players))
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		case <-ctx.Done():
			ticker.Stop()
			goto FINISH
		}
	}
FINISH:
}
