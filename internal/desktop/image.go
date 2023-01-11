package desktop

import (
	"bombom/resources"

	"bytes"
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type player struct {
	image *ebiten.Image
}

func newPlayer() (*player, error) {
	image, err := newImage(resources.PlayerImage)
	if err != nil {
		return nil, fmt.Errorf("failed to create the player image: %w", err)
	}

	return &player{
		image: image,
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
	op := &ebiten.DrawImageOptions{}

	screen.DrawImage(p.image, op)
}
