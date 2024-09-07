// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package model

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type Origin string

const (
	OriginTop  Origin = "top"
	OriginLeft Origin = "left"
)

func (e *Origin) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Origin(s)
	case string:
		*e = Origin(s)
	default:
		return fmt.Errorf("unsupported scan type for Origin: %T", src)
	}
	return nil
}

type NullOrigin struct {
	Origin Origin
	Valid  bool // Valid is true if Origin is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullOrigin) Scan(value interface{}) error {
	if value == nil {
		ns.Origin, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Origin.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullOrigin) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Origin), nil
}

type RequestStatus string

const (
	RequestStatusInit    RequestStatus = "init"
	RequestStatusRetry   RequestStatus = "retry"
	RequestStatusSuccess RequestStatus = "success"
	RequestStatusError   RequestStatus = "error"
)

func (e *RequestStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = RequestStatus(s)
	case string:
		*e = RequestStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for RequestStatus: %T", src)
	}
	return nil
}

type NullRequestStatus struct {
	RequestStatus RequestStatus
	Valid         bool // Valid is true if RequestStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullRequestStatus) Scan(value interface{}) error {
	if value == nil {
		ns.RequestStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.RequestStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullRequestStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.RequestStatus), nil
}

type RiverJobState string

const (
	RiverJobStateAvailable RiverJobState = "available"
	RiverJobStateCancelled RiverJobState = "cancelled"
	RiverJobStateCompleted RiverJobState = "completed"
	RiverJobStateDiscarded RiverJobState = "discarded"
	RiverJobStatePending   RiverJobState = "pending"
	RiverJobStateRetryable RiverJobState = "retryable"
	RiverJobStateRunning   RiverJobState = "running"
	RiverJobStateScheduled RiverJobState = "scheduled"
)

func (e *RiverJobState) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = RiverJobState(s)
	case string:
		*e = RiverJobState(s)
	default:
		return fmt.Errorf("unsupported scan type for RiverJobState: %T", src)
	}
	return nil
}

type NullRiverJobState struct {
	RiverJobState RiverJobState
	Valid         bool // Valid is true if RiverJobState is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullRiverJobState) Scan(value interface{}) error {
	if value == nil {
		ns.RiverJobState, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.RiverJobState.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullRiverJobState) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.RiverJobState), nil
}

type Game struct {
	Id        pgtype.UUID
	CreatedAt pgtype.Timestamptz
	EndedAt   pgtype.Timestamptz
	Requests  int32
}

type Request struct {
	Id        pgtype.UUID
	CreatedAt pgtype.Timestamptz
	EndedAt   pgtype.Timestamptz
	GameID    pgtype.UUID
	Colour    string
	Origin    Origin
	Speed     int32
	Width     int32
	Status    RequestStatus
}

type RiverClient struct {
	Id        string
	CreatedAt pgtype.Timestamptz
	Metadata  []byte
	PausedAt  pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

type RiverClientQueue struct {
	RiverClientID    string
	Name             string
	CreatedAt        pgtype.Timestamptz
	MaxWorkers       int64
	Metadata         []byte
	NumJobsCompleted int64
	NumJobsRunning   int64
	UpdatedAt        pgtype.Timestamptz
}

type RiverJob struct {
	Id          int64
	State       RiverJobState
	Attempt     int16
	MaxAttempts int16
	AttemptedAt pgtype.Timestamptz
	CreatedAt   pgtype.Timestamptz
	FinalizedAt pgtype.Timestamptz
	ScheduledAt pgtype.Timestamptz
	Priority    int16
	Args        []byte
	AttemptedBy []string
	Errors      [][]byte
	Kind        string
	Metadata    []byte
	Queue       string
	Tags        []string
	UniqueKey   []byte
}

type RiverLeader struct {
	ElectedAt pgtype.Timestamptz
	ExpiresAt pgtype.Timestamptz
	LeaderID  string
	Name      string
}

type RiverQueue struct {
	Name      string
	CreatedAt pgtype.Timestamptz
	Metadata  []byte
	PausedAt  pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}
