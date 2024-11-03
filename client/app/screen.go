package app

// @TODO: Add should rerender
type Screen interface {
	Render()
	Init(game *Game)
	GetInputHandlers(game *Game) []InputHandler
	GetNetworkHandlers(game *Game) []NetworkHandler
	// DisMount(game *Game)
	HandleEsc(game *Game)
}
