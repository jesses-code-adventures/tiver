package request

import (
	crand "crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jesses-code-adventures/tiver/model"
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

func MustGenerateRandHexColour() HexColour {
	b := make([]byte, 3)
	_, err := crand.Read(b)
	if err != nil {
		panic(err.Error())
	}
	return HexColour(fmt.Sprintf("#%02x%02x%02x", b[0], b[1], b[2]))
}

func (h HexColour) String() string {
	return string(h)
}

func (h HexColour) MarshalJSON() ([]byte, error) {
	return json.Marshal(h.String())
}

func (h *HexColour) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if !IsValidHexColour(s) {
		return errors.New("Not a valid hex colour")
	}
	*h = HexColour(s)
	return nil
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

func newRequest(id uuid.UUID, gameId uuid.UUID, createdAt time.Time, origin Origin, colour HexColour, speed int, width int, status Status) Request {
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

func RequestFromDbModel(r model.Request) Request {
	return Request{
		Id:        convertUUID(r.Id),
		GameId:    convertUUID(r.GameID),
		CreatedAt: r.CreatedAt.Time,
		Origin:    Origin(r.Origin),
		Colour:    HexColour(r.Colour),
		Speed:     int(r.Speed),
		Width:     int(r.Width),
		Status:    Status(r.Status),
	}
}

type IncomingRequest struct {
	GameId uuid.UUID `json:"game_id"`
	Origin Origin    `json:"origin"`
	Colour HexColour `json:"colour"`
	Speed  int       `json:"speed"`
	Width  int       `json:"width"`
}

func newIncomingRequest(gameId uuid.UUID, origin Origin, colour HexColour, speed int, width int) IncomingRequest {
	return IncomingRequest{
		GameId: gameId,
		Origin: origin,
		Colour: colour,
		Speed:  speed,
		Width:  width,
	}
}

func toPgtypeUuid(u uuid.UUID) pgtype.UUID {
	var pgUUID pgtype.UUID
	pgUUID.Scan(u)
	return pgUUID
}

func RequestDbParamsFromSenderBody(r *http.Request) (req model.CreateRequestParams, err error) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print("error reading body")
		return
	}
	defer r.Body.Close()
	var params IncomingRequest
	err = json.Unmarshal(bodyBytes, &params)
	if err != nil {
		return
	}
	return model.CreateRequestParams{
		GameID: toPgtypeUuid(params.GameId),
		Colour: params.Colour.String(),
		Origin: model.Origin(params.Origin),
		Speed:  int32(params.Speed),
		Status: model.RequestStatus(Init),
	}, err
}

func getRandomItem[T any](items ...T) T {
	return items[rand.Intn(len(items))]
}

func IncomingRequestDummy(gameId uuid.UUID) IncomingRequest {
	return newIncomingRequest(gameId, getRandomItem(Left, Top), MustGenerateRandHexColour(), 5, 5)
}

func Dummy(gameId uuid.UUID) Request {
	return newRequest(uuid.Must(uuid.NewV4()), gameId, time.Now(), getRandomItem(Left, Top), MustGenerateRandHexColour(), 5, 5, getRandomItem(Init, Success, Error, Retry))
}
