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
	IsActive    bool
	Origin      *Point
	Translation *Point
}

func NewBtnImg(img *ebiten.Image, origin *Point, translation *Point, message string, font *text.GoTextFace) (*BtnImg, *ebiten.DrawImageOptions) {
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
		Origin:      origin,
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

	// translate to the image's physical location
	minOriginX := float64(min.X) + b.Origin.X
	minOriginY := b.Origin.Y + float64(min.Y)

	maxOriginX := float64(max.X) + b.Origin.X
	maxOriginY := b.Origin.Y + float64(max.Y)

	// Check if the point (x, y) is within the bounds
	if pt.X >= minOriginX && pt.X <= maxOriginX && pt.Y >= minOriginY && pt.Y <= maxOriginY {
		fmt.Println("Click is within bounds")
		return true
	}

	fmt.Println("Click is out of bounds")
	return false
}
