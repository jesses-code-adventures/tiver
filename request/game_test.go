package request

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGameJSONKeys(t *testing.T) {
	id := uuid.Must(uuid.NewV4())
	createdAt := time.Now()
	var endedAt *time.Time
	requests := int32(5)

	game := NewGame(id, createdAt, endedAt, requests)
	gameJSON, err := json.Marshal(game)
	assert.NoError(t, err)

	var result map[string]interface{}
	err = json.Unmarshal(gameJSON, &result)
	assert.NoError(t, err)

	expectedKeys := []string{"id", "created_at", "ended_at", "requests"}
	for _, key := range expectedKeys {
		_, exists := result[key]
		assert.True(t, exists, "JSON key '%s' should be present", key)
	}
}
