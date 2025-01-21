package util

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Button interface {
	DidClick() bool
}

type BtnImg struct {
	Img         *ebiten.Image
	Translation *TPoint
	IsActive    bool
}

func NewBtnImg(img *ebiten.Image, translation *TPoint, message string, font *text.GoTextFace) (*BtnImg, *ebiten.DrawImageOptions) {
	// draw button
	img.Fill(color.White)
	bOpts := &ebiten.DrawImageOptions{}
	// Draw into center for now
	bOpts.GeoM.Translate(translation.X, translation.Y)
	// Draw inner text
	tOpts := &text.DrawOptions{}
	tOpts.ColorScale.Scale(0, 0, 0, 1)
	text.Draw(img, message, font, tOpts)

	return &BtnImg{
		Translation: translation,
		Img:         img,
	}, bOpts
}

func (b *BtnImg) DidClick(pt Point) bool {
	// TODO - I might be able to take advantage of this return false
	if b.Img == nil {
		return false
	}

	// Get the bounding points of the image
	min := b.Img.Bounds().Min
	max := b.Img.Bounds().Max

	minX := float64(min.X) + b.Translation.X
	minY := float64(min.Y) + b.Translation.Y

	maxX := float64(max.X) + b.Translation.X
	maxY := float64(max.Y) + b.Translation.Y

	// Check if the point (x, y) is within the bounds
	if float64(pt.X) >= minX && float64(pt.X) <= maxX && float64(pt.Y) >= minY && float64(pt.Y) <= maxY {
		fmt.Println("Click is within bounds")
		return true
	}

	fmt.Println("Click is out of bounds")
	return false
}
