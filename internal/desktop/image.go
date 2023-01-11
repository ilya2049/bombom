package desktop

import (
	"bombom/resources"

	"bytes"
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type corePlayer interface {
	X() float64
	Y() float64
	Move()
}

type player struct {
	corePlayer

	image *ebiten.Image
}

func newPlayer(coords corePlayer) (*player, error) {
	image, err := newImage(resources.PlayerImage)
	if err != nil {
		return nil, fmt.Errorf("failed to create the player image: %w", err)
	}

	return &player{
		corePlayer: coords,
		image:      image,
	}, nil
}

func newImage(imageBytes []byte) (*ebiten.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to decode an image: %w", err)
	}

	return ebiten.NewImageFromImage(img), nil
}

func (p *player) drawOnScreen(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(p.X(), p.Y())

	screen.DrawImage(p.image, options)
}
