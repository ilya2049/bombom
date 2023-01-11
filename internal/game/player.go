package game

type Player struct {
	x float64
	y float64

	speed     float64
	direction Direction
}

func NewPlayer() *Player {
	return &Player{
		speed: 1,
	}
}

func (p *Player) X() float64 {
	return p.x
}

func (p *Player) Y() float64 {
	return p.y
}

func (p *Player) StartMoving(d Direction) {
	p.direction = d
}

func (p *Player) StopMoving() {
	p.direction = DirectionUnknown
}

func (p *Player) Move() {
	switch p.direction {
	case DirectionUp:
		p.y -= p.speed
	case DirectionDown:
		p.y += p.speed
	case DirectionRight:
		p.x += p.speed
	case DirectionLeft:
		p.x -= p.speed
	}
}
