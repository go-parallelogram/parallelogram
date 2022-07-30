package gotelebot

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunPolling(t *testing.T) {
	t.Parallel()

	bot, _ := NewBot("", "token")

	assert.NoError(t, RunPolling(bot, HandlerFunc(func(upd Update) {
		updateJson, _ := json.Marshal(upd)
		log.Println(string(updateJson))
	})))
}
