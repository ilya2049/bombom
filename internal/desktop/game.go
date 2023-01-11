package desktop

import (
	"bombom/internal/input"
	"bombom/internal/pkg/event"
	"fmt"

	"context"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewGame() (*Game, error) {
	aPlayer, err := newPlayer()
	if err != nil {
		return nil, fmt.Errorf("failed to create the player: %w", err)
	}

	eventDispatcher := event.NewDispatcher()

	input.RegisterKeyHandlers(eventDispatcher)

	return &Game{
		player:          aPlayer,
		eventDispatcher: eventDispatcher,
	}, nil
}

type Game struct {
	eventDispatcher *event.Dispatcher

	player *player
}

func (g *Game) Update() error {
	var events []event.Event

	events = append(events, readInputEvents()...)

	return g.eventDispatcher.Dispatch(context.Background(), events...)
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.drawOnScreen(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
