package event

import "reflect"

type Event interface {
	Type() string
}

type autoTypedEvent struct {
	typeName string
	rawEvent any
}

func New(rawEvent any) Event {
	return autoTypedEvent{
		typeName: reflect.TypeOf(rawEvent).Name(),
		rawEvent: rawEvent,
	}
}

func (e autoTypedEvent) Type() string {
	return e.typeName
}

type jsonSerializedEvent struct {
	typeName string
	data     []byte
}

func NewJSON(typeName string, data []byte) Event {
	return jsonSerializedEvent{
		typeName: typeName,
		data:     data,
	}
}

func (e jsonSerializedEvent) Type() string {
	return e.typeName
}
