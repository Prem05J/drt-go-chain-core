package pubkeyConverter

import (
	"encoding/hex"
	"fmt"

	"github.com/Prem05J/drt-go-chain-core/core"
)

// hexPubkeyConverter encodes or decodes provided public key as/from hex
type hexPubkeyConverter struct {
	len int
}

// NewHexPubkeyConverter returns a hexPubkeyConverter instance
func NewHexPubkeyConverter(addressLen int) (*hexPubkeyConverter, error) {
	if addressLen < 1 {
		return nil, fmt.Errorf("%w when creating hex address converter, addressLen should have been greater than 0",
			ErrInvalidAddressLength)
	}
	if addressLen%2 == 1 {
		return nil, fmt.Errorf("%w when creating hex address converter, addressLen should have been an even number",
			ErrInvalidAddressLength)
	}

	return &hexPubkeyConverter{
		len: addressLen,
	}, nil
}

// Decode converts the provided public key string as hex decoded bytes
func (ppc *hexPubkeyConverter) Decode(humanReadable string) ([]byte, error) {
	buff, err := hex.DecodeString(humanReadable)
	if err != nil {
		return nil, err
	}

	if len(buff) != ppc.len {
		return nil, fmt.Errorf("%w when converting to address, expected length %d, received %d",
			ErrWrongSize, ppc.len, len(buff))
	}

	return buff, nil
}

// Encode converts the provided bytes in a form that this converter can decode. In this case it will encode to hex
func (ppc *hexPubkeyConverter) Encode(pkBytes []byte) (string, error) {
	return hex.EncodeToString(pkBytes), nil
}

// EncodeSlice converts the provided bytes slice in a form that this converter can decode. In this case it will encode to hex
func (ppc *hexPubkeyConverter) EncodeSlice(pkBytesSlice [][]byte) ([]string, error) {
	encodedSlice := make([]string, 0, len(pkBytesSlice))

	for _, item := range pkBytesSlice {
		encodedSlice = append(encodedSlice, hex.EncodeToString(item))
	}

	return encodedSlice, nil
}

// SilentEncode converts the provided bytes in a form that this converter can decode. In this case it will encode to hex
func (ppc *hexPubkeyConverter) SilentEncode(pkBytes []byte, log core.Logger) string {
	return hex.EncodeToString(pkBytes)
}

// Len returns the decoded address length
func (ppc *hexPubkeyConverter) Len() int {
	return ppc.len
}

// IsInterfaceNil returns true if there is no value under the interface
func (ppc *hexPubkeyConverter) IsInterfaceNil() bool {
	return ppc == nil
}
