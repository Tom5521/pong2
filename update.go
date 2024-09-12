package main

func (g *Game) Update() {
	// Update instances.
	for _, u := range g.Updaters {
		g.Frames++
		u.Update()
	}
}
