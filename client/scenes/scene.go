package scenes

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/jchapman63/neo-gkmn/client/config"
)

const padding float64 = 0.05

type BattleGUI struct {
	screen *ebiten.Image
	gConf  *config.GUI
	face   *text.GoTextFaceSource
}

func NewBattleGUI(screen *ebiten.Image, conf *config.GUI, face *text.GoTextFaceSource) *BattleGUI {
	return &BattleGUI{
		screen: screen,
		gConf:  conf,
		face:   face,
	}
}

func (b *BattleGUI) DrawBattleGUI() {
	bkg, ok := b.gConf.Sprites["temp-bkg.png"]
	if !ok {
		log.Fatal("could not find background image")
	}
	b.drawBackground(bkg)
	b.drawBattleBoxes()
}

func (b *BattleGUI) drawBox(name string, width int, height int, opts *ebiten.DrawImageOptions) {
	bBox := ebiten.NewImage(width, height)
	bBox.Fill(color.White)

	// draw text in box
	tOps := &text.DrawOptions{}
	tOps.ColorScale.Scale(0, 0, 0, 1)
	tOps.GeoM.Translate(float64(width)*padding, float64(height)*padding)
	text.Draw(bBox, name, &text.GoTextFace{Source: b.face, Size: 10}, tOps)

	// draw healthBar
	barW, barH := float64(width)*0.70, float64(height)*0.10
	healthBar := ebiten.NewImage(int(barW), int(barH))
	healthBar.Fill(color.RGBA{0x00, 0xff, 0x00, 0xff})
	hOps := &ebiten.DrawImageOptions{}
	barY := float64(height) - barH - float64(height)*padding
	hOps.GeoM.Translate(float64(width)*padding, barY)
	bBox.DrawImage(healthBar, hOps)

	// draw health numbers
	tOps = &text.DrawOptions{}
	tOps.GeoM.Translate(float64(width)*padding, barY-barH-10)
	tOps.ColorScale.Scale(0, 0, 0, 1)
	text.Draw(bBox, "25/25", &text.GoTextFace{Source: b.face, Size: 10}, tOps)

	// draw box
	b.screen.DrawImage(bBox, opts)
}

func (b *BattleGUI) drawBattleBoxes() {
	sw, sh := float64(b.screen.Bounds().Dx()), float64(b.screen.Bounds().Dy())
	horizontalPadding := sw * padding
	verticalPadding := sh * padding
	bw, bh := int(sw*0.30), int(sh*0.15)

	// opp box drawing
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(horizontalPadding, verticalPadding)
	b.drawBox("bulbasaur", bw, bh, op)

	// player box drawing
	playerX := sw - float64(bw) - horizontalPadding
	playerY := sh - float64(bh) - verticalPadding
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(playerX, playerY)
	b.drawBox("pikachu", bw, bh, op)

	// draw pokemon name in box
	tOps := &text.DrawOptions{}
	tOps.GeoM.Translate(float64(bw)*padding, float64(bh)*padding)
	tOps.ColorScale.Scale(0, 0, 0, 1)
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
