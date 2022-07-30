package gotelebot

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBot_GetMe(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		resp, _ := json.Marshal(ApiResponse[User]{
			Ok: true,
			Result: User{
				ID: 1,
			},
		})
		_, err := w.Write(resp)
		assert.NoError(t, err)
	}))

	bot, _ := NewBot(server.URL, "token")

	user, err := bot.GetMe()

	assert.NoError(t, err)
	assert.Equal(t, int64(1), user.ID)
}
