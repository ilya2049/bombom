package input

import (
	"bombom/internal/game"

	"context"
)

var directionMap = map[Key]game.Direction{
	KeyUp:    game.DirectionUp,
	KeyDown:  game.DirectionDown,
	KeyRight: game.DirectionRight,
	KeyLeft:  game.DirectionLeft,
}

type Player struct {
	*game.Player
}

func NewPlayer(p *game.Player) *Player {
	return &Player{
		Player: p,
	}
}

func (p *Player) handleKeyPressed(_ context.Context, e KeyPressed) error {
	if direction, ok := directionMap[e.Key]; ok {
		p.StartMoving(direction)
	}

	return nil
}

func (p *Player) handleKeyReleased(_ context.Context, e KeyReleased) error {
	if _, ok := directionMap[e.Key]; ok {
		p.StopMoving()
	}

	return nil
}
