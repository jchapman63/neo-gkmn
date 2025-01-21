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
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/jchapman63/neo-gkmn/client/config"
	"github.com/jchapman63/neo-gkmn/client/scenes"
	"github.com/jchapman63/neo-gkmn/client/util"
)

//go:embed sprites/gui/*.png
var gui embed.FS

//go:embed sprites/monsters/*.png
var mons embed.FS

type Game struct {
	Window *config.Window
	Scene  *scenes.BattleGUI
}

func NewGame() (*Game, error) {
	ts, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))
	if err != nil {
		log.Fatal(err)
	}
	w := &config.Window{
		Width:  640,
		Height: 480,
	}

	mons, err := fetchMonsterSprites()
	if err != nil {
		log.Fatal(err)
	}
	gui, err := fetchGUISprites()
	if err != nil {
		log.Fatal(err)
	}
	sprs := &config.Sprites{
		GUI: &config.GUI{
			Sprites: gui,
		},
		Monsters: &config.Monsters{
			Sprites: mons,
		},
	}
	scn := scenes.NewBattleGUI(w, sprs, ts)
	game := &Game{
		Window: w,
		Scene:  scn,
	}

	return game, nil
}

func fetchGUISprites() (map[string]*ebiten.Image, error) {
	sprites := map[string]*ebiten.Image{}
	// gui sprites
	gSprit, err := gui.ReadDir("sprites/gui")
	if err != nil {
		return nil, err
	}
	for _, sprite := range gSprit {
		filePath := "sprites/gui/" + sprite.Name()
		img, _, err := ebitenutil.NewImageFromFileSystem(gui, filePath)
		if err != nil {
			return nil, err
		}
		sprites[sprite.Name()] = img
	}
	return sprites, nil
}

func fetchMonsterSprites() (map[string]*ebiten.Image, error) {
	sprites := map[string]*ebiten.Image{}
	// monster sprites
	mSprites, err := mons.ReadDir("sprites/monsters")
	if err != nil {
		return nil, err
	}
	for _, sprite := range mSprites {
		filePath := "sprites/monsters/" + sprite.Name()
		img, _, err := ebitenutil.NewImageFromFileSystem(mons, filePath)
		if err != nil {
			return nil, err
		}
		sprites[sprite.Name()] = img
	}
	return sprites, nil
}

func (g *Game) Update() error {
	// check for registered button clicks
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		pt := util.Point{
			X: float64(x),
			Y: float64(y),
		}
		for _, btn := range g.Scene.Config.Buttons {
			btn.DidClick(pt)
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Scene.DrawBattleGUI(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Window.Width, g.Window.Height
}

func Run(game *Game) error {
	ebiten.SetWindowSize(game.Window.Width, game.Window.Height)
	if err := ebiten.RunGame(game); err != nil {
		return err
	}
	return nil
}
