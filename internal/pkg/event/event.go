package event

import "reflect"

type Event interface {
	Type() string
}

type baseEvent struct {
	typeName string
	rawEvent any
}

func New(rawEvent any) Event {
	return &baseEvent{
		typeName: reflect.TypeOf(rawEvent).Name(),
		rawEvent: rawEvent,
	}
}

func (e *baseEvent) Type() string {
	return e.typeName
}
