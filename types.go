package main

type Drawer interface {
	Draw()
}

type Updater interface {
	Update()
}

type UpdateDrawer interface {
	Updater
	Drawer
}