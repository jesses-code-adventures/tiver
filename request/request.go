package request

import (
	"errors"
	"regexp"
	"time"

	"github.com/gofrs/uuid"
)

type HexColour string

func IsValidHexColour(s string) bool {
	matched, _ := regexp.MatchString(`^#[0-9A-Fa-f]{6}$`, s)
	return matched
}

func HexColourFromString(s string) (hc HexColour, err error) {
	if !IsValidHexColour(s) {
		return hc, errors.New("Not a valid hex colour")
	}
	return HexColour(s), err
}

func (h HexColour) String() string {
	return string(h)
}

type Request struct {
	Id        uuid.UUID `json:"id"`
	GameId    uuid.UUID `json:"game_id"`
	CreatedAt time.Time `json:"created_at"`
	Origin    Origin    `json:"origin"`
	Colour    HexColour `json:"colour"`
	Speed     int       `json:"speed"`
	Width     int       `json:"width"`
	Status    Status    `json:"status"`
}

func NewRequest(id uuid.UUID, gameId uuid.UUID, createdAt time.Time, origin Origin, colour HexColour, speed int, width int, status Status, requests int) Request {
	return Request{
		Id:        id,
		GameId:    gameId,
		CreatedAt: createdAt,
		Origin:    origin,
		Colour:    colour,
		Speed:     speed,
		Width:     width,
		Status:    status,
	}
}
