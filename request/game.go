package request

import (
	"time"

	"github.com/gofrs/uuid"
)

type Game struct {
	Id        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	EndedAt   *time.Time `json:"ended_at"`
}

func (g Game) End() {
	now := time.Now()
	g.EndedAt = &now
}
