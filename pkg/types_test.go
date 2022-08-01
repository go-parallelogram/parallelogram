package gotelebot

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdate_Context(t *testing.T) {
	const (
		key   = "key"
		value = "value"
	)

	upd := &Update{}

	assert.Equal(t, upd.Context(), context.Background())

	valuedContext := context.WithValue(upd.Context(), key, value)

	upd = upd.WithContext(valuedContext)

	assert.Equal(t, upd.Context(), valuedContext)
}

func TestUpdate_Bot(t *testing.T) {
	bot := &Bot{}

	upd := &Update{}

	assert.Equal(t, (*Bot)(nil), upd.Bot())

	upd = upd.WithBot(bot)

	assert.Equal(t, upd.Bot(), bot)
}
