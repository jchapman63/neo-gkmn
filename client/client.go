package client

// a client implements the game logic for an end user
import (
	"bytes"
	"embed"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/jchapman63/neo-gkmn/client/config"
	"github.com/jchapman63/neo-gkmn/client/scenes"
)

//go:embed sprites/gui/*.png
var gui embed.FS

//go:embed sprites/monsters/*.png
var mons embed.FS

type Game struct {
	Window *config.Window
	Config *config.Game
	GUI    *config.GUI
	Face   *text.GoTextFaceSource
	// Mons config.Monsters
}

func NewGame() (*Game, error) {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))
	if err != nil {
		log.Fatal(err)
	}
	w := &config.Window{
		Width:  640,
		Height: 480,
	}
	game := &Game{
		Window: w,
		Config: &config.Game{
			GUI: &config.GUI{
				Sprites: map[string]*ebiten.Image{},
			},
			Monsters: &config.Monsters{
				Sprites: map[string]*ebiten.Image{},
			},
		},
		Face: s,
	}
	if err := game.fetchSprites(); err != nil {
		return nil, err
	}

	return game, nil
}

// fetches sprites from local dir to build battle scene
func (g *Game) fetchSprites() error {
	// gui sprites
	gSprit, err := gui.ReadDir("sprites/gui")
	if err != nil {
		return err
	}
	for _, sprite := range gSprit {
		filePath := "sprites/gui/" + sprite.Name()
		img, _, err := ebitenutil.NewImageFromFileSystem(gui, filePath)
		if err != nil {
			return err
		}
		g.Config.GUI.Sprites[sprite.Name()] = img
	}

	// monster sprites
	mSprites, err := mons.ReadDir("sprites/monsters")
	if err != nil {
		return err
	}
	for _, sprite := range mSprites {
		filePath := "sprites/monsters/" + sprite.Name()
		img, _, err := ebitenutil.NewImageFromFileSystem(mons, filePath)
		if err != nil {
			return err
		}
		g.Config.Monsters.Sprites[sprite.Name()] = img
	}
	return nil
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	bGUI := scenes.NewBattleGUI(screen, g.Config, g.Face)
	bGUI.DrawBattleGUI()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Window.Width, g.Window.Height
}

func Run(game *Game) error {
	ebiten.SetWindowSize(640, 480)
	if err := ebiten.RunGame(game); err != nil {
		return err
	}
	return nil
}
