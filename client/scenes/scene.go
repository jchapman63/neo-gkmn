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
	canvas  *Canvas
	screen  *ebiten.Image
	config  *config.Game
	textSrc *text.GoTextFaceSource
}

type Canvas struct {
	// Width of the Canvas
	Width float64
	// Height of the Canvas
	Height float64
	// Horizontal Padding - Global
	HorizontalPadding float64
	// Vertival Padding - Global
	VerticalPadding float64
}

// NewBattleGUI initializes a new BattleGUI instance.
func NewBattleGUI(screen *ebiten.Image, conf *config.Game, textSrc *text.GoTextFaceSource) *BattleGUI {
	w, h := float64(screen.Bounds().Dx()), float64(screen.Bounds().Dy())
	c := &Canvas{
		Width:             w,
		Height:            h,
		HorizontalPadding: w * padding,
		VerticalPadding:   h * padding,
	}
	return &BattleGUI{
		screen:  screen,
		config:  conf,
		textSrc: textSrc,
		canvas:  c,
	}
}

// DrawBattleGUI is the primary entry point to draw
// the background and all UI boxes for the battle scene.
func (b *BattleGUI) DrawBattleGUI() {
	bgImage, ok := b.config.GUI.Sprites["temp-bkg.png"]
	if !ok {
		log.Fatal("could not find background image")
	}
	// TODO - turn into iterable object of current monsters in battle
	mons, ok := b.config.Monsters.Sprites["bulbasaur.png"]
	if !ok {
		log.Fatal("cound not find monster image")
	}
	b.drawBackground(bgImage)
	b.drawBattleBoxes()
	b.drawMonsters(mons)
}

// drawMonsters draws pokemon adjacent to the respective entity's battle box
func (b *BattleGUI) drawMonsters(mons *ebiten.Image) {

	// pMon - flip then draw
	monH := mons.Bounds().Dy()
	monW := mons.Bounds().Dx()
	pOpts := &ebiten.DrawImageOptions{}
	pOpts.GeoM.Scale(-1, 1)
	pOpts.GeoM.Scale(2, 2)
	pOpts.GeoM.Translate(float64(monW)+(2*b.canvas.HorizontalPadding), b.canvas.Height-float64(monH)-2*b.canvas.VerticalPadding)
	b.screen.DrawImage(mons, pOpts)

	// oMon
	oOpts := &ebiten.DrawImageOptions{}
	oOpts.GeoM.Scale(2, 2)
	oOpts.GeoM.Translate(b.canvas.Width-2*b.canvas.HorizontalPadding-float64(monW), 0-b.canvas.VerticalPadding)
	// You want to see why these calculations are weird?
	// Fuck around and find out... uncomment this code.
	// blackColor := color.RGBA{0, 0, 0, 255} // Black with full opacity
	// mons.Fill(blackColor)
	b.screen.DrawImage(mons, oOpts)
}

// drawBattleBoxes draws the boxes for the opponent’s Pokémon and the player’s Pokémon.
func (b *BattleGUI) drawBattleBoxes() {
	boxW := b.canvas.Width * 0.30
	boxH := b.canvas.Height * 0.15

	// Opponent box
	oppOpts := &ebiten.DrawImageOptions{}
	oppOpts.GeoM.Translate(b.canvas.HorizontalPadding, b.canvas.VerticalPadding)
	b.drawBox("bulbasaur", boxW, boxH, oppOpts)

	// Player box
	playerX := b.canvas.Width - float64(boxW) - b.canvas.HorizontalPadding
	playerY := b.canvas.Height - float64(boxH) - b.canvas.VerticalPadding
	playerOpts := &ebiten.DrawImageOptions{}
	playerOpts.GeoM.Translate(playerX, playerY)
	b.drawBox("bulbasaur", boxW, boxH, playerOpts)
}

// drawBox creates and draws a single box containing a Pokémon’s name,
// health bar, and HP numbers.
func (b *BattleGUI) drawBox(name string, width, height float64, opts *ebiten.DrawImageOptions) {
	boxImg := ebiten.NewImage(int(width), int(height))
	boxImg.Fill(color.White)

	// Draw Pokémon name
	nameOpts := &text.DrawOptions{}
	nameOpts.ColorScale.Scale(0, 0, 0, 1)
	nameOpts.GeoM.Translate(float64(width)*padding, float64(height)*padding)
	text.Draw(boxImg, name, &text.GoTextFace{Source: b.textSrc, Size: 10}, nameOpts)

	// Draw health bar
	barW := width * 0.70
	barH := height * 0.10
	healthBar := ebiten.NewImage(int(barW), int(barH))
	healthBar.Fill(color.RGBA{0x00, 0xff, 0x00, 0xff}) // Green

	healthOpts := &ebiten.DrawImageOptions{}
	barY := height - barH - height*padding
	healthOpts.GeoM.Translate(width*padding, barY)
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

	bgW, bgH := float64(bgImage.Bounds().Dx()), float64(bgImage.Bounds().Dy())

	scaleW := b.canvas.Width / bgW
	scaleH := b.canvas.Height / bgH

	// Choose the larger scale to fill the screen.
	scale := scaleW
	if scale < scaleH {
		scale = scaleH
	}

	opts := &ebiten.DrawImageOptions{}
	// Start from image center
	opts.GeoM.Translate(-bgW/2, -bgH/2)
	opts.GeoM.Scale(scale, scale)
	opts.GeoM.Translate(b.canvas.Width/2, b.canvas.Height/2)
	opts.Filter = ebiten.FilterLinear

	b.screen.DrawImage(bgImage, opts)
}
