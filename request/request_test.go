package request

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRequestDbParamsJSONMarshalling(t *testing.T) {
	// Create a sample requestDbParams
	originalParams := requestDbParams{
		GameId: uuid.Must(uuid.NewV4()),
		Origin: Origin("sample_origin"),
		Colour: HexColour("#aabbcc"),
		Speed:  5,
		Width:  10,
		Status: Status("sample_status"),
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(originalParams)
	assert.NoError(t, err, "Marshalling should not produce an error")

	// Unmarshal back to requestDbParams
	var unmarshalledParams requestDbParams
	err = json.Unmarshal(jsonData, &unmarshalledParams)
	assert.NoError(t, err, "Unmarshalling should not produce an error")

	// Assert equality
	assert.Equal(t, originalParams, unmarshalledParams, "Original and unmarshalled params should be equal")
}

func TestUnmarshalFromBytes(t *testing.T) {
	byteArray := bytes.NewBuffer("{\"id\":\"f7172fa6-04c2-415c-ad4b-eb718740bda9\",\"game_id\":\"7188938e-57de-45db-ac1f-309bd9fb5002\",\"created_at\":\"2024-09-13T12:57:05.412681+10:00\",\"origin\":\"top\",\"colour\":\"#78c9e6\",\"speed\":5,\"width\":5,\"status\":\"success\"}")
}
