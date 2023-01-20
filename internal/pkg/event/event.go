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
	return baseEvent{
		typeName: reflect.TypeOf(rawEvent).Name(),
		rawEvent: rawEvent,
	}
}

func (e baseEvent) Type() string {
	return e.typeName
}

type encoded struct {
	typeName string
	data     []byte
}

func NewEncoded(typeName string, data []byte) Event {
	return encoded{
		typeName: typeName,
		data:     data,
	}
}

func (e encoded) Type() string {
	return e.typeName
}
