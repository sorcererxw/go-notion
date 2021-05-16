package notion

import (
	"errors"
)

type ErrCode string

const (
	// ErrCodeInvalidJson The request body could not be decoded as JSON.
	ErrCodeInvalidJson ErrCode = "invalid_json"
	// ErrCodeInvalidRequestUrl The request URL is not valid.
	ErrCodeInvalidRequestUrl ErrCode = "invalid_request_url"
	// ErrCodeInvalidRequest This request is not supported.
	ErrCodeInvalidRequest ErrCode = "invalid_request"
	// ErrCodeValidationError The request body does not match the schema for the expected parameters.
	// Check the "message" property for more details.
	ErrCodeValidationError ErrCode = "validation_error"
	// ErrCodeUnauthorized The bearer token is not valid.
	ErrCodeUnauthorized ErrCode = "unauthorized"
	// ErrCodeRestrictedResource Given the bearer token used, the client doesn't have permission to perform this operation.
	ErrCodeRestrictedResource ErrCode = "restricted_resource"
	// ErrCodeObjectNotFound Given the bearer token used, the resource does not exist.
	// This error can also indicate that the resource has not been shared with owner of the bearer token.
	ErrCodeObjectNotFound ErrCode = "object_not_found"
	// ErrCodeConflictError The transaction could not be completed, potentially due to a data collision.
	// Make sure the parameters are up to date and try again.
	ErrCodeConflictError ErrCode = "conflict_error"
	// ErrCodeRateLimited This request exceeds the number of requests allowed.
	// Slow down and try again.More details on rate limits
	ErrCodeRateLimited ErrCode = "rate_limited"
	// ErrCodeInternalServerError An unexpected error occurred.Reach out to Notion support.
	ErrCodeInternalServerError ErrCode = "internal_server_error"
	// ErrCodeServiceUnavailable Notion is unavailable. Try again later.
	// This can occur when the time to respond to a request takes longer than 60 seconds, the maximum request timeout.
	ErrCodeServiceUnavailable ErrCode = "service_unavailable"
)

type Error struct {
	Status  int     `json:"status,omitempty"`
	Code    ErrCode `json:"code,omitempty"`
	Message string  `json:"message,omitempty"`
}

func (e *Error) Error() string {
	return e.Message
}

// AsError tries casting the basic error to notion error
func AsError(e error) (err *Error, ok bool) {
	if e == nil {
		return nil, false
	}
	if errors.As(e, &err) {
		return err, true
	}
	return nil, false
}
