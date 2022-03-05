package player

type Object struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	W float64 `json:"w"`
	H float64 `json:"h"`
}

type Player struct {
	Object      `json:"object"`
	ID          string `json:"id"`
	Facing      string `josn:"facing"`
	Moving      bool   `json:"moving"`
	MovingCount int    `json:moving_count"`
}
