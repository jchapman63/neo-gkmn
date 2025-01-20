package scenes

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/jchapman63/neo-gkmn/client/config"
	"github.com/jchapman63/neo-gkmn/client/util"
)

const padding = 0.05

type BattleGUIConfig struct {
	Window  *config.Window
	Sprites *config.Sprites
	textSrc *text.GoTextFaceSource
	Buttons []*util.BtnImg
	canvas  *Canvas
	menu    *Menu
}

// BattleGUI handles drawing the background and status boxes
// for a Pokémon battle scene.
type BattleGUI struct {
	Config *BattleGUIConfig
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

type Menu struct {
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
func NewBattleGUI(window *config.Window, sprites *config.Sprites, textSrc *text.GoTextFaceSource) *BattleGUI {
	// Menu will take up 25% of the screen height, canvas will take up 75% of the screen height
	w, h := float64(window.Width), float64(window.Height)
	c := &Canvas{
		Width:             w,
		Height:            h * 0.66,
		HorizontalPadding: w * padding,
		VerticalPadding:   h * padding,
	}
	m := &Menu{
		Width:             w,
		Height:            h * 0.33,
		HorizontalPadding: w * padding,
		VerticalPadding:   h * padding,
	}
	config := &BattleGUIConfig{
		Window:  window,
		Sprites: sprites,
		textSrc: textSrc,
		canvas:  c,
		menu:    m,
	}
	return &BattleGUI{
		Config: config,
	}
}

// DrawBattleGUI is the primary entry point to draw
// the background and all UI boxes for the battle scene.
func (b *BattleGUI) DrawBattleGUI(screen *ebiten.Image) {
	bgImage, ok := b.Config.Sprites.GUI.Sprites["temp-bkg.png"]
	if !ok {
		log.Fatal("could not find background image")
	}
	// TODO - turn into iterable object of current monsters in battle
	mons, ok := b.Config.Sprites.Monsters.Sprites["bulbasaur.png"]
	if !ok {
		log.Fatal("cound not find monster image")
	}
	b.drawBackground(screen, bgImage)
	b.drawMenu(screen)
	b.drawBattleBoxes(screen)
	b.drawMonsters(screen, mons)
}

// drawMonsters draws pokemon adjacent to the respective entity's battle box
func (b *BattleGUI) drawMonsters(screen *ebiten.Image, mons *ebiten.Image) {

	// pMon - flip then draw
	monH := mons.Bounds().Dy()
	monW := mons.Bounds().Dx()
	pOpts := &ebiten.DrawImageOptions{}
	pOpts.GeoM.Scale(-1, 1)
	pOpts.GeoM.Scale(2, 2)
	pOpts.GeoM.Translate(float64(monW)+(2*b.Config.canvas.HorizontalPadding), b.Config.canvas.Height-float64(monH)-2*b.Config.canvas.VerticalPadding)
	screen.DrawImage(mons, pOpts)

	// oMon
	oOpts := &ebiten.DrawImageOptions{}
	oOpts.GeoM.Scale(2, 2)
	oOpts.GeoM.Translate(b.Config.canvas.Width-2*b.Config.canvas.HorizontalPadding-float64(monW), 0-b.Config.canvas.VerticalPadding)
	// You want to see why these calculations are weird?
	// Fuck around and find out... uncomment this code.
	// blackColor := color.RGBA{0, 0, 0, 255} // Black with full opacity
	// mons.Fill(blackColor)
	screen.DrawImage(mons, oOpts)
}

// drawBattleBoxes draws the boxes for the opponent’s Pokémon and the player’s Pokémon.
func (b *BattleGUI) drawBattleBoxes(screen *ebiten.Image) {
	boxW := b.Config.canvas.Width * 0.30
	boxH := b.Config.canvas.Height * 0.15

	// Opponent box
	oppOpts := &ebiten.DrawImageOptions{}
	oppOpts.GeoM.Translate(b.Config.canvas.HorizontalPadding, b.Config.canvas.VerticalPadding)
	b.drawBattleBox(screen, "bulbasaur", boxW, boxH, oppOpts)

	// Player box
	playerX := b.Config.canvas.Width - float64(boxW) - b.Config.canvas.HorizontalPadding
	playerY := b.Config.canvas.Height - float64(boxH) - b.Config.canvas.VerticalPadding
	playerOpts := &ebiten.DrawImageOptions{}
	playerOpts.GeoM.Translate(playerX, playerY)
	b.drawBattleBox(screen, "bulbasaur", boxW, boxH, playerOpts)
}

// drawBox creates and draws a single box containing a Pokémon’s name,
// health bar, and HP numbers.
func (b *BattleGUI) drawBattleBox(screen *ebiten.Image, name string, width, height float64, opts *ebiten.DrawImageOptions) {
	boxImg := ebiten.NewImage(int(width), int(height))
	boxImg.Fill(color.White)

	// Draw Pokémon name
	nameOpts := &text.DrawOptions{}
	nameOpts.ColorScale.Scale(0, 0, 0, 1)
	nameOpts.GeoM.Translate(float64(width)*padding, float64(height)*padding)
	text.Draw(boxImg, name, &text.GoTextFace{Source: b.Config.textSrc, Size: 10}, nameOpts)

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
	text.Draw(boxImg, "25/25", &text.GoTextFace{Source: b.Config.textSrc, Size: 10}, hpOpts)

	// Finally draw the box onto the screen
	screen.DrawImage(boxImg, opts)
}

// drawBackground fills the entire screen with white and then draws the
// background image scaled to fill the screen area.
func (b *BattleGUI) drawBackground(screen *ebiten.Image, bgImage *ebiten.Image) {
	screen.Fill(color.White)

	bgW, bgH := float64(bgImage.Bounds().Dx()), float64(bgImage.Bounds().Dy())

	scaleW := b.Config.canvas.Width / bgW
	scaleH := b.Config.canvas.Height / bgH

	// Choose the larger scale to fill the screen.
	scale := scaleW
	if scale < scaleH {
		scale = scaleH
	}

	opts := &ebiten.DrawImageOptions{}
	// Start from image center
	opts.GeoM.Translate(-bgW/2, -bgH/2)
	opts.GeoM.Scale(scale, scale)
	opts.GeoM.Translate(b.Config.canvas.Width/2, b.Config.canvas.Height/2)
	opts.Filter = ebiten.FilterLinear

	screen.DrawImage(bgImage, opts)
}

// draw menu implements a menu using the BattleGUI.menu attribute
// The menu is a rectangle that takes up the portion of the screen
// height as specified by BattleGUI.menu
func (b *BattleGUI) drawMenu(screen *ebiten.Image) {
	// cover specified height of the screen
	menuImg := ebiten.NewImage(int(b.Config.menu.Width), int(b.Config.menu.Height))
	menuImg.Fill(color.RGBA{144, 238, 144, 255})
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(0, float64(b.Config.Window.Height)-b.Config.menu.Height)

	// create button
	boxW := b.Config.menu.Width * 0.30
	boxH := b.Config.menu.Height * 0.15
	boxImgBtn := util.NewBtnImg(ebiten.NewImage(int(boxW), int(boxH)))

	// draw button
	boxImgBtn.Img.Fill(color.White)
	bOpts := &ebiten.DrawImageOptions{}
	// Draw into center for now
	bOpts.GeoM.Translate(b.Config.menu.Width/2, b.Config.menu.Height/2)
	// Draw inner text
	tOpts := &text.DrawOptions{}
	tOpts.ColorScale.Scale(0, 0, 0, 1)
	text.Draw(boxImgBtn.Img, "Tackle", &text.GoTextFace{Source: b.Config.textSrc, Size: 10}, tOpts)

	// register button
	b.Config.Buttons = append(b.Config.Buttons, boxImgBtn)

	// draw button into menu
	menuImg.DrawImage(boxImgBtn.Img, bOpts)
	// draw menu onto screen
	screen.DrawImage(menuImg, opts)
}
