package notion

import "time"

type Date struct {
	Start time.Time
	// If null, this property's date value is not a range.
	End *time.Time
}

type Error struct {
	Status  int    `json:"status,omitempty"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e *Error) Error() string {
	return e.Code + ":" + e.Message
}
