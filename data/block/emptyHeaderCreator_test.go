package block

import (
	"fmt"
	"testing"

	"github.com/Prem05J/drt-go-chain-core/core/check"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewEmptyHeaderCreator(t *testing.T) {
	t.Parallel()

	creator := NewEmptyHeaderCreator()
	require.False(t, check.IfNil(creator))
}

func TestEmptyHeaderCreator_CreateNewHeader(t *testing.T) {
	t.Parallel()

	creator := NewEmptyHeaderCreator()
	header := creator.CreateNewHeader()
	require.False(t, check.IfNil(header))
	assert.Equal(t, "*block.Header", fmt.Sprintf("%T", header))
}
