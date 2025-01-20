package util

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Button interface {
	DidClick() bool
}

type BtnImg struct {
	Img      *ebiten.Image
	IsActive bool
}

func NewBtnImg(img *ebiten.Image) *BtnImg {
	return &BtnImg{
		Img: img,
	}
}

func (b *BtnImg) DidClick() bool {
	fmt.Println("checked for click")
	return false
}
