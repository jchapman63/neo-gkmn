package config

import "github.com/hajimehoshi/ebiten/v2"

type Window struct {
	Length int
	Width  int
}
type GUI struct {
	Sprites map[string]*ebiten.Image
}

type Monsters struct {
	Sprites map[string]*ebiten.Image
}
