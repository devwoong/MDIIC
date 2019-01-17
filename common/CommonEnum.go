package common

type MSG_TYPE int8
type MSG_CODE int32

const (
	MSG_STRING MSG_TYPE = iota
	MSG_ALIVE
	MSG_EXIT
	MSG_KEYBOARD
	MSG_MOUSE
	MSG_SCREEN
)

const (
	KEYBOARD_DOWN MSG_CODE = iota
	KEYBOARD_UP
)

const (
	MOUSE_RIGHT_BTDOWN MSG_CODE = iota
	MOUSE_RIGHT_BTUP
	MOUSE_LEFT_BTDOWN
	MOUSE_LEFT_BTUP
	MOUSE_MOVE
)

const (
	SCREEN_FOCUS_CHANGE MSG_CODE = iota
)

// type Baser interface {
// 	Base() base
// }

// // every base must fulfill the Baser interface
// func (b base) Base() base {
// 	return b
// }
