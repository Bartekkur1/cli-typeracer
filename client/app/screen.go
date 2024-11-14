package app

// @TODO: Add should rerender
type Screen interface {
	Render()
	Init(game *Game)
	InitOnce(game *Game)
	GetInputHandlers() []InputHandler
	GetNetworkHandlers() []NetworkHandler
	HandleEsc(game *Game)
}
