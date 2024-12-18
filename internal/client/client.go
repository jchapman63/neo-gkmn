package client

// a client implements the game logic for an end user
import (
	"embed"
	"fmt"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed sprites/gui/*.png
var gui embed.FS

////go:embed sprites/monsters/*.png
//var mons embed.FS

type Game struct {
	guiSprites map[string]*ebiten.Image
	//monSprites map[string]*ebiten.Image
}

func NewGame() (*Game, error) {
	game := &Game{}

	if err := game.fetchSprites(); err != nil {
		return nil, err
	}

	return game, nil
}

func (g *Game) fetchSprites() error {
	gSprit, err := gui.ReadDir("sprites/gui")
	if err != nil {
		return err
	}
	for _, sprite := range gSprit {
		filePath := "sprites/gui" + sprite.Name()
		fmt.Println(sprite.Name())
		img, _, err := ebitenutil.NewImageFromFileSystem(gui, filePath)
		if err != nil {
			return nil
		}
		g.guiSprites[sprite.Name()] = img
		fmt.Println("gui sprites", g.guiSprites[sprite.Name()])
	}
	return nil
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
	// TODO - abstraction, only expecting one image for now
	screen.DrawImage(g.guiSprites["emptybox.png"], nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func Run(game *Game) error {
	ebiten.SetWindowSize(640, 480)
	if err := ebiten.RunGame(game); err != nil {
		return err
	}
	return nil
}
