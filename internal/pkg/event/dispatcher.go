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

func (r *Dispatcher) RegisterHandler(eventType string, h Handler) {
	handlerGroup, ok := r.handlerGroups[eventType]
	if !ok {
		handlerGroup = make([]Handler, 0, 1)
	}

	r.handlerGroups[eventType] = append(handlerGroup, h)
}

func (r *Dispatcher) Dispatch(ctx context.Context, e Event) error {
	handlerGroup, ok := r.handlerGroups[e.Type()]
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
