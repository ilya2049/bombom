package desktop

import (
	"image"
	"log"
	"os"
	"path"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewGame() *Game {
	bombermanImageFile, err := os.Open(path.Join("resources", "bomberman.png"))
	if err != nil {
		panic(err)
	}

	defer bombermanImageFile.Close()

	decodedBombermanImage, _, err := image.Decode(bombermanImageFile)
	if err != nil {
		log.Fatal(err)
	}
	bombermanImage := ebiten.NewImageFromImage(decodedBombermanImage)

	return &Game{
		bombermanImage: bombermanImage,
	}
}

type Game struct {
	bombermanImage *ebiten.Image
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.bombermanImage, &ebiten.DrawImageOptions{})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
