package gotelebot

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBot(t *testing.T) {
	t.Parallel()

	bot, err := NewBot(TelegramAPIServer, "token")

	assert.NoError(t, err)
	assert.NotNil(t, bot)
	assert.NotNil(t, bot.client)
}
