package api_test

import (
	"testing"

	"github.com/Prem05J/drt-go-chain-core/data/api"
	"github.com/stretchr/testify/require"
)

func TestAPIBlockFetchType(t *testing.T) {
	byNonceType := api.BlockFetchTypeByNonce
	require.Equal(t, "by-nonce", byNonceType.String())

	byHashType := api.BlockFetchTypeByHash
	require.Equal(t, "by-hash", byHashType.String())
}
