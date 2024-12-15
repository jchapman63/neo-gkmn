package client

// a client implements the game logic for an end user
import (
	"embed"
	"fmt"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed sprites/*.png
var sprites embed.FS

type Game struct {
	imgs []*ebiten.Image
}

func NewGame() (*Game, error) {
	imgs, err := fetchSprites()
	if err != nil {
		return nil, err
	}

	return &Game{
		imgs: imgs,
	}, nil
}

func fetchSprites() ([]*ebiten.Image, error) {
	sprites, err := sprites.ReadDir("sprites")
	if err != nil {
		return nil, err
	}
	fmt.Println(sprites)
	var imgs []*ebiten.Image
	for _, sprite := range sprites {
		filePath := "sprites/" + sprite.Name()
		img, _, err := ebitenutil.NewImageFromFile(filePath)
		if err != nil {
			return nil, err
		}
		imgs = append(imgs, img)
	}

	return imgs, nil
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
	// TODO - abstraction, only expecting one image for now
	screen.DrawImage(g.imgs[0], nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func Run(game *Game) error {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(game); err != nil {
		return err
	}
	return nil
}
