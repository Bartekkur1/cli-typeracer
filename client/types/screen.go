package types

type GameState int // GameState type

const (
	MainMenu GameState = iota
	JoinGame
	HostGame
	SetUsername
	SetServerURL
	Loading
	Exit
)

func (g GameState) String() string {
	return [...]string{
		"Main Menu",
		"Join Game",
		"Host Game",
		"Set Username",
		"Set Server URL",
		"Loading",
		"Exit"}[g]
}
