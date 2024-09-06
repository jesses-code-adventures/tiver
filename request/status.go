package request

import (
	"errors"
	"fmt"
	"log/slog"
)

type Status string

const (
	Init    Status = "init"
	Retry   Status = "retry"
	Success Status = "success"
	Error   Status = "error"
)

func StatusFromString(s string) (st Status, err error) {
	switch s {
	case "init":
		st = "init"
	case "retry":
		st = "retry"
	case "success":
		st = "success"
	case "error":
		st = "error"
	default:
		msg := fmt.Sprintf("invalid string passed to StatusFromString: %s", s)
		slog.Error(msg)
		return st, errors.New(msg)
	}
	return
}

func (s Status) String() string {
	switch s {
	case Init:
		return "init"
	case Retry:
		return "retry"
	case "success":
		return "success"
	case "error":
		return "error"
	default:
		slog.Error(fmt.Sprintf("invalid Status in s.Status: %v", s))
		return ""
	}
}
