package config

import "github.com/hajimehoshi/ebiten/v2"

type Window struct {
	Width  int
	Height int
}

type Game struct {
	GUI      *GUI
	Monsters *Monsters
}
type GUI struct {
	Sprites map[string]*ebiten.Image
}

type Monsters struct {
	Sprites map[string]*ebiten.Image
}
