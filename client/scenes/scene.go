package scenes

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jchapman63/neo-gkmn/client/config"
)

const padding float64 = 0.05

type BattleGUI struct {
	screen *ebiten.Image
	gConf  *config.GUI
}

func NewBattleGUI(screen *ebiten.Image, conf *config.GUI) *BattleGUI {
	return &BattleGUI{
		screen: screen,
		gConf:  conf,
	}
}

func (b *BattleGUI) DrawBattleGUI() {
	bkg, ok := b.gConf.Sprites["temp-bkg.png"]
	if !ok {
		log.Fatal("could not find background image")
	}

	battleBox, ok := b.gConf.Sprites["emptybox.png"]
	if !ok {
		log.Fatal("could not find background image")
	}
	b.drawBackground(bkg)
	b.drawBattleBox(battleBox)

	width := float64(b.screen.Bounds().Dx())
	height := float64(b.screen.Bounds().Dy())
	horizontalPadding := width * padding
	verticalPadding := height * padding

	// opp mon coordinates
	oppMonX := horizontalPadding + width*float64(0.7)
	oppMonY := verticalPadding + height*float64(0.0)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(3, 3)
	op.GeoM.Translate(oppMonX, oppMonY)
	b.screen.DrawImage(battleBox, op)

	// player mon coordinates
	pMonX := horizontalPadding + width*float64(0.0)
	pMonY := verticalPadding + height*float64(0.7)
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Scale(3, 3)
	op.GeoM.Translate(pMonX, pMonY)
	b.screen.DrawImage(battleBox, op)
}

func (b *BattleGUI) drawBattleBox(battleBox *ebiten.Image) {
	sw := float64(b.screen.Bounds().Dx())
	sh := float64(b.screen.Bounds().Dy())
	horizontalPadding := sw * padding
	verticalPadding := sh * padding

	boxWidth := float64(battleBox.Bounds().Dx())

	// opp box coordinates
	oppX := horizontalPadding
	oppY := verticalPadding
	// opp box drawing
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(6, 5)
	op.GeoM.Translate(oppX, oppY)
	b.screen.DrawImage(battleBox, op)

	// player box coordinates
	playerX := sw - boxWidth - horizontalPadding
	playerY := verticalPadding + sh*float64(0.6)
	// player box drawing
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Scale(6, 5)
	op.GeoM.Translate(playerX, playerY)
	b.screen.DrawImage(battleBox, op)
}

func (b *BattleGUI) drawBackground(bgImage *ebiten.Image) {
	b.screen.Fill(color.White)
	width, height := b.screen.Bounds().Dx(), b.screen.Bounds().Dy()

	bw, bh := bgImage.Bounds().Dx(), bgImage.Bounds().Dy()
	scaleW := float64(width) / float64(bw)
	scaleH := float64(height) / float64(bh)
	scale := scaleW
	if scale < scaleH {
		scale = scaleH
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(bw)/2, -float64(bh)/2)
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(float64(width)/2, float64(height)/2)
	op.Filter = ebiten.FilterLinear
	b.screen.DrawImage(bgImage, op)
}
