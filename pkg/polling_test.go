package gotelebot

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunPolling(t *testing.T) {
	t.Skip("Not implemented")
	t.Parallel()

	bot, _ := NewBot("", "token")

	assert.NoError(t, RunPolling(bot, HandlerFunc(func(upd *Update) {
		updateJSON, _ := json.Marshal(upd)
		log.Println(string(updateJSON))
	})))
}
