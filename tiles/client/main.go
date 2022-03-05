package main

import (
	"bytes"
	_ "embed"
	"flag"
	"image"
	_ "image/png"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xescugc/games/tiles/message"
)

//go:embed assets/TilesetFloor.png
var Tileset_png []byte

//go:embed assets/Walk.png
var Walk_png []byte

var (
	tilesetImg image.Image
	playerImg  image.Image
	wsc        *websocket.Conn

	wsHost string
	room   string
)

const (
	screenW = 240
	screenH = 240
)

func init() {
	tsi, _, err := image.Decode(bytes.NewReader(Tileset_png))
	if err != nil {
		log.Fatal(err)
	}
	tilesetImg = ebiten.NewImageFromImage(tsi)

	wi, _, err := image.Decode(bytes.NewReader(Walk_png))
	if err != nil {
		log.Fatal(err)
	}
	playerImg = ebiten.NewImageFromImage(wi)

	flag.StringVar(&wsHost, "ws-host", ":5555", "The host of the server, the format is 'host:port'")
	flag.StringVar(&room, "room", "", "The room to connect to")
}

func main() {
	flag.Parse()
	if room == "" {
		log.Fatal("The -room is required")
	}
	u := url.URL{Scheme: "ws", Host: wsHost, Path: "/ws"}

	// Establish connection
	var err error
	wsc, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer wsc.Close()

	ebiten.SetWindowTitle("Tiles")
	ebiten.SetWindowSize(screenW*2, screenH*2)
	g := &Game{
		Columns:  22,
		TileSize: 16,
		Tiles: []int{
			264, 264, 264, 264, 264, 264, 176, 178, 264, 264, 264, 264, 264, 264, 264,
			264, 264, 264, 264, 264, 264, 176, 178, 264, 264, 264, 264, 264, 264, 264,
			264, 264, 264, 264, 264, 264, 176, 178, 264, 264, 264, 264, 264, 264, 264,
			264, 264, 264, 264, 264, 264, 176, 178, 264, 264, 264, 264, 264, 264, 264,
			264, 264, 264, 264, 264, 264, 176, 178, 264, 264, 264, 264, 264, 264, 264,
			264, 264, 264, 264, 264, 264, 176, 178, 264, 264, 264, 264, 264, 264, 264,
			264, 264, 264, 264, 264, 154, 204, 203, 156, 264, 264, 264, 264, 264, 264,
			155, 155, 155, 155, 155, 204, 177, 177, 203, 155, 155, 155, 155, 155, 155,
			199, 199, 199, 199, 199, 182, 177, 177, 181, 199, 199, 199, 199, 199, 199,
			264, 264, 264, 264, 264, 198, 182, 181, 200, 264, 264, 264, 264, 264, 264,
			264, 264, 264, 264, 264, 264, 176, 178, 264, 264, 264, 264, 264, 264, 264,
			264, 264, 264, 264, 264, 264, 176, 178, 264, 264, 264, 264, 264, 264, 264,
			264, 264, 264, 264, 264, 264, 176, 178, 264, 264, 264, 264, 264, 264, 264,
			264, 264, 264, 264, 264, 264, 176, 178, 264, 264, 264, 264, 264, 264, 264,
			264, 264, 264, 264, 264, 264, 176, 178, 264, 264, 264, 264, 264, 264, 264,
		},
		Players: make(map[string]*Player),
	}
	go wsHandler(g)

	err = wsc.WriteJSON(message.NewOpenMessage())
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func wsHandler(g *Game) {
	for {
		var msg message.Message
		err := wsc.ReadJSON(&msg)
		if err != nil {
			// TODO remove from the Room
			log.Fatal(err)
		}

		switch msg.Type {
		case message.Connect:
			g.CurrentID = msg.Connect.SessionID
			err = wsc.WriteJSON(message.NewJoinRoom(g.CurrentID, room))
			if err != nil {
				log.Fatal(err)
			}
		case message.Update:
			seen := make(map[string]struct{})
			for k, up := range msg.Update.Players {
				if p, ok := g.Players[k]; ok {
					p.Player = *up
				} else {
					g.Players[k] = &Player{Player: *up}
				}
				seen[k] = struct{}{}
			}
			for k := range g.Players {
				if _, ok := seen[k]; !ok {
					delete(g.Players, k)
				}
			}
		}
	}
}
