package app

// @TODO: Add should rerender
type Screen interface {
	Render()
	Mount(game *Game)
	DisMount(game *Game)
	HandleEsc(game *Game)
}
