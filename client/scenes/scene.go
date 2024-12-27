package scenes

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jchapman63/neo-gkmn/client/config"
)

func BattleGUI(screen *ebiten.Image, window *config.Window, GUI *config.GUI) {
	drawBackground(screen, window, GUI.Sprites["temp-bkg.png"])

	// opp box coordinates
	oppX := float64(window.Width) * float64(0.0)
	oppY := float64(window.Height) * float64(0.0)
	// opp box drawing
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(6, 5)
	op.GeoM.Translate(oppX, oppY)
	screen.DrawImage(GUI.Sprites["emptybox.png"], op)

	// player mon coordinates
	oppMonX := float64(window.Width) * float64(0.7)
	oppMonY := float64(window.Height) * float64(0.0)
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Scale(3, 3)
	op.GeoM.Translate(oppMonX, oppMonY)
	screen.DrawImage(GUI.Sprites["emptybox.png"], op)

	// player box coordinates
	playerX := float64(window.Width) * float64(0.7)
	playerY := float64(window.Height) * float64(0.7)
	// player box drawing
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Scale(6, 5)
	op.GeoM.Translate(playerX, playerY)
	screen.DrawImage(GUI.Sprites["emptybox.png"], op)

	// player mon coordinates
	pMonX := float64(window.Width) * float64(0.0)
	pMonY := float64(window.Height) * float64(0.7)
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Scale(3, 3)
	op.GeoM.Translate(pMonX, pMonY)
	screen.DrawImage(GUI.Sprites["emptybox.png"], op)

}

func drawBackground(screen *ebiten.Image, window *config.Window, bgImage *ebiten.Image) {
	screen.Fill(color.White)

	w, h := bgImage.Bounds().Dx(), bgImage.Bounds().Dy()
	scaleW := float64(window.Width) / float64(w)
	scaleH := float64(window.Height) / float64(h)
	scale := scaleW
	if scale < scaleH {
		scale = scaleH
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(float64(window.Width)/2, float64(window.Height)/2)
	op.Filter = ebiten.FilterLinear
	screen.DrawImage(bgImage, op)
}
