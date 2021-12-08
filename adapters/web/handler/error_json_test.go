package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler_jsonError(t *testing.T) {
	msg := "Hello"
	result := jsonError(msg)
	require.Equal(t, []byte(`{"message":"Hello"}`), result)
}
