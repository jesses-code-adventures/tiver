package request

import (
	"encoding/hex"
	"errors"
)

type HexString string

func (hs HexString) Bytes() ([]byte, error) {
	return hex.DecodeString(string(hs))
}

func NewHexString(b []byte) HexString {
	return HexString(hex.EncodeToString(b))
}

func (hs HexString) Validate() error {
	_, err := hs.Bytes()
	if err != nil {
		return errors.New("invalid hex string")
	}
	return nil
}
