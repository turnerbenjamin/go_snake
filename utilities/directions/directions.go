package directions

type Direction byte

const (
	Up Direction = iota
	Down
	Left
	Right
	InvalidDirection
)