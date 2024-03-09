package db

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Test if the software throws an error if the Database get filled with the mockdata twice.
func TestMockCreate(t *testing.T) {
	provider := NewMockProvider()

	// Fill Database the first time and check if error
	err := prefillMockDB(provider)
	require.Nil(t, err)

	// Fill Database the second time and check if error
	err = prefillMockDB(provider)
	require.Nil(t, err)
}
