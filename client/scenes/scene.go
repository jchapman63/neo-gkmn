package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jchapman63/neo-gkmn/client/config"
)

func BattleGUI(screen *ebiten.Image, window *config.Window, GUI *config.GUI) {
	op := &ebiten.DrawImageOptions{}

	// opp box
	op.GeoM.Scale(6, 5)
	op.GeoM.Translate(25, 2)
	screen.DrawImage(GUI.Sprites["emptybox.png"], op)

	// player box
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Scale(6, 5)
	op.GeoM.Translate(25, 100)
	screen.DrawImage(GUI.Sprites["emptybox.png"], op)
}
