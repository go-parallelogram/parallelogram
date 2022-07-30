package gotelebot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBot(t *testing.T) {
	t.Parallel()

	bot, err := NewBot(TelegramAPIServer, "token")

	assert.NoError(t, err)
	assert.NotNil(t, bot)
	assert.NotNil(t, bot.client)
}
