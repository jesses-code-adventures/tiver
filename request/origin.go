package request

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jesses-code-adventures/tiver/model"
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

func OriginFromDbModel(dbModel model.Origin) (o Origin, err error) {
	var s string
	err = dbModel.Scan(&s)
	if err != nil {
		return o, err
	}
	return OriginFromString(s)
}

func (o Origin) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, o)), nil
}

func (o *Origin) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	var err error
	*o, err = OriginFromString(s)
	return err
}
