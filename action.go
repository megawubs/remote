package remote

type Command struct {
	Action Action
}

type Action int

const (
	Up Action = iota
	Down
	Power
	VolumeUp
	VolumeDown
)