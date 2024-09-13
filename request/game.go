package request

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jesses-code-adventures/tiver/model"
)

func convertUUID(u pgtype.UUID) uuid.UUID {
	uuidVal, err := uuid.FromBytes(u.Bytes[:])
	if err != nil {
		panic("UUID conversion failed")
	}
	return uuidVal
}

type Game struct {
	Id        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	EndedAt   *time.Time `json:"ended_at"`
	Requests  int32      `json:"requests"`
}

func NewGame(id uuid.UUID, createdAt time.Time, endedAt *time.Time, requests int32) Game {
	return Game{Id: id, CreatedAt: createdAt, EndedAt: endedAt, Requests: requests}
}

func GameFromDbModel(dbModel model.Game) Game {
	return NewGame(convertUUID(dbModel.Id), dbModel.CreatedAt.Time, &dbModel.EndedAt.Time, dbModel.Requests)
}

func GameFromResponse(r *http.Response) (g Game, error error) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("got error reading body")
		return
	}
	defer r.Body.Close()
	err = json.Unmarshal(bodyBytes, &g)
	if err != nil {
		return
	}
	return
}
