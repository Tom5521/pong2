package main

type Drawer interface {
	Draw()
}

type Updater interface {
	Update()
}

type UpdaterDrawer interface {
	Drawer
	Updater
}
