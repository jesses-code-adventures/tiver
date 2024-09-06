package request

import (
	"errors"
	"fmt"
	"strings"
)

type Origin string

const (
	Left Origin = "left"
	Top  Origin = "top"
)

func (o Origin) String() string {
	return string(o)
}

func OriginFromString(s string) (o Origin, err error) {
	s = strings.ToLower(s)
	switch s {
	case "left":
		return Left, err
	case "top":
		return Top, err
	default:
		return o, errors.New(fmt.Sprintf("unacceptable origin value: %s", s))
	}
}
