package scenes

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/jchapman63/neo-gkmn/client/config"
)

const padding = 0.05

// BattleGUI handles drawing the background and status boxes
// for a Pokémon battle scene.
type BattleGUI struct {
	screen  *ebiten.Image
	guiConf *config.GUI
	textSrc *text.GoTextFaceSource
}

// NewBattleGUI initializes a new BattleGUI instance.
func NewBattleGUI(screen *ebiten.Image, conf *config.GUI, textSrc *text.GoTextFaceSource) *BattleGUI {
	return &BattleGUI{
		screen:  screen,
		guiConf: conf,
		textSrc: textSrc,
	}
}

// DrawBattleGUI is the primary entry point to draw
// the background and all UI boxes for the battle scene.
func (b *BattleGUI) DrawBattleGUI() {
	bgImage, ok := b.guiConf.Sprites["temp-bkg.png"]
	if !ok {
		log.Fatal("could not find background image")
	}

	b.drawBackground(bgImage)
	b.drawBattleBoxes()
}

// drawBattleBoxes draws the boxes for the opponent’s Pokémon and the player’s Pokémon.
func (b *BattleGUI) drawBattleBoxes() {
	screenW, screenH := float64(b.screen.Bounds().Dx()), float64(b.screen.Bounds().Dy())
	hPad := screenW * padding
	vPad := screenH * padding
	boxW := int(screenW * 0.30)
	boxH := int(screenH * 0.15)

	// Opponent box
	oppOpts := &ebiten.DrawImageOptions{}
	oppOpts.GeoM.Translate(hPad, vPad)
	b.drawBox("bulbasaur", boxW, boxH, oppOpts)

	// Player box
	playerX := screenW - float64(boxW) - hPad
	playerY := screenH - float64(boxH) - vPad
	playerOpts := &ebiten.DrawImageOptions{}
	playerOpts.GeoM.Translate(playerX, playerY)
	b.drawBox("pikachu", boxW, boxH, playerOpts)
}

// drawBox creates and draws a single box containing a Pokémon’s name,
// health bar, and HP numbers.
func (b *BattleGUI) drawBox(name string, width, height int, opts *ebiten.DrawImageOptions) {
	boxImg := ebiten.NewImage(width, height)
	boxImg.Fill(color.White)

	// Draw Pokémon name
	nameOpts := &text.DrawOptions{}
	nameOpts.ColorScale.Scale(0, 0, 0, 1)
	nameOpts.GeoM.Translate(float64(width)*padding, float64(height)*padding)
	text.Draw(boxImg, name, &text.GoTextFace{Source: b.textSrc, Size: 10}, nameOpts)

	// Draw health bar
	barW := float64(width) * 0.70
	barH := float64(height) * 0.10
	healthBar := ebiten.NewImage(int(barW), int(barH))
	healthBar.Fill(color.RGBA{0x00, 0xff, 0x00, 0xff}) // Green

	healthOpts := &ebiten.DrawImageOptions{}
	barY := float64(height) - barH - float64(height)*padding
	healthOpts.GeoM.Translate(float64(width)*padding, barY)
	boxImg.DrawImage(healthBar, healthOpts)

	// Draw HP text
	hpOpts := &text.DrawOptions{}
	hpOpts.ColorScale.Scale(0, 0, 0, 1)
	hpOpts.GeoM.Translate(float64(width)*padding, barY-barH-10)
	text.Draw(boxImg, "25/25", &text.GoTextFace{Source: b.textSrc, Size: 10}, hpOpts)

	// Finally draw the box onto the screen
	b.screen.DrawImage(boxImg, opts)
}

// drawBackground fills the entire screen with white and then draws the
// background image scaled to fill the screen area.
func (b *BattleGUI) drawBackground(bgImage *ebiten.Image) {
	b.screen.Fill(color.White)

	screenW, screenH := float64(b.screen.Bounds().Dx()), float64(b.screen.Bounds().Dy())
	bgW, bgH := float64(bgImage.Bounds().Dx()), float64(bgImage.Bounds().Dy())

	scaleW := screenW / bgW
	scaleH := screenH / bgH

	// Choose the larger scale to fill the screen.
	scale := scaleW
	if scale < scaleH {
		scale = scaleH
	}

	opts := &ebiten.DrawImageOptions{}
	// Start from image center
	opts.GeoM.Translate(-bgW/2, -bgH/2)
	opts.GeoM.Scale(scale, scale)
	opts.GeoM.Translate(screenW/2, screenH/2)
	opts.Filter = ebiten.FilterLinear

	b.screen.DrawImage(bgImage, opts)
}
