package input

type Key string

const (
	KeyUnknown Key = ""
	KeyUp      Key = "up"
	KeyDown    Key = "down"
	KeyLeft    Key = "left"
	KeyRight   Key = "right"
	KeySpace   Key = "space"
)

type KeyPressed struct {
	Key Key
}

type KeyReleased struct {
	Key Key
}
