package message

import "github.com/xescugc/games/tiles/player"

type Message struct {
	Type Type `json:"type"`

	JoinRoom JoinRoomMessage `json:"join,omitempty"`
	Open     OpenMessage     `json:"open,omitempty"`
	Connect  ConnectMessage  `json:"connect,omitempty"`
	Move     MoveMessage     `json:"move,omitempty"`
	Update   UpdateMessage   `json:"update,omitempty"`
}

type JoinRoomMessage struct {
	SessionID string `json:"session_id"`
	Room      string `json:"room"`
}

func NewJoinRoom(sid, room string) *Message {
	return &Message{
		Type: JoinRoom,
		JoinRoom: JoinRoomMessage{
			Room:      room,
			SessionID: sid,
		},
	}
}

type OpenMessage struct {
}

func NewOpenMessage() *Message {
	return &Message{
		Type: Open,
		Open: OpenMessage{},
	}
}

type ConnectMessage struct {
	SessionID string
}

func NewConnectMessage(sid string) *Message {
	return &Message{
		Type: Connect,
		Connect: ConnectMessage{
			SessionID: sid,
		},
	}
}

type MoveMessage struct {
	SessionID string
	Direction string
}

func NewMoveMessage(sid, dir string) *Message {
	return &Message{
		Type: Move,
		Move: MoveMessage{
			SessionID: sid,
			Direction: dir,
		},
	}
}

type UpdateMessage struct {
	Players map[string]*player.Player
}

func NewUpdateMessage(players map[string]*player.Player) *Message {
	return &Message{
		Type: Update,
		Update: UpdateMessage{
			Players: players,
		},
	}
}
