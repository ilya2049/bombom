package desktop

import (
	"bombom/internal/game"
	"bombom/internal/input"
	"bombom/internal/pkg/event"
	"fmt"

	"context"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewGame() (*Game, error) {
	player := game.NewPlayer()

	desktopPlayer, err := newPlayer(player)
	if err != nil {
		return nil, fmt.Errorf("failed to create the player: %w", err)
	}

	eventDispatcher := event.NewDispatcher()

	input.RegisterKeyHandlers(input.NewPlayer(player), eventDispatcher)

	return &Game{
		desktopPlayer:   desktopPlayer,
		eventDispatcher: eventDispatcher,
	}, nil
}

type Game struct {
	eventDispatcher *event.Dispatcher

	desktopPlayer *player
}

func (g *Game) Update() error {
	g.desktopPlayer.Move()

	var events []event.Event

	events = append(events, readInputEvents()...)

	return g.eventDispatcher.Dispatch(context.Background(), events...)
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.desktopPlayer.drawOnScreen(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
