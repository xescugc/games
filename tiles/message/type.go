package message

type Type string

const (
	Open     Type = "open"
	Connect       = "connect"
	JoinRoom      = "join_room"
	Move          = "move"
	Update        = "update"
)
