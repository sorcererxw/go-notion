package notion

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAsError(t *testing.T) {
	err := &Error{
		Status:  400,
		Code:    ErrCodeRateLimited,
		Message: "rate limit",
	}
	err, ok := AsError(err)
	require.True(t, ok)
	require.Equal(t, "rate limit", err.Message)
}
