package notion

import "errors"

type Error struct {
	Code    string
	Message string
}

func (e *Error) Error() string {
	return e.Code + " " + e.Message
}

var ErrInvalidJSON = errors.New("invalid_json")
