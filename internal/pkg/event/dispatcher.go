package event

import (
	"context"
	"errors"
	"fmt"
)

var ErrNotRegistered = errors.New("event is not registered")

type Dispatcher struct {
	handlerGroups map[string][]Handler
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		handlerGroups: make(map[string][]Handler),
	}
}

func (d *Dispatcher) RegisterHandler(eventType string, h Handler) {
	handlerGroup, ok := d.handlerGroups[eventType]
	if !ok {
		handlerGroup = make([]Handler, 0, 1)
	}

	d.handlerGroups[eventType] = append(handlerGroup, h)
}

func (d *Dispatcher) Dispatch(ctx context.Context, events ...Event) error {
	for _, event := range events {
		if err := d.dispatch(ctx, event); err != nil {
			return err
		}
	}

	return nil
}

func (d *Dispatcher) dispatch(ctx context.Context, e Event) error {
	handlerGroup, ok := d.handlerGroups[e.Type()]
	if !ok {
		return fmt.Errorf("%w: %s", ErrNotRegistered, e.Type())
	}

	for i, h := range handlerGroup {
		if err := h.Handle(ctx, e); err != nil {
			return fmt.Errorf("failed to handle %s by handler %d: %w", e.Type(), i+1, err)
		}
	}

	return nil
}
