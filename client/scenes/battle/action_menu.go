package battle

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/jchapman63/neo-gkmn/client/util"
)

type Menu struct {
	Dimensions        util.Dimensions
	Origin            util.Point
	verticalPadding   float64
	horizontalPadding float64
	color             color.Color
	font              *text.GoTextFace
}

func NewMenu(d util.Dimensions, vPad float64, hPad float64) *Menu {
	return &Menu{
		Dimensions:        d,
		verticalPadding:   vPad,
		horizontalPadding: hPad,
	}
}

func (m *Menu) NewMenuImage() *ebiten.Image {
	return ebiten.NewImage(m.Dimensions.Width, m.Dimensions.Height)
}

func (m *Menu) ConstructMenuImage() *ebiten.Image {
	img := ebiten.NewImage(m.Dimensions.Width, m.Dimensions.Height)
	img.Fill(m.color)

	return img
}

func (m *Menu) OptionBox(wDiv float64, hDiv float64) *ebiten.Image {
	w, h := float64(m.Dimensions.Width)*wDiv, float64(m.Dimensions.Height)*hDiv
	img := ebiten.NewImage(int(w), int(h))
	return img
}
